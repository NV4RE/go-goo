package authentication

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/nv4re/go-goo/entity/authentication"
	"github.com/nv4re/go-goo/entity/errors"
	"github.com/nv4re/go-goo/repository"
)

const usernameKey = "username"
const roleNameKey = "name"

type mongoAuthenticationRepository struct {
	db      *mongo.Database
	userCol *mongo.Collection
	roleCol *mongo.Collection
}

func NewMongoAuthenticationRepository(database *mongo.Database) (repository.AuthRepository, error) {
	return &mongoAuthenticationRepository{
		database,
		database.Collection("user"),
		database.Collection("role"),
	}, nil
}

func (m *mongoAuthenticationRepository) CreateUser(ctx context.Context, user *authentication.User) error {
	_, err := m.userCol.InsertOne(ctx, user)
	return err
}

func (m *mongoAuthenticationRepository) UpdateUser(ctx context.Context, user *authentication.User) error {
	res, err := m.userCol.UpdateOne(ctx, bson.D{{usernameKey, user.Username}}, bson.D{{"$set", user}})
	if err != nil {
		return err
	}
	if res.UpsertedCount != 1 {
		return errors.NotFoundError
	}
	return nil
}

func (m *mongoAuthenticationRepository) GetUserByUsername(ctx context.Context, username string) (*authentication.User, error) {
	var u authentication.User
	res := m.userCol.FindOne(ctx, bson.D{{usernameKey, username}})
	if err := res.Decode(&u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (m *mongoAuthenticationRepository) SearchUser(ctx context.Context, username string, from, to int) ([]*authentication.User, error) {
	opts := options.Find()
	opts.SetSkip(int64(from))
	opts.SetLimit(int64(to - from))

	cur, err := m.userCol.Find(
		ctx,
		bson.D{{usernameKey, primitive.Regex{Pattern: username, Options: ""}}},
		opts,
	)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var users []*authentication.User
	for cur.Next(ctx) {
		var u authentication.User
		err := cur.Decode(&u)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (m *mongoAuthenticationRepository) CreateRole(ctx context.Context, role *authentication.Role) error {
	_, err := m.roleCol.InsertOne(ctx, role)
	return err
}

func (m *mongoAuthenticationRepository) UpdateRole(ctx context.Context, role *authentication.Role) error {
	res, err := m.roleCol.UpdateOne(ctx, bson.D{{roleNameKey, role.Name}}, bson.D{{"$set", role}})
	if err != nil {
		return err
	}
	if res.UpsertedCount != 1 {
		return errors.NotFoundError
	}
	return nil
}

func (m *mongoAuthenticationRepository) DeleteRoleByName(ctx context.Context, name string) error {
	res, err := m.roleCol.DeleteOne(ctx, bson.D{{roleNameKey, name}})
	if err != nil {
		return err
	}
	if res.DeletedCount != 1 {
		return errors.NotFoundError
	}
	return nil
}

func (m *mongoAuthenticationRepository) GetRoleByName(ctx context.Context, name string) (*authentication.Role, error) {
	var r authentication.Role
	res := m.roleCol.FindOne(ctx, bson.D{{roleNameKey, name}})
	if err := res.Decode(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (m *mongoAuthenticationRepository) ListRole(ctx context.Context, from, to int) ([]*authentication.Role, error) {
	opts := options.Find()
	opts.SetSkip(int64(from))
	opts.SetLimit(int64(to - from))

	cur, err := m.roleCol.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var Role []*authentication.Role
	for cur.Next(ctx) {
		var r authentication.Role
		err := cur.Decode(&r)
		if err != nil {
			return nil, err
		}
		Role = append(Role, &r)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return Role, nil
}

func (m *mongoAuthenticationRepository) IsHealthy(ctx context.Context) error {
	return m.db.Client().Ping(ctx, readpref.Primary())
}

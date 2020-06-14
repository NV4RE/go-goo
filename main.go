package main

import (
	"log"
	"time"

	_ "net/http/pprof"

	"github.com/caarlos0/env/v6"
	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"

	httpGin "github.com/nv4re/go-goo/interface/gin"
)

type config struct {
	AppName     string `env:"APP_NAME" envDefault:"go-goo"`
	Environment string `env:"ENVIRONMENT" envDefault:"local"`
	Server      struct {
		Host string `env:"SERVER_HOST"`
		Port int    `env:"SERVER_PORT" envDefault:"8801"`
	}
	Jaeger struct {
		Host string `env:"JAEGER_AGENT_HOST" envDefault:"localhost"`
		Port int    `env:"JAEGER_AGENT_PORT" envDefault:"6832"`
	}
	Redis struct {
		Address string `env:"REDIS_ADDR" envDefault:"localhost:16379"`
		DB      int    `env:"REDIS_DB" envDefault:"10"`
	}
	MONGODB struct {
		URI string `env:"MONGODB_URI" envDefault:"mongodb://localhost:27017"`
	}
}

func main() {
	// Load config
	c, err := setupEnv()
	if err != nil {
		log.Fatalln(err)
	}

	// Connect to DBs
	redisPool, err := setupDatabases(c)
	if err != nil {
		log.Fatalln(err)
	}

	// Setup repositories
	err = setupRepositories(c, redisPool) // jaegerTracer, reminderRepo, dLock, err
	if err != nil {
		log.Fatalln(err)
	}

	// Setup use-cases

	// Setup interfaces
	g := httpGin.NewGinServer(c.Server.Host, c.Server.Port)
	err = g.Start()
	if err != nil {
		log.Fatalln(err)
	}
}

func setupEnv() (*config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	c := &config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}

func setupDatabases(c *config) (*redis.Pool, error) {
	redisPool := &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", c.Redis.Address, redis.DialDatabase(c.Redis.DB))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return redisPool, nil
}

func setupRepositories(c *config, redisPool *redis.Pool) error {
	return nil
}

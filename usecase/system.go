package usecase

import (
	"context"

	"github.com/nv4re/go-goo/repository"
)

type systemUseCase struct {
	authRepo repository.AuthRepository
}

func NewSystemUseCase(authRepo repository.AuthRepository) SystemUseCase {
	return &systemUseCase{authRepo}
}

func (s *systemUseCase) IsAllHealthy(ctx context.Context) error {
	return s.authRepo.IsHealthy(ctx)
}

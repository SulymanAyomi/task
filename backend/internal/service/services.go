package service

import (
	"github.com/SulymanAyomi/task/internal/lib/job"
	"github.com/SulymanAyomi/task/internal/repository"
	"github.com/SulymanAyomi/task/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}

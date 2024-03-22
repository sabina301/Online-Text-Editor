package app

import (
	"Online-Text-Editor/server/internal/configuration"
	userImpl "Online-Text-Editor/server/internal/handler/user"
	userRepo "Online-Text-Editor/server/internal/repository/user"
	userService "Online-Text-Editor/server/internal/service/user"
	"github.com/jmoiron/sqlx"
	"log"
)

type serviceProvider struct {
	grpcConfiguration configuration.GrpcConfiguration
	userRepository    userService.Repository
	userService       userImpl.Service
	userImpl          *userImpl.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() configuration.GrpcConfiguration {
	if s.grpcConfiguration == nil {
		cfg, err := configuration.NewGrpcConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}
		s.grpcConfiguration = cfg
	}

	return s.grpcConfiguration
}

func (s *serviceProvider) UserRepository(db *sqlx.DB) userService.Repository {
	if s.userRepository == nil {
		s.userRepository = userRepo.NewUserRepository(db)
	}

	return s.userRepository
}

func (s *serviceProvider) UserService() userImpl.Service {
	if s.userService == nil {
		s.userService = userService.NewUserService(
			s.userRepository,
		)
	}

	return s.userService
}

func (s *serviceProvider) UserImpl() *userImpl.Implementation {
	if s.userImpl == nil {
		s.userImpl = userImpl.NewImplementation(s.UserService())
	}

	return s.userImpl
}

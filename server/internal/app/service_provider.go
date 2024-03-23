package app

import (
	"Online-Text-Editor/server/internal/configuration"
	userImpl "Online-Text-Editor/server/internal/handler/user"
	workspaceImpl "Online-Text-Editor/server/internal/handler/workspace"
	userRepo "Online-Text-Editor/server/internal/repository/user"
	workspaceRepo "Online-Text-Editor/server/internal/repository/workspace"
	userService "Online-Text-Editor/server/internal/service/user"
	workspaceService "Online-Text-Editor/server/internal/service/workspace"
	"github.com/jmoiron/sqlx"
	"log"
)

type serviceProvider struct {
	grpcConfiguration   configuration.GrpcConfiguration
	userRepository      userService.Repository
	userService         userImpl.Service
	userImpl            *userImpl.Implementation
	workspaceRepository workspaceService.Repository
	workspaceService    workspaceImpl.Service
	workspaceImpl       *workspaceImpl.Implementation
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

func (s *serviceProvider) UserService(db *sqlx.DB) userImpl.Service {
	if s.userService == nil {
		s.userService = userService.NewUserService(
			s.UserRepository(db),
		)
	}
	return s.userService
}

func (s *serviceProvider) UserImpl(db *sqlx.DB) *userImpl.Implementation {
	if s.userImpl == nil {
		s.userImpl = userImpl.NewImplementation(s.UserService(db))
	}
	return s.userImpl
}

func (s *serviceProvider) WorkspaceRepository(db *sqlx.DB) workspaceService.Repository {
	if s.workspaceRepository == nil {
		s.workspaceRepository = workspaceRepo.NewWorkspaceRepository(db)
	}

	return s.workspaceRepository
}

func (s *serviceProvider) WorkspaceService(db *sqlx.DB) workspaceImpl.Service {
	if s.workspaceService == nil {
		s.workspaceService = workspaceService.NewWorkspaceService(
			s.WorkspaceRepository(db),
		)
	}
	return s.workspaceService
}

func (s *serviceProvider) WorkspaceImpl(db *sqlx.DB) *workspaceImpl.Implementation {
	if s.workspaceImpl == nil {
		s.workspaceImpl = workspaceImpl.NewImplementation(s.WorkspaceService(db))
	}

	return s.workspaceImpl
}

package application

import (
	"fmt"
	"github.com/sergey4qb/mf1-test/delivery/grpc"
	"github.com/sergey4qb/mf1-test/repository"
	"github.com/sergey4qb/mf1-test/services"
)

type Application struct {
	repo     repository.Repository
	services services.Services
	grpc     *grpc.Server
}

func New() (*Application, error) {
	repo, err := repository.New()
	if err != nil {
		return nil, fmt.Errorf("Error initializing repo: %v", err)
	}

	svcs, err := services.New(repo)
	if err != nil {
		return nil, fmt.Errorf("Error initializing services: %v", err)
	}

	grpcSrv, err := grpc.New(svcs)
	if err != nil {
		return nil, fmt.Errorf("Error initializing grpc server: %v", err)
	}

	return &Application{
		repo:     repo,
		services: svcs,
		grpc:     grpcSrv,
	}, nil
}

func (app *Application) Run() error {
	return app.grpc.Start()
}

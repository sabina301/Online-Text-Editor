package app

import (
	desc "Online-Text-Editor/server/pkg/user_v1"
	"context"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}
	err := app.initDependencies(ctx)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (app *App) initDependencies(ctx context.Context) error {
	err := app.initConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load("server/.env")
	if err != nil {
		log.Fatal(err)
	}

	app.initServiceProvider()

	err = app.initGrpcServer()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (app *App) initConfiguration() error {
	viper.AddConfigPath("server/internal/configuration")
	viper.SetConfigName("configuration")
	return viper.ReadInConfig()
}

func (app *App) initServiceProvider() {
	app.serviceProvider = newServiceProvider()
}

func (app *App) initGrpcServer() error {
	app.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(app.grpcServer)
	desc.RegisterUserV1Server(app.grpcServer, app.serviceProvider.UserImpl())
	return nil
}

func (app *App) Run() error {

	listener, err := net.Listen("tcp", app.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}
	err = app.grpcServer.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}

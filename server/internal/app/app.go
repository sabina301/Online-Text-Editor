package app

import (
	desc "Online-Text-Editor/server/pkg/user_v1"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"sync"
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

func (app *App) Run(ctx context.Context) {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := startGrpcServer(app)
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()
		err := startHttpServer(ctx, app)
		if err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}

func startGrpcServer(app *App) error {
	listener, err := net.Listen("tcp", app.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}
	return app.grpcServer.Serve(listener)
}

func startHttpServer(ctx context.Context, app *App) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := desc.RegisterUserV1HandlerFromEndpoint(ctx, mux, app.serviceProvider.GRPCConfig().Address(), opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe("localhost:8080", mux)
}

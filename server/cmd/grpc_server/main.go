package main

import (
	"Online-Text-Editor/server/internal/app"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal("App cant start")
	}
	a.Run(ctx)
}

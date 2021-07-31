package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ianprogrammer/go-api-integration-test/internal/configuration"
	"github.com/ianprogrammer/go-api-integration-test/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type App struct{}

func (app *App) Run() error {

	config, err := configuration.LoadConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	_, err = database.NewDatabase(config.Database)

	if err != nil {
		return err
	}

	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	go func() {

		if err := e.Start(fmt.Sprintf(":%d", config.Server.Port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Server.GracefullShutdownTimeout)*time.Second)

	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	return nil
}

func main() {
	fmt.Println("Starting product API")
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}

package main

import (
	"TicTacToe/internal/di"
	"TicTacToe/internal/web"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		di.Module,
		fx.Invoke(registerRoutes),
	)

	go func() {
		if err := app.Start(context.Background()); err != nil {
			panic(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	if err := app.Stop(context.Background()); err != nil {
		panic(err)
	}
}
func registerRoutes(handler *web.GameHandler) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			handler.ServeUI(w, r)
			return
		}
	})

	http.HandleFunc("/game/", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleGame(w, r)
	})

	go func() {
		_ = http.ListenAndServe(":8080", nil)
	}()
}

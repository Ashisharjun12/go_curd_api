package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ashisharjun12/go_curd_api/internal/config"
	"github.com/Ashisharjun12/go_curd_api/internal/http/handler/student"
)

func main() {

	//config load
	cfg := config.Mustload()

	//db setup
	//router setup

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New() )


	//server setup
	server := &http.Server{
		Addr : cfg.Addr,
		Handler: router,

	}

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT,syscall.SIGALRM)


	go func(){
		err := server.ListenAndServe()
	if err != nil {
        slog.Info("server eror listen failed")
    }

	fmt.Println("server started...")

	}()

	<-done //wait for signal to stop the server

	slog.Info("shuttinng down server")

	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
       slog.Error("failed to gracefully shutdown server", slog.String("error", err.Error()))
    }

	slog.Info("server shutdown completed")
	
}
package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/tobiaskrok/plop/internal/db"
	"github.com/tobiaskrok/plop/internal/server"
)

func main() {

	db, err := db.NewSqlite(":memory:")
	if err != nil {
		//TODO: update connection string
		log.Fatal("failed to connect to SQLite database", err)
	}

	if err := db.Migrate(context.Background()); err != nil {
		log.Fatal("failed to run database migrations", err)
	}

	srv, err := server.New(db)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("Starting SSH server", "host", "localhost", "port", 2222)
		if err = srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("Could not start server", "error", err)
			done <- nil
		}
	}()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()

	log.Info("Stopping SSH server")
	if err := srv.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Could not stop server", "error", err)
	}
}

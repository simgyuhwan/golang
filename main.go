package main

import (
	"api/config"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("need port number\n")
		os.Exit(1)
	}

	port := os.Args[1]
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen port %s: %v", port, err)
	}
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %q", url)

	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to serve: %+v", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}
	return eg.Wait()
}

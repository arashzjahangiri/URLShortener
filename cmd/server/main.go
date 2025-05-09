package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"sync"
	"syscall"

	"go.uber.org/zap"

	"github.com/arashzjahangiri/URLShortener/cmd"
	"github.com/arashzjahangiri/URLShortener/internal/api/http"
	"github.com/arashzjahangiri/URLShortener/internal/config"
	"github.com/arashzjahangiri/URLShortener/internal/entities"
	"github.com/arashzjahangiri/URLShortener/internal/urls"
	"github.com/arashzjahangiri/URLShortener/pkg/observability/logger"
)

func main() {
	monitorPort := flag.Int("monitor-port", 8001, "The server port which handles monitoring endpoints (default: 8001)")
	requestPort := flag.Int("request-port", 8002, "The server port which handles http requests (default: 8002)")
	environmentRaw := flag.String("environment", "", "The environment (default: local)")
	flag.Parse() // Parse the command-line flags

	entities.LoadEnvironment(*environmentRaw)
	var cfg Config
	if err := config.Load(&cfg); err != nil {
		log.Panicf("failed to load config: \n%v", err)
	}

	logger, err := logger.New(cfg.Logger)
	if err != nil {
		log.Fatalf("failed to initialize logger: \n%v", err)
	}

	{ // print build informations
		fields := make([]zap.Field, 0, len(cmd.BuildInfo()))
		for k, v := range cmd.BuildInfo() {
			fields = append(fields, zap.String(k, v))
			logger.Warn("Build Information", fields...)
		}
	}

	urls, err := urls.NewService(cfg.URLs, logger)
	if err != nil {
		log.Fatalf("failed to initialize urls: \n%v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	var wg sync.WaitGroup

	wg.Add(1)
	go http.New(logger, urls).Serve(ctx, &wg, *monitorPort, *requestPort)

	<-ctx.Done()
	wg.Wait()
}

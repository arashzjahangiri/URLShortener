package main

import (
	"github.com/arashzjahangiri/URLShortener/internal/urls"
	"github.com/arashzjahangiri/URLShortener/pkg/observability/logger"
)

type Config struct {
	URLs   *urls.Config   `required:"true"`
	Logger *logger.Config `required:"true"`
}

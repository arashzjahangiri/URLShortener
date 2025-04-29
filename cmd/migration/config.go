package main

import "github.com/arashzjahangiri/URLShortener/pkg/databases/postgres"

type Config struct {
	Postgres *postgres.Config `required:"true"`
}

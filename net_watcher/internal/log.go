package internal

import (
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pushk1nn/netwatch/ent"
)

var (
	Client *ent.Client
	Ctx    context.Context
)

func init() {
	client, err := ent.Open("sqlite3", "file:data.sqlite?_loc=auto&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	Client = client

	ctx := context.Background()
	Ctx = ctx

	if err := Client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

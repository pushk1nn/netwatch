package internal

import (
	"context"
	"log"

	"github.com/pushk1nn/netwatch/ent"
	_ "github.com/mattn/go-sqlite3"
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

	defer Client.Close()
	ctx := context.Background()

	if err := Client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

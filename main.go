package main

import (
	"fmt"
	"log"
	"time"

	"github.com/supernova0730/rickandmortysync/client"
	"github.com/supernova0730/rickandmortysync/pkg/postgres"
	"github.com/supernova0730/rickandmortysync/pkg/rest"
	"github.com/supernova0730/rickandmortysync/processor"
	"github.com/supernova0730/rickandmortysync/repository"
)

func main() {
	db, err := postgres.Connect(postgres.Config{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "supernova",
		Password: "secret",
		Name:     "rickandmorty",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewCharacterRepository(db)
	restClient := rest.New()
	rickAndMortyClient := client.New(restClient)
	prc := processor.New(rickAndMortyClient, repo)

	started := time.Now()
	numOfWorkers := 48
	err = prc.SyncByNumWorkers(numOfWorkers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("prc.SyncByNumWorkers(%d) elapsed: %v\n", numOfWorkers, time.Since(started))

	started = time.Now()
	err = prc.Sync()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("prc.Sync elapsed: %v\n", time.Since(started))
}

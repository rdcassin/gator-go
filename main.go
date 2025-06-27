package main

import (
	"fmt"
	"log"
	"github.com/rdcassin/gator-go/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error starting program: %v", err)
	}

	err = cfg.SetUser("Jackal")
	if err != nil {
		log.Fatalf("Error setting user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal("Error starting program: %v", err)
	}
	fmt.Println(cfg)
}

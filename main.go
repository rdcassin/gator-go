package main

import (
	"log"
	"os"

	"github.com/rdcassin/gator-go/internal/config"
)


type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error starting program: %s", err)
	}
	runState := state{cfg: &cfg}

	cmds := commands{registeredCommands: make(map[string]func(*state, command) error)}
	err = cmds.register("login", handlerLogin)
	if err != nil {
		log.Fatalf("error registering login command: %s", err)
	}

	if len(os.Args) < 2 {
		log.Fatalf("please enter a valid command")
	}

	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmds.run(&runState, cmd)
	if err != nil {
		log.Fatalf("error running command: %s", err)
	}
}

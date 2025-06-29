package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/rdcassin/gator-go/internal/config"
	"github.com/rdcassin/gator-go/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

func main() {
	// Initializing configuration
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error starting program: %s", err)
	}
	runState := state{cfg: &cfg}

	db, err := sql.Open("postgres", runState.cfg.DBURL)
	if err != nil {
		log.Fatal("error connecting to database")
	}

	dbQueries := database.New(db)
	runState.db = dbQueries

	// Adding Commands
	cmds := commands{registeredCommands: make(map[string]func(*state, command) error)}

	// Adding Login Command
	cmdName := "login"
	cmds.register(cmdName, handlerLogin)
	
	// Adding Register Command
	cmdName = "register"
	cmds.register(cmdName, handlerRegister)
	
	// Adding ListUsers Command
	cmdName = "users"
	cmds.register(cmdName, middlewareLoggedIn(handlerListUsers))

	// Adding Aggregate Command
	cmdName = "agg"
	cmds.register(cmdName, handlerAggregate)

	// Adding AddFeed Command
	cmdName = "addfeed"
	cmds.register(cmdName, middlewareLoggedIn(handlerAddFeed))

	// Adding ListFeeds Command
	cmdName = "feeds"
	cmds.register(cmdName, handlerListFeeds)

	// Adding Follow Command
	cmdName = "follow"
	cmds.register(cmdName, middlewareLoggedIn(handlerFollow))

	// Adding ListFeedFollows Command
	cmdName = "following"
	cmds.register(cmdName, middlewareLoggedIn(handlerListFeedFollows))

	// Adding Unfollow Command
	cmdName = "unfollow"
	cmds.register(cmdName, middlewareLoggedIn(handlerUnfollow))

	// Adding Browse Command
	cmdName = "browse"
	cmds.register(cmdName, middlewareLoggedIn(handlerBrowse))

	// Adding Reset Command
	cmdName = "reset"
	cmds.register(cmdName, handlerReset)
	
	// Reading initialization inputs
	if len(os.Args) < 2 {
		log.Fatal("please enter a valid command")
	}
	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	// Running command inputed
	err = cmds.run(&runState, cmd)
	if err != nil {
		log.Fatalf("error running command <%s>: %s", cmd.Name, err)
	}
	
}

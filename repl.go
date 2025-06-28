package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// )

// func repl(s *state) {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for {
// 		fmt.Print("gator > ")
// 		scanner.Scan()
// 		input := scanner.Text()
// 		text := cleanInput(input)

// 		// Reading initialization inputs
// 		if len(os.Args) < 2 {
// 			log.Fatal("please enter a valid command")
// 		}
// 		cmd := command{
// 			Name: os.Args[1],
// 			Args: os.Args[2:],
// 		}

// 		// Running command inputed
// 		err = cmds.run(&runState, cmd)
// 		if err != nil {
// 			log.Fatalf("error running command <%s>: %s", cmd.Name, err)
// 		}
	

// 		command := text[0]
// 		commandParameter := []string{}
// 		if len(text) > 1 {
// 			commandParameter = text[1:]
// 		}

// 		cmd, exists := getCommands()[command]
// 		if !exists {
// 			fmt.Print("Unknown command\n")
// 		} else {
// 			err := cmd.callback(c, commandParameter...)
// 			if err != nil {
// 				fmt.Println(err)				
// 			}
// 		}
// 	}
// }

// func cleanInput(text string) []string {
// 	lowerCaseText:= strings.ToLower(text)
// 	splitText := strings.Fields(lowerCaseText)
// 	return splitText
// }
package main

import (
	"log"
	"os"

	"github.com/Suppur/gator/internal/config"
)

func main() {
	conf, err := config.ReadConf()
	if err != nil {
		log.Fatal(err)
	}

	s := state{
		conf: &conf,
	}

	var cmds commands
	cmds.cmdList = make(map[string]func(*state, command) error)

	cmds.register("login", handlerLogin)

	usrInput := os.Args
	if len(usrInput) < 2 {
		log.Fatal("error, please enter a command")
	}

	usrCmd := command{
		name: usrInput[1],
		args: usrInput[2:],
	}

	if err := cmds.run(&s, usrCmd); err != nil {
		log.Fatal(err)
	}

}

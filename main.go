package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Suppur/gator/internal/config"
	"github.com/Suppur/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {

	conf, err := config.ReadConf()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", conf.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	s := state{
		db:   dbQueries,
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

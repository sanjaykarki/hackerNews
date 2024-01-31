package main

import (
	"log"
	"os"
)

type application struct {
	appName string
	server  server
	debug   bool
	errLog  *log.Logger
	infoLog *log.Logger
}

type server struct {
	host string
	port string
	url  string
}

func main() {

	server := server{
		host: "localhost",
		port: "3000",
		url:  "http://localhost:3000",
	}
	app := &application{
		server:  server,
		appName: "Hacker News",
		debug:   true,
		infoLog: log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate|log.Lshortfile),
		errLog:  log.New(os.Stdout, "ERROR\t", log.Ltime|log.Ldate|log.Lshortfile),
	}

	if err := app.listenAndServer(); err != nil {
		log.Fatal(err)
	}
}

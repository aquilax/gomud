package main

import (
	"os"
	"log"
	"fmt"
	"flag"
)

func getLogger (logfile string, daemonized bool) (logger *log.Logger) {
	logfp, err := os.OpenFile(fmt.Sprintf("%s", logfile), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666);

	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't open logfile `%s': %s\n", logfile, err)
		os.Exit(1)
	}
	if daemonized {
		logger = log.New(logfp, "", log.Ldate|log.Ltime)
	} else {
		logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	}
	return
}

func main(){
	var port *int = flag.Int("port", 8099, "port to run server on");
	var daemon *bool = flag.Bool("daemonize", false, "whether or not to daemonize process");
	var logfile *string = flag.String("logfile", "gomud.log", "filename of the log file to write to");

	server := &Server{*port, getLogger(*logfile, *daemon), 0, 0}
	server.log.Print("test")
	server.Serve()
}

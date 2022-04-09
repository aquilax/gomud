package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func getLogger(logfile string, background bool) (*log.Logger, error) {
	f, err := os.OpenFile(fmt.Sprintf("%s", logfile), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		return nil, fmt.Errorf("couldn't open logfile `%s': %s\n", logfile, err)

	}
	if background {
		return log.New(f, "", log.Ldate|log.Ltime), nil
	}
	return log.New(os.Stdout, "", log.Ldate|log.Ltime), nil
}

func main() {
	var err error
	var l *log.Logger
	var port *int = flag.Int("port", 8099, "port to run server on")
	var background *bool = flag.Bool("background", false, "whether to run the process in the background")
	var logfile *string = flag.String("logfile", "gomud.log", "filename of the log file to write to")

	l, err = getLogger(*logfile, *background)
	server := NewServer(*port, l)

	err = NewGame(server).Start()
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

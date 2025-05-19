package main

import (
	"flag"
	"fmt"
	"marketflow/internal/config"
)

func main() {
	
	port := flag.Int("port", 8080, "port number")
	help := flag.Bool("help", false, "show usage")
	flag.Parse()

	if *help {
		flag.Usage()

		return
	}

	cfg, err := config.Load()
	if err != nil {
		return
	}
	// initial interfaces

	fmt.Printf("Starting on:%d ...\n", *port)
}

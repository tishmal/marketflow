package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.Int("port", 8080, "port number")
	help := flag.Bool("help", false, "show usage")
	flag.Parse()

	if *help {
		flag.Usage()

		return
	}
	// initial interfaces

	fmt.Printf("Starting on:%d ...\n", *port)
}

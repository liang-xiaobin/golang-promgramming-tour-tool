package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "hello", "help")
	flag.StringVar(&name, "n", "world", "info")
	flag.Parse()
	log.Printf("name: %s", name)
}

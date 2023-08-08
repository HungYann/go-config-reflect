package main

import "go-config-reflect/config"

// start the program
func main() {
	config.ReadFile("./file/config.json")
}

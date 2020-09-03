package main

import "github.com/bluesky2106/fun-bet/config"

func main() {
	conf := config.ParseConfig("config.json", "./config")
	conf.Print()
}

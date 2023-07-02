package main

import (
	"flag"
	"xss-css-injection/server"
)

func main() {
	var file_path = flag.String("p", "", "Conf file path")
	flag.Parse()

	server.RunHttpServer(*file_path)
}

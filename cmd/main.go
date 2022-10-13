package main

import "iniudin/belajar-golang-rest/internal/app"

func main() {
	server := app.NewServer()

	if err := server.Start(); err != nil {
		panic(err)
	}
}

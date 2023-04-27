package main

import "cordle/internal/app"

func main() {
	app.Run()
	defer app.Shutdown()
}

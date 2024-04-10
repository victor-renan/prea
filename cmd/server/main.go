package main

import (
	app "prea/internal"
)

func main() {
	app.MainServer{Port: ":3000"}.Run()
}

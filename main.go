package main

import "testgin/routes"

func main() {
	r := routes.SetupRouter()

	r.Run(":8000")
}

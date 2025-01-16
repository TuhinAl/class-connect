package main

import "golang-api/utility/configs"

func main() {
	server := configs.NewAPIServer(":8080")
	server.Run()
}

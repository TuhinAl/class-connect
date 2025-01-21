package main

import "golang-api/internal/utility/configs"

func main() {
	configs.NewAPIServer(":8080").Run()
}

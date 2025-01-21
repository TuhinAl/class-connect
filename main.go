package main

import "github.com/TuhinAl/class-connect/internal/utility/configs"

func main() {
	configs.NewAPIServer(":8080").Run()
}

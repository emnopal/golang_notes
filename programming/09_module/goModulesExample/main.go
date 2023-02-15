package main

import (
	"fmt"

	private_go_module "github.com/emnopal/private-go-module/v2"
)

func main() {
	result := private_go_module.TestPrivateGoModule("Hello")
	fmt.Println(result)
	private_go_module.PrintTestGoModule()
}

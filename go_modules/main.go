package main

import (
	"fmt"

	private_go_module "github.com/emnopal/private-go-module"
)

func main() {
	result := private_go_module.TestPrivateGoModule()
	fmt.Println(result)
	private_go_module.PrintTestGoModule()
}

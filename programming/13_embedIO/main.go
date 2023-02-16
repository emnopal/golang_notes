package main

// note: this is just for demostrate how compiled embed file works

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed string.*
//go:embed test.*
//go:embed readme.md
var fileMatchMultiple embed.FS

func main() {
	if dir, err := fileMatchMultiple.ReadDir("."); err == nil {
		for _, entry := range dir {
			if !entry.IsDir() {
				fmt.Println(entry.Name())
				if content, err := fileMatchMultiple.ReadFile(entry.Name()); err == nil {
					if strings.Contains(entry.Name(), "test") {
						fmt.Println("Content: ", content)
					} else {
						fmt.Println("Content: ", string(content))
					}
				} else {
					panic(err)
				}
			}
		}
	} else {
		panic(err)
	}

}

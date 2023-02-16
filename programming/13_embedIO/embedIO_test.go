package main

import (
	"embed" // we have to import this in order to use go:embed, to avoid no usage error, add _
	"fmt"
	"io/fs"
	"os"
	"testing"
)

// embed string/text

// since embed isn't work inside func, so i create it outside func
//
//go:embed string.txt
var exampleStrEmbed string

func TestEmbedString(t *testing.T) {
	fmt.Println(exampleStrEmbed)
}

// how about embed in nested directory? it's work!
//
//go:embed directory/string.txt
var nestedStrEmbed string

func TestEmbedNestedDir(t *testing.T) {
	fmt.Println(nestedStrEmbed)
}

// how about embed json? it's work!
//
//go:embed string.json
var embedJSON string

func TestEmbedJSON(t *testing.T) {
	fmt.Println(embedJSON)
}

// how about embed yaml? it's work!
//
//go:embed string.yaml
var embedYAML string

func TestEmbedYAML(t *testing.T) {
	fmt.Println(embedYAML)
}

// how about embed xml/html? it's work!
//
//go:embed string.xml
var embedXML string

func TestEmbedXML(t *testing.T) {
	fmt.Println(embedXML)
}

// how about embed markdown? it's work!
//
//go:embed readme.md
var embedMD string

func TestEmbedMD(t *testing.T) {
	fmt.Println(embedMD)
}

// how about embed byte/binary
//
//go:embed test.jpeg
var embedPic []byte

// to test if bytefile is work or not
func TestEmbedPic(t *testing.T) {
	fmt.Println(embedPic) // will be return matrix of picture
}

// test for stdout, convert jpeg to png
func TestCopyEmbedPic(t *testing.T) {
	if err := os.WriteFile("test.png", embedPic, fs.ModePerm); err != nil {
		panic(err)
	}
}

// multiple embed file
//
//go:embed string.json
//go:embed string.txt
//go:embed string.xml
//go:embed string.yaml
var embedMultiple embed.FS // we have to use embed.FS; embed.FS is embed FileSystem

func TestMultipleEmbed(t *testing.T) {
	if file, err := embedMultiple.ReadFile("string.json"); err == nil { // file return []byte, so we have to cast to string
		fmt.Println(string(file))
	} else {
		panic(err)
	}
	if file, err := embedMultiple.ReadFile("string.txt"); err == nil { // file return []byte, so we have to cast to string
		fmt.Println(string(file))
	} else {
		panic(err)
	}
	if file, err := embedMultiple.ReadFile("string.xml"); err == nil { // file return []byte, so we have to cast to string
		fmt.Println(string(file))
	} else {
		panic(err)
	}
	if file, err := embedMultiple.ReadFile("string.yaml"); err == nil { // file return []byte, so we have to cast to string
		fmt.Println(string(file))
	} else {
		panic(err)
	}
}

// multiple files. but using path matcher
//
//go:embed string.*
var embedPathMatch embed.FS

func TestPathMatch(t *testing.T) {
	if dir, err := embedPathMatch.ReadDir("."); err == nil {
		for _, entry := range dir {
			if !entry.IsDir() {
				fmt.Println(entry.Name())
				if content, err := embedPathMatch.ReadFile(entry.Name()); err == nil {
					fmt.Println("Content: ", string(content))
				} else {
					panic(err)
				}
			}
		}
	} else {
		panic(err)
	}
}

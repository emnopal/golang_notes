package main

import (
	"flag"
	"fmt"
)

func main() {
	// flag, masih ada kaitannya sama os.Args
	// buat parsing cli argument
	// read more: https://golang.org/pkg/flag/

	// contohnya
	// flag.tipe_data(nama_flag, default_value, keterangan)
	host := flag.String("host", "localhost", "Assign Host Name")
	port := flag.Int("port", 80, "Assign Port")
	username := flag.String("username", "root", "Assign Username")
	password := flag.String("password", "root", "Assign Username")

	flag.Parse() // for parsing argument above

	fmt.Println(*host, *port, *username, *password)

}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Println("need an address arg [ip:port] [directory]")
		return
	}
	addr := os.Args[1]
	path := os.Args[2]
	fs := http.FileServer(http.Dir(path))
	fmt.Println("http://" + addr)
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(addr, nil))
}

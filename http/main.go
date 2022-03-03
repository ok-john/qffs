package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	if len(os.Args) <= 1 {
		fmt.Println("need an address arg [ip:port]")
		return
	}
	addr := os.Args[1]
	fmt.Println("http://" + addr)
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(addr, nil))
}

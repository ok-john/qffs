package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) <= 3 {
		fmt.Println("[USAGE]\nhttps-server [ip:port] [sitename.com] [relative/directory]")
		return
	}

	addr := os.Args[1]
	certPath, keyPath := fmtLetsEncrypt(os.Args[2])
	fmt.Println("https://" + os.Args[2])
	fs := http.FileServer(http.Dir(os.Args[3]))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServeTLS(addr, certPath, keyPath, nil))
}

func fmtLetsEncrypt(sitename string) (string, string) {
	return fmt.Sprintf("/etc/letsencrypt/live/%s/fullchain.pem", sitename),
		fmt.Sprintf("/etc/letsencrypt/live/%s/privkey.pem", sitename)
}

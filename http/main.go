package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := "https://www.google.com"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error : unable to access url", url)
		os.Exit(1)
	}

	/* bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs)) */

	io.Copy(os.Stdout, resp.Body)

}

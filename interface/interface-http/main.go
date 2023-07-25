package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com/")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	} else {
		// fmt.Println(resp)

		// bs := make([]byte, 99999) //intiliaze byte slice with 99999 empty
		// resp.Body.Read(bs)
		// fmt.Println(string(bs)) //prints string of response body

		io.Copy(os.Stdout, resp.Body)
	}
}

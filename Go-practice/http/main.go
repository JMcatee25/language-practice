package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Below are 3 different ways to print response body and the last one uses the Writer interface to make its own Write function
	// ------------------------------

	// ------------------------------
	// bs := make([]byte, 999999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// ------------------------------

	// ------------------------------
	// io.Copy(os.Stdout, resp.Body)
	// ------------------------------

	// ------------------------------
	lw := logWriter{}
	io.Copy(lw, resp.Body)
	// ------------------------------

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

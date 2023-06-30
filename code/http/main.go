package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://dummyjson.com/products/1")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	//||
	//||
	io.Copy(os.Stdout, resp.Body)

	// lw := logWriter{}

	// io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs), "--------------------------------", len(bs))
	return len(bs), nil
}

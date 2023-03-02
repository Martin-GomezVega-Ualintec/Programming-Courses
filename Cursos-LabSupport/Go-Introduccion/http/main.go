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
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//fmt.Println(resp)

	//bs := make([]byte, 99999)
	// valor que implementa la intefaz del lector funcion de lectura:
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Cantidad de bytes escritos: ", len(bs))
	return len(bs), nil
}

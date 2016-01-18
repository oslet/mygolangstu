package main

import (
	"io"
	"net/http"
	"os"
)

var (
	url = "http://pub.oslet.net/php.ini"
)

func main() {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("php.ini")
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
}

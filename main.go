package main

import (
	"downloader/src"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: downloader <url>")
		os.Exit(1)
	}

	downloadLink := args[1]

	resp, err := http.Get(downloadLink)
	if err != nil {
		panic(err)
	}

	var filename string
	_, params, err := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))

	if err != nil {
		linkSplit := strings.Split(downloadLink, "/")

		filename = linkSplit[len(linkSplit)-1]
	} else {
		filename = params["filename"]
	}

	contentLength := resp.Header.Get("Content-Length")
	if contentLength == "" {
		contentLength = "0"
	}

	file := src.CreateFile(filename)

	fmt.Println("Downloading:", filename)
	fmt.Println("Size:", src.FormattedLength(contentLength))

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	go src.Downloaded(contentLength, file)

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		panic(err)
	}
}

package main

import (
	"downloader/src"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: downloader <url>")
		os.Exit(1)
	}

	resp, err := http.Get(args[1])
	if err != nil {
		panic(err)
	}

	mediaType, params, err := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
	if err != nil {
		panic(err)
	}

	contentLength := resp.Header.Get("Content-Length")
	if contentLength == "" {
		contentLength = "0"
	}

	if mediaType == "attachment" {
		file := src.CreateFile(params["filename"])

		fmt.Println("Downloading:", params["filename"])
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
	} else {
		panic("Unknown media type: " + mediaType)
	}
}

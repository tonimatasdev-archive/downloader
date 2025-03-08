package main

import (
	"downloader/src"
	"fmt"
	"io"
	"mime"
	"net/http"
)

func main() {
	downloadLink := "https://cloud.tonimatas.dev/s/rLBBNTFB9iGitKs/download/VanillaDona.zip"

	resp, err := http.Get(downloadLink)
	if err != nil {
		panic(err)
	}

	mediaType, params, err := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
	contentLength := resp.Header.Get("Content-Length")
	if err != nil {
		panic(err)
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

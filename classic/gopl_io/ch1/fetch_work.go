package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func FetchWork() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "https://") == false {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stdout, "fetch %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stdout, "fetch: reading%s:%v\n", url, err)
			os.Exit(1)
		}
		fmt.Println(resp.Status)
	}
}

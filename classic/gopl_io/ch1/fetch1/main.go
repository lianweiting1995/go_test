package fetch1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stdout, "fetch %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stdout, "fetch: reading%s:%v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("%s", string(b))
	}
}

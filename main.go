package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/snwfdhmp/errlog"
)

var (
	serverPort = 9900
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if errlog.Debug(err) {
			io.WriteString(w, "\n")
			return
		}

		fmt.Printf("--- At: %s\nURL: %s\nMethod: %v\nHeaders: %s\nBody:\n---START---\n%s\n--- END ---\n", time.Now(), r.URL.String(), r.Method, r.Header, string(body))
	})

	fmt.Printf("HTTPMirror server listening on port %d...\n", serverPort)
	panic(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", serverPort), nil))
}

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/tagirhamitov/services_practice_1/types"
)

const port = 2000

func main() {
	http.HandleFunc("/get_result", func(w http.ResponseWriter, r *http.Request) {
		formatStr := r.URL.Query().Get("format")
		format, err := types.ParseFormat(formatStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, fmt.Errorf("failed to get result: %w", err))
			return
		}
		hostname := string(format)
		getResult(w, hostname, port)
	})

	err := http.ListenAndServe("0.0.0.0:2000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getResult(w http.ResponseWriter, hostname string, port uint16) {
	resp, err := http.Get(fmt.Sprintf("http://%v:%v/get_result", hostname, port))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, fmt.Errorf("failed to connect to testing server: %w", err))
		return
	}

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
	io.Copy(w, resp.Body)
	resp.Body.Close()
}

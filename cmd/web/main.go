package main

import (
	"fmt"
	"mail-gif/cmd/web/html"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html.Home().Render(w)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

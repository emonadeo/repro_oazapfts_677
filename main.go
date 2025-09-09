package main

import "net/http"

func main() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		// Write binary data that isn't UTF-8
		w.Write([]byte{0xC3, 0x28})
		w.Header().Add("Content-Type", "application/octet-stream")
	})
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"
	"os"
)

func handleWatch(rw http.ResponseWriter, req *http.Request) {
	v := req.URL.Query().Get("v")
	if v == "" {
		rw.WriteHeader(404)
		return
	}

	rw.Write([]byte(fmt.Sprintf(`<html><head><title>%s</title></head><body>
	<iframe width="560" height="315" src="https://www.youtube.com/embed/%s" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
	</body></html>`, v, v)))
}

func main() {
	var addr = "127.0.0.1:8091"
	if len(os.Args) == 2 {
		addr = os.Args[1]
	}
	http.HandleFunc("/watch", handleWatch)
	http.ListenAndServe(addr, nil)
}

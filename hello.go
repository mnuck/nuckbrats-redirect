package hello

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    path := fmt.Sprintf("http://www.nuckbrats.org%v", r.URL.Path)
    http.Redirect(w, r, path, http.StatusMovedPermanently)
}

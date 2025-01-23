package main

import (
	"log"
	"mjpclab.dev/ghfs/src/param"
	"mjpclab.dev/ghfs/src/serverHandler"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world\n"))
	})

	p := &param.Param{
		Root:      "/tmp",
		IndexUrls: []string{"/"},
	}
	p.Normalize()
	fsHandler, _ := serverHandler.NewVhostHandler(p, nil, nil)
	http.Handle("/files/", http.StripPrefix("/files/", fsHandler))

	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

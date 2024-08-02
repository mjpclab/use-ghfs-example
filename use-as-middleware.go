package main

import (
	"log"
	"mjpclab.dev/ghfs/src/param"
	"mjpclab.dev/ghfs/src/serverHandler"
	"mjpclab.dev/ghfs/src/serverLog"
	"mjpclab.dev/ghfs/src/tpl/defaultTheme"
	"net/http"
	"os"
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
	logger, _ := serverLog.NewWriterMan().NewLogger(os.Stdout, os.Stderr)
	fsHandler, _ := serverHandler.NewVhostHandler(p, logger, defaultTheme.DefaultTheme)
	http.Handle("/files/", http.StripPrefix("/files/", fsHandler))

	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

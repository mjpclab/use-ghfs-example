package main

import (
	"fmt"
	"mjpclab.dev/ghfs/src/app"
	"mjpclab.dev/ghfs/src/param"
)

func main() {
	params, errs := param.NewParams([]param.Param{
		// virtual host 0
		{
			Root:      "/tmp",
			Listens:   []string{"8080"},
			IndexUrls: []string{"/"},
		},
		// virtual host 1
		{
			Root:         "/usr",
			Listens:      []string{"8443"},
			CertKeyPaths: [][2]string{{"cert/example.crt", "cert/example.key"}},
			IndexUrls:    []string{"/"},
		},
	})
	if len(errs) > 0 {
		fmt.Println(errs)
		return
	}

	// app
	appInst, errs := app.NewApp(params)
	if len(errs) > 0 {
		fmt.Println(errs)
		return
	}

	errs = appInst.Open()
	if len(errs) > 0 {
		fmt.Println(errs)
	}
}

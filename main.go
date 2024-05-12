package main

import (
	"crypto/tls"
	"fmt"
	"mjpclab.dev/ghfs/src/app"
	"mjpclab.dev/ghfs/src/param"
	"mjpclab.dev/ghfs/src/setting"
)

func main() {
	cert, err := tls.LoadX509KeyPair("cert/example.crt", "cert/example.key")
	if err != nil {
		fmt.Println(err)
		return
	}

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
			Certificates: []tls.Certificate{cert},
			IndexUrls:    []string{"/"},
		},
	})
	if len(errs) > 0 {
		fmt.Println(errs)
		return
	}

	// setting
	setting := &setting.Setting{
		Quiet:   false,
		PidFile: "",
	}

	// app
	appInst, errs := app.NewApp(params, setting)
	if len(errs) > 0 {
		fmt.Println(errs)
		return
	}

	errs = appInst.Open()
	if len(errs) > 0 {
		fmt.Println(errs)
	}
}

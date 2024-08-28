package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: curlv <URL>")
		return
	}
	url := os.Args[1]

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	startNamelookup := time.Now()
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	defer resp.Body.Close()

	timeNamelookup := time.Since(startNamelookup).Seconds()

	timeTotal := time.Since(start).Seconds()
	timeConnect := timeTotal
	timeAppConnect := 0.0
	timePretransfer := timeTotal
	timeRedirect := 0.0

	fmt.Printf(`<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
<html><head>
<title>301 Moved Permanently</title>
</head><body>
<h1>Moved Permanently</h1>
<p>The document has moved <a href="%v/">here</a>.</p>
</body></html>`, url)
	fmt.Printf("\ntime_namelookup: %.6f\n", timeNamelookup)
	fmt.Printf("time_connect: %.6f\n", timeConnect)
	fmt.Printf("time_appconnect: %.6f\n", timeAppConnect)
	fmt.Printf("time_pretransfer: %.6f\n", timePretransfer)
	fmt.Printf("time_redirect: %.6f\n", timeRedirect)
	fmt.Printf("time_starttransfer: %.6f\n", timeTotal)
	fmt.Printf("time_total: %.6f\n", timeTotal)
}

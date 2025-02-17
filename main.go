package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var serverID string
var serverInstanceType string
var blakListIPS []string
var permittedOrigins = map[string]string{
	"http://localhost:1234":  "1",
	"https://127.0.0.1:1234": "1",
	"http://localhost:3000":  "1",
}

var GOOGLE_ANALYTICS_CODE_SNIPPET = `
<!-- Global site tag (gtag.js) - Google Analytics -->
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-154567842-1"></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag() {
    dataLayer.push(arguments);
  }
  gtag("js", new Date());

  gtag("config", "UA-154567842-1");
</script>`

var STATIC_DIR = "dist" // this is to evade .gitignore tha is ignoring /build and docker follows it even if .dockeringore is not

func main() {
	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, os.Interrupt)

	go func() {
		for sig := range osSignalChan {
			// sig is a ^c, handle it
			if (sig == os.Interrupt) || (sig == os.Kill) {
				//msg := fmt.Sprintf("Goodbye Cruel world! signal: %#v \n I've just received this signal: ", sig)
				os.Exit(0)
			}
		}
	}()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("starting up, server: %s\n", serverID)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        headerSetter(vaultHandler{}),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	log.Printf("Start serving on port %s\n", port)
	log.Fatal(server.ListenAndServe())
	log.Print("Exiting dox server")
}

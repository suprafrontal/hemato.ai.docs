package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"strings"
	"time"
)

type vaultHandler struct{}

func (vh vaultHandler) ServeHTTP(w http.ResponseWriter, q *http.Request) {
	t1 := time.Now()
	log.Printf("[%s]-> %s\n", strings.ToUpper(q.Method), q.URL)
	switch q.URL.Path {
	case "/", "/index.html":
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0")
		http.ServeFile(w, q, fmt.Sprintf("%s/index.html", STATIC_DIR))
	case "/heartbeat":
		fmt.Fprint(w, RandStr(20))
	case "/version":
		w.Header().Add("Content-Type", string("application/json"))
		r := make(map[string]interface{})
		r["request_timestamp"] = time.Now().UnixNano()
		r["version"] = hematoDocsVersion()
		r["delta"] = fmt.Sprint(time.Since(t1))
		rjson, err := json.Marshal(r)
		if err != nil {
			// if we cannot do this we should not be considered up and running
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			fmt.Fprint(w, string(rjson))
		}
	default:
		serveStatic(w, q)
	}
}

// just for fun, serve beautiful error pages
func errorHandler(w http.ResponseWriter, q *http.Request, status int) {
	prettyRequest := "UnAvailable"
	b, err := httputil.DumpRequest(q, true)
	if err == nil && len(b) < 5000 {
		prettyRequest = string(b)
	}
	log.Printf("Serving error page %d for path %s \n %s", status, q.URL, prettyRequest)

	msg := ""
	if status == 404 {
		msg = fmt.Sprintf("Serving error page %d for path %s", status, q.URL)
	} else {
		msg = fmt.Sprintf("Serving error page %d for path %s \n %s", status, q.URL, prettyRequest)
	}
	if status != 404 {
		//slackThis(msg)
		log.Println(msg)
	}

	w.WriteHeader(status)
	switch status {
	case http.StatusNotFound:
		fmt.Fprint(w, `<body style='color:fcfcfc; background-color: 0379f9; '><div style='font-size: 350; width: 100%; height: 100%; text-align: center; >404</div>
            <div style="text-align:center; font-size=1em">`+hematoDocsVersion()+`</div>`+GOOGLE_ANALYTICS_CODE_SNIPPET+`</body>`)
	case http.StatusForbidden:
		fmt.Fprint(w, `<body style='color:fcfcfc; background-color: f9031c; '><div style='font-size: 350; width: 100%; height: 100%; text-align: center; >403</div>
            <div style="text-align:center; font-size=1em">`+hematoDocsVersion()+`</div>`+GOOGLE_ANALYTICS_CODE_SNIPPET+`</body>`)
	}
}

// the assumption here is all the static content are under a directory like ./static/
// and they can be addressed both like /index.html or /static/index.html
func serveStatic(w http.ResponseWriter, q *http.Request) {
	filename := ""
	if strings.Index(q.URL.Path, STATIC_DIR) == -1 {
		filename = STATIC_DIR + q.URL.Path
	} else {
		filename = q.URL.Path[1:]
	}

	// the default document is index.html
	// TODO: do we need to support default.html or index.htm and such? NOOO
	if strings.HasSuffix(filename, "/") {
		filename = filename + "index.html"
	}

	// we SHOULD NOT serve any arbitrary file extension
	extension := path.Ext(filename)
	log.Print(filename)

	// Added "" as valid extension
	// if extension == "" && q.URL.Path != "/" {
	// 	http.Redirect(w, q, "/", http.StatusMovedPermanently)
	// 	return
	// }
	if !isPermittedExtension(extension) {
		if strings.Index(filename, "apple-app-site-association") == -1 {
			errorHandler(w, q, http.StatusForbidden)
			return
		}
	}

	// if the file is there, serve it
	if _, err := os.Stat(filename); err == nil {
		http.ServeFile(w, q, filename)
		return
	}

	// if it is a path ending in / then serve /index.html
	if strings.HasSuffix(filename, "/") {
		if _, err := os.Stat(filename + "index.html"); err == nil {
			http.ServeFile(w, q, filename+"index.html")
			return
		}
	}

	// if there is an html file with that name serve that
	if _, err := os.Stat(filename + ".html"); err == nil {
		http.ServeFile(w, q, filename+".html")
		return
	}

	errorHandler(w, q, http.StatusNotFound)
	return
}

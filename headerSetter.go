package main

import (
	"net/http"
	"strings"
	"time"
)

func headerSetter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, q *http.Request) {
		w.Header().Set("X-Frame-Options", "SAMEORIGIN")
		// TODO: content security policy is complicated and seems to eb interfering with manual bundles
		//w.Header().Set("Content-Security-Policy", "default-src 'self' https://api.hemato.ai https://stimulator-www-oz7stxml5a-ue.a.run.app https://stimulator-api-oz7stxml5a-ue.a.run.app https://*.cloudflareinsights.com https://*.googletagmanager.com; object-src 'none'; require-trusted-types-for 'script'; img-src https://*; child-src 'none';")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")

		if w.Header().Get("Cache-Control") != "" {
			next.ServeHTTP(w, q)
			return
		}
		if !strings.Contains(q.URL.RawPath, "index.html") && !strings.HasSuffix(q.URL.RawPath, "/") {
			cacheUntil := time.Now().Add(time.Hour).Format(http.TimeFormat)
			w.Header().Set("Cache-Control", "max-age:360, public")
			w.Header().Set("Expires", cacheUntil)
			next.ServeHTTP(w, q)
		} else {
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			next.ServeHTTP(w, q)
		}
	})
}

func headerSetterFunc(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, q *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")

		if w.Header().Get("Cache-Control") != "" {
			next.ServeHTTP(w, q)
			return
		}
		if !strings.Contains(q.URL.RawPath, "index.html") && !strings.HasSuffix(q.URL.RawPath, "/") {
			cacheUntil := time.Now().Add(time.Hour).Format(http.TimeFormat)
			w.Header().Set("Cache-Control", "max-age:360, public")
			w.Header().Set("Expires", cacheUntil)
			next.ServeHTTP(w, q)
		} else {
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			next.ServeHTTP(w, q)
		}
	})
}

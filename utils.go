package main

import (
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func RandStr(length int) string {
	length = min(length, 100)
	length = max(length, 1)
	hash := sha512.New()
	rand.Seed(time.Now().UTC().UnixNano())
	i := rand.Int63()
	s := strconv.FormatInt(i, 10)
	b := hash.Sum([]byte(s))
	he := hex.EncodeToString(b)
	permu := rand.Perm(len(he))
	var shuf string
	for i := 0; i < len(he); i++ {
		shuf = shuf + string(he[permu[i]])
	}
	return shuf[0:length]
}

// from https://github.com/golang/example/blob/master/stringutil/reverse.go#L21
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// used to generate sessionid s
// uses the time
// reverses it, to avoid problems like "hot hashes" and busy shards
// and over kill but a good practice
func nextId() string {
	our_zero := time.Date(1978, 7, 20, 0, 0, 0, 0, time.UTC)
	delta := time.Now().Sub(our_zero)
	newid := strconv.FormatInt(int64(delta), 10)
	// revese it so that the haskeys on dynamo (or similarly shares on a db cluster and such) get evenly distributed over shards
	// or primary keys get evenly distributed over any datastore
	return reverse(newid)
}

// for security reasons we only support serving files with followin extensions
func isPermittedExtension(s string) bool {
	supportedExtensions := []string{
		".html",
		".css",
		".js",
		".json",
		".jpg",
		".jpeg",
		".png",
		".gif",
		".ico",
		".ttf",
		".woff",
		".woff2",
		".svg",
		".pdf",
		"apple-app-site-association",
		".txt", // robots.txt
		"",     // these are basically ids like podcasts/someId
		".php", // just to reduce the noise of bots probing
		".webmanifest",
	}
	s = strings.ToLower(s)
	for _, e := range supportedExtensions {
		if s == e {
			return true
		}
	}
	return false
}

var whitelist []func(*url.URL) bool = []func(*url.URL) bool{
	// robots
	func(u *url.URL) bool { return u.Path == "/robots.txt" },
	// heartbeat
	func(u *url.URL) bool { return u.Path == "/heartbeat" },
	func(u *url.URL) bool { return u.Path == "/version" },
	// google auth
	func(u *url.URL) bool { return strings.Index(u.Path, "/authcallback") == 0 },
	//favicon
	func(u *url.URL) bool {
		return strings.Index(u.Path, "/favicon") == 0 || u.Path == "/apple-touch-icon.png"
	},
	// password reset
	func(u *url.URL) bool {
		return strings.Index(u.Path, "/resetpassword.html") == 0 || strings.Index(u.Path, "/rpstyle.css") == 0 || strings.Index(u.Path, "/resetpasswordsuccess.html") == 0
	},
	// Apple
	func(u *url.URL) bool { return u.Path == "/apple-app-site-association" },
	// images
	func(u *url.URL) bool { return strings.Index(u.Path, "/images/") == 0 },
	// reports
	func(u *url.URL) bool { return strings.Index(u.Path, "/publicreports/") == 0 },
	// graphDeliverer until we move it to app
	func(u *url.URL) bool { return strings.Index(u.Path, "/graphdeliverer/") == 0 },
	func(u *url.URL) bool { return strings.Index(u.Path, "/graphdelivery/") == 0 },
	// public_utils
	func(u *url.URL) bool { return strings.Index(u.Path, "/public_utils/") == 0 },
	// subscriptions
	func(u *url.URL) bool { return strings.Index(u.Path, "/subscribe/") == 0 },
}

func isProtectedPath(path *url.URL) bool {
	for _, t := range whitelist {
		if t(path) {
			return false
		}
	}
	return true
}

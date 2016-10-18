package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func initStaticFiles(prefix string, urlPrefix string) {
	wf := func(path string, info os.FileInfo, err error) error {
		// log.Println(path, info, err)
		if path == prefix {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		urlpath := path[len(prefix):]
		if urlpath[0] != '/' {
			urlpath = "/" + urlpath
		}
		if urlPrefix != "" {
			if urlPrefix[len(urlPrefix)-1] == '/' {
				urlPrefix = strings.TrimSuffix(urlPrefix, "/")
			}
			urlpath = urlPrefix + urlpath
			if urlpath[0] != '/' {
				urlpath = "/" + urlpath
			}
		}
		log.Println("Registering", urlpath, path)
		f, err := os.Open(path)
		if err != nil {
			log.Println(err)
			return nil
		}
		content := make([]byte, info.Size())
		f.Read(content)
		f.Close()
		contentLength := strconv.Itoa(len(content))
		contentType := ""
		switch {
		case strings.HasSuffix(path, ".css"):
			contentType = "text/css"
		case strings.HasSuffix(path, ".js"):
			contentType = "application/javascript"
		case strings.HasSuffix(path, ".png"):
			contentType = "image/png"
		case strings.HasSuffix(path, ".gif"):
			contentType = "image/gif"
		case strings.HasSuffix(path, ".jpg"), strings.HasSuffix(path, ".jpeg"):
			contentType = "image/jpeg"
		case strings.HasSuffix(path, ".html"):
			contentType = "text/html"
		case strings.HasSuffix(path, ".json"):
			contentType = "application/json"
		}

		// TDOO: Show routing
		handler := func(w http.ResponseWriter, r *http.Request) {
			if contentType != "" {
				w.Header().Set("Content-Type", contentType)
			}
			w.Header().Set("Content-Length", contentLength)
			w.Write(content)
		}
		http.HandleFunc(urlpath, handler)
		return nil
	}
	filepath.Walk(prefix, wf)
}

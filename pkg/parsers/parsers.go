package parsers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type dict map[string]string

func ParseHeaders(req *http.Request) dict {
	var newHeaders = make(dict)

	for key, value := range req.Header {
		if strings.HasPrefix(strings.ToLower(key), "x-") {
			continue
		}
		newHeaders[key] = strings.Join(value, ";")
	}

	return newHeaders
}

func ParseParams(req *http.Request) dict {
	var newParam = make(dict)

	for key, value := range req.URL.Query() {
		newParam[key] = strings.Join(value, ";")
	}

	return newParam
}

func ParseURL(req *http.Request) string {
	var scheme string = "http"
	if req.TLS != nil {
		scheme = "https"
	}

	return scheme + "://" + req.Host + req.URL.Path
}

func ParseBodyData(req *http.Request, ct string) (dict, string) {
	var formData = make(dict)
	var textDate string

	if ct == "application/x-www-form-urlencoded" {
		req.ParseForm()
		for key, value := range req.PostForm {
			formData[key] = strings.Join(value, ";")
		}
	} else if ct == "multipart/form-data" {
		req.ParseMultipartForm(1000)
		for key, value := range req.PostForm {
			formData[key] = strings.Join(value, ";")
		}

	} else if ct == "text/plain" || ct == "text/html" {
		defer req.Body.Close()
		bdy, _ := io.ReadAll(req.Body)
		textDate = string(bdy)

	} else if ct == "application/json" {
		defer req.Body.Close()
		bdy, _ := io.ReadAll(req.Body)
		json.Unmarshal(bdy, &formData)
	}

	return formData, textDate
}

func ParseSupportedImage(req *http.Request, supportedTypes []string) string {
	headerAccept := req.Header["Accept"]

	if len(headerAccept) == 0 {
		return ""
	}

	for _, sup := range supportedTypes {
		if strings.Contains(headerAccept[0], sup) {
			return sup
		}
	}

	return ""

}

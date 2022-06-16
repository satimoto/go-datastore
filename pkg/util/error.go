package util

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func LogOnError(code string, message string, err error) {
	if err != nil {
		log.Printf("%s: %s", code, message)
		log.Printf("%s: %v", code, err)
	}
}

func PanicOnError(code string, message string, err error) {
	if err != nil {
		log.Printf("%s: %s", code, message)
		log.Fatalf("%s: %v", code, err)
	}
}

func LogHttpRequest(code, url string, request *http.Request, body bool) {
	if request != nil {
		bytes, err := httputil.DumpRequest(request, body)

		if err != nil {
			log.Printf("%s: %s", code, err)
		} else {
			log.Printf("%s: %s", code, url)
			log.Println(string(bytes))
		}
	}
}

func LogHttpResponse(code, url string, response *http.Response, body bool) {
	if response != nil {
		bytes, err := httputil.DumpResponse(response, body)

		if err != nil {
			log.Printf("%s: %s", code, err)
		} else {
			log.Printf("%s: %s", code, url)
			log.Println(string(bytes))
		}
	}
}

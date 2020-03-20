package main

import "net/http"

// the go crazy version
// replace the normal handler function by something that returns a response and an error
// this will be evaluated later and then the appropriate method will be called.
// this provides even more standardization but needs more work in general to work for all use cases.

func myWebHandler3(r *http.Request) (interface{}, error) {
	return getSomethingFromDB()
}

func myWebHandler3_usageExample() {
	http.Handle("/", JSONHandler(myWebHandler3))
}

type JSONHandler func(r *http.Request) (interface{}, error)

func (h JSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	response, err := h(r)
	respondWithJSON(w, response, err)
}
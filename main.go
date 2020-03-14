package main

import (
	"log"
	"net/http"
)

func MultiplePoliciesHandler1(w http.ResponseWriter, r *http.Request) {
	// Does the policy report to both report-uris?
	w.Header().Add("Content-Security-Policy", "script-src 'self'; report-uri https://5e6c1d6f99fa309cec74bd79.endpoint.csper.io;")
	w.Header().Add("Content-Security-Policy", "script-src https:; report-uri https://5e6c1d6f99fa309cec74bd79.endpoint.csper.io;")

	http.ServeFile(w, r, "static/multiple.html")
}

func MultiplePoliciesHandler2(w http.ResponseWriter, r *http.Request) {
	// What gets reported on one failure
	w.Header().Add("Content-Security-Policy", "script-src 'self'; report-uri https://5e6c1d6f99fa309cec74bd79.endpoint.csper.io;")
	w.Header().Add("Content-Security-Policy", "script-src https:; report-uri https://5e6c1d6f99fa309cec74bd79.endpoint.csper.io;")

	http.ServeFile(w, r, "static/multiple2.html")
}

func MultiplePoliciesHandler3(w http.ResponseWriter, r *http.Request) {
	// When using comma seperated policies, individual policies are reported
	w.Header().Add("Content-Security-Policy", "script-src https:; report-uri https://5e6c1d6f99fa309cec74bd79.endpoint.csper.io, script-src 'self'; report-uri https://5e6c1d6f99fa309cec74bd79.endpoint.csper.io")

	http.ServeFile(w, r, "static/multiple1.html")
}

func main() {
	http.HandleFunc("/multiple1", MultiplePoliciesHandler1)
	http.HandleFunc("/multiple2", MultiplePoliciesHandler2)
	http.HandleFunc("/multiple3", MultiplePoliciesHandler3)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

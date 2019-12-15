package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/SendSMS", sendSMS)
	http.ListenAndServe(":8081", nil)
}

func sendSMS(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		mobileNumber := r.FormValue("mobile")
		messageToSend := r.FormValue("message")
		fmt.Printf("Mobile number is: %s\nMessage to send is: %s", mobileNumber, messageToSend)
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	}
}

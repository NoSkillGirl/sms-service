package main

import (
	"net/http"

	"bitbucket.org/lightbulbwithswitch/sms-service.git/api"
)

func main() {
	http.HandleFunc("/SendSMS", api.SendSMS)
	http.ListenAndServe(":8081", nil)
}

package main

import (
	"net/http"
	
	"github.com/NoSkillGirl/sms-service/api"
)

func main() {
	http.HandleFunc("/SendSMS", api.SendSMS)
	http.ListenAndServe(":8081", nil)
}

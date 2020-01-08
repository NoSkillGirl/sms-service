package twillio

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

//SendSMS - public function
func SendSMS(mob string, msg string) int {
	// Set initial variables
	accountSid := "AC07a08a17e4573ebf6096882a6b1cb61d"
	authToken := "83b3cb55e000c775a70cfbcc74bbe8b8"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Build out the data for our message
	v := url.Values{}
	fmt.Println(mob, msg)
	v.Set("To", mob)
	v.Set("From", "+13374193975")
	v.Set("Body", msg)
	rb := *strings.NewReader(v.Encode())

	// Create client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	fmt.Println(req)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)
	// fmt.Println(resp.Status)
	return resp.StatusCode
}

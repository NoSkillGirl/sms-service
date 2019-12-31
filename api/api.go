package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"bitbucket.org/lightbulbwithswitch/sms-service.git/twillio"
	_ "github.com/go-sql-driver/mysql"
)

// SendSMSRequest - stuct
type SendSMSRequest struct {
	Message string
	Mobile  string
}

// SendSMSResponse - struct
type SendSMSResponse struct {
	Message string
}

// SendSMS - api send sms
func SendSMS(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sendSMSRequest := SendSMSRequest{}

		jsn, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error in reading request body", err)
		}

		err = json.Unmarshal(jsn, &sendSMSRequest)
		if err != nil {
			fmt.Println("Decoding Error", err)
		}

		fmt.Println("sendSMSRequest ....")
		fmt.Println(sendSMSRequest)

		w.Header().Set("Content-Type", "application/json")
		fmt.Printf("Mobile number is: %s\nMessage to send is: %s\n", sendSMSRequest.Mobile, sendSMSRequest.Message)

		//sending message through twillio
		respStatus := twillio.SendSMS(sendSMSRequest.Mobile, sendSMSRequest.Message)
		fmt.Println(respStatus)

		if respStatus != 201 {
			fmt.Fprintf(w, "SMS FAILED, %s!", r.URL.Path[1:])
			return
		}

		// initialize db connection
		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/smsDetails")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		smsInsertQuery := `insert into messages (mobile, message, delivery_status, provider) values ('%s', '%s', 'SENT', 'TWILLIO')`
		smsInsertQueryString := fmt.Sprintf(smsInsertQuery, sendSMSRequest.Mobile, sendSMSRequest.Message)
		fmt.Println(smsInsertQueryString)

		// perform a db.Query insert
		insert, err := db.Query(smsInsertQueryString)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()

		sendSMSResponse := SendSMSResponse{
			Message: "SMS Sent Successfull",
		}

		json.NewEncoder(w).Encode(sendSMSResponse)
	}
}

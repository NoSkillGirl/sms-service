package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"bitbucket.org/lightbulbwithswitch/sms-service.git/twillio"
	_ "github.com/go-sql-driver/mysql"
)

// SendSMS - api send sms
func SendSMS(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		mobileNumber := "+91" + r.FormValue("mobile")
		messageToSend := r.FormValue("message")
		fmt.Printf("Mobile number is: %s\nMessage to send is: %s\n", mobileNumber, messageToSend)

		//sending message through twillio
		respStatus := twillio.SendSMS(mobileNumber, messageToSend)
		fmt.Println(respStatus)

		if respStatus != 201 {
			fmt.Fprintf(w, "SMS FAILED, %s!", r.URL.Path[1:])
			return
		}

		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/smsDetails")

		// if there is an error opening the connection, handle it
		if err != nil {
			panic(err.Error())
		}

		// defer the close till after the main function has finished
		// executing
		defer db.Close()

		smsInsertQuery := `insert into messages (mobile, message, delivery_status, provider) values ('%s', '%s', 'SENT', 'TWILLIO')`

		smsInsertQueryString := fmt.Sprintf(smsInsertQuery, mobileNumber, messageToSend)
		fmt.Println(smsInsertQueryString)

		// perform a db.Query insert
		insert, err := db.Query(smsInsertQueryString)

		// if there is an error inserting, handle it
		if err != nil {
			panic(err.Error())
		}
		// be careful deferring Queries if you are using transactions
		defer insert.Close()
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	}
}

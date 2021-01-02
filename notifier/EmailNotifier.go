package notifier

import (
	"larsdebruijn.nl/holla/task"
	"log"
	"net/smtp"
)

type EmailNotifier struct{}

func (n EmailNotifier) Notify(recv <-chan []task.Result) {
	for {
		select {
		case results := <-recv:
			errors := findErrors(results)
			if len(errors) > 0 {
				sendEmail(errors)
			}
		}
	}
}

func sendEmail(errors []task.Result) {
	errorString := ""
	for _, result := range errors {
		errorString = "[" + result.Timestamp.String() + "] " + result.Target.Uri + " failed due to " + result.Error.Error() + " \r\n"
	}

	// Choose auth method and set it up
	auth := smtp.PlainAuth("", "holla@ldebruijn.nl", "extremely_secret_pass", "smtp.mailtrap.io")

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{"larsdebruijn12@gmail.com"}
	msg := []byte("Holla: larsdebruijn12@gmail.com\r\n" +
		"One or more failures reported in liveness checks\r\n" +
		errorString)
	err := smtp.SendMail("smtp.mailtrap.io:25", auth, "piotr@mailtrap.io", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func findErrors(results []task.Result) []task.Result {
	errors := []task.Result{}
	for _, result := range results {
		if result.Error != nil {
			errors = append(errors, result)
		}
	}
	return errors
}

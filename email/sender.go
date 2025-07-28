package email

import (
	"github.com/50-Course/simple-mail-server/model"
	"log"
	"time"
)

// Uses the stdout to simulate sending an email jjust like Django's ConsoleEmailBackend
func sendMail(job model.EmailJob) {
	log.Printf("Sending email to %s | Subject: %s", job.To, job.Subject)
	time.Sleep(1 * time.Second)
	log.Printf("Email sent to %s", job.To)
}

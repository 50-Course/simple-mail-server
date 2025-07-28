// houses the everything related to the job struct

package model

import (
	"errors"
	"regexp"
)

var ErrRequiredFields = errors.New("all fields are required")
var ErrInvalidEmailFormat = errors.New("invalid email format")

// Represents the incoming email request payload
type JobRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// Represents structured API response
type JobResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Represents a single job in the queue
type EmailJob struct {
	JobRequest
	Retries int
}

// Validate checks if the email job request is valid
// this ould help us match emails like: word@word.word
var emailPattern = regexp.MustCompile(`^[\w\.-]+@[\w\.-]+\.\w+$`)

// Ensures required fields are present and email is valid
func ValidateJob(req JobRequest) error {
	if req.To == "" || req.Subject == "" || req.Body == "" {
		return ErrRequiredFields
	}
	if !emailPattern.MatchString(req.To) {
		return ErrInvalidEmailFormat
	}
	return nil
}

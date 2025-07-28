// houses everything related to custom worker pool
package worker

import (
	"github.com/50-Course/simple-mail-server/email"
	"github.com/50-Course/simple-mail-server/model"
	"log"
	"sync"
)

// Spins up a goroutine that listens for jobs and processes them
// using the provided worker ID and job channel.
func StartWorker(id int, jobs <-chan model.EmailJob, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for job := range jobs {
			log.Printf("Worker %d picked up a job", id)
			email.SendMail(job)
		}
		log.Printf("Worker %d shutting down", id)
	}()
}

// Launches `StartWorker` with a default number -3
func StartDefaultWorker(jobs <-chan model.EmailJob, wg *sync.WaitGroup) {
	StartWorker(3, jobs, wg)
}

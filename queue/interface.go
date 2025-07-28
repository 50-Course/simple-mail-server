package queue

import "github.com/50-Course/simple-mail-server/model"

// Top-level interface for the queue system
//
// This interface defines the methods that allow us to swap
// out different queue implementations.
type Queue interface {
	// AddJob adds a new job to the queue
	Enqueue(job model.EmailJob) error

	// Gracefully shutdown the queue
	Close() error
}

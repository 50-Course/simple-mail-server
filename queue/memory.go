package queue

import (
	"errors"
	"github.com/50-Course/simple-mail-server/model"
	"sync"
)

var ErrQueueClosed = errors.New("queue is closed")
var ErrQueueFull = errors.New("queue is full")

// contians the in-memory queue implementation
type InMemoryQueue struct {

	// A mutex to ensure thread-safe operations
	mu sync.Mutex

	// A flag to indicate if the queue is open for new jobs
	open bool

	// A channel to hold the email jobs
	jobs chan model.EmailJob
}

// adds a job to the in-memory queue
func (q *InMemoryQueue) Enqueue(job model.EmailJob) error {

	// approach:
	// first we lock the queue to ensure thread safety
	// then we check if the queue is open
	// if it is not open, we return an error
	// if it is open, we try to send the job to the channel

	q.mu.Lock()
	defer q.mu.Unlock()

	if !q.open {
		return ErrQueueClosed
	}

	select {
	case q.jobs <- job:
		return nil
	default:
		return ErrQueueFull
	}
}

// Marks the queue as closed and prevents further enqueueing
// of jobs for processing. It also closes the jobs channel.
func (q *InMemoryQueue) Close() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.open = false
	close(q.jobs)
}

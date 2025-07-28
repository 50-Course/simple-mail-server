package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/50-Course/simple-mail-server/model"
	"github.com/50-Course/simple-mail-server/queue"
	"github.com/50-Course/simple-mail-server/worker"
)

// Handles incoming HTTP requests to send emails.
func SendEmailHandler(q *queue.InMemoryQueue) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.JobRequest

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if err := model.ValidateJob(req); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		job := model.EmailJob{JobRequest: req}

		if err := q.Enqueue(job); err != nil {
			http.Error(w, "Queue full", http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(model.JobResponse{
			Status:  "accepted",
			Message: "Job enqueued successfully",
		})
	}
}

// Gracefully shuts down the server and stop accepting new jobs.
func Shutdown(server *http.Server, q *queue.InMemoryQueue, workerWg *sync.WaitGroup) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down gracefully...")

	q.Close() // stop accepting jobs

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}

	workerWg.Wait()
	log.Println("All workers finished. Shutdown complete.")
}

func main() {
	var (
		port      = flag.String("port", "8080", "HTTP port to listen on")
		queueSize = flag.Int("queue-size", 10, "Size of the job queue")
		workerNum = flag.Int("workers", 3, "Number of concurrent workers to use")
		queueType = flag.String("queue-type", "memory", "Queue type: memory or rabbitmq")
	)
	flag.Parse()

	log.Printf("Starting Email Service on port %s with %d workers and queue size %d", *port, *workerNum, *queueSize)

	if *queueType != "memory" && *queueType != "rabbitmq" {
		//TODO: Implement RabbitMQ connector later
		// this would connect to RabbitMQ and use it as a queue through the queue.rabbitmq package
		log.Fatalf("Unsupported queue type: %s. Supported types are: memory, rabbitmq", *queueType)
	}

	// for now, we are using an in-memory queue
	jobQueue := &queue.InMemoryQueue{
		// we are using a buffered channel allows us to limit queue size^
		Jobs: make(chan model.EmailJob, *queueSize),
		Open: true,
	}

	var workerWg sync.WaitGroup
	workerWg.Add(*workerNum)

	// here, we start for the specified number of workers
	for i := 0; i < *workerNum; i++ {
		worker.StartWorker(i+1, jobQueue.Jobs, &workerWg)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/send-email", SendEmailHandler(jobQueue))

	server := &http.Server{
		Addr:    ":" + *port,
		Handler: mux,
	}

	go Shutdown(server, jobQueue, &workerWg)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server error: %v", err)
	}
}

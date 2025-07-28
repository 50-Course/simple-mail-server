# Simple Email Queue Service

## Architecture Design

```sh
.
├── main.go                # program entrypoint
├── server.go
├── queue/
│   ├── interface.go       # simple abstraction layer for queue, allowing reuse across our layers  (in-memory, RabbitMQ)
│   ├── memory.go          # buffered channel queue implementation
│   └── rabbitmq.go        # connector to rabbitmq
├── worker/
│   └── pool.go            # worker pool that reads from queue
├── email/
│   └── sender.go          # notification service
├── model/
│   └── job.go             # base model for jobs, and validation logic
└── go.mod

```

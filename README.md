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

## Usage

### Run the service

You can run the service using the following commands:

```sh
go run main.go --port=<preffered-port> --queue-size=<queue-size> --workers=<workers>
```

OR

```sh
go build -o email-service
./email-service
```

### Installation

The project uses built-in Go modules from the stdlib, therefore does not require any additional dependencies.
However, it does have a single dependency used during development for Live Server Reload (Air).
You can install the project using the following commands:

Clone the repository:

```sh
git clone github.com:50-Course/simple-mail-server.git
```

Then navigate to the project directory and run the following commands to install the dependencies:

```sh
go mod tidy
```

## Contributing

If you would like to contribute to this project, please follow these steps:

1. Fork the repository
2. Create a new branch for your feature or bug fix
3. Make your changes and commit them with a clear message
4. Push your changes to your forked repository
5. Create a pull request to the main repository

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

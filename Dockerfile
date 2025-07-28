FROM golang:1.24-alpine
WORKDIR /app
COPY . /app/

RUN go build -o email-service
CMD ["./email-service"]

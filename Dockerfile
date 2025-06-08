FROM golang:latest

WORKDIR  /app

RUN go mod download

COPY . .

RUN go build -o myapp

EXPOSE 9091

CMD["./myapp"]
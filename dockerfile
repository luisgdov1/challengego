FROM golang:latest
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
EXPOSE 8000
RUN go build -o main 
CMD ["./main"]

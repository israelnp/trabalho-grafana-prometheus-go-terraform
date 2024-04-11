FROM golang:1.22.1

WORKDIR /app
COPY . ./
RUN go mod download
RUN GOOS=linux go build -o /app/app

CMD ["/app/app"]



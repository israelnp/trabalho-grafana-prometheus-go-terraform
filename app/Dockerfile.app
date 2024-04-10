FROM golang:1.22.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . ./


RUN ls -la

COPY app.sqlite ./
RUN GOOS=linux go build -o /app/app

CMD ["/app/app"]



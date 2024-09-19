FROM ubuntu:latest
LABEL authors="poulin"

ENTRYPOINT ["top", "-b"]

FROM golang:1.18-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
##todo
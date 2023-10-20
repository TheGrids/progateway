FROM golang:1.20-alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build main.go
CMD ["/app/main"]
LABEL Name=serv Version=0.0.1
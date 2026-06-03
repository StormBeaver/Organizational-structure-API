FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o ./orgApi ./cmd/main.go

EXPOSE 8080

CMD [ "./orgApi" ]

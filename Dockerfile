FROM golang:1.18

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /my-app

EXPOSE 8080

CMD [ "/my-app" ]

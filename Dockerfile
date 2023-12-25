FROM golang:1.21.4-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY /src .

RUN go build -o main .

CMD [ "./main" ]
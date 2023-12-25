FROM golang:1.21.4-alpine

WORKDIR /app

COPY /src .

RUN go mod download
RUN go build -o main .

CMD [ "./main" ]
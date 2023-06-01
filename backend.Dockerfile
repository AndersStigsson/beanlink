FROM golang:1.16-alpine

WORKDIR /app

COPY ./backend/go.mod .
COPY ./backend/go.sum .
COPY ./backend/.env .
COPY ./backend/protos/ ./protos

RUN go mod download

COPY ./backend/*.go ./

RUN go build -o backend

ENTRYPOINT [ "./backend" ]

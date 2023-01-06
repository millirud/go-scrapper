FROM golang:1.20rc2-buster as dev

WORKDIR /home/app


COPY . .
RUN go mod download
RUN go install github.com/cosmtrek/air@v1.40.4
CMD air

FROM scratch as prod

ENV GIN_MODE=release

FROM golang:1.20rc2-buster as dev

WORKDIR /home/app


COPY . .
RUN go mod download
RUN go install github.com/cosmtrek/air@v1.40.4
CMD air

FROM golang:1.20rc2-alpine3.17 as build
WORKDIR /home/app
COPY . .
RUN go mod download
RUN go build -ldflags "-s -w" -o main ./cmd/app


FROM alpine as prod
WORKDIR /home/app
ENV GIN_MODE=release
COPY --from=build /home/app/main main
COPY --from=build /home/app/config/config.yml ./config/config.yml
CMD ./main
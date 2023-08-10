FROM golang:latest

ENV APP_HOME=/fashionshop
RUN mkdir ${APP_HOME}
WORKDIR ${APP_HOME}

COPY . ${APP_HOME}

RUN go mod tidy
ENTRYPOINT go run main.go
FROM golang:latest

ENV APP_HOME=/fashionshop
RUN mkdir ${APP_HOME}
WORKDIR ${APP_HOME}

COPY . ${APP_HOME}

RUN go mod tidy
RUN go build -o fashionshop
CMD [ "./fashionshop" ]
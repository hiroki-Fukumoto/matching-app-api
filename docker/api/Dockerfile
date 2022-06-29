FROM golang:1.17.7-buster

ENV GO111MODULE=auto

WORKDIR /go/api

COPY ./api .

RUN go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air && \
  go get -u github.com/gin-gonic/gin

COPY ./api/go.mod ./api/go.sum ./
RUN go mod download

RUN ["apt-get", "update"]
RUN ["apt-get", "install", "-y", "vim"]

CMD ["air", "-c", ".air.toml"]

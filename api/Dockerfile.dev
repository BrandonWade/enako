FROM golang:latest
WORKDIR /app

RUN go get github.com/pilu/fresh

COPY ./ /go/src/github.com/BrandonWade/enako/
WORKDIR /go/src/github.com/BrandonWade/enako/

CMD ["fresh"]

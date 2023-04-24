FROM golang:1.19.2-alpine

RUN apk update && apk --no-cache add git build-base

WORKDIR /go/src/app

COPY . .

RUN go install golang.org/x/tools/gopls@latest && \ 
  go install github.com/cosmtrek/air@latest

CMD ["air"]
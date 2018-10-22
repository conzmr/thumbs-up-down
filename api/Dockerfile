FROM golang:1.8
MAINTAINER Constanza

WORKDIR /go/src/app
COPY . .

# ENV GOPATH /go

RUN go get gopkg.in/mgo.v2/bson
RUN go get -u google.golang.org/api/classroom/v1
RUN go get -u golang.org/x/oauth2/...
RUN go get github.com/pilu/fresh
RUN go get github.com/gorilla/mux

ADD . /go/src/

CMD [ "fresh" ]

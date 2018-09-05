FROM golang:1.8
MAINTAINER Constanza
RUN adduser --disabled-password --gecos '' api
USER api

WORKDIR /go/src/app
COPY . .

# ENV GOPATH /go
# Install beego & bee
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee
RUN go get gopkg.in/mgo.v2/bson
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/garyburd/redigo/redis
ADD . /go/src/
RUN go get github.com/pilu/fresh
RUN go-wrapper download
RUN go-wrapper install

CMD [ "fresh" ]

FROM golang:1.6

RUN go get github.com/dgrijalva/jwt-go && \
    go get github.com/SkygearIO/skygear-server
RUN mkdir /go/src/app
ADD main.go /go/src/app
WORKDIR /go/src/app
RUN go install app
CMD /go/bin/app




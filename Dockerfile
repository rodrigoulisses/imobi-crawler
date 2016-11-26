FROM golang:1.6.3

RUN go get gopkg.in/mgo.v2
RUN go get github.com/PuerkitoBio/goquery

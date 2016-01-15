FROM r.fds.so:5000/golang1.5.3

ADD . /go/src/github.com/MISingularity/logging
WORKDIR /go/src/github.com/MISingularity/logging
RUN go get github.com/tools/godep
RUN godep go build -o mislogd /go/src/github.com/MISingularity/logging/be/cmd/main.go

#ENV MONGO_SERVICE_HOST 127.0.0.1
#ENV MONGO_SERVICE_PORT 27017

CMD /go/src/github.com/MISingularity/logging/mislogd -mongo-host $MONGO_SERVICE_HOST -mongo-port $MONGO_SERVICE_PORT
EXPOSE 50051

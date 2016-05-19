FROM r.fds.so:5000/golang1.5.3

ADD . /go/src/github.com/MISingularity/logging
WORKDIR /go/src/github.com/MISingularity/logging
#RUN go get github.com/bradfitz/http2
#RUN go get github.com/golang/protobuf/proto
#RUN go get golang.org/x/net/context
#RUN go get golang.org/x/oauth2
#RUN go get google.golang.org/cloud/compute/metadata
#RUN go get google.golang.org/cloud/internal
#RUN go get google.golang.org/grpc
#RUN go get gopkg.in/mgo.v2
RUN go get github.com/tools/godep
RUN godep go build -o logging_server /go/src/github.com/MISingularity/logging/be/cmd/main.go

ENV MONGO_SERVICE_HOST 42.159.133.35
ENV MONGO_SERVICE_PORT 27017

VOLUME /go/src/github.com/MISingularity/logging/log

CMD /go/src/github.com/MISingularity/logging/logging_server -mongo-host $MONGO_SERVICE_HOST -mongo-port $MONGO_SERVICE_PORT
EXPOSE 50051
EXPOSE 8080

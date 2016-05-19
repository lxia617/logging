sudo docker build -t logging .
sudo docker rm -f loggingx
sudo docker run -d -p 50051:50051 -p 8082:8080 -v /home/dsadmin/docker/logging_server/log:/go/src/github.com/MISingularity/logging/log --name loggingx logging
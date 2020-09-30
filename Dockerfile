# Dockerfile from https://github.com/hugobarzano/nlp.git
FROM    golang:1.13  
MAINTAINER    hugobarzano  
WORKDIR    /usr/src/app  
COPY    . /usr/src/app  
RUN    go get -d -v ./...  
EXPOSE    4343  
ENTRYPOINT    ["go", "run", "server.go"]  

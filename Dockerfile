FROM    golang:1.14
MAINTAINER    neuron
WORKDIR    /usr/src/app  
COPY    . /usr/src/app  
RUN    go get -d -v ./...  
EXPOSE    8080
ENTRYPOINT    ["go", "run", "cmd/neuron-orch/main.go"]

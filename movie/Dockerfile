FROM golang:1.22.5

RUN apt update && apt upgrade -y &&\
    apt install -y git\
    make openssh-client

WORKDIR /go/src/app

COPY . ./
RUN go mod tidy \
    && go mod verify


RUN go build -o  /main

EXPOSE 8083:8083

CMD ["/main"]
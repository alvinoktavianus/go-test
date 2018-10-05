FROM golang:1.11-alpine3.8

# Download and install the latest release of dep and git
ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR $GOPATH/src/backend_go
ADD . $GOPATH/src/backend_go/
RUN dep ensure
RUN go build -o main cmd/business/main.go
EXPOSE 9999
CMD [ "./main" ]

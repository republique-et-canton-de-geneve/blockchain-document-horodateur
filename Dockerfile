FROM golang:alpine
MAINTAINER Sylvain Laurent

RUN apk add texlive texlive-xetex
RUN	apk add --no-cache  bash \
    		gcc \
    		musl-dev \
    		openssl \
    		go

ENV GOBIN $GOPATH/bin
ENV PROJECT_DIR github.com/geneva_horodateur/
ENV PROJECT_NAME r-c-g-horodatage-server

ADD vendor /usr/local/go/src
ADD cmd /go/src/${PROJECT_DIR}/cmd
ADD models /go/src/${PROJECT_DIR}/models
ADD restapi /go/src/${PROJECT_DIR}/restapi
ADD merkle /go/src/${PROJECT_DIR}/merkle
ADD internal /go/src/${PROJECT_DIR}/internal
ADD template /go/src/${PROJECT_DIR}/template
ADD myservice.cert /go/src/${PROJECT_DIR}/
ADD myservice.key /go/src/${PROJECT_DIR}/

WORKDIR /go/src/${PROJECT_DIR}

RUN go build -v -o /go/bin/main /go/src/${PROJECT_DIR}/cmd/${PROJECT_NAME}/main.go
ADD run.sh /go/src/${PROJECT_DIR}/
ENTRYPOINT /go/src/${PROJECT_DIR}/run.sh

EXPOSE 8090

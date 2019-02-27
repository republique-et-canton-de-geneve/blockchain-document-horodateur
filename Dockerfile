FROM golang
MAINTAINER Sylvain Laurent

RUN apt-get update && apt-get -y install \
        texlive texlive-xetex texlive-lang-french && \
        rm -rf /var/lib/apt/lists/*

ENV GOBIN $GOPATH/bin
ENV PROJECT_DIR /go/src/github.com/geneva_horodateur/
ENV PROJECT_NAME r-c-g-horodatage-server

ADD vendor /usr/local/go/src

ADD cmd ${PROJECT_DIR}/cmd
ADD models ${PROJECT_DIR}/models
ADD restapi ${PROJECT_DIR}/restapi
ADD merkle ${PROJECT_DIR}/merkle
ADD internal ${PROJECT_DIR}/internal
ADD template ${PROJECT_DIR}/template
ADD myservice.cert ${PROJECT_DIR}/
ADD myservice.key ${PROJECT_DIR}/

WORKDIR ${PROJECT_DIR}

RUN go build -v -o /go/bin/main ${PROJECT_DIR}/cmd/${PROJECT_NAME}/main.go
ADD run.sh ${PROJECT_DIR}/
ENTRYPOINT ${PROJECT_DIR}/run.sh

EXPOSE 8090

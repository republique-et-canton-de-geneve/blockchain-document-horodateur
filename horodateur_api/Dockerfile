FROM golang
MAINTAINER Sylvain Laurent

RUN apt-get update && apt-get -y install \
        texlive texlive-xetex texlive-lang-french && \
        rm -rf /var/lib/apt/lists/*

# Create a group and user
RUN groupadd -g 999 appgroup && \
    useradd -r -u 999 -g appgroup appuser

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
WORKDIR ${PROJECT_DIR}

RUN go build -v -o /go/bin/main ${PROJECT_DIR}/cmd/${PROJECT_NAME}/main.go
ADD run.sh ${PROJECT_DIR}/

RUN chown appuser:appgroup -R .
# Tell docker that all future commands should run as the appuser user
USER appuser

ENTRYPOINT ${PROJECT_DIR}/run.sh

EXPOSE 8090

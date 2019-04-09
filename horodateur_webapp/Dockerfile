FROM golang:alpine AS build-env

ENV GOBIN $GOPATH/bin
ENV PROJECT_DIR /go/src/github.com/geneva_horodateur

ADD vendor ${PROJECT_DIR}/vendor

ADD main.go ${PROJECT_DIR}/
ADD mockup ${PROJECT_DIR}/mockup

WORKDIR ${PROJECT_DIR}

RUN go build -o webapp


FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/geneva_horodateur/webapp /app/
COPY --from=build-env /go/src/github.com/geneva_horodateur/mockup /app/mockup

EXPOSE 8080

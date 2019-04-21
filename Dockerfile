FROM golang:1.12.4-stretch
ENV APP_NAME=ghz-web
WORKDIR /usr/src
ADD ./ /usr/src/${APP_NAME}
WORKDIR /usr/src/${APP_NAME}
ENV CGO_ENABLED=1
RUN go build -mod vendor -o out/${APP_NAME}

FROM debian:stretch
ENV APP_NAME=ghz-web
WORKDIR /opt/${APP_NAME}
COPY --from=0 usr/src/${APP_NAME}/out/${APP_NAME} .
ADD ./ui/dist /opt/${APP_NAME}/ui/dist
CMD ["./ghz-web"]

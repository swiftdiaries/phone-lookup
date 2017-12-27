FROM golang
LABEL maintainer="Adhita Selvaraj <adhita.selvaraj@gmail.com>"

ADD . ~/go/src/github.com/swiftdiaries/phone-lookup/frontend

RUN go install github.com/swiftdiaries/phone-lookup/frontend

ENV PORT_1=9090
ENV REDIS_URL=127.0.0.1:6379

ENTRYPOINT /go/bin/frontend

EXPOSE 9090
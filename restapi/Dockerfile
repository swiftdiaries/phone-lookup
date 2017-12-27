FROM golang
LABEL maintainer="Adhita Selvaraj <adhita.selvaraj@gmail.com>"

ADD . ~/go/src/github.com/swiftdiaries/phone-lookup/restapi

RUN go install github.com/swiftdiaries/phone-lookup/restapi

ENV PORT_2=4040
ENV REDIS_URL=127.0.0.1:6379

ENTRYPOINT /go/bin/restapi

EXPOSE 4040
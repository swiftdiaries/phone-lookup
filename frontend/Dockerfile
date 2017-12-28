FROM golang
LABEL maintainer="Adhita Selvaraj <adhita.selvaraj@gmail.com>"

COPY . /go/src/github.com/swiftdiaries/phone-lookup

RUN go install github.com/swiftdiaries/phone-lookup/frontend

ENV PORT_1=9090
ENV REDIS_URL=redis://localhost:6379

ENTRYPOINT /go/bin/frontend

EXPOSE 9090
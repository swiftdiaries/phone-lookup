FROM golang
LABEL maintainer="Adhita Selvaraj <adhita.selvaraj@gmail.com>"

COPY . /go/src/github.com/swiftdiaries/phone-lookup

RUN go install github.com/swiftdiaries/phone-lookup/frontend

ENV PORT_1=8080

WORKDIR /go/src/github.com/swiftdiaries/phone-lookup

ENTRYPOINT /go/bin/frontend

EXPOSE 8080
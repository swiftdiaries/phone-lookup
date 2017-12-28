FROM golang
LABEL maintainer="Adhita Selvaraj <adhita.selvaraj@gmail.com>"

COPY . /go/src/github.com/swiftdiaries/phone-lookup

RUN go install github.com/swiftdiaries/phone-lookup/restapi

ENV PORT_2=4040
ENV REDIS_URL=redis://localhost:6379

ENTRYPOINT /go/bin/restapi

EXPOSE 4040
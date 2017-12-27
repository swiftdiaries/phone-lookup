FROM iron/go:dev
LABEL maintainer="Adhita Selvaraj <adhita.selvaraj@gmail.com>"

COPY . /app/src/github.com/swiftdiaries/phone-lookup
WORKDIR /app/src/github.com/swiftdiaries/phone-lookup

ENV HOME /app
ENV GOVERSION=1.9
ENV GOROOT $HOME/.go/$GOVERSION/go
ENV GOPATH $HOME
ENV PATH $PATH:$HOME/bin:$GOROOT/bin:$GOPATH/bin

RUN mkdir -p $HOME/.go/$GOVERSION
RUN cd $HOME/.go/$GOVERSION; curl -s https://storage.googleapis.com/golang/go$GOVERSION.linux-amd64.tar.gz | tar zxf -
RUN go install -v github.com/swiftdiaries/phone-lookup/restapi
RUN go install -v github.com/swiftdiaries/phone-lookup/frontend

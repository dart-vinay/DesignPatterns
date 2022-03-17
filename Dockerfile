FROM golang:1.16
#ARG _ENV=prod
#RUN echo $_ENV
WORKDIR $GOPATH/src/DesignPattern
COPY . $GOPATH/src/DesignPattern
RUN go mod vendor
RUN go build main.go
RUN pwd
RUN ls
EXPOSE 32007
CMD ./main
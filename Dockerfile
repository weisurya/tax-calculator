FROM golang:alpine
MAINTAINER Soen, Surya Soenaryo<wei.surya@gmail.com>

ARG APP_ENV
ENV GOBIN /go/bin
ENV GOPATH /app
ENV PATH=$GOPATH/bin:$PATH

RUN mkdir -p /app/src/tax-calculator
ADD . /app/src/tax-calculator
WORKDIR /app/src/tax-calculator

RUN go build -o ./output ./main.go

EXPOSE 9090

CMD ["./main"]
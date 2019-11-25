FROM golang:latest AS build

WORKDIR /go/csc482
COPY . .

RUN go get github.com/jamespearly/loggly
RUN go get github.com/aws/aws-sdk-go/aws
RUN go get github.com/aws/aws-sdk-go/service


ENV AWS_ACCESS_KEY_ID=
ENV AWS_SECRET_ACCESS_KEY=
ENV AWS_DEFAULT_REGION=

# Start the application
CMD ["go", "run","main.go"]

FROM golang:latest AS build

WORKDIR /go/csc482
COPY . .

RUN go get github.com/jamespearly/loggly
RUN go get github.com/aws/aws-sdk-go/aws
RUN go get github.com/aws/aws-sdk-go/service


ENV AWS_ACCESS_KEY_ID=AKIA34XNLPJYOAFQI4TV
ENV AWS_SECRET_ACCESS_KEY=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
ENV AWS_DEFAULT_REGION=us-east-1

# Start the application
CMD ["go", "run","main.go"]

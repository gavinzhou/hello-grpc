FROM golang:1.12 as builder

RUN mkdir /hello
WORKDIR /hello

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

#build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/greeter

# STEP 2 build a small image
# start from scratch
FROM scratch
# Copy our static executable
COPY --from=builder /go/bin/greeter /go/bin/greeter
EXPOSE 8972
ENTRYPOINT ["/go/bin/greeter"]

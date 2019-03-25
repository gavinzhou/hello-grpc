FROM golang:alpine as builder

RUN apk update && apk add git

RUN mkdir /hello
WORKDIR /hello

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

#build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/greeter
RUN adduser -D -g '' appuser

# STEP 2 build a small image
# start from scratch
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
COPY --from=builder /go/bin/greeter /go/bin/greeter
EXPOSE 8972
USER appuser
ENTRYPOINT ["/go/bin/greeter"]

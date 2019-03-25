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
RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
RUN adduser -D -g '' appuser

# STEP 2 build a small image
# start from scratch
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
COPY --from=builder /go/bin/greeter /go/bin/greeter
COPY --from=builder /bin/grpc_health_probe /bin/grpc_health_probe
EXPOSE 8972
USER appuser
ENTRYPOINT ["/go/bin/greeter"]

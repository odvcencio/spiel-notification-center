FROM golang:1.9-alpine as builder

# Installing dependencies
RUN apk add --no-cache \
      git \
      ca-certificates \
      curl \
      tzdata

# Downloading dep
RUN curl -sSL \
         -o $GOPATH/bin/dep \
         https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 \
 && chmod +x $GOPATH/bin/dep

# Setting work directory
WORKDIR /go/src/spiel/notification-center

# Populating vendor directory and building
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only

# Copying rest of the code and building
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ./notification-center

FROM scratch
COPY --from=builder /go/src/spiel/notification-center/notification-center ./notification-center
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo/
CMD ["./notification-center"]

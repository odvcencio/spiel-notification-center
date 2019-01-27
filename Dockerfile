###############
# Build Stage #
###############
FROM golang:1.9-alpine as builder

# Installing dependencies
RUN apk add --no-cache \
      git \
      ca-certificates \
      curl \
      tzdata

# Installing dep
RUN curl -sSL https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Setting work directory
WORKDIR ${GOPATH}/src/spiel/notification-center

# Populating vendor directory
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only

# Copying rest of the code and building
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ./notification-center

#################
# Release Stage #
#################
FROM scratch

# Copying files and folders from builder stage
COPY --from=builder /go/src/spiel/notification-center/notification-center ./notification-center
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo/

# Exposing ports
EXPOSE 8080/TCP

# Setting docker command
CMD ["./notification-center"]

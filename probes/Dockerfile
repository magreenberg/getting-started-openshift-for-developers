FROM registry.access.redhat.com/ubi8/ubi-minimal:latest as builder

RUN microdnf -y install go ca-certificates

WORKDIR /app

# The default build uses the "vendor" directory.
# Update the "vendor" directory with: go mod vendor
# Bypass the "vendor" directory during the build with: -build-arg VENDOR_CACHE=false
ARG VENDOR_CACHE=true
# load static dependencies to speed build
COPY go.mod ./
# preload go modules as cache
RUN set -x;[[ "${VENDOR_CACHE}" != "true" ]] && go mod download || true

COPY . .

RUN set -x;[[ "${VENDOR_CACHE}" != "true" ]] && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /app/server || \
    GOFLAGS=-mod=vendor CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /app/server

# Runtime image
FROM scratch AS base
COPY --from=builder /etc/pki /etc/ssl /etc/
COPY --from=builder /app/server /app/server
ENTRYPOINT ["/app/server"]

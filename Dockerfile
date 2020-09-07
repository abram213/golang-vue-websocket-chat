# golang alpine 1.14.1
FROM golang:alpine as builder

RUN apk update && apk add --update alpine-sdk gcc musl-dev bash

# Enable go mod
ENV GO111MODULE=on

# Set work dir
WORKDIR /go/webapp/

COPY . .

RUN go mod verify

# Build the binary.
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"'

FROM scratch

# Import from builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable and config file
COPY --from=builder go/webapp/chat webapp/chat
COPY --from=builder go/webapp/config.env webapp/config.env
COPY --from=builder go/webapp/dist/ webapp/dist/

EXPOSE 8080
# Run the webapp binary.
ENTRYPOINT ["webapp/chat", "--migrate"]
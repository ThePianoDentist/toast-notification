FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/ThePianoDentist/toast-notification/
WORKDIR /go/src/github.com/ThePianoDentist/toast-notification
RUN go mod download
COPY . /go/src/github.com/ThePianoDentist/toast-notification
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/toast-notification github.com/ThePianoDentist/toast-notification
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/toast-notification /go/src/github.com/ThePianoDentist/toast-notification

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/ThePianoDentist/toast-notification/build/toast-notification /usr/bin/toast-notification
COPY --from=builder /go/src/github.com/ThePianoDentist/toast-notification/firebaseServiceAccountKey.json /opt/firebaseServiceAccountKey.json
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/toast-notification"]

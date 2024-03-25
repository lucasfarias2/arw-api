FROM golang:latest as gobuilder

ENV GO111MODULE=on

WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o arw-api .

FROM alpine:latest

WORKDIR /root/

COPY --from=gobuilder /go/src/app/arw .

EXPOSE 8080

ENV APP_ENV=production

ENTRYPOINT ["/root/arw-api"]
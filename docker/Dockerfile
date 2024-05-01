# stage 1 build
FROM golang:1.21.5-alpine as builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o uploader .

# stage 2 migrations
FROM migrate/migrate as migrate


# stage 3 build
FROM alpine:latest as host

RUN apk update
RUN apk upgrade

COPY --from=builder /app/uploader /
COPY --from=migrate /migrate /usr/local/bin/migrate
COPY ./migrations /migrations
COPY ./docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

WORKDIR /

ENV AWS_REGION=us-east-1
ENV INPUT_S3_BUCKET=bucket-name
ENV PORT=4000

ENTRYPOINT ["/entrypoint.sh"]
CMD ["/uploader"]
FROM golang:1.21.5-alpine as builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o uploader .

FROM alpine:latest as host

RUN apk update
RUN apk upgrade

COPY --from=builder /app/uploader /
COPY ./videos/templates /videos/templates

WORKDIR /

ENV AWS_REGION=us-east-1
ENV INPUT_S3_BUCKET=bucket-name
ENV PORT=4000

CMD ["/uploader"]
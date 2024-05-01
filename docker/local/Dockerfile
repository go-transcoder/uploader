
# stage 2 migrations
FROM migrate/migrate as migrate


FROM golang:1.21.5-alpine as builder

ENV PORT="4000"
ENV AWS_REGION="us-east-1"
ENV INPUT_S3_BUCKET="bucket-name"
ENV DBUSER="test"
ENV DBPASS="password"
ENV DBHOST="localhost"
ENV DBPORT="5432"
ENV DBNAME="test"
ENV SSLMODE="disable"

RUN mkdir /app
WORKDIR /app
COPY . .

RUN export GOPATH="$HOME/go"
RUN export PATH="$GOPATH/bin:$PATH"

# copy migrations and migrate
COPY --from=migrate /migrate /usr/local/bin/migrate
COPY ./migrations /migrations
COPY ./docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# install the gow package to watch changes and rerun
RUN go install github.com/mitranim/gow@latest

ENTRYPOINT ["/entrypoint.sh"]

CMD ["gow", "run", "."]
#!/bin/sh

urlencode() {
  string="$1"
  encoded=""
  len="${#string}"
  i=0
  while [ "$i" -lt "$len" ]; do
    char="${string:$i:1}"
    case "$char" in
      [a-zA-Z0-9.~_-]) encoded="$encoded$char" ;;
      *) encoded="$encoded$(printf '%%%02X' "'$char")" ;;
    esac
    i=$((i + 1))
  done
  echo "$encoded"
}
escaped_password=$(urlencode "$DBPASS")

#echo "$escaped_password"
#echo postgres://"$DB_USER":"$escaped_password"@"$DB_HOST":"$DB_PORT"/"$DB_NAME"?sslmode=disable
migrate -path=/migrations -database postgres://"$DBUSER":"$escaped_password"@"$DBHOST":"$DBPORT"/"$DBNAME"?sslmode="$SSLMODE" up

exec "$@"
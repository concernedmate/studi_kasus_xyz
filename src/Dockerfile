# 1st stage build go binary
FROM golang:1.23.4-alpine AS build-go

WORKDIR /usr/src

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o ./app_bin64
RUN mkdir /usr/app

RUN mv ./app_bin64 /usr/app/app_bin64
RUN mv ./certs /usr/app/certs
RUN mv ./.env /usr/app/.env

RUN mkdir /usr/app/logs
RUN mkdir /usr/app/uploads

# 2nd stage
FROM alpine:latest
RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

WORKDIR /usr/app

COPY --from=build-go /usr/app /usr/app

# uncomment below to use non root container
# RUN adduser -u 1000 -D app
# RUN chown -R app:app /usr/app
# USER app

CMD ["./app_bin64"]
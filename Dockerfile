FROM golang:1.20 as builder

WORKDIR /go/src/github.com/onlyfeed/www
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o www
RUN pwd
RUN ls -lah out

FROM alpine
RUN apk add --no-cache ca-certificates

RUN pwd
RUN ls -lah
COPY out /
COPY --from=builder /go/src/github.com/onlyfeed/www/www /www
COPY --from=builder /go/src/github.com/onlyfeed/www/out /dist
RUN ls -lah
RUN ls -la dist

CMD ["/www"]
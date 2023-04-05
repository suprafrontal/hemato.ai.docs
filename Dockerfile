FROM golang:1.20

WORKDIR /go/src/github.com/suprafrontal/hemato.ai.docs
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o www

CMD ["./www"]
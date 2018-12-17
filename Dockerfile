FROM golang:1.11-alpine

WORKDIR /

COPY app /

CMD ["./app"]

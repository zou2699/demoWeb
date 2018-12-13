FROM golang:1.10-alpine

WORKDIR /

COPY app /

CMD ["./app"]

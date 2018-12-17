FROM golang:1.11

WORKDIR /

COPY app /

CMD ["./app"]

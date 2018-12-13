FROM golang

WORKDIR /

COPY app /

CMD ["./app"]

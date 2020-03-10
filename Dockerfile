FROM golang:1.14-alpine
WORKDIR /src

COPY . .
RUN go build -o /src/bin/arp .

CMD /src/bin/arp
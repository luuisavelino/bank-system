FROM golang:1.21.1 AS builder

WORKDIR /go/src/github.com/luuisavelino/system-bank/

COPY ./bank .

RUN go get -d -v ./...

RUN go install -v ./...

COPY ./bank/cmd/bank/main.go ./

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /go/src/github.com/luuisavelino/system-bank/app ./

CMD ["./app"]

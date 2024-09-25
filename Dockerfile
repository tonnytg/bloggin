#Build Builder aplication
FROM golang
WORKDIR /var/app
COPY . .

RUN go mod tidy
RUN go build -o bloggin cmd/webserver/main.go

CMD ["./bloggin"]
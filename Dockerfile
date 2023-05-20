#Build Builder aplication
FROM golang
WORKDIR /var/app
COPY . .
RUN go build -o /var/app/bloggin

CMD ["tail", "-f", "/dev/nul"]
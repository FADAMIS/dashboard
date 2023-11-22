FROM golang:1.21.3
WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o dashboard .
EXPOSE 8080

CMD ./dashboard

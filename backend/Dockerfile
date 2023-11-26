FROM golang:1.21.3-alpine
WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o dashboard .
EXPOSE 8080

CMD ./dashboard

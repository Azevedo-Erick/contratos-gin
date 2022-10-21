FROM golang:1.19-alpine3.16
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod ./
COPY  go.sum ./
RUN go mod download
COPY src  ./
RUN go mod tidy
RUN go build -o ./test src/cmd/main.go
EXPOSE 3000

CMD [ "/cmd/test" ]
FROM golang:1.22.2-bullseye
RUN go install github.com/cosmtrek/air@latest
WORKDIR /go/src/app
COPY . .
RUN go mod download
EXPOSE 3000
CMD ["air", "-c", ".air.toml"]

LABEL  version="1.0-Alpha1"

WORKDIR /backend

COPY . .
RUN go mod init tea 
RUN go get . 
RUN GOOS=linux GOARCH=amd64 go build -o backend-bubbletea

FROM alpine:3.15

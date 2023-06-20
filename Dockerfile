FROM golang:1.20-alpine3.16
COPY . /
WORKDIR /
RUN go mod download
RUN go build -o libre .
CMD ./libre 

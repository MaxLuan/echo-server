FROM golang:alpine
ENV ServerName=test
ENV http=false
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o main .
WORKDIR /dist
RUN cp /build/main .
EXPOSE 3333/tcp
CMD ["/dist/main"]
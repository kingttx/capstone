FROM golang:1.16-alpine

WORKDIR /app
COPY go.mod ./

COPY *.go ./
RUN go build -o /webapp

EXPOSE 8080

CMD [ "/webapp" ]

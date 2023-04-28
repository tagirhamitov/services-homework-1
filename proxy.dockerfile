FROM golang:1.20

WORKDIR /code

COPY src ./
RUN go mod download && go mod verify
RUN go build -o /app app/proxy/main.go

EXPOSE 2000

CMD /app

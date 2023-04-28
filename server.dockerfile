FROM golang:1.20

ARG SERIALIZATION_FORMAT

WORKDIR /code

RUN apt update && apt install --yes --no-install-recommends protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
ENV PATH="$PATH:$GOROOT/bin/protoc-gen-go"

COPY src ./
RUN protoc -I=proto proto/*.proto --go_out=pb
RUN go mod download && go mod verify
RUN go build -o /app app/server/main.go

EXPOSE 2000

ENV SERIALIZATION_FORMAT=${SERIALIZATION_FORMAT}
CMD /app "${SERIALIZATION_FORMAT}"

# Run Health Check Api

FROM golang:latest
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ../.. .

RUN go build -v -o /usr/local/bin/app


# Install CompileDaemon
RUN go get github.com/githubnemo/CompileDaemon

COPY . .

COPY ./entrypoint.sh /entrypoint.sh

ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh

ENTRYPOINT["sh","/entrypoint.sh"]

#CMD ["app"]

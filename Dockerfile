FROM golang

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

RUN mkdir -p /usr/local/bin/xmapi
RUN go build -v -o /usr/local/bin/xmapi ./...

CMD ["/usr/local/bin/xmapi/api"]
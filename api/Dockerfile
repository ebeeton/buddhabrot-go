FROM golang:1.20

WORKDIR /usr/src/app

COPY ./api/go.mod ./api/go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/ ./api/...

CMD [ "buddhabrot-go" ]

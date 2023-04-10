FROM golang:1.20

WORKDIR /usr/src/app

COPY ./plotter/go.mod ./plotter/go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/ ./plotter/...

CMD [ "plotter" ]

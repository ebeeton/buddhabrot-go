FROM golang:1.20.3-alpine3.17 AS build

WORKDIR /usr/src/app

COPY ./plotter/go.mod ./plotter/go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o plotter ./plotter/...

FROM alpine:3.17

COPY --from=build /usr/src/app/plotter /usr/local/bin/

CMD [ "plotter" ]

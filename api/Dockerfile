FROM golang:1.20.3-alpine3.17 AS build

WORKDIR /usr/src/app

COPY ./api/go.mod ./api/go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o api ./api/...

FROM alpine:3.17

COPY --from=build /usr/src/app/api /usr/local/bin/

CMD [ "buddhabrot-go" ]
FROM golang:1.25.5-alpine3.23 AS build

WORKDIR /usr/src/app

COPY ./api/go.mod ./api/go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o api ./api/...

FROM alpine:3.23

COPY --from=build /usr/src/app/api /usr/local/bin/

CMD [ "api" ]
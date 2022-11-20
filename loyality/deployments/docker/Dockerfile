FROM golang:1.19-alpine as build
WORKDIR /go/bin/main
COPY ../../../loyality ./loyality
COPY ../../../libs ./libs
WORKDIR /go/bin/main/loyality
#RUN go mod tidy
RUN go mod download
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/app/main.go

FROM scratch
COPY --from=build /go/bin/main/loyality/main /main
COPY --from=build /go/bin/main/loyality/.env /.env
ENTRYPOINT ["/main"]
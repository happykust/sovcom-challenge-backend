FROM golang:1.19-alpine as build
WORKDIR /go/bin/main
COPY ../../../api-gateway ./api-gateway
COPY ../../../libs ./libs
WORKDIR /go/bin/main/api-gateway
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/app/main.go

FROM scratch
COPY --from=build /go/bin/main/api-gateway/main /main
COPY --from=build /go/bin/main/api-gateway/.env /.env
ENTRYPOINT ["/main"]

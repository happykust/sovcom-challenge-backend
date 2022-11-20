FROM golang:1.19-alpine as build
WORKDIR /go/bin/main
COPY ../../../deals ./deals
COPY ../../../libs ./libs
WORKDIR /go/bin/main/deals
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/app/main.go

FROM scratch
COPY --from=build /go/bin/main/deals/main /main
COPY --from=build /go/bin/main/deals/.env /.env
ENTRYPOINT ["/main"]




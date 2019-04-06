# GoLang build stage
FROM golang:1.12 as gobuilder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o entryserver

# Final stage
FROM scratch

WORKDIR /app

COPY --from=gobuilder /app/entryserver /app/entryserver

ENTRYPOINT ["/app/entryserver"]
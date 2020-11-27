# Build stage
FROM golang:alpine as builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tzserver

# Final stage
FROM scratch as final

COPY --from=builder /app/tzserver .
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /

#ENV TZ=Asia/Baghdad
ENV ZONEINFO=/zoneinfo.zip

ENTRYPOINT ["./tzserver"]
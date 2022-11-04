## STAGE-1 Build golang app
FROM golang:1.18 as BUILDER
WORKDIR /app
COPY . .
RUN go build

## STAGE-2 Copy the binary to busybox
FROM busybox
WORKDIR /app
COPY --from=BUILDER /app/hello /app
CMD ["hello"]
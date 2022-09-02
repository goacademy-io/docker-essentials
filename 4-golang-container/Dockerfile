## STAGE-1 Build golang app
FROM golang:1.18 as BUILDER
WORKDIR /usr/src/app
COPY . .
RUN go build

## STAGE-2 Copy the binary to busybox
FROM busybox
COPY --from=BUILDER /usr/src/app/hello /usr/share/nginx/html/
CMD ["hello"]
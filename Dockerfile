FROM golang:1.17

WORKDIR /go/src/gitlab.com/tokend/subgroup/tokenproject

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/tokenproject gitlab.com/tokend/subgroup/tokenproject


###

FROM alpine:3.9

COPY --from=0 /usr/local/bin/tokenproject /usr/local/bin/tokenproject
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["tokenproject"]

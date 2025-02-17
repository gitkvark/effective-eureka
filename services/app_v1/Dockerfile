FROM cgr.dev/chainguard/go AS builder
COPY go.mod go.sum app_v1.go /app/
RUN cd /app && go build

FROM cgr.dev/chainguard/glibc-dynamic
COPY --from=builder /app/app_v1 /usr/bin/

EXPOSE 7001

ENTRYPOINT ["/usr/bin/app_v1"]

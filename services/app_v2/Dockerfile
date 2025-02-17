FROM cgr.dev/chainguard/go AS builder
COPY go.mod go.sum app_v2.go /app/
RUN cd /app && go build

FROM cgr.dev/chainguard/glibc-dynamic
COPY --from=builder /app/app_v2 /usr/bin/

EXPOSE 7002

ENTRYPOINT ["/usr/bin/app_v2"]

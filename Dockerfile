FROM golang@sha256:8e02eb337d9e0ea459e041f1ee5eece41cbb61f1d83e7d883a3e2fb4862063fa AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o currencyAPI .

# ==-==-==-==-==-==-==-==-==

FROM alpine@sha256:c3f8e73fdb79deaebaa2037150150191b9dcbfba68b4a46d70103204c53f4709 AS runtime

ENV PORT=8000
ENV AUTHOR=a.bezpyatko
ENV VERSION=1.1.0

RUN addgroup -g 1001 appuser && \ 
    adduser -u 1001 -G appuser -S appuser
COPY --from=builder --chown=appuser:appuser /app/currencyAPI /app/currencyAPI

USER appuser
WORKDIR /app

EXPOSE 8000

HEALTHCHECK --interval=30s --timeout=10s --start-period=10s --retries=3 \
            CMD wget --spider --quiet http://localhost:$PORT/info || exit 1

CMD ["./currencyAPI"]

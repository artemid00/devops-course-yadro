FROM golang@sha256:c2a1f7b2095d046ae14b286b18413a05bb82c9bca9b25fe7ff5efef0f0826166 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o currencyAPI .

# ==-==-==-==-==-==-==-==-==

FROM gcr.io/distroless/static:nonroot AS runtime

WORKDIR /app
COPY --from=builder /app/currencyAPI /app/currencyAPI

USER nonroot

EXPOSE 8000

CMD ["/app/currencyAPI"]

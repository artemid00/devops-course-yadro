FROM golang@sha256:8e02eb337d9e0ea459e041f1ee5eece41cbb61f1d83e7d883a3e2fb4862063fa AS builder

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

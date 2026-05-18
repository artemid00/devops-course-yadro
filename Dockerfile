FROM golang@sha256:313faae491b410a35402c05d35e7518ae99103d957308e940e1ae2cfa0aac29b AS builder

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

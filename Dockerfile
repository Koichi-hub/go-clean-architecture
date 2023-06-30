FROM golang:1.20-alpine3.18 as builder
RUN apk update && apk upgrade && \
    apk --update add make build-base
WORKDIR /app
COPY . .
RUN make build

FROM golang:1.20-alpine3.18
WORKDIR /app
COPY --from=builder /app/.env /app/
COPY --from=builder /app/main /app/
EXPOSE 8080
CMD /app/main
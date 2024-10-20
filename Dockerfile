FROM public.ecr.aws/docker/library/golang:1.23.1-alpine3.18 AS builder
WORKDIR /nova
COPY . .
RUN go mod download
RUN go test -race -count=1 ./...
RUN mkdir build
RUN go build -o build/nova .

FROM public.ecr.aws/docker/library/alpine:latest AS runner
WORKDIR /nova
COPY --from=builder /nova/build/nova /nova/build/nova
COPY --from=builder /nova/.env.production /nova/.env
EXPOSE 8080
CMD ["/nova/build/nova"]

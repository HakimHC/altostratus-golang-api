FROM public.ecr.aws/docker/library/golang:alpine3.20 as builder

WORKDIR /app
COPY go.* /app/
RUN go mod download

COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/api_service -ldflags "-X main.build=." ./cmd

FROM public.ecr.aws/docker/library/alpine:3.18

COPY --from=builder /app/bin/api_service /usr/bin/api_service
RUN chmod +x /usr/bin/api_service

EXPOSE 80

CMD ["/usr/bin/api_service"]


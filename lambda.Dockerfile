FROM golang:1.21 AS build-stage

WORKDIR /app

# cache dependencies
COPY go.mod go.sum ./
RUN go mod download -x
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint cmd/main.go

# copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2 as release-stage
COPY --from=build-stage /entrypoint /entrypoint
ENTRYPOINT ["/entrypoint"]


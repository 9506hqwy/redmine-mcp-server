FROM golang:1.24.3-bookworm AS build

WORKDIR /build

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,target=. \
    VERSION=$(git tag --sort -version:refname | head -n 1); \
    COMMIT=$(git rev-parse HEAD); \
    CGO_ENABLED=0 go build \
        -trimpath \
        -ldflags "-s -w -X main.version=${VERSION#v} -X main.commit=${COMMIT}" \
        -o /bin/redmine-mcp-server \
        ./cmd/redmine-mcp-server/main.go

FROM gcr.io/distroless/static-debian12

COPY --from=build /bin/redmine-mcp-server /

ENTRYPOINT ["/redmine-mcp-server"]

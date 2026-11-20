default:
    @just --list

sync:
    uv sync
    cd fetcher && go mod tidy

fetch:
    cd fetcher && go run ./cmd
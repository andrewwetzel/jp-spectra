default:
    @just --list

sync:
    uv sync
    cd fetcher && go mod tidy

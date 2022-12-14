@echo off  
set "command=%1"

IF not exist bin\ (
    mkdir bin
)

IF "%command%"=="clean" (
    go mod tidy
)

IF "%command%"=="test" (
    go test ./tests/...
)

IF "%command%"=="run" (
    go run src/main.go
)

IF "%command%"=="build" (
    go build -o bin/trackr.exe src/main.go
) ELSE (
    IF "%command%"=="" (
        go build -o bin/trackr.exe src/main.go
    )
)
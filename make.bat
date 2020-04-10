@ECHO OFF

IF "%1" == "" GOTO rohan
IF "%1" == "rohan" GOTO rohan

:gondor
    del gondor.exe
    go.exe build -ldflags="-s -w" ./cmd/gondor/
:rohan
    del rohan.exe
    go.exe build -ldflags="-s -w" ./cmd/rohan/

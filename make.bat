@ECHO OFF

IF "%1" == "all" GOTO gondor
IF "%1" == "" GOTO moria
IF "%1" == "moria" GOTO moria
IF "%1" == "rohan" GOTO rohan

:gondor
    del gondor.exe
    go.exe build -ldflags="-s -w" ./cmd/gondor/
:moria
    del moria.exe
    go.exe build -ldflags="-s -w" ./cmd/moria/
:rohan
    del rohan.exe
    go.exe build -ldflags="-s -w" ./cmd/rohan/

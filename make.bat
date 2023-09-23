@echo off

setlocal

set dist=dist
set target=wzry-skin

rem build for Windows
go build -o %dist%/windows/%target%.exe

rem cross platform compile
set CGO_ENABLED=0
set GOARCH=amd64

rem build for Linux
set GOOS=linux
go build -o %dist%/linux/%target%

rem build for macOS
set GOOS=darwin
go build -o %dist%/darwin/%target%

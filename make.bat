@echo off

setlocal

set dist=dist
set target=wzry-skin
set build=go.exe build -ldflags="-s -w"

rem build for Windows
%build% -o %dist%/windows/%target%.exe

rem cross platform compile
set CGO_ENABLED=0
set GOARCH=amd64

rem build for Linux
set GOOS=linux
%build% -o %dist%/linux/%target%

rem build for macOS
set GOOS=darwin
%build% -o %dist%/darwin/%target%

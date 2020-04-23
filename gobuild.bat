@echo off
rem BUILD WEB APPLICATION 
rem (the resulting executables are in the current folder)
@echo on
go clean -cache ./...
go build -o . -v ./...
@echo off
if %ERRORLEVEL% GEQ 1 echo !!!!! ERROR !!!!!
rem dir *.exe
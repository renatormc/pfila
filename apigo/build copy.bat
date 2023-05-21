@echo off
go1.20 build -o .\dist\pfila.exe && cd runner && go1.20 build -o ..\dist\pfila_runner.exe
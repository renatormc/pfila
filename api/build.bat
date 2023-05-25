@echo off
go1.20 build -o .\dist\pfila.exe && cd updater && go1.20 build -o ..\dist\update.exe
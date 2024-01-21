@echo off
setlocal enabledelayedexpansion

for %%I in (%0) do set SCRIPT_DIR=%%~dpI

"%SCRIPT_DIR%\pfila_runner.exe" --port %1 --proc-id %2 > %3 2>&1

endlocal

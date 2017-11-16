@echo off
REM ----------------------------------------
if /i "%~1" == "" goto :help

rem https://www.microsoft.com/resources/documentation/windows/xp/all/proddocs/en-us/batch.mspx?mfr=true
set filename=%~n1
set fileextension=%~x1

echo Building file %filename%

rem Remove the symbol and debug info with go build -ldflags "-s -w"
go build -ldflags="-s -w" %1

if /i "%~2" == "-u" call :upx
goto :exit

REM ----------------------------------------
:upx
    echo UPX
    REM for %%i in ("%1%") do (
    REM     echo filename=%%~ni
    REM     echo fileextension=%%~xi
    REM )

    echo Compressing file %filename%.exe
    upx -9 --best --ultra-brute -k %filename%.exe

    REM EXIT /B will return to the position where you used CALL
    exit /b

REM ----------------------------------------
:help
    echo Filename is needed !!!
    echo.
    echo $ cmpl source_file.go

REM ----------------------------------------
:exit
    echo Done
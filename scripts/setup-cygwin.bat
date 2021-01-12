@ECHO OFF

REM -- Downloading cygwin
curl "https://www.cygwin.com/setup-x86_64.exe" -O setup-x86_64.exe

SET LOCALDIR=%CD%
SET ROOTDIR=C:/cygwin
SET PACKAGES=bash,bash-completion,openssh,rxvt,python,git,wget,curl,vim,gcc-core,make

REM -- Running cygwin setup
echo Installing
setup-x86_64.exe -R "%ROOTDIR%" -P "%PACKAGES%"
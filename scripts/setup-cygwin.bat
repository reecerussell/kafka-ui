@ECHO OFF

REM -- Downloading cygwin
curl "https://www.cygwin.com/setup-x86_64.exe" -O setup-x86_64.exe

SET ROOTDIR=C:/cygwin
SET PACKAGES=bash,bash-completion,openssh,rxvt,python,git,wget,curl,vim,gcc-core,make
SET SITE=https://cygwin.mirror.uk.sargasso.net

REM -- Running cygwin setup
echo Installing
setup-x86_64.exe -q -s "%SITE%" -R "%ROOTDIR%" -P "%PACKAGES%"
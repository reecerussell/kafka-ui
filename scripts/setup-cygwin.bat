@ECHO OFF
REM

SETLOCAL
REM
cd %HOMEPATH%\Downloads

SET SITE=http://cygwin.mirrors.pair.com/
SET LOCALDIR=%CD%
SET ROOTDIR=C:/cygwin

SET PACKAGES=bash,bash-completion,openssh,rxvt,python,git,wget,curl,vim,gcc-core,make

REM
echo Installing
setup-x86.exe --disable-buggy-antivirus -q -s %SITE% -l "%LOCALDIR%" -R "%ROOTDIR%" -P "%PACKAGES%"
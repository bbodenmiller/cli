ECHO Hello

SET GOPATH=%CD%\go
SET CF_DIAL_TIMEOUT=15

SET PATH=C:\Go\bin;%PATH%
SET PATH=%GOPATH%\bin;%PATH%
SET PATH=C:\Program Files\GnuWin32\bin;%PATH%
SET PATH=%CD%;%PATH%

SET SKIP_SSL_VALIDATION=false

go get -v -u github.com/onsi/ginkgo/ginkgo

cd %GOPATH%\src\code.cloudfoundry.org\cli
go build -tags="forceposix" -o cf.exe
MOVE %CD%\cf.exe %GOPATH%\bin\cf.exe

echo "CF_INT_CLIENT_CREDENTIALS_TEST_MODE: $CF_INT_CLIENT_CREDENTIALS_TEST_MODE"

ginkgo.exe -r -randomizeAllSpecs -slowSpecThreshold 60 -nodes 8 %CF_TEST_SUITE% || exit 1


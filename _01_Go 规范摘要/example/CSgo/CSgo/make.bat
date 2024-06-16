set GOARCH=amd64
set CGO_ENABLED=1
go build -ldflags "-s -w" -buildmode=c-shared -o CSgo.Interop.dll netCgo.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/geep-linux
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o dist/geep-mac-arm
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/geep-windows-x64.exe
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o dist/geep-windows-x32.exe
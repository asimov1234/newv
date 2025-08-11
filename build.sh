CGO_ENABLED=0 go build -o xray -trimpath -ldflags "-s -w -buildid=" ./main


CGO_ENABLED=0 go build  -o build/wemebox-server main.go
rsync -az ./build wemebox: --progress

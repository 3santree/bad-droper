all:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -H=windowsgui" -o ./pivottemplate/a.exe ./pivottemplate/main.go

# upx --best template/a.exe

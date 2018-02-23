buildBin:
	rm tpl/bindata.go
	go-bindata -nometadata -o tpl/bindata.go -ignore bindata.go -pkg tpl tpl/*
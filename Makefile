buildBin:
	rm tpl/bindata.go
	go-bindata -nometadata -o tpl/bindata.go -ignore bindata.go -pkg tpl tpl/*


build: buildBin
	go build -o bin/bag cmd/*.go

genLib: build
	rm -r lib
	./bin/bag vomit --onlyBuiltIn=true --to=lib
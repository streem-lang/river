all:
	go build src/main.go
	mv ./main bin/river

clean:
	rm -rf bin/*


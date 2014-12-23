all:
	gcc src/main.c -o bin/river
	gcc src/update.c -o bin/rivup
	gcc src/doc.c -o bin/rivdoc

clean:
	rm -rf bin/*


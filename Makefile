OUT=build/main

all:
	go build -o $(OUT)

.PHONY: run
run: all
	$(OUT)

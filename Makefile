SRC=main.go
OUT=build/main

$(OUT): $(SRC)
	go build -o $(OUT)

.PHONY: run
run: $(OUT)
	$(OUT)

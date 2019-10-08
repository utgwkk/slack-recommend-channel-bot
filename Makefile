SRC=main.go slack.go
OUT=build/main

$(OUT): $(SRC)
	env GO111MODULE=on go build -o $(OUT)

.PHONY: run
run: $(OUT)
	$(OUT)

.PHONY: clean
clean:
	rm -f $(OUT)

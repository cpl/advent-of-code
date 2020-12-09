.PHONY: morning today clean

all: morning today

TODAY_PATH = $$(date +year%Y/day%d)

morning:
	@mkdir -p $(TODAY_PATH)
	@cp -n ./utils/main.go.template $(TODAY_PATH)/main.go

today:
	@go run $(TODAY_PATH)/*.go

clean:
	@rm ./*/*/input.txt

.PHONY: morning today clean

morning:
	mkdir -p $$(date +year%Y/day%d)
	cp -n ./utils/main.go.template $$(date +year%Y/day%d)/main.go

today:
	go run $$(date +year%Y/day%d)/*.go

clean:
	rm ./*/*/input.txt

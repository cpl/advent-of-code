

COOKIE := $(shell cat .cookie)

YEAR := $(shell date +%Y)
DAY := $(shell date +%-d)
DAY_F := $(shell date +%d)

puzzle:
	@mkdir -p "data/years/$(YEAR)/"
	@curl -s https://adventofcode.com/$(YEAR)/day/$(DAY)/input -b .cookie -H "Cookie: session=$(COOKIE)" > pkg/years/$(YEAR)/day$(DAY_F).txt

solve:
	go test -v -run TestSolveDay$(DAY_F) ./pkg/years/$(YEAR)



AOC_COOKIE ?= $(shell cat .cookie)

YEAR ?= $(shell date +%Y)
DAY ?= $(shell date +%-d)
DAY_F ?= $(shell date +%d)

.PHONY: puzzle solve test

puzzle:
	@mkdir -p "pkg/years/$(YEAR)/data"
	@curl -s https://adventofcode.com/$(YEAR)/day/$(DAY)/input -b .cookie -H "Cookie: session=$(AOC_COOKIE)" > pkg/years/$(YEAR)/data/$(YEAR)_$(DAY_F).txt

solve:
	go test -v -run TestSolveDay$(DAY_F) ./pkg/years/$(YEAR)

test:
	AOC_COOKIE=$(AOC_COOKIE) go test -v -run Test$(TEST) ./pkg/years/$(YEAR)

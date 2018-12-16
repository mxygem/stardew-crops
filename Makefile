.PHONY: unit
unit:
	TEST=true
	if [ ! -d coverage ]; then mkdir coverage; fi	
	goverage -v -race -coverprofile=coverage/stardew-crops.profile -timeout 30s ./...

.PHONY: godog
godog:
	TEST=true
	godog --strict

.PHONY: coverage_results
coverage_results: 
	go tool cover -html=coverage/stardew-crops.profile

.PHONY: all
all: unit godog

.PHONY: dep
dep:
	dep ensure -v

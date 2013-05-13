Name := defer

clean:
	@go clean
	@rm -f go.out temp.js out.js
	@rm -f *~
	@rm -f go-result.txt js-result.txt

build: clean
	@go build
	@gotojs ./ > temp.js
	@js-beautify temp.js > out.js

test:
	@echo $(Name)
	@./$(Name) > go-result.txt
	@node out.js > js-result.txt
	@diff go-result.txt js-result.txt

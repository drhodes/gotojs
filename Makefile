.PHONY: all test clean

build: 
	go build

test:
	python test.py

clean:
	bash zap.sh	

check:
	bash ./check-deps.sh

commit: clean
	git add .
	git commit -a

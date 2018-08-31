help:
	@echo
	@echo "Parameters"
	@echo "- filepath: Used to specify the input file path to evalulate."
	@echo
	@echo "Tasks:"
	@echo "- run: Runs the evaluation script. Takes the \"filepath\" parameter."
	@echo "- shell": Runs the bash shell at the current directory. Useful for debugging.

build-dev:
	docker build -t solver-test/dev .

run: build-dev
	test -n "$(filepath)" # must specify filepath
	docker run -i --rm \
		-v $$(pwd):/go/src/github.com/gregmankes/solver \
		solver-test/dev \
		go run main.go $(filepath)

shell: build-dev
	docker run -it --rm \
		-v $$(pwd):/go/src/github.com/gregmankes/solver \
		solver-test/dev \
		bash

test: build-dev
	docker run -it --rm \
		-v $$(pwd):/go/src/github.com/gregmankes/solver \
		solver-test/dev \
		go test $$(go list ./... | grep -v vendor)

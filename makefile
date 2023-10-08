unit_test:
	go test -v ./reader

e2e_test:
	go test -v ./e2e

test_all:
	go test -v -cover ./...

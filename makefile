all:
	go run cmd/alu/main.go
	#cd cmd/alu/ && go run main.go

test:
	go test alu_test.go

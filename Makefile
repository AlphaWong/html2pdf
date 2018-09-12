cover: 
	go test -coverprofile=coverage.out -covermode=atomic ./...

coverage:
	go test -tags=coverage -coverpkg=./handlers/... ./... -coverprofile coverage.out && go tool cover -func=coverage.out

coverhtml: cover
	go tool cover -html=coverage.out

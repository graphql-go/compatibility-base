go clean -testcache && go test -cover -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html; sed -i'' -e's/black/whitesmoke/g' coverage.html && open coverage.html

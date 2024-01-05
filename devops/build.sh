#!/bin/bash

# Clean up the build location
rm -rf ./build && GOBIN=$(pwd)/build go install

# Download packages
go mod download

# Run the tests and generate coverage profile
go test -coverpkg=./... -coverprofile=./build/coverage.out -covermode=count ./...

# Generate test summary
go tool cover -func=./build/coverage.out > ./build/summary.log

# Generate html test report
go tool cover -html=./build/coverage.out -o ./build/coverage.html

# Use go tool cover to parse the coverage percentages
coverage=$(cat ./build/summary.log | grep total | awk '{print $3}' | sed -e 's/[%]//g')

# Define the minimum coverage threshold
threshold=85.00

# Compare coverage with the threshold
if (( $(echo "$coverage < $threshold" | bc -l) )); then
    echo "Coverage ($coverage%) is below the threshold ($threshold%)."
    exit 1
else
    echo "Coverage is above or equal to the threshold ($threshold%)."
fi
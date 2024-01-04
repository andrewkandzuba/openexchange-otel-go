#!/bin/bash

# Clean up the build location
rm -rf ./.out && GOBIN=$(pwd)/.out go install

# Run the tests and generate coverage profile
go test -coverprofile=./.out/coverage.out -covermode=atomic  ./... > ./.out/test.logs

# Generate test summary
go tool cover -func=./.out/coverage.out > ./.out/summary.log

# Generate html test report
go tool cover -html=./.out/coverage.out -o ./.out/coverage.html

# Use go tool cover to parse the coverage percentages
coverage=$(cat ./.out/summary.log | grep total | awk '{print $3}' | sed -e 's/[%]//g')

# Define the minimum coverage threshold
threshold=85.00

# Compare coverage with the threshold
if (( $(echo "$coverage < $threshold" | bc -l) )); then
    echo "Coverage is below the threshold ($threshold%)."
    exit 1
else
    echo "Coverage is above or equal to the threshold ($threshold%)."
fi

# Discover not covered code 
notest=$(cat ./.out/test.logs | grep -c 'no test files')

# Check for module with no tests coverage
if (( $(echo "$notest > 0" | bc -l) )); then
    echo "Modules with no tests are detected. Abort the build!!!"
    exit 1
else
    echo "No modules without tests. Success!"
fi

env GOOS=linux  GOARCH=386 go build 

#go test -coverprofile=./.out/coverage.out -covermode=atomic  ./...
#go tool cover -html=./.out/coverage.out -o ./.out/coverage.html

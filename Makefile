.PHONY: build clean

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/LambdaFramework/web LambdaFramework/web.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/LambdaFramework/kinesis LambdaFramework/kinesis.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/LambdaFramework/authoriser LambdaFramework/authoriser.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/Common/main Common/main.go

clean:
	rm -rf ./bin
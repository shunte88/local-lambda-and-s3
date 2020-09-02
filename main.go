package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.S3Event) {
	fmt.Println("S3 trigger Lambda")
}

func main() {
	lambda.Start(handler)
}

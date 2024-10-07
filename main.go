package main

import (
	"hello/router"
	"log"

    "context"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/awslabs/aws-lambda-go-api-proxy/gin"
    "github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init(){
    // stdout and stderr are sent to AWS CloudWatch Logs
    log.Printf("Gin cold start")
    r := gin.Default()

	// ルーティング設定を呼び出す
    router.RouteAPI(r)

    ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
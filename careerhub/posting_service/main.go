package main

import (
	"context"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/app"
)

func main() {
	app.Run(context.Background())
}

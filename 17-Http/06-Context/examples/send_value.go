package examples

import "context"

func SendValueExample() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "api-key", "1234567890")
	ctx = context.WithValue(ctx, "api-key", "1234567891")
	ctx = context.WithValue(ctx, "validation-code", "1020304050")

	Print(ctx)

}

func Print(ctx context.Context) {
	var apiKey = ctx.Value("api-key")
	var validationCode = ctx.Value("validation-code")

	println(apiKey.(string), validationCode.(string))
}

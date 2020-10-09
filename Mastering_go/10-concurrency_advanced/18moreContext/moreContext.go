package main

import (
	"context"
	"fmt"
)

type Key string

func searchKey(ctx context.Context, key Key) {
	val := ctx.Value(key)
	if val != nil {
		fmt.Println("Found value:", val)
		return
	} else {
		fmt.Println("Key not found:", key)
	}
}

func main() {
	key := Key("theKey")

	ctx := context.Background()
	ctx = context.WithValue(ctx, key, "theKey")

	searchKey(ctx, key)

	searchKey(ctx, "someKey")
	tmpCtx := context.TODO()
	searchKey(tmpCtx, "someKey")
}

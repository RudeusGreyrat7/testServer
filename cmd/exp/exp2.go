package main

import (
	stdctx "context"
	"fmt"

	"github.com/LENSLOCKED/context"
	"github.com/LENSLOCKED/models"
)

type ctxKey string

const (
	colorKey ctxKey = "yellow"
)

func main() {
	ctx := stdctx.Background()

	user := models.User{
		Email: "jon@snow.com",
	}

	ctx = context.WithUser(ctx, &user)
	ctxUser := context.User(ctx)
	fmt.Println(ctxUser.Email)

}

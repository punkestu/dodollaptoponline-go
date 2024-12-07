package main

import (
	"github.com/punkestu/dodollaptoponline-go/features/product"
	"github.com/punkestu/dodollaptoponline-go/features/user"
)

func main() {
	wait := make(chan struct{})

	go func() {
		user := user.Init()
		user.Listen(":3000")
		wait <- struct{}{}
	}()

	go func() {
		product := product.Init()
		product.Listen(":3001")
		wait <- struct{}{}
	}()

	<-wait
}

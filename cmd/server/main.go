package main

import (
	"github.com/maivankien/go-ecommerce-api/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run(":8002")
}

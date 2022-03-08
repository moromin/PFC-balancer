package main

import (
	_ "github.com/moromin/go-svelte/backend/config"
	"github.com/moromin/go-svelte/backend/infrastructure"
)

func main() {
	infrastructure.Serve()
}

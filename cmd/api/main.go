package main

import (
	"fmt"

	"github.com/JueViGrace/clo-backend/internal/server"
)

func main() {
	server := server.New()

	if err := server.Init(); err != nil {
		panic(fmt.Sprintf("cannot start server at: %s", err))
	}
}

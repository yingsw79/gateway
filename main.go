package main

import (
	"gateway/router"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	// h.Use(cors.New(cors.Config{
	// 	AllowAllOrigins:  true,
	// 	AllowCredentials: true,
	// 	AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
	// 	AllowHeaders:     []string{"*"},
	// 	MaxAge:           12 * time.Hour,
	// }))

	// h.Use(handler.Gateway)

	router.Register(h)
	h.Spin()
}

package fserver

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type FServer struct {
	FiberApp *fiber.App
}

func NewFServer() *FServer {
	return &FServer{FiberApp: fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Employee2",
	})}
}

func (fsrv *FServer) Run(addr string) error {
	if err := fsrv.FiberApp.Listen(addr); err != nil {
		log.Fatalln("Fiber: failed to listen: ", err)
	}

	return nil
}

func (fsrv *FServer) Shutdown() error {
	if err := fsrv.FiberApp.Shutdown(); err != nil {
		log.Fatalln("Fiber: failed to shutdown: ", err)
	}

	return nil
}

package main

import (
	"OctavianoRyan25/GinGo/routers"
)

func main() {
	PORT := ":8080"
	routers.StartServer().Run(PORT)
}
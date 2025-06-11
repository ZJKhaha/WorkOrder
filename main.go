package main

import "JeeTron/Routers"

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}

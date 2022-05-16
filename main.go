package main

import (
	"github.com/eight-corner/learngo/routes"
	"github.com/gin-gonic/gin"
	"os"
)

/*func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("Done...")
	length = len(name)
	uppercase = strings.ToUpper(name)
	fmt.Println("함수")
	return
}*/

func main() {
	//	totalLength, up := lenAndUpper("corner")
	//	fmt.Println("----------")
	//	fmt.Println(totalLength, up)
	//
	server := gin.Default()

	routes.SetupRouter(server)

	listenAddress := ":8080"
	if envListenAddress, exists := os.LookupEnv("LISTEN_PORT"); exists {
		listenAddress = envListenAddress
	}

	// server run
	if err := server.Run(listenAddress); err != nil {
		panic(err)
	}
}

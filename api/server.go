package api

import (
	"fmt"
	"log"
	"net/http"

	httplogger "github.com/jesseokeya/go-httplogger"
	"github.com/obasajujoshua31/blogos/api/router"
	"github.com/obasajujoshua31/blogos/auto"
	"github.com/obasajujoshua31/blogos/config"
)

func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("Listening at Port :%d\n", config.PORT)
	Listen(config.PORT)

}

func Listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), httplogger.Golog(r)))
}

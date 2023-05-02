package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/renatormc/pfila/api/config"
	"github.com/renatormc/pfila/api/database"
	"github.com/renatormc/pfila/api/mods/procmod"
	"github.com/renatormc/pfila/api/processes"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cf := config.GetConfig()
	database.Migrate()
	r := gin.Default()

	api := r.Group("/api")
	procmod.ConfigRoutes(api)

	go func() {
		for {
			fmt.Println("checking processes")
			if err := processes.CheckProcesses(); err != nil {
				log.Fatal(err)
			}
			time.Sleep(30 * time.Second)
		}

	}()
	log.Fatal(r.Run(fmt.Sprintf(":%s", cf.Port)))
}

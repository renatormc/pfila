package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/akamensky/argparse"
	"github.com/gin-gonic/gin"
	"github.com/renatormc/pfila/api/config"
	"github.com/renatormc/pfila/api/database"
	"github.com/renatormc/pfila/api/mods/procmod"
	"github.com/renatormc/pfila/api/processes"
)

func Test() {
	cmd := exec.Command("D:\\tests\\pfila\\iped\\iped-4.1.2\\jre\\bin\\java.exe", "-jar",
		"D:\\tests\\pfila\\iped\\iped-4.1.2\\iped.jar", "-profile", "fastmode",
		"-d", "D:\\tests\\pfila\\pen.E01", "-o", "D:\\tests\\pfila\\result", "--nogui")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

func Serve() {
	cf := config.GetConfig()
	database.Migrate()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.File(filepath.Join(cf.AppDir, "app", "index.html"))
	})

	r.Static("/static", filepath.Join(cf.AppDir, "app"))

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

func main() {
	args := os.Args
	if len(args) == 1 {
		args = append(args, "serve")
	}
	parser := argparse.NewParser("PFila", "PFila")
	serveCmd := parser.NewCommand("serve", "Start server")
	testCmd := parser.NewCommand("test", "Run test function")

	err := parser.Parse(args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	switch {
	case testCmd.Happened():
		Test()
	case serveCmd.Happened():
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Press any key to exit.")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
			}
		}()
		Serve()
	}
}

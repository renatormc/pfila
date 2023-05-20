package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// cf := config.GetConfig()
	// database.Migrate()
	// r := gin.Default()

	// api := r.Group("/api")
	// procmod.ConfigRoutes(api)

	// go func() {
	// 	for {
	// 		fmt.Println("checking processes")
	// 		if err := processes.CheckProcesses(); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		time.Sleep(30 * time.Second)
	// 	}

	// }()
	// log.Fatal(r.Run(fmt.Sprintf(":%s", cf.Port)))
	cmd := exec.Command("D:\\tests\\pfila\\iped\\iped-4.1.2\\jre\\bin\\java.exe", "-jar",
		"D:\\tests\\pfila\\iped\\iped-4.1.2\\iped.jar", "-profile", "fastmode",
		"-d", "D:\\tests\\pfila\\pen.E01", "-o", "D:\\tests\\pfila\\result", "--nogui")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

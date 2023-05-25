package procmod

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/renatormc/pfila/api/config"
	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/database/repo"
	"github.com/renatormc/pfila/api/external"
	"github.com/renatormc/pfila/api/helpers"
	"github.com/renatormc/pfila/api/processes"
)

func ConfigRoutes(group *gin.RouterGroup) {
	group.GET("/proc", ListProcs)
	group.POST("/proc", CreateProc)
	group.PUT("/proc/:id", UpdateProc)
	group.PUT("/queue-proc/:id", QueueProc)
	group.PUT("/stop-proc/:id", StopProc)
	group.GET("/proc-console/:id", GetProcConsole)
	group.DELETE("/proc/:id", DeleteProc)
	group.GET("/disks", GetDisks)
	group.GET("/iped-profiles", IpedProfiles)
}

func ListProcs(c *gin.Context) {
	procs := repo.GetAllProcesses()
	c.JSON(http.StatusOK, SerializeManyProc(procs))
}

func CreateProc(c *gin.Context) {
	m := models.Process{}
	sl := ProcSchemaLoad{}
	if !helpers.LoadFromBody[models.Process](c, &sl, &m) {
		return
	}

	m.CreatedAt = time.Now().Local()
	m.RandomID = uuid.NewString()
	m.Finish = time.Time{}
	m.Status = "ADICIONADO"
	if !helpers.SaveModel(c, &m) {
		return
	}
	c.JSON(http.StatusOK, SerializeProc(&m))
}

func UpdateProc(c *gin.Context) {
	m, ok := GetProcFromRequest(c)
	if !ok {
		return
	}
	sl := ProcSchemaLoad{}
	if !helpers.LoadFromBody[models.Process](c, &sl, m) {
		return
	}
	if !helpers.SaveModel(c, &m) {
		return
	}
	c.JSON(http.StatusOK, SerializeProc(m))
}

func DeleteProc(c *gin.Context) {
	m, ok := GetProcFromRequest(c)
	if !ok {
		return
	}
	if !helpers.DeleteModel(c, m) {
		return
	}

	c.JSON(http.StatusOK, SerializeProc(m))
}

func QueueProc(c *gin.Context) {

	m, ok := GetProcFromRequest(c)
	if !ok {
		return
	}

	m.StartWaiting = time.Now().Local()
	m.Finish = time.Time{}
	m.Status = "AGUARDANDO"
	if err := repo.SaveProc(m); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "interval server error"})
		return
	}
	if err := processes.CheckProcesses(); err != nil {
		log.Println(err)
	}
	m, err := repo.GetProcessById(int64(m.ID))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "interval server error"})
		return
	}
	c.JSON(http.StatusOK, SerializeProc(m))
}

func StopProc(c *gin.Context) {
	m, ok := GetProcFromRequest(c)
	if !ok {
		return
	}
	if err := processes.StopProcess(m); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "interval server error"})
		return
	}
	m.Status = "CANCELADO"
	m.Finish = time.Now().Local()
	if err := repo.SaveProc(m); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "interval server error"})
		return
	}
	c.JSON(http.StatusOK, SerializeProc(m))
}

func GetProcConsole(c *gin.Context) {
	m, ok := GetProcFromRequest(c)
	if !ok {
		return
	}
	text := strings.ReplaceAll(processes.GetProcConsole(m, 10), "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	c.JSON(http.StatusOK, gin.H{"console": text})
}

func GetDisks(c *gin.Context) {
	disks, err := external.GetDisks()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "interval server error"})
		return
	}
	c.JSON(http.StatusOK, disks)
}

func IpedProfiles(c *gin.Context) {
	cf := config.GetConfig()
	log.Println(cf.IpedProfileFolder)
	entries, err := os.ReadDir(cf.IpedProfileFolder)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "interval server error"})
		return
	}
	profiles := []string{}
	for _, entry := range entries {
		if entry.IsDir() {
			profiles = append(profiles, entry.Name())
		}
	}
	c.JSON(http.StatusOK, profiles)
}

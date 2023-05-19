package procmod

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/database/repo"
	"github.com/renatormc/pfila/api/helpers"
	"github.com/renatormc/pfila/api/processes"
	"github.com/renatormc/pfila/api/processes/ftkimager"
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
	// html := strings.ReplaceAll(processes.GetProcConsole(m, 20), "\n", "<br>")
	// c.JSON(http.StatusOK, gin.H{"console": html})
	// html := strings.ReplaceAll(processes.GetProcConsole(m, 20), "\n", "<br>")
	log.Print(processes.GetProcConsole(m, 20))
	c.JSON(http.StatusOK, gin.H{"console": processes.GetProcConsole(m, 20)})
}

func GetDisks(c *gin.Context) {
	disks, err := ftkimager.GetDisks()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "interval server error"})
		return
	}
	c.JSON(http.StatusOK, disks)
}

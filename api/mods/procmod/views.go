package procmod

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/database/repo"
	"github.com/renatormc/pfila/api/helpers"
	"github.com/renatormc/pfila/api/processes"
)

func ConfigRoutes(group *gin.RouterGroup) {
	group.GET("/proc", ListProcs)
	group.POST("/proc", CreateProc)
	group.PUT("/proc/:id", UpdateProc)
	group.PUT("/queue-proc/:id", QueueProc)
	group.PUT("/stop-proc/:id", QueueProc)
	group.GET("/proc-console/:id", GetProcConsole)
	group.DELETE("/proc/:id", DeleteProc)
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
	m.CreatedAt = time.Now()
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
	if !helpers.DeleteModel(c, &m) {
		return
	}

	c.JSON(http.StatusOK, SerializeProc(m))
}

func QueueProc(c *gin.Context) {
	m, ok := GetProcFromRequest(c)
	if !ok {
		return
	}
	m.StartWaiting = time.Now()
	if err := repo.SaveProc(m); err != nil {
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
	c.JSON(http.StatusOK, gin.H{"console": processes.GetProcConsole(m, 20)})
}

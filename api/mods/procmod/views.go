package procmod

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reantormc/pfila/api/database/models"
	"github.com/reantormc/pfila/api/database/repo"
	"github.com/reantormc/pfila/api/helpers"
)

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
	if !helpers.SaveModel(c, &m) {
		return
	}
	c.JSON(http.StatusOK, SerializeProc(&m))
}

func UpdateProc(c *gin.Context) {
	id, ok := helpers.GetIDFromParam(c)
	if !ok {
		return
	}
	m, err := repo.GetProcessById(id)
	if helpers.CheckDBError(c, err) {
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
	id, ok := helpers.GetIDFromParam(c)
	if !ok {
		return
	}
	m, err := repo.GetProcessById(id)
	if helpers.CheckDBError(c, err) {
		return
	}
	if !helpers.DeleteModel(c, &m) {
		return
	}

	c.JSON(http.StatusOK, SerializeProc(m))
}

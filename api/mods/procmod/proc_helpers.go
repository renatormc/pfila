package procmod

import (
	"github.com/gin-gonic/gin"
	"github.com/renatormc/pfila/api/database/models"
	"github.com/renatormc/pfila/api/database/repo"
	"github.com/renatormc/pfila/api/helpers"
)

func GetProcFromRequest(c *gin.Context) (*models.Process, bool) {
	id, ok := helpers.GetIDFromParam(c)
	if !ok {
		return nil, false
	}
	m, err := repo.GetProcessById(id)
	if helpers.CheckDBError(c, err) {
		return nil, false
	}

	return m, true
}

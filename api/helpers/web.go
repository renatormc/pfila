package helpers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/renatormc/pfila/api/database"
	"gorm.io/gorm"
)

func GetIDFromParam(c *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return 0, false
	}
	return id, true
}

func CheckDBError(c *gin.Context, err error) bool {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			return true
		}
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		return true
	}
	return false
}

func LoadFromBody[M any](c *gin.Context, schema LoadSchema[M], model *M) bool {
	if err := c.ShouldBindJSON(schema); err != nil {
		ve := NewValidationError()
		ve.ParseError(err, schema)
		c.JSON(http.StatusUnprocessableEntity, ve.Messages)
		return false
	}
	if ve := schema.Fill(model); ve != nil {
		c.JSON(http.StatusUnprocessableEntity, ve.Messages)
		return false
	}
	return true
}

func SaveModel(c *gin.Context, model interface{}) bool {
	db := database.GetDatabase()
	if err := db.Save(model).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		return false
	}
	return true
}

func DeleteModel(c *gin.Context, model interface{}) bool {
	db := database.GetDatabase()
	if err := db.Delete(model).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		return false
	}
	return true
}

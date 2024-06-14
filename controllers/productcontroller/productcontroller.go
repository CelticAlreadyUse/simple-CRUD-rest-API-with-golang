package productcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CelticAlreadyUse/Social-Media-Project/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context){
	var products []models.Product
	models.DB.Find(&products)
	c.JSON(http.StatusOK,gin.H{"products":products})
}
func Show(c *gin.Context){
	var product models.Product
	id := c.Param("id")

	if err :=models.DB.First(&product,id).Error; err != nil{
		switch err{
		case gorm.ErrRecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound,gin.H{"message": "Data was not found"})
				return 
				default:
				c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"message": err.Error()})
				return 
			
		}
	}
	c.JSON(http.StatusOK,gin.H{"product":product})
}
func Create(c *gin.Context){
	var product models.Product

	if err :=c.ShouldBindJSON(&product); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message": err.Error()})
		return 
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK,gin.H{"product":product})


}
func Update(c *gin.Context){
	var product models.Product
	id := c.Param("id")

	if err :=c.ShouldBindJSON(&product); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message": err.Error()})
		return 
	}

	if models.DB.Model(&product).Where("id = ?",id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message": "Product can't be updated,there something wrong"})
		return 
	}

	c.JSON(http.StatusOK,gin.H{"message":"Data sucessfully updated"})
}
func Delete(c *gin.Context){
	var product models.Product
	
	var input struct{
		Id json.Number
	}

	if err :=c.ShouldBindJSON(&input); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message": err.Error()})
		return 
	}
	id,_ := input.Id.Int64()
	if models.DB.Delete(&product,id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message":"Something went wrong when deleting data"})
		return 
	}

	c.JSON(http.StatusOK,gin.H{"message":"Data sucessfully Deleted"})

}
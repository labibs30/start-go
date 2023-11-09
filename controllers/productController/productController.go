package productController 

import (
	"net/http"
	"github.com/tentangkode/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context){
	
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func Show(c *gin.Context){
	// Get data by id
	var product models.Product
	id:=c.Param("id")

	if err:= models.DB.First(&product,id).Error; err != nil {
		switch err {
			case gorm.ErrRecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message":"Data tidak ditemukan"})
				return 
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
				return 
		}
	}
	c.JSON(http.StatusOK, gin.H{"data":product})

}

func Create(c *gin.Context){
	var product models.Product

	if err:= c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":err.Error()})
		return
	}
	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data":product})
}

func Update(c *gin.Context){
	var product models.Product
	id:=c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":err.Error()})
	}
	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":"Data tidak ditemukan"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message":"Data berhasil diupdate"})
	
}

func Delete(c *gin.Context){
	
}
package seeders

import (
	"github.com/gin-gonic/gin"
)

func RunSeeders(c *gin.Context) {
	S_2021_01_31_01_Level()
	S_2021_02_01_01_Menu()

}

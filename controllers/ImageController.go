package Controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ImageController struct{}

func GetImage(c *gin.Context) {
	parts := fmt.Sprint("storage/", c.Param("folder"), "/", c.Param("id"))
	// filename := parts
	// file, err := os.Open(filename)

	// // helper.Log(typefile)
	// if err != nil {
	// 	log.Println("Error opening file", err)
	// 	c.JSON(http.StatusNotFound, gin.H{"status": "Error opening file " + filename})
	// 	return
	// }
	// defer file.Close()

	// print(traits.GetTypeImageFile(file))
	// switch traits.GetTypeImageFile(file) {
	// case "image/jpeg", "image/jpg":
	// 	img, err := jpeg.Decode(file)
	// case "image/gif":
	// 	img, err := gif.Decode(file)

	// case "image/png":
	// 	img, err := png.Decode(file)
	// default:
	// 	img, err := jpeg.Decode(file)
	// }

	// var newImage *image.NRGBA
	// png.Encode(file, newImage)
	c.File(parts)

	// if err != nil {
	// 	log.Println("Error decoding file", err)
	// 	c.JSON(http.StatusNotFound, gin.H{"status": "Error decoding file " + filename})
	// 	return
	// }

	// fmt.Println(img)
	// m := imaging.Fit(img, width, height, imaging.Lanczos)

}

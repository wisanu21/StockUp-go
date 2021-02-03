package seeders

import (
	"stockup-go/config"
	"stockup-go/helper"
	"stockup-go/models"
)

func S_2021_01_31_01_Level() {
	name_file := "S_2021_01_31_01_Level"
	println(name_file, " ---->seeders loading")
	arrlevels := []string{"เจ้าของร้าน", "หัวหน้าลูกจ้าง", "ลูกจ้าง"}
	for id := 1; id <= len(arrlevels); id++ {
		var levels models.Level
		levels.Id = id
		levels.Name = arrlevels[id-1]
		var err = config.GetDB().Save(&levels)
		if err.Error != nil {
			helper.Log(err.Error.Error())
		} else {
			println(name_file, " : ", arrlevels[id-1], " ----> seeders No Error")
		}
	}
	println(name_file, " ----> seeders successfully ")
}

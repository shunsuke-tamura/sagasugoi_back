package model

import (
	"fmt"
)

type Carp struct {
	GormModel
	// CreateCarpReq
	Word    string  `json:"word" binding:"required"`
	Comment string  `json:"comment" binding:"required"`
	Url     *string `json:"url"`
	Looks
}

func ReadAllCarps() []Carp {
	carps := []Carp{}
	result := db.Find(&carps)
	if result.Error != nil {
		panic(result.Error)
	}
	return carps
}

func CreateCarp(carp Carp) bool {
	fmt.Printf("%+v\n", carp)
	result := db.Create(&carp)
	if result.Error != nil {
		panic(result.Error)
	}
	return true
}

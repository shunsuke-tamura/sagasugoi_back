package model

type Carp struct {
	GormModel
	// CreateCarpReq
	Word      string  `json:"word" binding:"required"`
	Comment   string  `json:"comment" binding:"required"`
	Url       *string `json:"url"`
	ImageName *string `json:"iamgeName"`
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

func CreateCarp(carp Carp) Carp {
	result := db.Create(&carp)
	if result.Error != nil {
		panic(result.Error)
	}
	return carp
}

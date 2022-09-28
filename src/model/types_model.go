package model

type Types struct {
	// gorm.Model
	Id   string  `json:"id"`
	Size float64 `json:"size"`
	R    int     `json:"r"`
	G    int     `json:"g"`
	B    int     `json:"b"`
}

func GetAll() (types []Types) {
	result := db.Find(&types)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}
func GetOne(id string) (data Types) {
	result := db.First(&data, id)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}
func (b *Types) Create() {
	result := db.Create(b)
	if result.Error != nil {
		panic(result.Error)
	}
}
func (b *Types) Update() {
	result := db.Save(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *Types) Delete() {
	result := db.Delete(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

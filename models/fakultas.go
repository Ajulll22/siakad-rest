package models

type Fakultas struct {
	Id_fakultas   uint        `json:"id_fakultas" gorm:"primaryKey"`
	Nama_fakultas string      `json:"nama_fakultas" gorm:"type:varchar(100)"`
	Jurusan       []Jurusan   `gorm:"foreignKey:Id_fakultas;references:Id_fakultas;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Mahasiswa     []Mahasiswa `gorm:"foreignKey:Id_fakultas;references:Id_fakultas;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

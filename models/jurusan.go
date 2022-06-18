package models

type Jurusan struct {
	Id_jurusan   uint        `json:"id_jurusan" gorm:"primaryKey"`
	Nama_jurusan string      `json:"nama_jurusan" gorm:"type:varchar(100)"`
	Id_fakultas  uint        `json:"id_fakultas"`
	Mahasiswa    []Mahasiswa `gorm:"foreignKey:Id_jurusan;references:Id_jurusan;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

package models

import "gorm.io/gorm"

type KRS struct {
    IdKrs        uint   `gorm:"primaryKey" json:"id_krs"`
    Nim          string `json:"nim"`
    KodeMatakuliah string `json:"kode matakuliah"`
    Matakuliah   string `json:"matakuliah"`
    Semester     string `json:"semester"`
    TahunAkademik string `json:"tahunakademik"`
}

func Migrate(db *gorm.DB) {
    db.AutoMigrate(&KRS{})
}

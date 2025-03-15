package models

import "gorm.io/gorm"

type KRS struct {
    IdKrs          uint   `gorm:"primaryKey" json:"id_krs"`
    Nim            string `json:"nim"`
    KodeMatakuliah string `gorm:"column:kode matakuliah" json:"kode_matakuliah"`
    Matakuliah     string `json:"matakuliah"`
    Semester       string `json:"semester"`
    TahunAkademik  string `gorm:"column:tahunakademik" json:"tahun_akademik"`
}


func Migrate(db *gorm.DB) {
    db.AutoMigrate(&KRS{})
}

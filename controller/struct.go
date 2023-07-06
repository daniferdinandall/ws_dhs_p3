package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ================================================================================================
// DHS
type Dhs struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Mahasiswa  Mahasiswa          `bson:"mahasiswa,omitempty" json:"mahasiswa,omitempty"`
	MataKuliah []NilaiMataKuliah  `bson:"mata_kuliah,omitempty" json:"mata_kuliah,omitempty"`
	CreatedAt  primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Npm          int                `bson:"npm,omitempty" json:"npm,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Fakultas     Fakultas           `bson:"fakultas,omitempty" json:"fakultas,omitempty"`
	ProgramStudi ProgramStudi       `bson:"program_studi,omitempty" json:"program_studi,omitempty"`
	DosenWali    Dosen              `bson:"dosen_wali,omitempty" json:"dosen_wali,omitempty"`
}

type MataKuliah struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	KodeMatkul string             `bson:"kode_matkul,omitempty" json:"kode_matkul,omitempty"`
	Nama       string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Dosen      Dosen              `bson:"dosen,omitempty" json:"dosen,omitempty"`
	Sks        int                `bson:"sks,omitempty" json:"sks,omitempty"`
}

type NilaiMataKuliah struct {
	KodeMatkul string `bson:"kode_matkul,omitempty" json:"kode_matkul,omitempty"`
	Nama       string `bson:"nama,omitempty" json:"nama,omitempty"`
	Dosen      Dosen  `bson:"dosen,omitempty" json:"dosen,omitempty"`
	Sks        int    `bson:"sks,omitempty" json:"sks,omitempty"`
	Nilai      string `bson:"nilai,omitempty" json:"nilai,omitempty"`
}

type Dosen struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	KodeDosen   string             `bson:"kode_dosen,omitempty" json:"kode_dosen,omitempty"`
	Nama        string             `bson:"nama,omitempty" json:"nama,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
}

type Fakultas struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	KodeFakultas string             `bson:"kode_fakultas,omitempty" json:"kode_fakultas,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
}

type ProgramStudi struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	KodeProgramStudi string             `bson:"kode_program_studi,omitempty" json:"kode_program_studi,omitempty"`
	Nama             string             `bson:"nama,omitempty" json:"nama,omitempty"`
}

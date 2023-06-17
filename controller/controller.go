package controller

import (
	// "github.com/aiteung/musik"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	module "github.com/daniferdinandall/be_dhs_p3/module"
	"github.com/daniferdinandall/ws-dhs-p3/config"
)

func Homepage2(c *fiber.Ctx) error {
	// ipaddr := musik.GetIPaddress()
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})
}

// DHS
func GetAllDHS(c *fiber.Ctx) error {
	dhs := module.GetDhsAll(config.Ulbimongoconn)
	return c.JSON(dhs)
}

func GetDHSByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := module.GetDhsFromID(config.Ulbimongoconn, objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

func GetDHSByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	if npm == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	npm2, err := strconv.Atoi(npm)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid npm parameter",
		})
	}

	ps, err := module.GetDhsFromNPM(config.Ulbimongoconn, npm2)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for npm %s", npm),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for npm %s", npm),
		})
	}
	return c.JSON(ps)
}

func CreateDHS(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "Data created",
	})
}
func UpdateDHS(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})
}

func DeleteDHS(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})
}

// MAHASISWA
func GetAllMahasiswa(c *fiber.Ctx) error {
	mhs := module.GetMhsAll(config.Ulbimongoconn)
	return c.JSON(mhs)
}
func GetMahasiswaByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := module.GetMhsFromID(config.Ulbimongoconn, objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}
func GetMahasiswaByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	if npm == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	npm2, err := strconv.Atoi(npm)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid npm parameter",
		})
	}

	ps, err := module.GetMhsFromNPM(config.Ulbimongoconn, npm2)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for npm %s", npm),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for npm %s", npm),
		})
	}
	return c.JSON(ps)
}

func CreateMahasiswa(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})
}

func UpdateMahasiswa(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})
}

func DeleteMahasiswa(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})
}

// DOSEN
func GetAllDosen(c *fiber.Ctx) error {
	dosen := module.GetDosenAll(config.Ulbimongoconn)
	return c.JSON(dosen)
}
func GetDosenByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := module.GetDosenFromID(config.Ulbimongoconn, objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

func GetDosenByKodeDosen(c *fiber.Ctx) error {
	kodeDosen := c.Params("kodeDosen")
	if kodeDosen == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	ps, err := module.GetDosenFromKodeDosen(config.Ulbimongoconn, kodeDosen)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for kodeDosen %s", kodeDosen),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for kodeDosen %s", kodeDosen),
		})
	}
	return c.JSON(ps)
}

func CreateDosen(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})
}

func UpdateDosen(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})
}

func DeleteDosen(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})
}

// MATA KULIAH
func GetAllMataKuliah(c *fiber.Ctx) error {
	matakuliah := module.GetMatkulAll(config.Ulbimongoconn)
	return c.JSON(matakuliah)
}

func GetMataKuliahByID(c *fiber.Ctx) error {
	matkul := c.Params("id")
	if matkul == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(matkul)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := module.GetMatkulFromID(config.Ulbimongoconn, objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", matkul),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", matkul),
		})
	}
	return c.JSON(ps)
}

func GetMataKuliahByKodeMataKuliah(c *fiber.Ctx) error {
	kodeMataKuliah := c.Params("kodeMataKuliah")
	if kodeMataKuliah == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	ps, err := module.GetMatkulFromKodeMatkul(config.Ulbimongoconn, kodeMataKuliah)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for kodeMataKuliah %s", kodeMataKuliah),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for kodeMataKuliah %s", kodeMataKuliah),
		})
	}
	return c.JSON(ps)
}

func CreateMataKuliah(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})

}

func UpdateMataKuliah(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})

}

func DeleteMataKuliah(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})

}

// PROGRAM STUDI
func GetAllProgramStudi(c *fiber.Ctx) error {
	programstudi := module.GetProdiAll(config.Ulbimongoconn)
	return c.JSON(programstudi)
}
func GetProgramStudiByID(c *fiber.Ctx) error {
	prodi := c.Params("id")
	if prodi == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(prodi)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := module.GetProdiFromID(config.Ulbimongoconn, objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", prodi),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", prodi),
		})
	}
	return c.JSON(ps)
}

func GetProgramStudiByKodeProgramStudi(c *fiber.Ctx) error {
	kodeProgramStudi := c.Params("kodeProgramStudi")
	if kodeProgramStudi == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	ps, err := module.GetProdiFromKodeProdi(config.Ulbimongoconn, kodeProgramStudi)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for kodeProgramStudi %s", kodeProgramStudi),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for kodeProgramStudi %s", kodeProgramStudi),
		})
	}
	return c.JSON(ps)
}
func CreateProgramStudi(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})
}

func UpdateProgramStudi(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})

}

func DeleteProgramStudi(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})

}

// FAKULTAS

func GetAllFakultas(c *fiber.Ctx) error {
	fakultas := module.GetFakultasAll(config.Ulbimongoconn)
	return c.JSON(fakultas)
}

func GetFakultasByID(c *fiber.Ctx) error {
	fakultas := c.Params("id")
	if fakultas == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(fakultas)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := module.GetFakultasFromID(config.Ulbimongoconn, objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", fakultas),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", fakultas),
		})
	}
	return c.JSON(ps)
}
func GetFakultasByKodeFakultas(c *fiber.Ctx) error {
	kodeFakultas := c.Params("kodeFakultas")
	if kodeFakultas == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	ps, err := module.GetFakultasFromKodeFakultas(config.Ulbimongoconn, kodeFakultas)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for kodeFakultas %s", kodeFakultas),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for kodeFakultas %s", kodeFakultas),
		})
	}
	return c.JSON(ps)
}

func CreateFakultas(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})

}

func UpdateFakultas(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})

}

func DeleteFakultas(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})

}

package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/daniferdinandall/ws-dhs-p3/config"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	model "github.com/daniferdinandall/be_dhs2/model"
	module "github.com/daniferdinandall/be_dhs2/module"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

// GetAllPresensi godoc
// @Summary Get All Data Presensi.
// @Description Mengambil semua data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Success 200 {object} Presensi
// @Router /presensi [get]
func GetAllPresensi(c *fiber.Ctx) error {
	ps := module.GetAllPresensi(config.Ulbimongoconn, "presensi")
	return c.JSON(ps)
}

// GetPresensiID godoc
// @Summary Get By ID Data Presensi.
// @Description Ambil per ID data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /presensi/{id} [get]
func GetPresensiID(c *fiber.Ctx) error {
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
	ps, err := module.GetPresensiFromID(objID, config.Ulbimongoconn, "presensi")
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
func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(ps)
}

// InsertData godoc
// @Summary Insert data presensi.
// @Description Input data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /ins [post]
func InsertData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var presensi model.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := module.InsertPresensi(db, "presensi",
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// UpdateData godoc
// @Summary Update data presensi.
// @Description Ubah data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Karyawan
// @Failure 400
// @Failure 500
// @Router /upd/{id} [put]
func UpdateData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var presensi model.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = module.UpdatePresensi(db, "presensi",
		objectID,
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// DeletePresensiByID godoc
// @Summary Delete data presensi.
// @Description Hapus data presensi.
// @Tags Dani
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeletePresensiByID(c *fiber.Ctx) error {
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

	err = module.DeletePresensiByID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}

// dhs

func GetDhs(c *fiber.Ctx) error {
	queryValue := c.Query("npm")
	i, err := strconv.Atoi(queryValue)
	if err != nil {
		// ... handle error
		// panic(err)
		return c.SendString("format params kurang tepat")
	}
	ps := module.GetDhsFromNPM(int(i))

	return c.JSON(ps)
}

func GetAllDhs(c *fiber.Ctx) error {

	ps := module.GetDhsAll()
	return c.JSON(ps)
}

//mahasiswa

func GetMhs(c *fiber.Ctx) error {
	queryValue := c.Query("npm")
	i, err := strconv.Atoi(queryValue)
	if err != nil {
		// ... handle error
		// panic(err)
		return c.SendString("format params kurang tepat")
	}
	ps := module.GetMhsFromNPM(int(i))

	return c.JSON(ps)
}

func GetAllMhs(c *fiber.Ctx) error {

	ps := module.GetMhsAll()
	return c.JSON(ps)
}

// Dosen

func GetDosen(c *fiber.Ctx) error {
	queryValue := c.Query("kode")

	ps := module.GetDosenFromKodeDosen(queryValue)

	return c.JSON(ps)
}

func GetAllDosen(c *fiber.Ctx) error {

	ps := module.GetDosenAll()
	return c.JSON(ps)
}

func GetMatkul(c *fiber.Ctx) error {
	queryValue := c.Query("kode")

	ps := module.GetMatkulFromKodeMatkul(queryValue)

	return c.JSON(ps)
}

func GetAllMatkul(c *fiber.Ctx) error {

	ps := module.GetMatkulAll()
	return c.JSON(ps)
}

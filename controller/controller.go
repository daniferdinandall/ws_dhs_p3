package controller

import (
	// "github.com/aiteung/musik"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	model "github.com/daniferdinandall/be_dhs_p3/model"
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

// Auth
var jwtSecret = []byte("secret-key")

func Login(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var data_login model.User
	if err := c.BodyParser(&data_login); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	user, message, err := module.ValidateUserFromEmail(db, data_login.Email, data_login.Password)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
	}
	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["nama"] = data_login.Nama
	claims["email"] = data_login.Email
	claims["role"] = data_login.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Set token expiration time to 24 hours

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     message,
		"nama":        user.Nama,
		"email":       user.Email,
		"role":        user.Role,
		"tokenString": tokenString,
	})
}

func ValidateToken(c *fiber.Ctx) error {
	var token model.Token
	if err := c.BodyParser(&token); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	tkn := token.Token_String
	// Check if token exists
	if tkn == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	initoken, err := jwt.Parse(tkn, func(initoken *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := initoken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return jwtSecret, nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	// Validate token claims
	if claims, ok := initoken.Claims.(jwt.MapClaims); ok && initoken.Valid {
		// Check if token has expired
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		if time.Now().After(expirationTime) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Token has expired",
			})
		}

		// c.Locals("username", claims["username"])
		// return c.Next()
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status": http.StatusOK,
			"email":  claims["email"],
			"role":   claims["role"],
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Invalid token",
	})
}


// GetAllDHS godoc
// @Summary Get All Data Dhs.
// @Description Mengambil semua data Dhs.
// @Tags Dhs
// @Accept json
// @Produce json
// @Success 200 {object} Dhs
// @Router /dhs [get]
func GetAllDHS(c *fiber.Ctx) error {
	dhs := module.GetDhsAll(config.Ulbimongoconn)
	return c.JSON(dhs)
}

// GetDHSByID godoc
// @Summary Get By ID Data Dhs.
// @Description Ambil per ID data Dhs.
// @Tags Dhs
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Dhs
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /dhs/{id} [get]
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

// GetDHSByNPM godoc
// @Summary Get By NPM Data Dhs.
// @Description Ambil per NPM data Dhs.
// @Tags Dhs
// @Accept json
// @Produce json
// @Param npm path string true "Masukan NPM"
// @Success 200 {object} Dhs
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /dhs/{npm} [get]
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

// CreateDHS godoc
// @Summary Create data Dhs.
// @Description Input data Dhs.
// @Tags Dhs
// @Accept json
// @Produce json
// @Param request body Dhs true "Payload Body [RAW]"
// @Success 200 {object} Dhs
// @Failure 400
// @Failure 500
// @Router /dhs [post]
func CreateDHS(c *fiber.Ctx) error {

	db := config.Ulbimongoconn
	var data_dhs model.Dhs
	if err := c.BodyParser(&data_dhs); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	insertedID, err := module.InsertDHS(db,
		data_dhs.Mahasiswa,
		data_dhs.MataKuliah,
	)
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

// UpdateDHS godoc
// @Summary Update data Dhs.
// @Description Ubah data Dhs.
// @Tags Dhs
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Dhs true "Payload Body [RAW]"
// @Success 200 {object} Dhs
// @Failure 400
// @Failure 500
// @Router /dhs/{id} [put]
func UpdateDHS(c *fiber.Ctx) error {
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

	// Parse the request body into a Dhs object
	var data_dhs model.Dhs
	if err := c.BodyParser(&data_dhs); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdateDhsById function with the parsed ID and the Presensi object
	err = module.UpdateDhsById(db,
		objectID,
		data_dhs.Mahasiswa,
		data_dhs.MataKuliah)
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

// DeleteDeleteDHS godoc
// @Summary Delete data Dhs.
// @Description Hapus data Dhs.
// @Tags Dhs
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /dhs/{id} [delete]
func DeleteDHS(c *fiber.Ctx) error {
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

	err = module.DeleteDhsByID(config.Ulbimongoconn, objID)
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

// GetAllMahasiswa godoc
// @Summary Get All Data Mahasiswa.
// @Description Mengambil semua data mahasiswa.
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Success 200 {object} Mahasiswa
// @Router /mahasiswa [get]
func GetAllMahasiswa(c *fiber.Ctx) error {
	mhs := module.GetMhsAll(config.Ulbimongoconn)
	return c.JSON(mhs)
}

// GetMahasiswaByID godoc
// @Summary Get By ID Data Mahasiswa.
// @Description Ambil per ID data Mahasiswa.
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Mahasiswa
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /mahasiswa/{id} [get]
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

// GetMahasiswaByNPM godoc
// @Summary Get By NPM Data Mahasiswa.
// @Description Ambil per NPM data Mahasiswa.
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param npm path string true "Masukan NPM"
// @Success 200 {object} Mahasiswa
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /mahasiswa/{npm} [get]
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

// CreateMahasiswa godoc
// @Summary Create data Mahasiswa.
// @Description Input data Mahasiswa.
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param request body Mahasiswa true "Payload Body [RAW]"
// @Success 200 {object} Mahasiswa
// @Failure 400
// @Failure 500
// @Router /mahasiswa [post]
func CreateMahasiswa(c *fiber.Ctx) error {

	db := config.Ulbimongoconn
	var data_mhs model.Mahasiswa
	if err := c.BodyParser(&data_mhs); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	insertedID, err := module.InsertMhs(db,
		data_mhs.Npm,
		data_mhs.Nama,
		data_mhs.Fakultas,
		data_mhs.DosenWali,
		data_mhs.ProgramStudi,
	)
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

// UpdateMahasiswa godoc
// @Summary Update data Mahasiswa.
// @Description Ubah data Mahasiswa.
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Mahasiswa true "Payload Body [RAW]"
// @Success 200 {object} Mahasiswa
// @Failure 400
// @Failure 500
// @Router /mahasiswa/{id} [put]
func UpdateMahasiswa(c *fiber.Ctx) error {
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

	// Parse the request body into a Dhs object
	var data_mhs model.Mahasiswa
	if err := c.BodyParser(&data_mhs); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdateDhsById function with the parsed ID and the Presensi object
	err = module.UpdateMhsById(db,
		objectID,
		data_mhs.Npm,
		data_mhs.Nama,
		data_mhs.Fakultas,
		data_mhs.DosenWali,
		data_mhs.ProgramStudi)
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


// DeleteDeleteMahasiswa godoc
// @Summary Delete data Mahasiswa.
// @Description Hapus data Mahasiswa.
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /mahasiswa/{id} [delete]
func DeleteMahasiswa(c *fiber.Ctx) error {
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

	// Call the DeleteDhsById function with the parsed ID
	err = module.DeleteMhsByID(db, objectID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully deleted",
	})
}

// GetAllDosen godoc
// @Summary Get All Data Dosen.
// @Description Mengambil semua data Dosen.
// @Tags Dosen
// @Accept json
// @Produce json
// @Success 200 {object} Dosen
// @Router /Dosen [get]
func GetAllDosen(c *fiber.Ctx) error {
	dosen := module.GetDosenAll(config.Ulbimongoconn)
	return c.JSON(dosen)
}

// GetDosenByID godoc
// @Summary Get By ID Data Dosen.
// @Description Ambil per ID data Dosen.
// @Tags Dosen
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Dosen
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /Dosen/{id} [get]
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

// GetDosenBykodeDosen godoc
// @Summary Get By Kode Dosen Data Dosen.
// @Description Ambil per Kode Dosen data dosen.
// @Tags Dosen
// @Accept json
// @Produce json
// @Param kode path string true "Masukan Kode Dosen"
// @Success 200 {object} Dosen
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /dosen/{kode} [get]
func GetDosenByKodeDosen(c *fiber.Ctx) error {
	kodeDosen := c.Params("kode")
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

// CreateDosen godoc
// @Summary Create data Dosen.
// @Description Input data Dosen.
// @Tags Dosen
// @Accept json
// @Produce json
// @Param request body Dosen true "Payload Body [RAW]"
// @Success 200 {object} Dosen
// @Failure 400
// @Failure 500
// @Router /dosen [post]
func CreateDosen(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var data_dosen model.Dosen
	if err := c.BodyParser(&data_dosen); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	insertedID, err := module.InsertDosen(db,
		data_dosen.KodeDosen,
		data_dosen.Nama,
		data_dosen.PhoneNumber,
	)
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

// UpdateDosen godoc
// @Summary Update data Dosen.
// @Description Ubah data Dosen.
// @Tags Dosen
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Dosen true "Payload Body [RAW]"
// @Success 200 {object} Dosen
// @Failure 400
// @Failure 500
// @Router /dosen/{id} [put]
func UpdateDosen(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	var data_dosen model.Dosen
	if err := c.BodyParser(&data_dosen); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	err = module.UpdateDosenByID(db,
		objectID,
		data_dosen.KodeDosen,
		data_dosen.Nama,
		data_dosen.PhoneNumber,
	)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data berhasil diupdate.",
	})
}

// DeleteDeleteDosen godoc
// @Summary Delete data Dosen.
// @Description Hapus data Dosen.
// @Tags Dosen
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /dosen/{id} [delete]
func DeleteDosen(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = module.DeleteDosenByID(db, objectID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data berhasil dihapus.",
	})

}


// GetAllMataKuliah godoc
// @Summary Get All Data MataKuliah.
// @Description Mengambil semua data MataKuliah.
// @Tags MataKuliah
// @Accept json
// @Produce json
// @Success 200 {object} MataKuliah
// @Router /matakuliah [get]
func GetAllMataKuliah(c *fiber.Ctx) error {
	matakuliah := module.GetMatkulAll(config.Ulbimongoconn)
	return c.JSON(matakuliah)
}

// GetMataKuliahByID godoc
// @Summary Get By ID Data MataKuliah.
// @Description Ambil per ID data MataKuliah.
// @Tags MataKuliah
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} MataKuliah
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /matakuliah/{id} [get]
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


// GetMataKuliahByKodeMataKuliah godoc
// @Summary Get By Kode Mata Kuliah Data MataKuliah.
// @Description Ambil per Kode Dosen data MataKuliah.
// @Tags MataKuliah
// @Accept json
// @Produce json
// @Param kode path string true "Masukan Kode MataKuliah"
// @Success 200 {object} MataKuliah
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /matakuliah/{kode} [get]
func GetMataKuliahByKodeMataKuliah(c *fiber.Ctx) error {
	kodeMataKuliah := c.Params("kode")
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

// CreateMataKuliah godoc
// @Summary Create data MataKuliah.
// @Description Input data MataKuliah.
// @Tags MataKuliah
// @Accept json
// @Produce json
// @Param request body MataKuliah true "Payload Body [RAW]"
// @Success 200 {object} MataKuliah
// @Failure 400
// @Failure 500
// @Router /matakuliah [post]
func CreateMataKuliah(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var data_matkul model.MataKuliah
	if err := c.BodyParser(&data_matkul); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	insertedID, err := module.InsertMatkul(db,
		data_matkul.KodeMatkul,
		data_matkul.Nama,
		data_matkul.Sks,
		data_matkul.Dosen,
	)
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

// UpdateMataKuliah godoc
// @Summary Update data MataKuliah.
// @Description Ubah data MataKuliah.
// @Tags MataKuliah
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body MataKuliah true "Payload Body [RAW]"
// @Success 200 {object} MataKuliah
// @Failure 400
// @Failure 500
// @Router /matakuliah/{id} [put]
func UpdateMataKuliah(c *fiber.Ctx) error {
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
	var data_matkul model.MataKuliah
	if err := c.BodyParser(&data_matkul); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = module.UpdateMatkulFromID(db,
		objectID,
		data_matkul.KodeMatkul,
		data_matkul.Nama,
		data_matkul.Sks,
		data_matkul.Dosen,
	)
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


// DeleteDeleteMataKuliah godoc
// @Summary Delete data MataKuliah.
// @Description Hapus data MataKuliah.
// @Tags MataKuliah
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /matakuliah/{id} [delete]
func DeleteMataKuliah(c *fiber.Ctx) error {
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

	// Call the DeletePresensi function with the parsed ID
	err = module.DeleteMatkulByID(db, objectID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully deleted",
	})
}


// GetAllProgramStudi godoc
// @Summary Get All Data ProgramStudi.
// @Description Mengambil semua data ProgramStudi.
// @Tags ProgramStudi
// @Accept json
// @Produce json
// @Success 200 {object} ProgramStudi
// @Router /programstudi [get]
func GetAllProgramStudi(c *fiber.Ctx) error {
	programstudi := module.GetProdiAll(config.Ulbimongoconn)
	return c.JSON(programstudi)
}


// GetProgramStudiByID godoc
// @Summary Get By ID Data ProgramStudi.
// @Description Ambil per ID data ProgramStudi.
// @Tags ProgramStudi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} ProgramStudi
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /programstudi/{id} [get]
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

// GetProgramStudiByKodeProgramStudi godoc
// @Summary Get By Program Studi Data ProgramStudi.
// @Description Ambil per Program Studi data ProgramStudi.
// @Tags ProgramStudi
// @Accept json
// @Produce json
// @Param kode path string true "Masukan Program Studi"
// @Success 200 {object} ProgramStudi
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /programstudi/{kode} [get]
func GetProgramStudiByKodeProgramStudi(c *fiber.Ctx) error {
	kodeProgramStudi := c.Params("kode")
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


// CreateProgramStudi godoc
// @Summary Create data ProgramStudi.
// @Description Input data ProgramStudi.
// @Tags ProgramStudi
// @Accept json
// @Produce json
// @Param request body ProgramStudi true "Payload Body [RAW]"
// @Success 200 {object} ProgramStudi
// @Failure 400
// @Failure 500
// @Router /programstudi [post]
func CreateProgramStudi(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var data_prodi model.ProgramStudi
	if err := c.BodyParser(&data_prodi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	insertedID, err := module.InsertProdi(db,
		data_prodi.KodeProgramStudi,
		data_prodi.Nama,
	)
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

// UpdateProgramStudi godoc
// @Summary Update data ProgramStudi.
// @Description Ubah data ProgramStudi.
// @Tags ProgramStudi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body ProgramStudi true "Payload Body [RAW]"
// @Success 200 {object} ProgramStudi
// @Failure 400
// @Failure 500
// @Router /programstudi/{id} [put]
func UpdateProgramStudi(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	prodi := c.Params("id")
	if prodi == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objectID, err := primitive.ObjectIDFromHex(prodi)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	// Parse the request body into a Presensi object
	var data_prodi model.ProgramStudi
	if err := c.BodyParser(&data_prodi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = module.UpdateProdiByID(db,
		objectID,
		data_prodi.KodeProgramStudi,
		data_prodi.Nama,
	)
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

// DeleteDeleteProgramStudi godoc
// @Summary Delete data ProgramStudi.
// @Description Hapus data ProgramStudi.
// @Tags ProgramStudi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /programstudi/{id} [delete]
func DeleteProgramStudi(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	prodi := c.Params("id")
	if prodi == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objectID, err := primitive.ObjectIDFromHex(prodi)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = module.DeleteProdiByID(db, objectID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully deleted",
	})
}

// FAKULTAS

// GetAllFakultas godoc
// @Summary Get All Data Fakultas.
// @Description Mengambil semua data Fakultas.
// @Tags Fakultas
// @Accept json
// @Produce json
// @Success 200 {object} Fakultas
// @Router /Fakultas [get]
func GetAllFakultas(c *fiber.Ctx) error {
	fakultas := module.GetFakultasAll(config.Ulbimongoconn)
	return c.JSON(fakultas)
}


// GetFakultasByID godoc
// @Summary Get By ID Data Fakultas.
// @Description Ambil per ID data Fakultas.
// @Tags Fakultas
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Fakultas
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /fakultas/{id} [get]
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

// GetFakultasByKodeFaklutas godoc
// @Summary Get By Kode Faklutas Data Fakultas.
// @Description Ambil per Kode Faklutas data Fakultas.
// @Tags Fakultas
// @Accept json
// @Produce json
// @Param kode path string true "Masukan Kode Faklutas"
// @Success 200 {object} Fakultas
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /fakultas/{kode} [get]
func GetFakultasByKodeFakultas(c *fiber.Ctx) error {
	kodeFakultas := c.Params("kode")
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

// CreateFakultas godoc
// @Summary Create data Fakultas.
// @Description Input data Fakultas.
// @Tags Fakultas
// @Accept json
// @Produce json
// @Param request body Fakultas true "Payload Body [RAW]"
// @Success 200 {object} Fakultas
// @Failure 400
// @Failure 500
// @Router /fakultas [post]
func CreateFakultas(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var data_fakultas model.Fakultas
	if err := c.BodyParser(&data_fakultas); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	insertedID, err := module.InsertFakultas(db,
		data_fakultas.KodeFakultas,
		data_fakultas.Nama,
	)
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

// UpdateFakultas godoc
// @Summary Update data Fakultas.
// @Description Ubah data Fakultas.
// @Tags Fakultas
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Fakultas true "Payload Body [RAW]"
// @Success 200 {object} Fakultas
// @Failure 400
// @Failure 500
// @Router /fakultas/{id} [put]
func UpdateFakultas(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	fakultas := c.Params("id")
	if fakultas == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objectID, err := primitive.ObjectIDFromHex(fakultas)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	// Parse the request body into a Presensi object
	var data_fakultas model.Fakultas
	if err := c.BodyParser(&data_fakultas); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = module.UpdateFakultasFromID(db,
		objectID,
		data_fakultas.KodeFakultas,
		data_fakultas.Nama,
	)
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

// DeleteDeleteFakultas godoc
// @Summary Delete data Fakultas.
// @Description Hapus data Fakultas.
// @Tags Fakultas
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /fakultas/{id} [delete]
func DeleteFakultas(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	fakultas := c.Params("id")
	if fakultas == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objectID, err := primitive.ObjectIDFromHex(fakultas)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = module.DeleteFakultasFromID(db, objectID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", fakultas),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully deleted",
	})
}

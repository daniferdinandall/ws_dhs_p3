package url

import (
	"github.com/daniferdinandall/ws-dhs-p3/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" 
)

func Web(page *fiber.App) {
	page.Get("/", controller.Homepage2)

	page.Get("/dhs", controller.GetAllDHS)
	page.Get("/dhs/:id", controller.GetDHSByID)
	page.Get("/dhsnpm/:npm", controller.GetDHSByNPM)
	page.Post("/dhs", controller.CreateDHS)
	page.Put("/dhs/:id", controller.UpdateDHS)
	page.Delete("/dhs/:id", controller.DeleteDHS)

	page.Get("/mahasiswa", controller.GetAllMahasiswa)
	page.Get("/mahasiswa/:id", controller.GetMahasiswaByID)
	page.Get("/mahasiswanpm/:npm", controller.GetMahasiswaByNPM)
	page.Post("/mahasiswa", controller.CreateMahasiswa)
	page.Put("/mahasiswa/:id", controller.UpdateMahasiswa)
	page.Delete("/mahasiswa/:id", controller.DeleteMahasiswa)

	page.Get("/dosen", controller.GetAllDosen)
	page.Get("/dosen/:id", controller.GetDosenByID)
	page.Get("/dosenkode/:kode", controller.GetDosenByKodeDosen)
	page.Post("/dosen", controller.CreateDosen)
	page.Put("/dosen/:id", controller.UpdateDosen)
	page.Delete("/dosen/:id", controller.DeleteDosen)

	page.Get("/matakuliah", controller.GetAllMataKuliah)
	page.Get("/matakuliah/:id", controller.GetMataKuliahByID)
	page.Get("/matakuliahkode/:kode", controller.GetMataKuliahByKodeMataKuliah)
	page.Post("/matakuliah", controller.CreateMataKuliah)
	page.Put("/matakuliah/:id", controller.UpdateMataKuliah)
	page.Delete("/matakuliah/:id", controller.DeleteMataKuliah)

	page.Get("/programstudi", controller.GetAllProgramStudi)
	page.Get("/programstudi/:id", controller.GetProgramStudiByID)
	page.Get("/programstudikode/:kode", controller.GetProgramStudiByKodeProgramStudi)
	page.Post("/programstudi", controller.CreateProgramStudi)
	page.Put("/programstudi/:id", controller.UpdateProgramStudi)
	page.Delete("/programstudi/:id", controller.DeleteProgramStudi)

	page.Get("/fakultas", controller.GetAllFakultas)
	page.Get("/fakultas/:id", controller.GetFakultasByID)
	page.Get("/fakultaskode/:kode", controller.GetFakultasByKodeFakultas)
	page.Post("/fakultas", controller.CreateFakultas)
	page.Put("/fakultas/:id", controller.UpdateFakultas)
	page.Delete("/fakultas/:id", controller.DeleteFakultas)
	page.Get("/docs/*", swagger.HandlerDefault)
}

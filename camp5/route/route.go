package route

import (
	"tugas-sds/camp5/database"
	"tugas-sds/camp5/models"

	"github.com/gofiber/fiber/v2"
)

// buatlah endpoint Insert Data sesuai dengan Colection Postman
func InsertData(c *fiber.Ctx) error {

	//Tulis Jawaban Code di Sini :))

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"Pesan": "Data telah berhasil di tambahkan",
	})
}

// Lengkapi Code Berikut untuk untuk Mengambil data untuk semua user - user
func GetAllData(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.JSON(fiber.Map{
		"data": users,
	})

}

//Lengkapi Code berikut Untuk Mengambil data dari id_user berdasarkan Parameter

func GetUserByid(c *fiber.Ctx) error {

	id := c.Params(`id_user`)
	var user models.User

	result := database.DB.Find(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.JSON(fiber.Map{
		"data": &user,
	})
}

func Delete(c *fiber.Ctx) error {

	var user models.User

	id_user := c.Params("id_user")

	database.DB.Where("id_user = ?", id_user).Delete(user)

	return c.JSON(fiber.Map{
		"Pesan": "Data telah di hapus",
	})
}

// mengupdate data user berdasarkan id_user yang di tempatkan di parameter
func Update(c *fiber.Ctx) error {

	id_user := c.Params("id_user")

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var users models.User
	database.DB.Find(&users)
	//data yang di ubah
	//membuat variable user berdasarkan model user
	var user models.User

	update := models.User{
		Nama:     data["nama"],
		Email:    data["email"],
		Username: data["username"],
		Password: data["password"],
	}
	//mengambil database untuk di update

	database.DB.Model(&user).Where("id_user = ?", id_user).Updates(update)

	return c.JSON(fiber.Map{
		"Pesan": "Data User telah di Update",
	})
}

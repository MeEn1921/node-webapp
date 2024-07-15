package controllers

import (
	"fmt"
	"log"
	"strconv"

	m "go-fiber-test/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func HelloTestV2(c *fiber.Ctx) error {
	return c.SendString("Hello, World V2👋!")
}

func BodyParserTest(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe
	str := p.Name + p.Pass
	return c.JSON(str)
}

func ParamsTest(c *fiber.Ctx) error {

	str := "hello ==> " + c.Params("name")
	return c.JSON(str)
}

func QueryTest(c *fiber.Ctx) error {
	a := c.Query("search")
	str := "my search is  " + a
	return c.JSON(str)
}

func ValidationTest(c *fiber.Ctx) error {
	//Connect to database

	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	return c.JSON(user)
}

// 5.1
func FactorialTest(c *fiber.Ctx) error {
	n := c.Params("n")
	fact := 1
	nInt, err := strconv.Atoi(n)
	if err != nil {
		return err
	}
	for i := 1; i <= nInt; i++ {
		fact = fact * i
	}
	return c.JSON(fact)
}

// 5.2
func QueryTestId(c *fiber.Ctx) error {
	a := c.Query("tax_id")
	ascii := []int{}
	for _, v := range a {
		ascii = append(ascii, int(v))
	}
	return c.JSON(ascii)
}

// 6
func RegisterTest(c *fiber.Ctx) error {
	fmt.Println("eiei")
	user := new(m.Usernames)
	fmt.Println("1")
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	fmt.Println("2")
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errors.Error(),
		})
	}
	fmt.Println("User: ", user)
	return c.JSON(user)
}

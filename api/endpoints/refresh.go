package endpoints

import (
	"strings"

	"github.com/drew138/games/api/authorization"
	"github.com/gofiber/fiber/v2"
)

// RefreshJWT - function handle to refresh jwts
func RefreshJWT(c *fiber.Ctx) {
	rToken := strings.Split(c.Get("Authorization"), " ")[1]
	parsedRToken, err := authorization.ParseJWT(rToken)
	if err != nil {
		c.Status(401).Send([]byte(err.Error()))
		return
	}
	tokenPair, err := authorization.RefreshToken(parsedRToken)
	if err != nil {
		c.Status(401).Send([]byte("Error: Invalid token"))
	}
	c.Status(201).JSON(tokenPair)
}

package middleware

import (
	"boilerplate/pkg/exception"
	"strings"
	"time"
	as "code.chakra.uno/crm/go-library/authservices"
	pd "code.chakra.uno/crm/go-library/portal-decryptor"

	"github.com/gofiber/fiber/v2"
)

func CheckAccessAndJWTValidator(linkUrl string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		init := exception.InitException(c, initData.Conf, initData.Log)

		authorizationHeader := c.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			return exception.CreateResponse_Log(init, fiber.StatusUnauthorized, "Invalid token format", "Format token tidak valid", nil)
		}

		magicToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		dataDecToken, err := pd.DecryptPortalToken(magicToken)
		if err != nil {
			return exception.CreateResponse_Log(init, fiber.StatusUnauthorized, err.Error(), err.Error(), nil)
		}

		//log.Println("DATA DEC TOKEN APP NAME : ", dataDecToken.ApplicationName)

		if dataDecToken.ApplicationName != initData.Conf.App.ApiGwName {
			return exception.CreateResponse_Log(init, fiber.StatusUnauthorized, "Youre not allowed to acces this application", "Anda tidak diizinkan untuk mengakses aplikasi ini", nil)
		}

		now := time.Now().UTC()

		isExpired := now.After(dataDecToken.ExpiredDate)

		if isExpired == true {
			return exception.CreateResponse_Log(init, fiber.StatusUnauthorized, "Portal Token is expired, please try to relogin in Portal", "Portal Token telah kadaluarsa, silahkan login kembali di Portal", nil)
		}

		applicationName := initData.Conf.App.ApiGwName

		roleName := "User"

		baseUrl := initData.Conf.App.BaseUrl

		IsProduction := initData.Conf.App.IsProduction

		errorData, validateAccess := as.GetAccessMethod(applicationName, roleName, baseUrl, linkUrl, IsProduction)
		if errorData != nil {
			return exception.CreateResponse_Log(init, fiber.StatusUnauthorized, errorData.Message, errorData.MessageInd, nil)
		} else {
			if validateAccess == false {
				return exception.CreateResponse_Log(init, fiber.StatusUnauthorized, "Roles "+roleName+" doesn't has access to method "+baseUrl+linkUrl, "Roles "+roleName+" tidak memiliki untuk method "+baseUrl+linkUrl, nil)
			} else {
				return c.Next()
			}
		}
	}
}

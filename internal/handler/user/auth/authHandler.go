package auth

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ReneKroon/ttlcache/v2"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/vinicius.csantos/fawkes-api/config"
	"gitlab.com/vinicius.csantos/fawkes-api/database"
	"gitlab.com/vinicius.csantos/fawkes-api/internal/model"
	constants "gitlab.com/vinicius.csantos/fawkes-api/internal/util/constant"
	"gitlab.com/vinicius.csantos/fawkes-api/internal/util/jwt"
	"golang.org/x/oauth2/twitch"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/oauth2/clientcredentials"
)

const (
	stateCallbackKey    = "oauth-state-callback"
	broadcasterTokenKey = "broadcaster-token"
)

var (
	clientID       = config.GetSecretKey("TWITCH_CLIENT_ID")
	clientSecret   = config.GetSecretKey("TWITCH_CLIENT_SECRET")
	oauth2Config   *clientcredentials.Config
	scopes         = strings.Split(config.GetSecretKey("TWITCH_SCOPES"), ";")
	twitchHelixUrl = config.GetSecretKey("TWITCH_HELIX_URL")
)

var AuthCache ttlcache.SimpleCache = ttlcache.NewCache()

func CreateState(c *fiber.Ctx) error {
	authCache := AuthCache

	var tokenBytes [255]byte

	if _, err := rand.Read(tokenBytes[:]); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constants.StatusInternalServerError, "message": constants.GenericInternalServerErrorMessage, "data": err.Error()})
	}

	state := hex.EncodeToString(tokenBytes[:])

	err := authCache.SetTTL(8 * time.Hour)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constants.StatusInternalServerError, "message": constants.SetCacheFailed, "data": err.Error()})
	}

	err = authCache.Set(stateCallbackKey, state)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constants.StatusInternalServerError, "message": constants.SetCacheFailed, "data": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": constants.StatusSuccess, "message": "Success", "data": state})
}

func LoginOrRegisterUser(c *fiber.Ctx) error {
	db := database.DB
	var userBody *model.UserFind
	var dataTwitch *model.DataTwitch
	var user *model.User

	err := c.BodyParser(&userBody)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": constants.StatusBadRequest, "message": constants.GenericInternalServerErrorMessage, "data": err.Error()})
	}

	userToken := userBody.TwitchToken

	err = getTwitchUserFromToken(&dataTwitch, userToken)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constants.StatusInternalServerError, "message": constants.GenericInternalServerErrorMessage, "data": err.Error()})
	}

	if dataTwitch != nil {
		if dataTwitch.Data == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": constants.StatusUnauthorized, "message": constants.GenericUnauthorizedMessage, "data": nil})
		} else if dataTwitch.Data[0].ID == "" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": constants.StatusNotFound, "message": model.MessageUser(constants.GenericNotFoundMessage), "data": nil})
		}
	}

	err = db.Find(&user, constants.IdCondition, dataTwitch.Data[0].ID).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constants.StatusInternalServerError, "message": constants.GenericInternalServerErrorMessage, "data": err.Error()})
	}

	if user.ID == 0 {
		twitchUser := dataTwitch.Data[0]
		user.ID, _ = strconv.ParseInt(twitchUser.ID, 10, 64)
		user.Login = twitchUser.Login
		user.DisplayName = twitchUser.DisplayName
		user.Email = twitchUser.Email
		user.ProfileImageUrl = twitchUser.ProfileImageUrl
		user.InRedemptionCooldown = false
		user.Role = model.RoleUser
		user.RedemptionCooldownEndsAt = time.Now()

		err = db.Create(&user).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constants.StatusInternalServerError, "message": constants.GenericInternalServerErrorMessage, "data": err.Error()})
		}
	}

	validToken, err := jwt.GenerateToken(user.Login, user.DisplayName, user.Email, strconv.FormatInt(user.ID, 10))

	data := fiber.Map{"user": user, "token": validToken}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": constants.StatusSuccess, "message": model.MessageUser(constants.GenericFoundSuccessMessage), "data": data})
}

func GetBroadcasterToken() (string, error) {
	authCache := AuthCache

	broadcasterToken, err := authCache.Get(broadcasterTokenKey)

	if err == nil {
		if err == ttlcache.ErrNotFound {
			oauth2Config = &clientcredentials.Config{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				TokenURL:     twitch.Endpoint.TokenURL,
				Scopes:       scopes,
			}

			err := authCache.SetTTL(8 * time.Hour)

			if err != nil {
				return "", err
			}

			token, err := oauth2Config.Token(context.Background())

			if err != nil {
				return "", err
			}

			err = authCache.Set(broadcasterTokenKey, token.TokenType)

			if err != nil {
				return "", err
			}

			return token.TokenType, nil

		}

		return "", nil
	}

	return broadcasterToken.(string), nil

}

var myClient = &http.Client{Timeout: 30 * time.Second}

func getTwitchUserFromToken(target interface{}, userToken string) error {

	url := fmt.Sprintf("%s/users", twitchHelixUrl)

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	request.Header.Add("Authorization", "Bearer "+userToken)
	request.Header.Add("client-id", clientID)

	response, err := myClient.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&target)

	if err != nil {
		return err
	}
	/*body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)*/

	return nil
}

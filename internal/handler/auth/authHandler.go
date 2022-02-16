package auth

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ReneKroon/ttlcache/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/vinicarvalhosantos/fawkes-api/config"
	"github.com/vinicarvalhosantos/fawkes-api/database"
	"github.com/vinicarvalhosantos/fawkes-api/internal/model"
	constants "github.com/vinicarvalhosantos/fawkes-api/internal/util/constant"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/oauth2/clientcredentials"
)

const (
	stateCallbackKey    = "oauth-state-callback"
	broadcasterTokenKey = "broadcaster-token"
	oauthSessionName    = "cookie:oauth-session"
	oauthTokenKey       = "oauth-token"
	sessionSaveFailed   = "Could not save the request session. Please contact an administrator!"
	stateGenerateFailed = "it was not possible to create a new state. Please contact an administrator!"
	invalidSession      = "This is a invalid session, generating a new one!"
	invalidState        = "This state is invalid!"
)

var (
	clientID                = config.GetSecretKey("TWITCH_CLIENT_ID")
	clientSecret            = config.GetSecretKey("TWITCH_CLIENT_SECRET")
	oauth2ClientCredentials *clientcredentials.Config
	redirectURL             = config.GetSecretKey("TWITCH_LOGIN_REDIRECT_URL")
	oauth2Config            *oauth2.Config
	scopes                  = strings.Split(config.GetSecretKey("TWITCH_SCOPES"), ";")
	twitchHelixUrl          = config.GetSecretKey("TWITCH_HELIX_URL")
	baseSuccessUrl          = fmt.Sprintf("%s/success", config.Config("REDIRECT_URL", ""))
	baseErrorUrl            = fmt.Sprintf("%s/error", config.Config("REDIRECT_URL", ""))
	cookieStore             = session.New(session.Config{
		CookieSecure: true,
		Expiration:   24 * time.Hour,
		KeyLookup:    oauthSessionName,
		KeyGenerator: utils.UUID,
	})
)

var Cache ttlcache.SimpleCache = ttlcache.NewCache()

func LoginOrRegisterTwitchUser(c *fiber.Ctx) error {
	sessionStorage, err := cookieStore.Get(c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constants.StatusInternalServerError, "message": invalidSession, "data": err.Error()})
	}

	state, err := generateState()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constants.StatusInternalServerError, "message": stateGenerateFailed, "data": err.Error()})
	}

	sessionStorage.Set(stateCallbackKey, state)

	if err = sessionStorage.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constants.StatusInternalServerError, "message": sessionSaveFailed, "data": err.Error()})
	}

	oauth2Config = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint:     twitch.Endpoint,
		RedirectURL:  redirectURL,
	}

	return c.Redirect(oauth2Config.AuthCodeURL(state), fiber.StatusTemporaryRedirect)
}

func UserLoginCallback(c *fiber.Ctx) error {
	sessionStorage, err := cookieStore.Get(c)

	if err != nil {
		urlParameters := fmt.Sprintf(constants.RedirectFrontErrorParams, constants.StatusInternalServerError, invalidSession, err.Error())
		return c.Redirect(fmt.Sprintf("%s%s", baseErrorUrl, url.QueryEscape(urlParameters)))
	}

	switch stateChallenge, state := sessionStorage.Get(stateCallbackKey), c.FormValue("state"); {

	case state == "", stateChallenge == "":
		urlParameters := fmt.Sprintf(constants.RedirectFrontErrorParams, constants.StatusBadRequest, "State is missing!", nil)
		return c.Redirect(fmt.Sprintf("%s%s", baseErrorUrl, url.QueryEscape(urlParameters)), fiber.StatusBadRequest)

	case state != stateChallenge:
		urlParameters := fmt.Sprintf(constants.RedirectFrontErrorParams, constants.StatusBadRequest, invalidState, nil)
		return c.Redirect(fmt.Sprintf("%s%s", baseErrorUrl, url.QueryEscape(urlParameters)), fiber.StatusBadRequest)

	}

	token, err := oauth2Config.Exchange(context.Background(), c.FormValue("code"))

	if err != nil {
		urlParameters := fmt.Sprintf(constants.RedirectFrontErrorParams, constants.StatusInternalServerError, invalidSession, err.Error())
		return c.Redirect(fmt.Sprintf("%s%s", baseErrorUrl, url.QueryEscape(urlParameters)), fiber.StatusInternalServerError)
	}

	sessionStorage.Set(oauthTokenKey, token)

	db := database.DB
	var dataTwitch *model.DataTwitch
	var user *model.User

	err = getTwitchUserFromToken(&dataTwitch, token.AccessToken)

	if err != nil {
		urlParameters := fmt.Sprintf(constants.RedirectFrontErrorParams, constants.StatusInternalServerError, constants.GenericInternalServerErrorMessage, err.Error())
		return c.Redirect(fmt.Sprintf("%s%s", baseErrorUrl, url.QueryEscape(urlParameters)), fiber.StatusInternalServerError)
	}

	if dataTwitch != nil {
		if dataTwitch.Data == nil {
			urlParameters := fmt.Sprintf("?status=%s&message=%s&data=%s", constants.StatusUnauthorized, constants.GenericUnauthorizedMessage, nil)
			return c.Redirect(fmt.Sprintf("%s%s", baseErrorUrl, url.QueryEscape(urlParameters)), fiber.StatusUnauthorized)
		} else if dataTwitch.Data[0].ID == "" {
			urlParameters := fmt.Sprintf(constants.RedirectFrontErrorParams, constants.StatusNotFound, model.MessageUser(constants.GenericNotFoundMessage), nil)
			return c.Redirect(fmt.Sprintf("%s%s", baseErrorUrl, url.QueryEscape(urlParameters)), fiber.StatusNotFound)
		}
	}

	err = db.Find(&user, constants.IdCondition, dataTwitch.Data[0].ID).Error

	if err != nil {
		urlParameters := fmt.Sprintf(constants.RedirectFrontErrorParams, constants.StatusInternalServerError, constants.GenericInternalServerErrorMessage, err.Error())
		return c.Redirect(fmt.Sprintf("%s%s", baseErrorUrl, url.QueryEscape(urlParameters)), fiber.StatusInternalServerError)
	}

	if user.ID == 0 {
		twitchUser := dataTwitch.Data[0]
		user.ID, _ = strconv.ParseInt(twitchUser.ID, 10, 64)
		user.Login = twitchUser.Login
		user.DisplayName = twitchUser.DisplayName
		user.Email = twitchUser.Email
		user.ProfileImageUrl = twitchUser.ProfileImageUrl
		user.InRedemptionCooldown = false
		user.Role = model.UserRole
		user.RedemptionCooldownEndsAt = time.Now()

		err = db.Create(&user).Error
		if err != nil {
			urlParameters := fmt.Sprintf(constants.RedirectFrontErrorParams, constants.StatusInternalServerError, constants.GenericInternalServerErrorMessage, err.Error())
			return c.Redirect(fmt.Sprintf("%s%s", baseErrorUrl, url.QueryEscape(urlParameters)), fiber.StatusInternalServerError)
		}
	}

	paramEncoded := fmt.Sprintf("user_id=%d", user.ID)
	return c.Redirect(fmt.Sprintf("%s?%s", baseSuccessUrl, url.QueryEscape(paramEncoded)), fiber.StatusPermanentRedirect)

}

func GetBroadcasterToken() (string, error) {
	authCache := Cache

	broadcasterToken, err := authCache.Get(broadcasterTokenKey)

	if err == nil {
		if err == ttlcache.ErrNotFound {
			oauth2ClientCredentials = &clientcredentials.Config{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				TokenURL:     twitch.Endpoint.TokenURL,
				Scopes:       scopes,
			}

			err := authCache.SetTTL(8 * time.Hour)

			if err != nil {
				return "", err
			}

			token, err := oauth2ClientCredentials.Token(context.Background())

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

func generateState() (string, error) {

	var tokenBytes [255]byte

	if _, err := rand.Read(tokenBytes[:]); err != nil {
		return "", err
	}

	state := hex.EncodeToString(tokenBytes[:])

	return state, nil
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

	return nil
}

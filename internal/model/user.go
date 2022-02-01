package model

import (
	stringUtil "github.com/vinicarvalhosantos/fawkes-api/internal/util/string"
	"time"
)

type UserRole string

const (
	RoleBroadcaster UserRole = "broadcaster" // Permissão Full
	RoleAdmin       UserRole = "admin"       // Permissão Full Exceto Editar os cargos
	RoleModerator   UserRole = "moderator"   //Permissão Básica Somente Leitura
	RoleUser        UserRole = "user"        //Permissão de Viewer
)

type User struct {
	ID                       int64  `gorm:"primaryKey"`
	Login                    string `gorm:"index;unique;not null;"`
	DisplayName              string `gorm:"not null;"`
	Email                    string `gorm:"index;unique;not null;"`
	ProfileImageUrl          string
	CpfCnpj                  string    `gorm:"index;"`
	Address                  []Address `gorm:"foreignKey:UserID"`
	Birthdate                string
	Balance                  int64 `gorm:"default:0"`
	InRedemptionCooldown     bool
	RedemptionCooldownEndsAt time.Time
	Role                     UserRole `gorm:"not null;"`
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

type TwitchUser struct {
	ID              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	ProfileImageUrl string `json:"profile_image_url"`
	Email           string `json:"email"`
}

type DataTwitch struct {
	Data []TwitchUser `json:"data"`
}

type UserFind struct {
	TwitchToken string
	State       string
}

func MessageUser(genericMessage string) string {
	return stringUtil.FormatGenericMessagesString(genericMessage, "User")
}

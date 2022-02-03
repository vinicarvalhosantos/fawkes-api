package model

import (
	stringUtil "github.com/vinicarvalhosantos/fawkes-api/internal/util/string"
	"time"
)

type Role string

const (
	BroadcasterRole Role = "broadcaster" // Permissão Full
	AdminRole       Role = "admin"       // Permissão Full Exceto Editar os cargos
	ModeratorRole   Role = "moderator"   //Permissão Básica Somente Leitura
	UserRole        Role = "user"        //Permissão de Viewer
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
	Balance                  int32 `gorm:"default:0"`
	InRedemptionCooldown     bool
	RedemptionCooldownEndsAt time.Time
	Role                     Role `gorm:"not null;"`
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

type UpdateUser struct {
	CpfCnpj   string
	Birthdate string
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

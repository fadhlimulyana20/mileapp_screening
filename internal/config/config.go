package config

import (
	"github.com/mileapp_screening/database"
	mail "github.com/mileapp_screening/utils/mailer"
	"github.com/mileapp_screening/utils/minio"

	"gorm.io/gorm"
)

type Config struct {
	ENV      string
	DB       *gorm.DB
	SMTP     *mail.Mailer
	Secret   string
	Minio    minio.MinioStorageContract
	Mongo    database.MongoDB
	HTTPPort int
}

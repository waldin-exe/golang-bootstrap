package seed

import (
	"context"
	"log"

	"github.com/waldin-exe/golang-bootstrap/internal/shared/contract"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSeeder struct {
	tx contract.DBProvider
}

func NewUserSeeder(tx contract.DBProvider) *UserSeeder {
	return &UserSeeder{
		tx: tx,
	}
}

func (s *UserSeeder) Run(ctx context.Context) {
	log.Println("[Seeder] running user seeder...")

	err := s.tx.WithTx(ctx, func(tx *gorm.DB) error {
		var count int64

		if err := tx.
			Table("users").
			Where("username = ?", "owner@bumexs99.com").
			Count(&count).Error; err != nil {
			return err
		}

		if count > 0 {
			log.Println("[Seeder] user already exists, skipping...")
			return nil
		}

		hashedPwd, err := bcrypt.GenerateFromPassword([]byte("owner123"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user := map[string]interface{}{
			"username":   "owner@mail.com",
			"password":   string(hashedPwd),
			"level":      "owner",
			"pegawai_id": 1,
		}

		if err := tx.Table("users").Create(&user).Error; err != nil {
			return err
		}

		log.Println("✅ user seed inserted")
		return nil
	})

	if err != nil {
		log.Printf("[Seeder] failed: %v", err)
	}
}

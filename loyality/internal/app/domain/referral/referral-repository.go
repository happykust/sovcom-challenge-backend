package referral

import (
	"github.com/google/uuid"
	"loyality/pkg/core/database"
)

func ReferralCreateRepository(referral Referral) Referral {
	database.PG.Create(&referral)
	return referral
}

func ReferralReadRepository(referralUUID uuid.UUID) []Referral {
	var referralFromDB []Referral
	database.PG.Find(&referralFromDB, "uuid = ?", referralUUID)
	return referralFromDB
}

func ReferralReadRepositoryByUserID(userID uint) []Referral {
	var referralFromDB []Referral
	database.PG.Find(&referralFromDB, "user_id = ?", userID)
	return referralFromDB
}

func ReferralUpdateRepository(referral Referral) Referral {
	database.PG.Model(&referral).Where("uuid = ?", referral.UUID).Updates(referral)
	return referral
}

func ReferralDeleteRepository(referralUUID uuid.UUID) uuid.UUID {
	database.PG.Delete(&Referral{}, "uuid = ?", referralUUID)
	return referralUUID
}

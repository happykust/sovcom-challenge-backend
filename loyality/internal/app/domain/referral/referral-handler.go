package referral

import (
	"github.com/google/uuid"
	"libs/contracts/loyality/referral"
)

func ReferralMathDepositBonusHandler(referralsCount uint) uint {
	if referralsCount >= 10 {
		return 5
	}
	return 0
}

func ReferralCreateHandler(request referral.CreateRequest) uuid.UUID {
	isReferralExist := ReferralReadRepositoryByUserID(request.UserID)
	if len(isReferralExist) > 0 {
		return uuid.UUID{}
	}
	createdReferral := ReferralCreateRepository(Referral{
		UUID:          uuid.New(),
		UserID:        request.UserID,
		Ticker:        request.Ticker,
		Amount:        request.Amount,
		ReferralCount: 0,
	})
	return createdReferral.UUID
}

func ReferralReadHandler(request referral.ReadRequest) Referral {
	var isReferralExist []Referral
	if request.UUID != uuid.Nil {
		isReferralExist = ReferralReadRepository(request.UUID)
	} else if request.UserID != 0 {
		isReferralExist = ReferralReadRepositoryByUserID(request.UserID)
	}
	if len(isReferralExist) == 0 {
		return Referral{}
	}
	return isReferralExist[0]
}

func ReferralDeleteHandler(request referral.DeleteRequest) uuid.UUID {
	isReferralExist := ReferralReadRepository(request.UUID)
	if len(isReferralExist) == 0 {
		return uuid.UUID{}
	}
	ReferralDeleteRepository(request.UUID)
	return request.UUID
}

func ReferralIncreaseReferralCountHandler(request referral.AddRequest) string {
	isReferralExist := ReferralReadRepository(request.UUID)
	if len(isReferralExist) == 0 {
		return "Referral not found"
	}
	ReferralUpdateRepository(Referral{
		UUID:          request.UUID,
		ReferralCount: isReferralExist[0].ReferralCount + 1,
	})
	return "Referral count increased"
}

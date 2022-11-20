package promocodes

import "libs/contracts/loyality/promocodes"

func PromocodesCreateHandler(request promocodes.CreateRequest) Promocode {
	isPromocodeExist := PromocodesReadRepositoryByPromocodeName(request.Promocode)
	if len(isPromocodeExist) > 0 {
		return Promocode{}
	}
	createdPromocode := PromocodesCreateRepository(Promocode{
		Promocode:            request.Promocode,
		Ticker:               request.Ticker,
		Amount:               request.Amount,
		ActivationCountLimit: request.ActivationCountLimit,
		ActivationCount:      0,
	})
	return createdPromocode
}

func PromocodesReadHandler(request promocodes.ReadRequest) Promocode {
	isPromocodeExist := PromocodesReadRepositoryByPromocodeName(request.Promocode)
	if len(isPromocodeExist) == 0 {
		return Promocode{}
	}
	return isPromocodeExist[0]
}

func PromocodesUpdateHandler(request Promocode) Promocode {
	isPromocodeExist := PromocodesReadRepositoryByPromocodeName(request.Promocode)
	if len(isPromocodeExist) == 0 {
		return Promocode{}
	}
	updatedPromocode := PromocodesUpdateRepository(request)
	return updatedPromocode
}

func PromocodesDeleteHandler(request promocodes.DeleteRequest) string {
	isPromocodeExist := PromocodesReadRepositoryByPromocodeName(request.Promocode)
	if len(isPromocodeExist) == 0 {
		return "Promocode not found"
	}
	deletedPromocodeName := PromocodesDeleteRepository(request.Promocode)
	return deletedPromocodeName
}

func PromocodesUseHandler(request promocodes.UseRequest) string {
	isPromocodeExist := PromocodesReadRepositoryByPromocodeName(request.Promocode)
	if len(isPromocodeExist) == 0 {
		return "Promocode not found"
	}
	if isPromocodeExist[0].ActivationCountLimit <= isPromocodeExist[0].ActivationCount {
		return "Max activation count limit reached"
	}
	isPromocodeExist[0].ActivationCount++
	PromocodesUpdateRepository(isPromocodeExist[0])
	return "Promocode used"
}

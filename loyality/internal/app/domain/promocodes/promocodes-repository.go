package promocodes

import "loyality/pkg/core/database"

func PromocodesReadRepositoryByPromocodeName(promocode string) []Promocode {
	var promocodeModel []Promocode
	database.PG.Find(&promocodeModel, "promocode = ?", promocode)
	return promocodeModel
}

func PromocodesCreateRepository(promocode Promocode) Promocode {
	database.PG.Create(&promocode)
	return promocode
}

func PromocodesUpdateRepository(promocode Promocode) Promocode {
	database.PG.Model(&promocode).Where("promocode = ?", promocode.Promocode).Updates(promocode)
	updatedPromocode := PromocodesReadRepositoryByPromocodeName(promocode.Promocode)
	return updatedPromocode[0]
}

func PromocodesDeleteRepository(promocode string) string {
	database.PG.Delete(&Promocode{}, "promocode = ?", promocode)
	return promocode
}

package simpleDeals

import "deals/pkg/core/database"

func SimpleDealCreateRepository(deal SimpleDeal) {
	database.PG.Create(&deal)
}

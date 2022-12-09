package offer

import "github.com/andyj29/wannabet/internal/domain/offer"

type Repository interface {
	Load(string) offer.Offer
	Save(offer.Offer) error
}

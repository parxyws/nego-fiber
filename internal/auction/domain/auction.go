package domain

type Auction struct{}

func (a *Auction) TableName() string {
	return "auctions"
}

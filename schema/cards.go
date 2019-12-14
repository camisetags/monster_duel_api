package schema

// BanList struct
type BanList struct {
	BanTcg *string `json:"ban_tcg"`
	BanOcg *string `json:"ban_ocg"`
}

// Card struct
type Card struct {
	ID          string     `json:"id"`
	Name        *string    `json:"name"`
	Type        *string    `json:"type"`
	Desc        *string    `json:"desc"`
	Atk         *string    `json:"atk"`
	Def         *string    `json:"def"`
	Level       *string    `json:"level"`
	Race        *string    `json:"race"`
	Attribute   *string    `json:"attribute"`
	CardSets    []*CardSet `json:"card_sets"`
	CardPrices  *CardPrice `json:"card_prices"`
	Archetype   *string    `json:"archetype"`
	Linkval     *string    `json:"linkval"`
	Fname       *string    `json:"fname"`
	Rank        *string    `json:"rank"`
	Linkmarkers []*string  `json:"linkmarkers"`
	BanlistInfo *BanList   `json:"banlist_info"`
	Format      *string    `json:"format"`
	Sort        *string    `json:"sort"`
	La          *string    `json:"la"`
	Scale       *string    `json:"scale"`
}

// CardPrice struct
type CardPrice struct {
	CardmarketPrice *float64 `json:"cardmarket_price"`
	TcgplayerPrice  *float64 `json:"tcgplayer_price"`
	EbayPrice       *float64 `json:"ebay_price"`
	AmazonPrice     *float64 `json:"amazon_price"`
}

// CardSet struct
type CardSet struct {
	SetName   *string `json:"set_name"`
	SetCode   *string `json:"set_code"`
	SetRarity *string `json:"set_rarity"`
	SetPrice  *string `json:"set_price"`
}

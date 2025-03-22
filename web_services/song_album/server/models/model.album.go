package models

// album represents data about a record album.
// Album (A) is exported while album (a) not exported in Go.
type Album struct {
	ID        string  `json:"id"`
	Title     string  `json:"title" bson:"title"`
	Artist    string  `json:"artist" bson:"artist"`
	Price     float64 `json:"price" bson:"price"`
	Image     string  `json:"image" bson:"image"`
	IsDeleted bool    `json:"isdeleted" bson:"isdeleted"`
	CreatedBy string  `json:"createdby" bson:"createdby"`
	CreatedAt string  `json:"createdat" bson:"createdat,default:currentTimestamp"`
	UpdatedBy string  `json:"updatedby" bson:"updatedby"`
	UpdatedAt string  `json:"updatedat" bson:"updatedat,default:currentTimestamp"`
}

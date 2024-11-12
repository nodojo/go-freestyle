package types

type User struct {
	// if we don't specify the json tag `bson:"_id"` here,
	// it will be marshalled as ID -> how do we want this in our DB?
	// bson is for mongodb and we want ID to appear there as _id
	ID        string `bson:"_id" json:"id"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}

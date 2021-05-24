package model

type Person struct {
	FirstName string `json:"FirstName" bson:"FirstName" binding:"required"`
	LastName  string `json:"LastName" bson:"LastName" binding:"required"`
	Age       int8   `json:"Age" bson:"Age" binding:"gte=1,lte=130"`
	Email     string `json:"Email" bson:"Email" binding:"required,email"`
}

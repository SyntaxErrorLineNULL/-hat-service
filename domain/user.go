package domain

type User struct {
	// User identifier
	ID string `json:"id" bson:"id"`
	// User first name
	FirstName string `json:"firstname,omitempty" bson:"firstname"`
	// User last name
	LastName string `json:"lastname,omitempty" bson:"lastname"`
	// User name (nickname)
	UserName string `json:"username" bson:"username"`
	// User email
	Email string `json:"email" bson:"email"`
}

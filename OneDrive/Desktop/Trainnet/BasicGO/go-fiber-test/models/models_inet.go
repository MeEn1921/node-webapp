package models

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

type Usernames struct {
	Email        string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
	UserName     string `json:"name" validate:"required,min=3,max=32,regexp=^[a-zA-Z0-9\\-_]*$"`
	Password     string `json:"password" validate:"required,min=6,max=20"`
	LineID       string `json:"lineid" validate:"required,min=3,max=32"`
	PhoneNumber  string `json:"phonenumber" validate:"required,min=3,max=9,regexp=^[0-9]*$"`
	TypeBusiness string `json:"typebusiness" validate:"required,min=3,max=32"`
	WebSite      string `json:"website,omitempty" validate:"required,min=2,max=30,regexp=^[a-z0-9\\-]*$"`
}

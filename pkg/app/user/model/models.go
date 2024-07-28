package model

import ("time")

type CreateUserModel struct{
    Username  string `json:"username" validate:"min=4"`
    Email     string `json:"email" validate:"nonzero"`
    Password string `json:"password" validate:"min=6"`
    PasswordVerified string `json:"password_valid" validate:"min=6"`
	CreatedAt time.Time `json:"-"`
}

type LoginUserModel struct{
    Email     string `json:"email"`
    Password string `json:"password"`
}

type GetUserModel struct{
    Username  string `json:"username"`
    Email     string `json:"email"`
}
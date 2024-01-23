package lapak

import (
	"time"
)

type User struct {
	Username     string `json:"username" bson:"username"`
	Password     string `json:"password" bson:"password"`
	PasswordHash string `json:"passwordhash" bson:"passwordhash"`
	Role         string `json:"role,omitempty" bson:"role,omitempty"`
	Token        string `json:"token,omitempty" bson:"token,omitempty"`
	Private      string `json:"private,omitempty" bson:"private,omitempty"`
	Public       string `json:"public,omitempty" bson:"public,omitempty"`
}

type RegisterStruct struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Response struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}

type Payload struct {
	User        string    `json:"user"`
	Pendapatan  string    `json:"pendapatan"`
	Pengeluaran string    `json:"pengeluaran"`
	Role        string    `json:"role"`
	Exp         time.Time `json:"exp"`
	Iat         time.Time `json:"iat"`
	Nbf         time.Time `json:"nbf"`
}

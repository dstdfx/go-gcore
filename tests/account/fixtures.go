package account

import "github.com/dstdfx/go-gcore/gcore"

var TestAccountDetailResponse = `{
    "currentUser": 511,
    "id": 505,
    "users": [
        {
            "client": 5,
            "company": "Your company",
            "deleted": false,
            "email": "user@yourcompany.com",
            "id": 513,
            "lang": "en",
            "name": "user",
            "phone": "+79882233443",
           "groups": [
               {
                "id": 2,
                "name": "users"
               }
           ]
        }
    ]
}
     `

var TestAccountDetailExpected = &gcore.Account{
	CurrentUser: 511,
	ID:          505,
	Users: []gcore.User{
		{
			Client:  5,
			Company: "Your company",
			Deleted: false,
			Email:   "user@yourcompany.com",
			ID:      513,
			Lang:    "en",
			Name:    "user",
			Phone:   "+79882233443",
			Groups: []*gcore.Group{
				{ID: 2, Name: "users"},
			},
		},
	},
}

package clients

import (
	"time"

	"github.com/dstdfx/go-gcore/gcore"
)

// Mocked responses
var (
	TestGetClientResponse = `{
  "id": 2,
  "users": [
    {
      "id": 6,
      "deleted": false,
      "email": "common2@gcore.lu",
      "name": "Client 2 Name",
      "client": 2,
      "company": "Client 2 Company Name",
      "lang": "en",
      "phone": "1232323",
      "groups": [
        {
          "id": 2,
          "name": "Administrators"
        }
      ]
    }
  ],
  "currentUser": 7,
  "email": "common2@gcore.lu",
  "phone": "Client 2 Company Phone",
  "name": "Client 2 Name",
  "status": "trial",
  "created": "2018-04-09T11:31:40.000000Z",
  "updated": "2018-04-09T11:32:31.000000Z",
  "companyName": "Client 2 Company Name",
  "utilization_level": 0,
  "reseller": 1,
  "cname": "example.gcdn.co"
}`
	TestCreateClientResponse = `{
  "id": 2,
  "users": [],
  "currentUser": 7,
  "client": 5,
  "email": "common2@gcore.lu",
  "phone": "Client 2 Company Phone",
  "name": "Client 2 Name",
  "status": "active",
  "created": "2018-04-09T11:31:40.000000Z",
  "updated": "2018-04-09T11:32:31.000000Z",
  "companyName": "Client 2 Company Name",
  "utilization_level": 0,
  "reseller": 1
}`
	TestListClientsResponse = `[{
  "id": 2,
  "users": [
    {
      "id": 6,
      "deleted": false,
      "email": "common2@gcore.lu",
      "name": "Client 2 Name",
      "client": 2,
      "company": "Client 2 Company Name",
      "lang": "en",
      "phone": "1232323",
      "groups": [
        {
          "id": 2,
          "name": "Administrators"
        }
      ]
    }
  ],
  "currentUser": 7,
  "email": "common2@gcore.lu",
  "phone": "Client 2 Company Phone",
  "name": "Client 2 Name",
  "status": "trial",
  "created": "2018-04-09T11:31:40.000000Z",
  "updated": "2018-04-09T11:32:31.000000Z",
  "companyName": "Client 2 Company Name",
  "utilization_level": 0,
  "reseller": 1,
  "cname": "example.gcdn.co"
}]`
	TestUpdateClientResponse = `{
  "id": 2,
  "users": [
    {
      "id": 6,
      "deleted": false,
      "email": "common2@gcore.lu",
      "name": "Client 2 Name",
      "client": 2,
      "company": "Client 2 Company Name",
      "lang": "en",
      "phone": "1232323",
      "groups": [
        {
          "id": 2,
          "name": "Administrators"
        }
      ]
    }
  ],
  "currentUser": 7,
  "email": "common2@gcore.lu",
  "phone": "Client 2 Company Phone",
  "name": "Another Name",
  "status": "trial",
  "created": "2018-04-09T11:31:40.000000Z",
  "updated": "2018-04-09T11:32:31.000000Z",
  "companyName": "Client 2 Company Name",
  "utilization_level": 0,
  "reseller": 1
}`

	TestUserTokenResponse = `{"token": "123123123ololo"}`
)

// Expected results
var (
	TestCreateClientExpected = &gcore.ClientAccount{
		ID:               2,
		Users:            []*gcore.User{},
		CurrentUser:      7,
		Status:           "active",
		Created:          gcore.NewGCoreTime(time.Date(2018, time.April, 9, 11, 31, 40, 0, time.UTC)),
		Updated:          gcore.NewGCoreTime(time.Date(2018, time.April, 9, 11, 32, 31, 0, time.UTC)),
		CompanyName:      "Client 2 Company Name",
		UtilizationLevel: 0,
		Reseller:         1,
		Email:            "common2@gcore.lu",
		Phone:            "Client 2 Company Phone",
		Name:             "Client 2 Name",
		Client:           5}

	TestGetClientExpected = &gcore.ClientAccount{
		ID: 2,
		Users: []*gcore.User{
			{
				ID:      6,
				Name:    "Client 2 Name",
				Deleted: false,
				Email:   "common2@gcore.lu",
				Client:  2,
				Company: "Client 2 Company Name",
				Lang:    "en",
				Phone:   "1232323",
				Groups: []*gcore.Group{
					{
						ID:   2,
						Name: "Administrators",
					},
				},
			},
		},
		CurrentUser:      7,
		Status:           "trial",
		Created:          gcore.NewGCoreTime(time.Date(2018, time.April, 9, 11, 31, 40, 0, time.UTC)),
		Updated:          gcore.NewGCoreTime(time.Date(2018, time.April, 9, 11, 32, 31, 0, time.UTC)),
		CompanyName:      "Client 2 Company Name",
		UtilizationLevel: 0,
		Reseller:         1,
		Email:            "common2@gcore.lu",
		Phone:            "Client 2 Company Phone",
		Name:             "Client 2 Name",
		Cname:            "example.gcdn.co",
	}

	TestListClientsExpected = []*gcore.ClientAccount{
		{
			ID: 2,
			Users: []*gcore.User{
				{
					ID:      6,
					Name:    "Client 2 Name",
					Deleted: false,
					Email:   "common2@gcore.lu",
					Client:  2,
					Company: "Client 2 Company Name",
					Lang:    "en",
					Phone:   "1232323",
					Groups: []*gcore.Group{
						{
							ID:   2,
							Name: "Administrators",
						},
					},
				},
			},
			CurrentUser:      7,
			Status:           "trial",
			Created:          gcore.NewGCoreTime(time.Date(2018, time.April, 9, 11, 31, 40, 0, time.UTC)),
			Updated:          gcore.NewGCoreTime(time.Date(2018, time.April, 9, 11, 32, 31, 0, time.UTC)),
			CompanyName:      "Client 2 Company Name",
			UtilizationLevel: 0,
			Reseller:         1,
			Email:            "common2@gcore.lu",
			Phone:            "Client 2 Company Phone",
			Name:             "Client 2 Name",
			Cname:            "example.gcdn.co",
		},
	}

	TestUpdateClientExpected = &gcore.ClientAccount{
		ID: 2,
		Users: []*gcore.User{
			{
				ID:      6,
				Name:    "Client 2 Name",
				Deleted: false,
				Email:   "common2@gcore.lu",
				Client:  2,
				Company: "Client 2 Company Name",
				Lang:    "en",
				Phone:   "1232323",
				Groups: []*gcore.Group{
					{
						ID:   2,
						Name: "Administrators",
					},
				},
			},
		},
		CurrentUser:      7,
		Status:           "trial",
		Created:          gcore.NewGCoreTime(time.Date(2018, time.April, 9, 11, 31, 40, 0, time.UTC)),
		Updated:          gcore.NewGCoreTime(time.Date(2018, time.April, 9, 11, 32, 31, 0, time.UTC)),
		CompanyName:      "Client 2 Company Name",
		UtilizationLevel: 0,
		Reseller:         1,
		Email:            "common2@gcore.lu",
		Phone:            "Client 2 Company Phone",
		Name:             "Client 2 Name",
	}

	TestUserTokenExpected = &gcore.Token{
		Value:  "123123123ololo",
		Expire: nil,
	}
)

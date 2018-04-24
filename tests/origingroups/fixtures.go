package origingroups

import "github.com/dstdfx/go-gcore/gcore"

var (
	TestGetOriginGroupResponse = `{
   "useNext" : false,
   "id" : 7272,
   "path" : "",
   "name" : "whatever.ru_wiggly.gcdn.co",
   "origins" : [
      {
         "enabled" : true,
         "backup" : false,
         "source" : "whatever.ru"
      }
   ],
   "origin_ids" : [
      {
         "id" : 9257,
         "source" : "whatever.ru",
         "enabled" : true,
         "backup" : false
      }
   ]
}`
	TestCreateOriginGroupResponse = `{
   "useNext" : false,
   "id" : 7272,
   "path" : "",
   "name" : "whatever.ru_wiggly.gcdn.co",
   "origins" : [
      {
         "enabled" : true,
         "backup" : false,
         "source" : "whatever.ru"
      }
   ],
   "origin_ids" : [
      {
         "id" : 9257,
         "source" : "whatever.ru",
         "enabled" : true,
         "backup" : false
      }
   ]
}`
	TestListOriginGroupsResponse = `[{
   "useNext" : false,
   "id" : 7272,
   "path" : "",
   "name" : "whatever.ru_wiggly.gcdn.co",
   "origins" : [
      {
         "enabled" : true,
         "backup" : false,
         "source" : "whatever.ru"
      }
   ],
   "origin_ids" : [
      {
         "id" : 9257,
         "source" : "whatever.ru",
         "enabled" : true,
         "backup" : false
      }
   ]
}]`
	TestUpdateOriginGroupResponse = `{
   "useNext" : false,
   "id" : 7272,
   "path" : "",
   "name" : "whatever2.ru_wiggly.gcdn.co",
   "origins" : [
      {
         "enabled" : true,
         "backup" : false,
         "source" : "whatever.ru"
      }
   ],
   "origin_ids" : [
      {
         "id" : 9257,
         "source" : "whatever.ru",
         "enabled" : true,
         "backup" : false
      }
   ]
}`
)

var (
	TestGetOriginGroupExpected = &gcore.OriginGroup{
		UseNext: false,
		ID:      7272,
		Name:    "whatever.ru_wiggly.gcdn.co",
		Origins: []gcore.Origin{
			{
				Backup:  false,
				Source:  "whatever.ru",
				Enabled: true,
			},
		},
		OriginIDs: []gcore.Origin{
			{
				ID:      9257,
				Source:  "whatever.ru",
				Enabled: true,
				Backup:  false,
			},
		},
	}

	TestListOriginGroupsExpected = []*gcore.OriginGroup{
		{
			UseNext: false,
			ID:      7272,
			Name:    "whatever.ru_wiggly.gcdn.co",
			Origins: []gcore.Origin{
				{
					Backup:  false,
					Source:  "whatever.ru",
					Enabled: true,
				},
			},
			OriginIDs: []gcore.Origin{
				{
					ID:      9257,
					Source:  "whatever.ru",
					Enabled: true,
					Backup:  false,
				},
			},
		},
	}

	TestCreateOriginGroupExpected = &gcore.OriginGroup{
		UseNext: false,
		ID:      7272,
		Name:    "whatever.ru_wiggly.gcdn.co",
		Origins: []gcore.Origin{
			{
				Backup:  false,
				Source:  "whatever.ru",
				Enabled: true,
			},
		},
		OriginIDs: []gcore.Origin{
			{
				ID:      9257,
				Source:  "whatever.ru",
				Enabled: true,
				Backup:  false,
			},
		},
	}

	TestUpdateOriginGroupExpected = &gcore.OriginGroup{
		UseNext: false,
		ID:      7272,
		Name:    "whatever2.ru_wiggly.gcdn.co",
		Origins: []gcore.Origin{
			{
				Backup:  false,
				Source:  "whatever.ru",
				Enabled: true,
			},
		},
		OriginIDs: []gcore.Origin{
			{
				ID:      9257,
				Source:  "whatever.ru",
				Enabled: true,
				Backup:  false,
			},
		},
	}
)

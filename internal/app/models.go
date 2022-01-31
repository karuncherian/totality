package app

const jsonData = `
[
	{
		"id": 1,
		"fname": "Steve",
		"city": "LA",
		"phone": 123465,
		"height": 5.8,
		"married": true
	},
	{
		"id": 2,
		"fname": "Lopez",
		"city": "LA",
		"phone": 45613,
		"height": 6.2,
		"married": true
	},
	{
		"id": 3,
		"fname": "Ave",
		"city": "CT",
		"phone": 789654,
		"height": 5.5,
		"married": false
	}
]`

type User struct {
	ID      int     `json:"id"`
	FName   string  `json:"fname"`
	City    string  `json:"city"`
	Phone   int64   `json:"phone"`
	Height  float32 `json:"height"`
	Married bool    `json:"married"`
}

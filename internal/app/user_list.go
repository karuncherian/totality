package app

import (
	"encoding/json"
	"net/http"
)

func UserList(w http.ResponseWriter, ids []int) {
	if len(ids) == 0 {
		fail(w, http.StatusBadRequest, 0, "Bad request")
		return
	}
	var req []User
	err := json.Unmarshal([]byte(jsonData), &req)
	if err != nil {
		fail(w, http.StatusBadRequest, 400, "Invalid data")
		return
	}

	var res []User
	for _, a := range ids {
		for _, b := range req {
			if a == b.ID {
				res = append(res, b)
			}
		}
	}
	if len(ids) > 1 {
		if len(res) > 0 {
			send(w, 200, res)
			return
		} else {
			send(w, 200, "No records")
			return
		}
	} else {
		if len(res) > 0 {
			send(w, 200, res[0])
			return
		} else {
			send(w, 200, "No records")
		}
	}
}

package config

import (
	"fmt"
)

const AuthId = 1

func CreateState() {
	e := db.Create(&State{
		ID: AuthId,
	})
	fmt.Println("e: ", e)
}

func (s *State) SetAuth(username, token string) {
	s.Username = username
	s.AuthToken = token
	s.Update()
}
func (s *State) Update() {
	s.ID = AuthId
	db.Save(s)
}

func GetState() *State {
	return db.Find(&State{}, AuthId).Value.(*State)
}

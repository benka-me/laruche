package config

import (
	"fmt"
)

const AuthId = 1

func CreateCredential() {
	e := db.Create(&State{
		ID:        AuthId,
		Username:  "haha",
		AuthToken: "jhkasdfg;asgfariuegiu3hh59ghp",
	})
	fmt.Println("e: ", e)
}

func (s State) UpdateCredential() {

}
func (s *State) GetCredential() *State {
	if s == nil {
		return nil
	}
	return db.Find(s, AuthId).Value.(*State)
}

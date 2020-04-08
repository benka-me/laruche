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

func (s State) UpdateState() {

}

func GetState() *State {
	return db.Find(&State{}, AuthId).Value.(*State)
}

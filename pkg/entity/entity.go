package entity

import "fmt"

type User struct {
	Id      int64
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

func (u *User) ToString() string {
	return fmt.Sprintf("name is %s and age is %d and friends %v \n", u.Name, u.Age, u.Friends)
}

func NewUser(name string, age int, friends []int) *User {
	return nil
}

type MakeFriends struct {
	SourceId string `json:"source_id"`
	TargetId string `json:"target_id"`
}

type DeleteUser struct {
	TargetId string `json:"target_id"`
}

type UpdateUser struct {
	Target string `json:"target_id"`
	NewAge string `json:"new_age"`
}

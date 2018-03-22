package models

type User struct {
	Mid string `json:"mid"`
	Address string `json:"address"`
	Name string `json:"name"`
	Sex byte `json:"sex"`
	Avatar string `json:"avatar"`
}

func (u *User) AddUser(){

}

func (u *User) GetUserByMid(mid string){

}

func (u *User) GetBalance(){

}

func (u *User) DeleteUserById(mid string){
	
}
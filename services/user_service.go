pakcage service

import (

)
// hubungan sama repo
type UserService struct{
	UserRepository *repository.UserRepository
}

type Request struct{
	errCode string `json:errcode`
	errMsg string `json:errmsg`
	Data RequestData `json:data`
}

type RequestData struct{
	id, 
	pass, 
	etc
}

type Response struct{
	errCode, 
	ErrMsg, 
	Token,
	Token,
	asdasdlsajdlksajf
	amsdas
	d;alsd
}

func NewUserRepository(userRepository *repository.UserRepository) *UserService{
	return &UserService{
		userRepository: userRepository
	}
}

func (us *UserService) GetUser(id string) (*Model.User, error){
	Request := &model.Request{}
	err := us.userRepository.GetUserById(Request.data.id,user)
	if err!= nil{
		return nil,err
	}
	return user, nil
}

func (us *UserService) GetUsers() (*[]Model.User, error){
	Request := &model.Request{}
	Response, err := us.userRepository.GetUserById(id,user)
	if err!= nil{
		return nil,err
	}
	return Response, nil
}
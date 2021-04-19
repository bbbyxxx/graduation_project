package rpc

import (
	"context"
	"lab_device_management_api/dal"
	"lab_device_management_api/model"
	"lab_device_management_api/proto/person/person"
	"log"

	"google.golang.org/grpc"
)

const (
	//address grpc服务地址
	Address = "127.0.0.1:8000"
)

func Login(ctx context.Context, conn *grpc.ClientConn, login *model.Login) (*person.LoginResponse, error) {
	var (
		err  error
		resp *person.LoginResponse
	)
	if conn == nil {
		conn, err = grpc.Dial(Address, grpc.WithInsecure())
	}
	if err != nil {
		log.Printf("[Login] Dial is failed,err:%v\n", err)
		return resp, err
	}
	c := person.NewPersonClient(conn)
	req := &person.LoginRequest{
		MultiId:  login.MultiId,
		Password: login.Password,
		Code:     login.Code,
	}
	resp, err = c.Login(ctx, req)
	if err != nil {
		log.Printf("[Login] call rpc Login is failed,err:%v\n", err)
		return resp, err
	}
	return resp, nil
}

func UpdatePerson(ctx context.Context, conn *grpc.ClientConn, modelPerson *model.Person) (*person.UpdateResponse, error) {
	var (
		err  error
		resp *person.UpdateResponse
	)
	if conn == nil {
		conn, err = grpc.Dial(Address, grpc.WithInsecure())
	}
	if err != nil {
		log.Printf("[UpdatePerson] Dial is failed,err:%v\n", err)
		return resp, err
	}
	c := person.NewPersonClient(conn)
	req := &person.UpdateRequest{
		Person: &person.Person{
			MultiId:       modelPerson.MultiId,
			Name:          modelPerson.Name,
			Sex:           modelPerson.Sex,
			Password:      modelPerson.Password,
			Phone:         modelPerson.Phone,
			Major:         modelPerson.Major,
			Grade:         modelPerson.Grade,
			Class:         modelPerson.Class,
			RegistTime:    modelPerson.RegistTime,
			UpdateTime:    modelPerson.UpdateTime,
			LoginTime:     modelPerson.LoginTime,
			LastLoginTime: modelPerson.LastLoginTime,
			Indentity:     modelPerson.Indentity,
		},
		IsDeleted: modelPerson.IsDeleted,
	}
	resp, err = c.UpdatePerson(ctx, req)
	if err != nil {
		log.Printf("[UpdatePerson] call rpc UpdatePerson is failed,err:%v\n", err)
		return resp, err
	}
	//删除请求的话，还要删除掉token
	if modelPerson.IsDeleted == 1 {
		err = dal.DelKey(nil, modelPerson.MultiId)
		if err != nil {
			return resp, nil
		}
	}
	return resp, nil
}

func RegisterPerson(ctx context.Context, conn *grpc.ClientConn, modelPerson *model.Person) (*person.RegisterResponse, error) {
	var (
		err  error
		resp *person.RegisterResponse
	)
	if conn == nil {
		conn, err = grpc.Dial(Address, grpc.WithInsecure())
	}
	if err != nil {
		log.Printf("[RegisterPerson] Dial is failed,err:%v\n", err)
		return resp, err
	}
	c := person.NewPersonClient(conn)
	req := &person.RegisterRequest{
		Person: &person.Person{
			MultiId:       modelPerson.MultiId,
			Name:          modelPerson.Name,
			Sex:           modelPerson.Sex,
			Password:      modelPerson.Password,
			Phone:         modelPerson.Phone,
			Major:         modelPerson.Major,
			Grade:         modelPerson.Grade,
			Class:         modelPerson.Class,
			RegistTime:    modelPerson.RegistTime,
			UpdateTime:    modelPerson.UpdateTime,
			LoginTime:     modelPerson.LoginTime,
			LastLoginTime: modelPerson.LastLoginTime,
			Indentity:     modelPerson.Indentity,
		},
	}

	resp, err = c.RegisterPerson(ctx, req)
	if err != nil {
		log.Printf("[RegisterPerson] call rpc RegisterPerson is failed,err:%v\n", err)
		return resp, err
	}
	return resp, nil
}

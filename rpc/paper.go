package rpc

import (
	"context"
	"lab_device_management_api/model"
	"lab_device_management_api/proto/paper/paper"
	"log"

	"google.golang.org/grpc"
)

func MGetPaper(ctx context.Context, conn *grpc.ClientConn, modelPersonDevicePaper *model.PersonDevicePaper) (*paper.MGetPaperResponse, error) {
	var (
		err  error
		resp *paper.MGetPaperResponse
	)
	if conn == nil {
		conn, err = grpc.Dial(AddressDevice, grpc.WithInsecure())
	}
	if err != nil {
		log.Printf("[MGetPaper] Dial is failed,err:%v\n", err)
		return resp, err
	}
	c := paper.NewPaperClient(conn)
	req := &paper.MGetPaperRequest{
		MultiId:             modelPersonDevicePaper.MultiId,
		PaperNumber:         modelPersonDevicePaper.PaperNumber,
		DeviceNumberModelId: modelPersonDevicePaper.DeviceNumberModelId,
	}
	resp, err = c.MGetPaper(ctx, req)
	if err != nil {
		log.Printf("[MGetPaper] is failed,err:%v\n", err)
		return resp, nil
	}
	return resp, nil
}

func AddPaper(ctx context.Context, conn *grpc.ClientConn, modelPaper *model.Paper, multiId string) (*paper.AddPaperResponse, error) {
	var (
		err  error
		resp *paper.AddPaperResponse
	)
	if conn == nil {
		conn, err = grpc.Dial(AddressDevice, grpc.WithInsecure())
	}
	if err != nil {
		log.Printf("[AddPaper] Dial is failed,err:%v\n", err)
		return resp, err
	}
	c := paper.NewPaperClient(conn)
	modelPaperReq := &paper.Paper{
		PaperNumber:         modelPaper.PaperNumber,
		PaperName:           modelPaper.PaperName,
		PaperTopic:          modelPaper.PaperTopic,
		PaperContent:        modelPaper.PaperContent,
		RelatedCode:         modelPaper.RelatedCode,
		DeviceNumberModelId: modelPaper.DeviceNumberModelId,
	}
	req := &paper.AddPaperRequest{
		MultiId: multiId,
		Paper:   modelPaperReq,
	}
	resp, err = c.AddPaper(ctx, req)
	if err != nil {
		log.Printf("[AddPaper] call rpc AddPaper is failed,err:%v\n", err)
		return resp, err
	}
	return resp, nil
}

func UpdatePaper(ctx context.Context, conn *grpc.ClientConn, modelPaper *model.Paper, multiId string) (*paper.UpdatePaperResponse, error) {
	var (
		err  error
		resp *paper.UpdatePaperResponse
	)
	if conn == nil {
		conn, err = grpc.Dial(AddressDevice, grpc.WithInsecure())
	}
	if err != nil {
		log.Printf("[UpdatePaper] Dial is failed,err:%v\n", err)
		return resp, err
	}
	c := paper.NewPaperClient(conn)
	req := &paper.UpdatePaperRequest{
		PaperNumber:         modelPaper.PaperNumber,
		PaperName:           modelPaper.PaperName,
		PaperTopic:          modelPaper.PaperTopic,
		PaperContent:        modelPaper.PaperContent,
		RelatedCode:         modelPaper.RelatedCode,
		DeviceNumberModelId: modelPaper.DeviceNumberModelId,
		IsDeleted:           modelPaper.IsDeleted,
		MultiId:             multiId,
	}
	resp, err = c.UpdatePaper(ctx, req)
	if err != nil {
		log.Printf("[UpdatePaper] call rpc UpdatePaper is failed,err:%v\n", err)
		return resp, err
	}
	return resp, nil
}

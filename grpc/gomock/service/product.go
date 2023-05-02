package service

import (
	"context"
	"fmt"
	"io"
	"time"
)

var ProductService = &productService{}

type productService struct {
}

func (p *productService) GetProductStock(context.Context, *ProductRequest) (*ProductResponse, error) {
	return nil, nil
}

func (p *productService) mustEmbedUnimplementedProdServiceServer() {

}

func (p *productService) UpdateProductStockClientStream(stream ProdService_UpdateProductStockClientStreamServer) error {
	count := 0
	for {
		// 源源不断的去接受客户端发来的信息
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		fmt.Println("服务端接收到的流", recv.ProdId, count)
		count++
		if count > 10 {
			rsp := &ProductResponse{ProdStock: recv.ProdId}
			err := stream.SendAndClose(rsp)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func (p *productService) SayHelloServerStream(stream ProdService_SayHelloServerStreamServer) error {
	for {
		recv, err := stream.Recv()
		if err != nil {
			return nil
		}
		fmt.Println("服务端收到客户端的消息", recv.ProdId)
		time.Sleep(time.Second)
		rsp := &ProductResponse{ProdStock: recv.ProdId}
		err = stream.Send(rsp)
		if err != nil {
			return nil
		}
	}
}

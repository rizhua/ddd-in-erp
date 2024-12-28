package rpc

import (
	"context"
	"log"
	"net"
	"reflect"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type Server struct {
	listen net.Listener
}

func NewServer(listen net.Listener) *Server {
	return &Server{listen: listen}
}

func (t *Server) Run() {
	// 注册一元 interceptor
	// var opts []grpc.ServerOption
	// opts = append(opts, grpc.UnaryInterceptor(interceptor))

	// 创建grpc句柄
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor),
	)
	// defer srv.GracefulStop()
	RegisterRizhuaServer(srv, &Server{})

	// 监听服务
	err := srv.Serve(t.listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (t *Server) Dial(ctx context.Context, in *Query) (*Reply, error) {
	var (
		err   error
		reply *Reply
	)

	reply = &Reply{}
	arr := strings.Split(in.Method, ".")
	if 2 > len(arr) {
		err = status.Error(0, "Method is invalid")
		return reply, err
	}

	// 获取逻辑层
	obj := map[string]interface{}{
		// "order": facade.NewOrderFacade(ctx),
	}
	if _, ok := obj[arr[0]]; !ok {
		err = status.Error(404, "Struct is invalid")
		return reply, err
	}

	// 动态调方法
	f := reflect.ValueOf(obj[arr[0]]).MethodByName(arr[1])
	if f.IsValid() {
		args := []reflect.Value{reflect.ValueOf(in.Data)}
		value := f.Call(args)
		for _, v := range value {
			if v.Kind() == reflect.Slice && !v.IsNil() {
				reply.Data = v.Interface().([]byte)
			}

			if v.Kind() == reflect.Int64 {
				reply.Total = v.Interface().(int64)
			}

			if v.Kind() == reflect.Interface && !v.IsNil() {
				err = v.Interface().(error)
			}
		}
	} else {
		err = status.Error(2041, "Method does not exist")
	}

	return reply, err
}

// interceptor 一元拦截器
func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// err := auth(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// 继续处理请求
	return handler(ctx, req)
}

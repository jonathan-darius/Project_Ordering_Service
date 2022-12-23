// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: product.proto

package ProductService

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProductManagementServiceClient is the client API for ProductManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductManagementServiceClient interface {
	RegisterProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductId, error)
	UpdateProduct(ctx context.Context, in *UpdateProductMSG, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteProduct(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*empty.Empty, error)
	AddStock(ctx context.Context, in *Stock, opts ...grpc.CallOption) (*empty.Empty, error)
	AddCategory(ctx context.Context, in *UpdateProductMSG, opts ...grpc.CallOption) (*empty.Empty, error)
	RemoveCatogory(ctx context.Context, in *UpdateProductMSG, opts ...grpc.CallOption) (*empty.Empty, error)
	// ToDo Image Management Array
	UploadProductImage(ctx context.Context, opts ...grpc.CallOption) (ProductManagementService_UploadProductImageClient, error)
}

type productManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductManagementServiceClient(cc grpc.ClientConnInterface) ProductManagementServiceClient {
	return &productManagementServiceClient{cc}
}

func (c *productManagementServiceClient) RegisterProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductId, error) {
	out := new(ProductId)
	err := c.cc.Invoke(ctx, "/product.ProductManagementService/RegisterProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManagementServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductMSG, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/product.ProductManagementService/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManagementServiceClient) DeleteProduct(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/product.ProductManagementService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManagementServiceClient) AddStock(ctx context.Context, in *Stock, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/product.ProductManagementService/AddStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManagementServiceClient) AddCategory(ctx context.Context, in *UpdateProductMSG, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/product.ProductManagementService/AddCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManagementServiceClient) RemoveCatogory(ctx context.Context, in *UpdateProductMSG, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/product.ProductManagementService/RemoveCatogory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManagementServiceClient) UploadProductImage(ctx context.Context, opts ...grpc.CallOption) (ProductManagementService_UploadProductImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductManagementService_ServiceDesc.Streams[0], "/product.ProductManagementService/UploadProductImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &productManagementServiceUploadProductImageClient{stream}
	return x, nil
}

type ProductManagementService_UploadProductImageClient interface {
	Send(*UploadImageRequest) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type productManagementServiceUploadProductImageClient struct {
	grpc.ClientStream
}

func (x *productManagementServiceUploadProductImageClient) Send(m *UploadImageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *productManagementServiceUploadProductImageClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductManagementServiceServer is the server API for ProductManagementService service.
// All implementations must embed UnimplementedProductManagementServiceServer
// for forward compatibility
type ProductManagementServiceServer interface {
	RegisterProduct(context.Context, *Product) (*ProductId, error)
	UpdateProduct(context.Context, *UpdateProductMSG) (*empty.Empty, error)
	DeleteProduct(context.Context, *ProductId) (*empty.Empty, error)
	AddStock(context.Context, *Stock) (*empty.Empty, error)
	AddCategory(context.Context, *UpdateProductMSG) (*empty.Empty, error)
	RemoveCatogory(context.Context, *UpdateProductMSG) (*empty.Empty, error)
	// ToDo Image Management Array
	UploadProductImage(ProductManagementService_UploadProductImageServer) error
	mustEmbedUnimplementedProductManagementServiceServer()
}

// UnimplementedProductManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductManagementServiceServer struct {
}

func (UnimplementedProductManagementServiceServer) RegisterProduct(context.Context, *Product) (*ProductId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterProduct not implemented")
}
func (UnimplementedProductManagementServiceServer) UpdateProduct(context.Context, *UpdateProductMSG) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductManagementServiceServer) DeleteProduct(context.Context, *ProductId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductManagementServiceServer) AddStock(context.Context, *Stock) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStock not implemented")
}
func (UnimplementedProductManagementServiceServer) AddCategory(context.Context, *UpdateProductMSG) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategory not implemented")
}
func (UnimplementedProductManagementServiceServer) RemoveCatogory(context.Context, *UpdateProductMSG) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCatogory not implemented")
}
func (UnimplementedProductManagementServiceServer) UploadProductImage(ProductManagementService_UploadProductImageServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadProductImage not implemented")
}
func (UnimplementedProductManagementServiceServer) mustEmbedUnimplementedProductManagementServiceServer() {
}

// UnsafeProductManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductManagementServiceServer will
// result in compilation errors.
type UnsafeProductManagementServiceServer interface {
	mustEmbedUnimplementedProductManagementServiceServer()
}

func RegisterProductManagementServiceServer(s grpc.ServiceRegistrar, srv ProductManagementServiceServer) {
	s.RegisterService(&ProductManagementService_ServiceDesc, srv)
}

func _ProductManagementService_RegisterProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagementServiceServer).RegisterProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductManagementService/RegisterProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagementServiceServer).RegisterProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManagementService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductMSG)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagementServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductManagementService/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagementServiceServer).UpdateProduct(ctx, req.(*UpdateProductMSG))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManagementService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagementServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductManagementService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagementServiceServer).DeleteProduct(ctx, req.(*ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManagementService_AddStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagementServiceServer).AddStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductManagementService/AddStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagementServiceServer).AddStock(ctx, req.(*Stock))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManagementService_AddCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductMSG)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagementServiceServer).AddCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductManagementService/AddCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagementServiceServer).AddCategory(ctx, req.(*UpdateProductMSG))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManagementService_RemoveCatogory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductMSG)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagementServiceServer).RemoveCatogory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductManagementService/RemoveCatogory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagementServiceServer).RemoveCatogory(ctx, req.(*UpdateProductMSG))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManagementService_UploadProductImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProductManagementServiceServer).UploadProductImage(&productManagementServiceUploadProductImageServer{stream})
}

type ProductManagementService_UploadProductImageServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*UploadImageRequest, error)
	grpc.ServerStream
}

type productManagementServiceUploadProductImageServer struct {
	grpc.ServerStream
}

func (x *productManagementServiceUploadProductImageServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *productManagementServiceUploadProductImageServer) Recv() (*UploadImageRequest, error) {
	m := new(UploadImageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductManagementService_ServiceDesc is the grpc.ServiceDesc for ProductManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductManagementService",
	HandlerType: (*ProductManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterProduct",
			Handler:    _ProductManagementService_RegisterProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductManagementService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductManagementService_DeleteProduct_Handler,
		},
		{
			MethodName: "AddStock",
			Handler:    _ProductManagementService_AddStock_Handler,
		},
		{
			MethodName: "AddCategory",
			Handler:    _ProductManagementService_AddCategory_Handler,
		},
		{
			MethodName: "RemoveCatogory",
			Handler:    _ProductManagementService_RemoveCatogory_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadProductImage",
			Handler:       _ProductManagementService_UploadProductImage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "product.proto",
}

// OrderingServiceClient is the client API for OrderingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderingServiceClient interface {
	SearchProduct(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (OrderingService_SearchProductClient, error)
	DecreaseStock(ctx context.Context, in *Stock, opts ...grpc.CallOption) (*empty.Empty, error)
	AddSoldProduct(ctx context.Context, in *Stock, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateRating(ctx context.Context, in *Rating, opts ...grpc.CallOption) (*empty.Empty, error)
	GetProductByID(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*Product, error)
	GetAllProduct(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (OrderingService_GetAllProductClient, error)
}

type orderingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderingServiceClient(cc grpc.ClientConnInterface) OrderingServiceClient {
	return &orderingServiceClient{cc}
}

func (c *orderingServiceClient) SearchProduct(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (OrderingService_SearchProductClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderingService_ServiceDesc.Streams[0], "/product.OrderingService/SearchProduct", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderingServiceSearchProductClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderingService_SearchProductClient interface {
	Recv() (*Product, error)
	grpc.ClientStream
}

type orderingServiceSearchProductClient struct {
	grpc.ClientStream
}

func (x *orderingServiceSearchProductClient) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderingServiceClient) DecreaseStock(ctx context.Context, in *Stock, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/product.OrderingService/DecreaseStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderingServiceClient) AddSoldProduct(ctx context.Context, in *Stock, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/product.OrderingService/AddSoldProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderingServiceClient) UpdateRating(ctx context.Context, in *Rating, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/product.OrderingService/UpdateRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderingServiceClient) GetProductByID(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/product.OrderingService/GetProductByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderingServiceClient) GetAllProduct(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (OrderingService_GetAllProductClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderingService_ServiceDesc.Streams[1], "/product.OrderingService/GetAllProduct", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderingServiceGetAllProductClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderingService_GetAllProductClient interface {
	Recv() (*Product, error)
	grpc.ClientStream
}

type orderingServiceGetAllProductClient struct {
	grpc.ClientStream
}

func (x *orderingServiceGetAllProductClient) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderingServiceServer is the server API for OrderingService service.
// All implementations must embed UnimplementedOrderingServiceServer
// for forward compatibility
type OrderingServiceServer interface {
	SearchProduct(*SearchQuery, OrderingService_SearchProductServer) error
	DecreaseStock(context.Context, *Stock) (*empty.Empty, error)
	AddSoldProduct(context.Context, *Stock) (*empty.Empty, error)
	UpdateRating(context.Context, *Rating) (*empty.Empty, error)
	GetProductByID(context.Context, *ProductId) (*Product, error)
	GetAllProduct(*Pagination, OrderingService_GetAllProductServer) error
	mustEmbedUnimplementedOrderingServiceServer()
}

// UnimplementedOrderingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderingServiceServer struct {
}

func (UnimplementedOrderingServiceServer) SearchProduct(*SearchQuery, OrderingService_SearchProductServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchProduct not implemented")
}
func (UnimplementedOrderingServiceServer) DecreaseStock(context.Context, *Stock) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecreaseStock not implemented")
}
func (UnimplementedOrderingServiceServer) AddSoldProduct(context.Context, *Stock) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSoldProduct not implemented")
}
func (UnimplementedOrderingServiceServer) UpdateRating(context.Context, *Rating) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRating not implemented")
}
func (UnimplementedOrderingServiceServer) GetProductByID(context.Context, *ProductId) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductByID not implemented")
}
func (UnimplementedOrderingServiceServer) GetAllProduct(*Pagination, OrderingService_GetAllProductServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllProduct not implemented")
}
func (UnimplementedOrderingServiceServer) mustEmbedUnimplementedOrderingServiceServer() {}

// UnsafeOrderingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderingServiceServer will
// result in compilation errors.
type UnsafeOrderingServiceServer interface {
	mustEmbedUnimplementedOrderingServiceServer()
}

func RegisterOrderingServiceServer(s grpc.ServiceRegistrar, srv OrderingServiceServer) {
	s.RegisterService(&OrderingService_ServiceDesc, srv)
}

func _OrderingService_SearchProduct_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderingServiceServer).SearchProduct(m, &orderingServiceSearchProductServer{stream})
}

type OrderingService_SearchProductServer interface {
	Send(*Product) error
	grpc.ServerStream
}

type orderingServiceSearchProductServer struct {
	grpc.ServerStream
}

func (x *orderingServiceSearchProductServer) Send(m *Product) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderingService_DecreaseStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderingServiceServer).DecreaseStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.OrderingService/DecreaseStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderingServiceServer).DecreaseStock(ctx, req.(*Stock))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderingService_AddSoldProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderingServiceServer).AddSoldProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.OrderingService/AddSoldProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderingServiceServer).AddSoldProduct(ctx, req.(*Stock))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderingService_UpdateRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rating)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderingServiceServer).UpdateRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.OrderingService/UpdateRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderingServiceServer).UpdateRating(ctx, req.(*Rating))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderingService_GetProductByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderingServiceServer).GetProductByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.OrderingService/GetProductByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderingServiceServer).GetProductByID(ctx, req.(*ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderingService_GetAllProduct_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Pagination)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderingServiceServer).GetAllProduct(m, &orderingServiceGetAllProductServer{stream})
}

type OrderingService_GetAllProductServer interface {
	Send(*Product) error
	grpc.ServerStream
}

type orderingServiceGetAllProductServer struct {
	grpc.ServerStream
}

func (x *orderingServiceGetAllProductServer) Send(m *Product) error {
	return x.ServerStream.SendMsg(m)
}

// OrderingService_ServiceDesc is the grpc.ServiceDesc for OrderingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.OrderingService",
	HandlerType: (*OrderingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DecreaseStock",
			Handler:    _OrderingService_DecreaseStock_Handler,
		},
		{
			MethodName: "AddSoldProduct",
			Handler:    _OrderingService_AddSoldProduct_Handler,
		},
		{
			MethodName: "UpdateRating",
			Handler:    _OrderingService_UpdateRating_Handler,
		},
		{
			MethodName: "GetProductByID",
			Handler:    _OrderingService_GetProductByID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchProduct",
			Handler:       _OrderingService_SearchProduct_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAllProduct",
			Handler:       _OrderingService_GetAllProduct_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "product.proto",
}

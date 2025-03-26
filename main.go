package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/xavicci/taxisgo/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Aquí simulamos una verificación básica de credenciales
	// En un caso real, verificarías contra una base de datos
	if req.Email == "usuario@ejemplo.com" && req.Password == "password123" {
		user := &pb.User{
			Id:        "usr_123",
			Email:     req.Email,
			Name:      "Usuario Ejemplo",
			Phone:     "+1234567890",
			CreatedAt: time.Now().Format(time.RFC3339),
		}

		return &pb.LoginResponse{
			Token: "jwt_token_simulado_123",
			User:  user,
		}, nil
	}

	return nil, grpc.Errorf(codes.Unauthenticated, "Credenciales inválidas")
}

func (s *server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Aquí implementarías la lógica de registro
	// Por ahora, simulamos un registro exitoso
	user := &pb.User{
		Id:        "usr_" + time.Now().Format("20060102150405"),
		Email:     req.Email,
		Name:      req.Name,
		Phone:     req.Phone,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	return &pb.RegisterResponse{
		Token: "jwt_token_nuevo_usuario_123",
		User:  user,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("falló al escuchar: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server{})
	log.Printf("servidor escuchando en %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("falló al servir: %v", err)
	}
}

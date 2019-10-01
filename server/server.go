package main

import (
	pb "github.com/castaneai/grpc-broadcast-example"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"sync"
)

type User struct {
	chat pb.ChatRoom_ChatServer
	move pb.ChatRoom_ChatServer
}

type server struct{
	// 接続中のクライアントリスト
	clients map[string]pb.ChatRoom_ChatServer

	mu sync.RWMutex
}

func (s *server) addClient(uid string, srv pb.ChatRoom_ChatServer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[uid] = srv
}

func (s *server) removeClient(uid string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, uid)
}

func (s *server) getClients() []pb.ChatRoom_ChatServer {
	var cs []pb.ChatRoom_ChatServer

	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, c := range s.clients {
		cs = append(cs, c)
	}
	return cs
}

func (s *server) Chat(srv pb.ChatRoom_ChatServer) error {
	uid := uuid.Must(uuid.NewRandom()).String()
	log.Printf("new user: %s", uid)

	// 接続クライアントリストに登録
	s.addClient(uid, srv)
	// 関数を抜けるときはリストから削除
	defer s.removeClient(uid)

	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic: %v", err)
			os.Exit(1)
		}
	}()

	for {
		resp, err := srv.Recv()
		if err != nil {
			log.Printf("recv err: %v", err)
			break
		}
		log.Printf("broadcast: %s", resp.Message)
		for _, ss := range s.getClients() {
			if err := ss.Send(&pb.ChatResponse{Message: resp.Message}); err != nil {
				log.Printf("broadcast err: %v", err)
			}
		}
	}
	return nil
}

func main() {
	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatRoomServer(s, &server{
		clients: make(map[string]pb.ChatRoom_ChatServer),
		mu: sync.RWMutex{},
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

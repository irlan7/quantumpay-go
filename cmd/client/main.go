package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/irlan/quantumpay-go/internal/grpc/proto"
)

func main() {
	// Target: Node Lokal Anda
	target := "localhost:9090"
	
	fmt.Println("========================================")
	fmt.Printf("🔌 CLIENT: Mencoba koneksi ke %s...\n", target)
	
	// 1. Dial (Telepon) ke Server
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("❌ Gagal menelepon: %v", err)
	}
	defer conn.Close()
	client := pb.NewNodeServiceClient(conn)

	// 2. Siapkan Data Sapaan (Handshake)
	// Kita coba pakai Chain ID yang BENAR (77077)
	req := &pb.HandshakeRequest{
		ChainId:     77077, // <--- KUNCI UTAMA
		GenesisHash: "0x1d58599424f1159828236111f1f9e83063f66345091a99540c4989679269491a",
		NodeVersion: "v1.0-client-tester",
	}

	// 3. Kirim Sapaan
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("📨 CLIENT: Mengirim Handshake (ChainID: 77077)...")
	resp, err := client.Handshake(ctx, req)
	if err != nil {
		log.Fatalf("❌ Error Handshake: %v", err)
	}

	// 4. Lihat Balasan Server
	if resp.Success {
		fmt.Println("\n✅ SERVER MEMBALAS:")
		fmt.Printf("   Pesan: \"%s\"\n", resp.Message)
		fmt.Println("   Status: TERVERIFIKASI! Pintu Terbuka.")
	} else {
		fmt.Println("\n⛔ SERVER MENOLAK:")
		fmt.Printf("   Alasan: %s\n", resp.Message)
	}
	fmt.Println("========================================")
}

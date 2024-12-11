package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// ログインサービスを起動
	loginCmd := exec.Command("go", "run", "./apps/login/main.go")
	loginCmd.Stdout = os.Stdout
	loginCmd.Stderr = os.Stderr
	if err := loginCmd.Start(); err != nil {
		log.Fatalf("ログインサービスの起動に失敗しました: %v", err)
	}

	// ホームサービスを起動
	homeCmd := exec.Command("go", "run", "./apps/home/main.go")
	homeCmd.Stdout = os.Stdout
	homeCmd.Stderr = os.Stderr
	if err := homeCmd.Start(); err != nil {
		log.Fatalf("ホームサービスの起動に失敗しました: %v", err)
	}

	// サービスが終了するのを待つ
	if err := loginCmd.Wait(); err != nil {
		log.Printf("ログインサービスが終了しました: %v", err)
	}
	if err := homeCmd.Wait(); err != nil {
		log.Printf("ホームサービスが終了しました: %v", err)
	}
}

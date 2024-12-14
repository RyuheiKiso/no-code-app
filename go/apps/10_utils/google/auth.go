package google

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// GetService は指定されたスコープでGoogle APIサービスを取得します
func GetService(credentialsFile string, tokenFile string, scope ...string) (*http.Client, error) {

	// 認証情報を読み込む
	b, err := os.ReadFile(credentialsFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		return nil, err
	}

	// 認証情報を使用してOAuth2コンフィグを取得
	config, err := google.ConfigFromJSON(b, scope...)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
		return nil, err
	}

	// トークンを取得
	client := getClient(config, tokenFile)
	return client, nil
}

// getClient はOAuth2クライアントを取得します
func getClient(config *oauth2.Config, tokenFile string) *http.Client {
	tok, err := tokenFromFile(tokenFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokenFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// tokenFromFile はファイルからトークンを取得します
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// getTokenFromWeb はウェブからトークンを取得します
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// saveToken はトークンをファイルに保存します
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

package google

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/api/forms/v1"
	"google.golang.org/api/option"
)

// FormClient はGoogle Forms APIクライアントを表します
type FormClient struct {
	Service *forms.Service
}

// NewFormClient は新しいGoogle Forms APIクライアントを作成します
// client: 認証済みのHTTPクライアント
func NewFormClient(client *http.Client) (*FormClient, error) {
	srv, err := forms.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("formsクライアントの作成に失敗しました: %v", err)
	}
	return &FormClient{Service: srv}, nil
}

// CreateForm は新しいフォームを作成します
// title: フォームのタイトル
func (client *FormClient) CreateForm(title string) (*forms.Form, error) {
	form := &forms.Form{
		Info: &forms.Info{
			Title: title,
		},
	}
	createdForm, err := client.Service.Forms.Create(form).Do()
	if err != nil {
		return nil, fmt.Errorf("フォームの作成に失敗しました: %v", err)
	}
	return createdForm, nil
}

// GetForm は指定されたフォームIDのフォームを取得します
// formID: フォームのID
func (client *FormClient) GetForm(formID string) (*forms.Form, error) {
	form, err := client.Service.Forms.Get(formID).Do()
	if err != nil {
		return nil, fmt.Errorf("フォームの取得に失敗しました: %v", err)
	}
	return form, nil
}

// UpdateForm は指定されたフォームを更新します
// formID: フォームのID
// requests: 更新リクエストのスライス
func (client *FormClient) UpdateForm(formID string, requests []*forms.Request) (*forms.Form, error) {
	updateRequest := &forms.BatchUpdateFormRequest{
		Requests: requests,
	}
	_, err := client.Service.Forms.BatchUpdate(formID, updateRequest).Do()
	if err != nil {
		return nil, fmt.Errorf("フォームの更新に失敗しました: %v", err)
	}
	// 更新後のフォームを取得
	updatedForm, err := client.Service.Forms.Get(formID).Do()
	if err != nil {
		return nil, fmt.Errorf("更新後のフォームの取得に失敗しました: %v", err)
	}
	return updatedForm, nil
}

// AddQuestion は指定されたフォームに質問を追加します
// formID: フォームのID
// questionItem: 追加する質問のアイテム
// index: 質問を追加する位置
func (client *FormClient) AddQuestion(formID string, questionItem *forms.Item, index int) error {
	requests := []*forms.Request{
		{
			CreateItem: &forms.CreateItemRequest{
				Item: questionItem,
				Location: &forms.Location{
					Index: int64(index), // 追加する位置を指定
				},
			},
		},
	}
	updateRequest := &forms.BatchUpdateFormRequest{
		Requests: requests,
	}
	_, err := client.Service.Forms.BatchUpdate(formID, updateRequest).Do()
	if err != nil {
		return fmt.Errorf("質問の追加に失敗しました: %v", err)
	}
	return nil
}

// GetFormResponses は指定されたフォームの回答を取得します
// formID: フォームのID
func (client *FormClient) GetFormResponses(formID string) ([]*forms.FormResponse, error) {
	response, err := client.Service.Forms.Responses.List(formID).Do()
	if err != nil {
		return nil, fmt.Errorf("フォームの回答の取得に失敗しました: %v", err)
	}
	return response.Responses, nil
}

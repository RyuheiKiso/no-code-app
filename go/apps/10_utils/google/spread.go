package google

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// GoogleAPIClient はGoogle Sheets APIクライアントを表します
type GoogleAPIClient struct {
	Service *sheets.Service
}

// NewGoogleAPIClient は新しいGoogle Sheets APIクライアントを作成します
// client: HTTPクライアント
func NewGoogleAPIClient(client *http.Client) (*GoogleAPIClient, error) {
	srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Printf("Unable to retrieve Sheets client: %v", err)
		return nil, err
	}

	return &GoogleAPIClient{Service: srv}, nil
}

// ReadSheet は指定されたスプレッドシートIDとシート名、範囲からデータを読み取ります
// spreadsheetID: スプレッドシートのID
// sheetName: シート名
// readRange: 読み取る範囲
func (client *GoogleAPIClient) ReadSheet(spreadsheetID string, sheetName string, readRange string) ([][]interface{}, error) {
	fullRange := fmt.Sprintf("%s!%s", sheetName, readRange)
	resp, err := client.Service.Spreadsheets.Values.Get(spreadsheetID, fullRange).Do()
	if err != nil {
		log.Printf("Unable to retrieve data from sheet: %v", err)
		return nil, err
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
		return nil, nil
	} else {
		fmt.Println("Data retrieved successfully.")
		return resp.Values, nil
	}
}

// WriteSheet は指定されたスプレッドシートIDとシート名、範囲にデータを書き込みます
// spreadsheetID: スプレッドシートのID
// sheetName: シート名
// writeRange: 書き込む範囲
// values: 書き込むデータ
func (client *GoogleAPIClient) WriteSheet(spreadsheetID string, sheetName string, writeRange string, values [][]interface{}) error {
	fullRange := fmt.Sprintf("%s!%s", sheetName, writeRange)
	valueRange := &sheets.ValueRange{
		Values: values,
	}

	_, err := client.Service.Spreadsheets.Values.Update(spreadsheetID, fullRange, valueRange).ValueInputOption("RAW").Do()
	if err != nil {
		log.Printf("Unable to write data to sheet: %v", err)
		return err
	}

	fmt.Println("Data written successfully.")
	return nil
}

// AppendSheet は指定されたスプレッドシートIDとシート名、範囲にデータを追加します
// spreadsheetID: スプレッドシートのID
// sheetName: シート名
// writeRange: 追加する範囲
// values: 追加するデータ
func (client *GoogleAPIClient) AppendSheet(spreadsheetID string, sheetName string, writeRange string, values [][]interface{}) error {
	fullRange := fmt.Sprintf("%s!%s", sheetName, writeRange)
	valueRange := &sheets.ValueRange{
		Values: values,
	}

	_, err := client.Service.Spreadsheets.Values.Append(spreadsheetID, fullRange, valueRange).ValueInputOption("RAW").InsertDataOption("INSERT_ROWS").Do()
	if err != nil {
		log.Printf("Unable to append data to sheet: %v", err)
		return err
	}

	fmt.Println("Data appended successfully.")
	return nil
}

// AddSheet は指定されたスプレッドシートIDに新しいシートを追加します
// spreadsheetID: スプレッドシートのID
// sheetName: 新しいシートの名前
func (client *GoogleAPIClient) AddSheet(spreadsheetID string, sheetName string) error {
	request := &sheets.Request{
		AddSheet: &sheets.AddSheetRequest{
			Properties: &sheets.SheetProperties{
				Title: sheetName,
			},
		},
	}

	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{request},
	}

	_, err := client.Service.Spreadsheets.BatchUpdate(spreadsheetID, batchUpdateRequest).Do()
	if err != nil {
		log.Printf("Unable to add sheet: %v", err)
		return err
	}

	fmt.Println("Sheet added successfully.")
	return nil
}

// DeleteSheet は指定されたスプレッドシートIDからシートを削除します
// spreadsheetID: スプレッドシートのID
// sheetID: 削除するシートのID
func (client *GoogleAPIClient) DeleteSheet(spreadsheetID string, sheetID int64) error {
	request := &sheets.Request{
		DeleteSheet: &sheets.DeleteSheetRequest{
			SheetId: sheetID,
		},
	}

	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{request},
	}

	_, err := client.Service.Spreadsheets.BatchUpdate(spreadsheetID, batchUpdateRequest).Do()
	if err != nil {
		log.Printf("Unable to delete sheet: %v", err)
		return err
	}

	fmt.Println("Sheet deleted successfully.")
	return nil
}

// RenameSheet は指定されたスプレッドシートIDのシートの名前を変更します
// spreadsheetID: スプレッドシートのID
// sheetID: 名前を変更するシートのID
// newSheetName: 新しいシートの名前
func (client *GoogleAPIClient) RenameSheet(spreadsheetID string, sheetID int64, newSheetName string) error {
	request := &sheets.Request{
		UpdateSheetProperties: &sheets.UpdateSheetPropertiesRequest{
			Properties: &sheets.SheetProperties{
				SheetId: sheetID,
				Title:   newSheetName,
			},
			Fields: "title",
		},
	}

	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{request},
	}

	_, err := client.Service.Spreadsheets.BatchUpdate(spreadsheetID, batchUpdateRequest).Do()
	if err != nil {
		log.Printf("Unable to rename sheet: %v", err)
		return err
	}

	fmt.Println("Sheet renamed successfully.")
	return nil
}

// ClearSheet は指定されたスプレッドシートIDとシート名、範囲のデータをクリアします
// spreadsheetID: スプレッドシートのID
// sheetName: シート名
// clearRange: クリアする範囲
func (client *GoogleAPIClient) ClearSheet(spreadsheetID string, sheetName string, clearRange string) error {
	fullRange := fmt.Sprintf("%s!%s", sheetName, clearRange)
	_, err := client.Service.Spreadsheets.Values.Clear(spreadsheetID, fullRange, &sheets.ClearValuesRequest{}).Do()
	if err != nil {
		log.Printf("Unable to clear data from sheet: %v", err)
		return err
	}

	fmt.Println("Data cleared successfully.")
	return nil
}

// CopySheet は指定されたスプレッドシートIDから別のスプレッドシートIDにシートをコピーします
// spreadsheetID: コピー元のスプレッドシートのID
// sheetID: コピーするシートのID
// destinationSpreadsheetID: コピー先のスプレッドシートのID
func (client *GoogleAPIClient) CopySheet(spreadsheetID string, sheetID int64, destinationSpreadsheetID string) error {
	copyRequest := &sheets.CopySheetToAnotherSpreadsheetRequest{
		DestinationSpreadsheetId: destinationSpreadsheetID,
	}

	_, err := client.Service.Spreadsheets.Sheets.CopyTo(spreadsheetID, sheetID, copyRequest).Do()
	if err != nil {
		log.Printf("Unable to copy sheet: %v", err)
		return err
	}

	fmt.Println("Sheet copied successfully.")
	return nil
}

// GetSheetProperties は指定されたスプレッドシートIDのシートプロパティを取得します
// spreadsheetID: スプレッドシートのID
func (client *GoogleAPIClient) GetSheetProperties(spreadsheetID string) ([]*sheets.SheetProperties, error) {
	resp, err := client.Service.Spreadsheets.Get(spreadsheetID).Do()
	if err != nil {
		log.Printf("Unable to retrieve sheet properties: %v", err)
		return nil, err
	}

	var sheetProperties []*sheets.SheetProperties
	for _, sheet := range resp.Sheets {
		sheetProperties = append(sheetProperties, sheet.Properties)
	}

	fmt.Println("Sheet properties retrieved successfully.")
	return sheetProperties, nil
}

// BatchUpdate は指定されたスプレッドシートIDに対してバッチ更新を実行します
// spreadsheetID: スプレッドシートのID
// requests: バッチ更新リクエストのリスト
func (client *GoogleAPIClient) BatchUpdate(spreadsheetID string, requests []*sheets.Request) error {
	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: requests,
	}

	_, err := client.Service.Spreadsheets.BatchUpdate(spreadsheetID, batchUpdateRequest).Do()
	if err != nil {
		log.Printf("Unable to execute batch update: %v", err)
		return err
	}

	fmt.Println("Batch update executed successfully.")
	return nil
}

// SortSheet は指定されたスプレッドシートIDとシートIDのデータをソートします
// spreadsheetID: スプレッドシートのID
// sheetID: ソートするシートのID
// sortSpecs: ソート仕様のリスト
func (client *GoogleAPIClient) SortSheet(spreadsheetID string, sheetID int64, sortSpecs []*sheets.SortSpec) error {
	request := &sheets.Request{
		SortRange: &sheets.SortRangeRequest{
			Range: &sheets.GridRange{
				SheetId: sheetID,
			},
			SortSpecs: sortSpecs,
		},
	}

	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{request},
	}

	_, err := client.Service.Spreadsheets.BatchUpdate(spreadsheetID, batchUpdateRequest).Do()
	if err != nil {
		log.Printf("Unable to sort sheet: %v", err)
		return err
	}

	fmt.Println("Sheet sorted successfully.")
	return nil
}

// MergeCells は指定されたスプレッドシートIDとシートIDのセルをマージします
// spreadsheetID: スプレッドシートのID
// sheetID: マージするシートのID
// startRowIndex: マージ開始行インデックス
// endRowIndex: マージ終了行インデックス
// startColumnIndex: マージ開始列インデックス
// endColumnIndex: マージ終了列インデックス
func (client *GoogleAPIClient) MergeCells(spreadsheetID string, sheetID int64, startRowIndex int64, endRowIndex int64, startColumnIndex int64, endColumnIndex int64) error {
	request := &sheets.Request{
		MergeCells: &sheets.MergeCellsRequest{
			Range: &sheets.GridRange{
				SheetId:          sheetID,
				StartRowIndex:    startRowIndex,
				EndRowIndex:      endRowIndex,
				StartColumnIndex: startColumnIndex,
				EndColumnIndex:   endColumnIndex,
			},
			MergeType: "MERGE_ALL",
		},
	}

	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{request},
	}

	_, err := client.Service.Spreadsheets.BatchUpdate(spreadsheetID, batchUpdateRequest).Do()
	if err != nil {
		log.Printf("Unable to merge cells: %v", err)
		return err
	}

	fmt.Println("Cells merged successfully.")
	return nil
}

// ApplyConditionalFormatting は指定されたスプレッドシートIDとシートIDに条件付き書式を適用します
// spreadsheetID: スプレッドシートのID
// sheetID: シートのID
// startRowIndex: 開始行インデックス
// endRowIndex: 終了行インデックス
// startColumnIndex: 開始列インデックス
// endColumnIndex: 終了列インデックス
// conditionType: 条件の種類
// conditionValues: 条件の値のリスト
// format: 適用する書式
func (client *GoogleAPIClient) ApplyConditionalFormatting(spreadsheetID string, sheetID int64, startRowIndex int64, endRowIndex int64, startColumnIndex int64, endColumnIndex int64, conditionType string, conditionValues []string, format *sheets.CellFormat) error {
	request := &sheets.Request{
		AddConditionalFormatRule: &sheets.AddConditionalFormatRuleRequest{
			Rule: &sheets.ConditionalFormatRule{
				Ranges: []*sheets.GridRange{
					{
						SheetId:          sheetID,
						StartRowIndex:    startRowIndex,
						EndRowIndex:      endRowIndex,
						StartColumnIndex: startColumnIndex,
						EndColumnIndex:   endColumnIndex,
					},
				},
				BooleanRule: &sheets.BooleanRule{
					Condition: &sheets.BooleanCondition{
						Type:   conditionType,
						Values: []*sheets.ConditionValue{{UserEnteredValue: conditionValues[0]}},
					},
					Format: format,
				},
			},
			Index: 0,
		},
	}

	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{request},
	}

	_, err := client.Service.Spreadsheets.BatchUpdate(spreadsheetID, batchUpdateRequest).Do()
	if err != nil {
		log.Printf("Unable to apply conditional formatting: %v", err)
		return err
	}

	fmt.Println("Conditional formatting applied successfully.")
	return nil
}

// ApplyFilter は指定されたスプレッドシートIDとシートIDにフィルタを適用します
// spreadsheetID: スプレッドシートのID
// sheetID: シートのID
// startRowIndex: 開始行インデックス
// endRowIndex: 終了行インデックス
// startColumnIndex: 開始列インデックス
// endColumnIndex: 終了列インデックス
func (client *GoogleAPIClient) ApplyFilter(spreadsheetID string, sheetID int64, startRowIndex int64, endRowIndex int64, startColumnIndex int64, endColumnIndex int64) error {
	request := &sheets.Request{
		SetBasicFilter: &sheets.SetBasicFilterRequest{
			Filter: &sheets.BasicFilter{
				Range: &sheets.GridRange{
					SheetId:          sheetID,
					StartRowIndex:    startRowIndex,
					EndRowIndex:      endRowIndex,
					StartColumnIndex: startColumnIndex,
					EndColumnIndex:   endColumnIndex,
				},
			},
		},
	}

	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{request},
	}

	_, err := client.Service.Spreadsheets.BatchUpdate(spreadsheetID, batchUpdateRequest).Do()
	if err != nil {
		log.Printf("Unable to apply filter: %v", err)
		return err
	}

	fmt.Println("Filter applied successfully.")
	return nil
}

// ExportSheetToCSV は指定されたスプレッドシートIDとシート名、範囲のデータをCSVファイルにエクスポートします
// spreadsheetID: スプレッドシートのID
// sheetName: シート名
// readRange: 読み取る範囲
// filePath: エクスポートするCSVファイルのパス
func (client *GoogleAPIClient) ExportSheetToCSV(spreadsheetID string, sheetName string, readRange string, filePath string) error {
	data, err := client.ReadSheet(spreadsheetID, sheetName, readRange)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Unable to create CSV file: %v", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range data {
		stringRow := make([]string, len(row))
		for i, cell := range row {
			stringRow[i] = fmt.Sprintf("%v", cell)
		}
		if err := writer.Write(stringRow); err != nil {
			log.Printf("Unable to write row to CSV: %v", err)
			return err
		}
	}

	fmt.Println("Data exported to CSV successfully.")
	return nil
}

// AddCommentToCell は指定されたセルにコメントを追加します
// service: Google Sheets APIサービス
// spreadsheetId: スプレッドシートのID
// sheetId: シートのID
// cell: セルのアドレス（例: "A1"）
// comment: 追加するコメント
func AddCommentToCell(service *sheets.Service, spreadsheetId string, sheetId int64, cell string, comment string) error {
	request := &sheets.Request{
		UpdateCells: &sheets.UpdateCellsRequest{
			Range: &sheets.GridRange{
				SheetId:          sheetId,
				StartRowIndex:    cellRow(cell),
				EndRowIndex:      cellRow(cell) + 1,
				StartColumnIndex: cellColumn(cell),
				EndColumnIndex:   cellColumn(cell) + 1,
			},
			Rows: []*sheets.RowData{
				{
					Values: []*sheets.CellData{
						{
							Note: comment,
						},
					},
				},
			},
			Fields: "note",
		},
	}

	_, err := service.Spreadsheets.BatchUpdate(spreadsheetId, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{request},
	}).Context(context.Background()).Do()
	if err != nil {
		return fmt.Errorf("unable to add comment to cell: %v", err)
	}
	return nil
}

// GetCommentFromCell は指定されたセルからコメントを取得します
// service: Google Sheets APIサービス
// spreadsheetId: スプレッドシートのID
// sheetId: シートのID
// cell: セルのアドレス（例: "A1"）
func GetCommentFromCell(service *sheets.Service, spreadsheetId string, sheetId int64, cell string) (string, error) {
	resp, err := service.Spreadsheets.Get(spreadsheetId).Ranges(cell).IncludeGridData(true).Context(context.Background()).Do()
	if err != nil {
		return "", fmt.Errorf("unable to retrieve comment from cell: %v", err)
	}

	for _, sheet := range resp.Sheets {
		for _, rowData := range sheet.Data[0].RowData {
			for _, cellData := range rowData.Values {
				if cellData.Note != "" {
					return cellData.Note, nil
				}
			}
		}
	}
	return "", nil
}

// SetCellBackgroundColor は指定されたセルの背景色を設定します
// service: Google Sheets APIサービス
// spreadsheetId: スプレッドシートのID
// sheetId: シートのID
// cell: セルのアドレス（例: "A1"）
// red: 背景色の赤成分（0.0〜1.0）
// green: 背景色の緑成分（0.0〜1.0）
// blue: 背景色の青成分（0.0〜1.0）
func SetCellBackgroundColor(service *sheets.Service, spreadsheetId string, sheetId int64, cell string, red, green, blue float64) error {
	request := &sheets.Request{
		UpdateCells: &sheets.UpdateCellsRequest{
			Range: &sheets.GridRange{
				SheetId:          sheetId,
				StartRowIndex:    cellRow(cell),
				EndRowIndex:      cellRow(cell) + 1,
				StartColumnIndex: cellColumn(cell),
				EndColumnIndex:   cellColumn(cell) + 1,
			},
			Rows: []*sheets.RowData{
				{
					Values: []*sheets.CellData{
						{
							UserEnteredFormat: &sheets.CellFormat{
								BackgroundColor: &sheets.Color{
									Red:   red,
									Green: green,
									Blue:  blue,
								},
							},
						},
					},
				},
			},
			Fields: "userEnteredFormat.backgroundColor",
		},
	}

	_, err := service.Spreadsheets.BatchUpdate(spreadsheetId, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{request},
	}).Context(context.Background()).Do()
	if err != nil {
		return fmt.Errorf("unable to set cell background color: %v", err)
	}
	return nil
}

// Helper functions to get row and column indices from cell notation (e.g., "A1").
// cellRow はセルの行インデックスを取得します
// cell: セルのアドレス（例: "A1"）
func cellRow(cell string) int64 {
	row, _ := strconv.Atoi(strings.TrimLeft(cell, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
	return int64(row - 1)
}

// cellColumn はセルの列インデックスを取得します
// cell: セルのアドレス（例: "A1"）
func cellColumn(cell string) int64 {
	column := strings.TrimRight(cell, "0123456789")
	col := 0
	for i := 0; i < len(column); i++ {
		col = col*26 + int(column[i]-'A'+1)
	}
	return int64(col - 1)
}

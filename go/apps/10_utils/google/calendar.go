package google

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// CalendarClient はGoogle Calendar APIクライアントを表します
type CalendarClient struct {
	Service *calendar.Service
}

// NewCalendarClient は新しいGoogle Calendar APIクライアントを作成します
// client: 認証済みのHTTPクライアント
func NewCalendarClient(client *http.Client) (*CalendarClient, error) {
	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
		return nil, err
	}

	return &CalendarClient{Service: srv}, nil
}

// CreateEvent は新しいカレンダーイベントを作成します
// calendarID: カレンダーのID
// event: 作成するイベント
func (client *CalendarClient) CreateEvent(calendarID string, event *calendar.Event) (*calendar.Event, error) {
	createdEvent, err := client.Service.Events.Insert(calendarID, event).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to create event: %v", err)
	}
	fmt.Println("Event created successfully.")
	return createdEvent, nil
}

// GetEvent は指定されたイベントIDのイベントを取得します
// calendarID: カレンダーのID
// eventID: 取得するイベントのID
func (client *CalendarClient) GetEvent(calendarID string, eventID string) (*calendar.Event, error) {
	event, err := client.Service.Events.Get(calendarID, eventID).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve event: %v", err)
	}
	return event, nil
}

// UpdateEvent は指定されたイベントを更新します
// calendarID: カレンダーのID
// eventID: 更新するイベントのID
// event: 更新するイベントの内容
func (client *CalendarClient) UpdateEvent(calendarID string, eventID string, event *calendar.Event) (*calendar.Event, error) {
	updatedEvent, err := client.Service.Events.Update(calendarID, eventID, event).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to update event: %v", err)
	}
	fmt.Println("Event updated successfully.")
	return updatedEvent, nil
}

// DeleteEvent は指定されたイベントを削除します
// calendarID: カレンダーのID
// eventID: 削除するイベントのID
func (client *CalendarClient) DeleteEvent(calendarID string, eventID string) error {
	err := client.Service.Events.Delete(calendarID, eventID).Do()
	if err != nil {
		return fmt.Errorf("unable to delete event: %v", err)
	}
	fmt.Println("Event deleted successfully.")
	return nil
}

// ListEvents は指定されたカレンダーのイベントをリストします
// calendarID: カレンダーのID
// timeMin: 開始時間の最小値
// timeMax: 終了時間の最大値
func (client *CalendarClient) ListEvents(calendarID string, timeMin time.Time, timeMax time.Time) ([]*calendar.Event, error) {
	events, err := client.Service.Events.List(calendarID).ShowDeleted(false).
		SingleEvents(true).TimeMin(timeMin.Format(time.RFC3339)).TimeMax(timeMax.Format(time.RFC3339)).OrderBy("startTime").Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve events: %v", err)
	}

	return events.Items, nil
}

// WatchEvents はカレンダーのイベントの変更を監視します
// calendarID: カレンダーのID
// channelID: チャンネルのID
// resourceID: リソースのID
// token: トークン
// expiration: 有効期限
func (client *CalendarClient) WatchEvents(calendarID string, channelID string, resourceID string, token string, expiration int64) (*calendar.Channel, error) {
	channel := &calendar.Channel{
		Id:         channelID,
		ResourceId: resourceID,
		Token:      token,
		Expiration: expiration,
	}

	watchCall := client.Service.Events.Watch(calendarID, channel)
	watchChannel, err := watchCall.Do()
	if err != nil {
		return nil, fmt.Errorf("unable to watch events: %v", err)
	}

	fmt.Println("Watching events successfully.")
	return watchChannel, nil
}

// StopWatchingEvents はカレンダーのイベントの監視を停止します
// channelID: チャンネルのID
// resourceID: リソースのID
func (client *CalendarClient) StopWatchingEvents(channelID string, resourceID string) error {
	channel := &calendar.Channel{
		Id:         channelID,
		ResourceId: resourceID,
	}

	err := client.Service.Channels.Stop(channel).Do()
	if err != nil {
		return fmt.Errorf("unable to stop watching events: %v", err)
	}

	fmt.Println("Stopped watching events successfully.")
	return nil
}

// CreateRecurringEvent は新しい繰り返しカレンダーイベントを作成します
// calendarID: カレンダーのID
// event: 作成するイベント
func (client *CalendarClient) CreateRecurringEvent(calendarID string, event *calendar.Event) (*calendar.Event, error) {
	createdEvent, err := client.Service.Events.Insert(calendarID, event).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to create recurring event: %v", err)
	}
	fmt.Println("Recurring event created successfully.")
	return createdEvent, nil
}

// GetCalendarList はカレンダーリストを取得します
func (client *CalendarClient) GetCalendarList() ([]*calendar.CalendarListEntry, error) {
	calendarList, err := client.Service.CalendarList.List().Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve calendar list: %v", err)
	}

	return calendarList.Items, nil
}

// AddAttendee は指定されたイベントに参加者を追加します
// calendarID: カレンダーのID
// eventID: イベントのID
// attendeeEmail: 追加する参加者のメールアドレス
func (client *CalendarClient) AddAttendee(calendarID string, eventID string, attendeeEmail string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	attendee := &calendar.EventAttendee{Email: attendeeEmail}
	event.Attendees = append(event.Attendees, attendee)

	updatedEvent, err := client.UpdateEvent(calendarID, eventID, event)
	if err != nil {
		return nil, err
	}

	fmt.Println("Attendee added successfully.")
	return updatedEvent, nil
}

// RemoveAttendee は指定されたイベントから参加者を削除します
// calendarID: カレンダーのID
// eventID: イベントのID
// attendeeEmail: 削除する参加者のメールアドレス
func (client *CalendarClient) RemoveAttendee(calendarID string, eventID string, attendeeEmail string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	for i, attendee := range event.Attendees {
		if attendee.Email == attendeeEmail {
			event.Attendees = append(event.Attendees[:i], event.Attendees[i+1:]...)
			break
		}
	}

	updatedEvent, err := client.UpdateEvent(calendarID, eventID, event)
	if err != nil {
		return nil, err
	}

	fmt.Println("Attendee removed successfully.")
	return updatedEvent, nil
}

// GetEventInstances は指定された繰り返しイベントのインスタンスを取得します
// calendarID: カレンダーのID
// eventID: 繰り返しイベントのID
// timeMin: 開始時間の最小値
// timeMax: 終了時間の最大値
func (client *CalendarClient) GetEventInstances(calendarID string, eventID string, timeMin time.Time, timeMax time.Time) ([]*calendar.Event, error) {
	instances, err := client.Service.Events.Instances(calendarID, eventID).TimeMin(timeMin.Format(time.RFC3339)).TimeMax(timeMax.Format(time.RFC3339)).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve event instances: %v", err)
	}

	return instances.Items, nil
}

// MoveEvent は指定されたイベントを別のカレンダーに移動します
// calendarID: 元のカレンダーのID
// eventID: 移動するイベントのID
// destinationCalendarID: 移動先のカレンダーのID
func (client *CalendarClient) MoveEvent(calendarID string, eventID string, destinationCalendarID string) (*calendar.Event, error) {
	movedEvent, err := client.Service.Events.Move(calendarID, eventID, destinationCalendarID).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to move event: %v", err)
	}

	fmt.Println("Event moved successfully.")
	return movedEvent, nil
}

// QuickAddEvent はクイック追加文字列を使用して新しいイベントを作成します
// calendarID: カレンダーのID
// text: クイック追加文字列
func (client *CalendarClient) QuickAddEvent(calendarID string, text string) (*calendar.Event, error) {
	event, err := client.Service.Events.QuickAdd(calendarID, text).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to quick add event: %v", err)
	}

	fmt.Println("Event added successfully using quick add.")
	return event, nil
}

// SetEventReminder は指定されたイベントにリマインダーを設定します
// calendarID: カレンダーのID
// eventID: イベントのID
// method: リマインダーの方法（例: "email", "popup"）
// minutes: リマインダーを送信する時間（分単位）
func (client *CalendarClient) SetEventReminder(calendarID string, eventID string, method string, minutes int64) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	reminder := &calendar.EventReminder{
		Method:  method,
		Minutes: minutes,
	}
	event.Reminders = &calendar.EventReminders{
		UseDefault: false,
		Overrides:  []*calendar.EventReminder{reminder},
	}

	updatedEvent, err := client.UpdateEvent(calendarID, eventID, event)
	if err != nil {
		return nil, err
	}

	fmt.Println("Reminder set successfully.")
	return updatedEvent, nil
}

// GetPrimaryCalendarID はプライマリカレンダーのIDを取得します
func (client *CalendarClient) GetPrimaryCalendarID() (string, error) {
	calendarList, err := client.GetCalendarList()
	if err != nil {
		return "", err
	}

	for _, calendar := range calendarList {
		if calendar.Primary {
			return calendar.Id, nil
		}
	}

	return "", fmt.Errorf("primary calendar not found")
}

// CreateCalendar は新しいカレンダーを作成します
// summary: カレンダーの概要
func (client *CalendarClient) CreateCalendar(summary string) (*calendar.Calendar, error) {
	calendar := &calendar.Calendar{
		Summary: summary,
	}

	createdCalendar, err := client.Service.Calendars.Insert(calendar).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to create calendar: %v", err)
	}

	fmt.Println("Calendar created successfully.")
	return createdCalendar, nil
}

// DeleteCalendar は指定されたカレンダーを削除します
// calendarID: 削除するカレンダーのID
func (client *CalendarClient) DeleteCalendar(calendarID string) error {
	err := client.Service.Calendars.Delete(calendarID).Do()
	if err != nil {
		return fmt.Errorf("unable to delete calendar: %v", err)
	}

	fmt.Println("Calendar deleted successfully.")
	return nil
}

// UpdateCalendar は指定されたカレンダーを更新します
// calendarID: 更新するカレンダーのID
// summary: 新しいカレンダーの概要
func (client *CalendarClient) UpdateCalendar(calendarID string, summary string) (*calendar.Calendar, error) {
	calendar := &calendar.Calendar{
		Summary: summary,
	}

	updatedCalendar, err := client.Service.Calendars.Update(calendarID, calendar).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to update calendar: %v", err)
	}

	fmt.Println("Calendar updated successfully.")
	return updatedCalendar, nil
}

// GetFreeBusy は指定された時間範囲内の空き時間と忙しい時間を取得します
// calendarIDs: カレンダーのIDのリスト
// timeMin: 開始時間の最小値
// timeMax: 終了時間の最大値
func (client *CalendarClient) GetFreeBusy(calendarIDs []string, timeMin time.Time, timeMax time.Time) (*calendar.FreeBusyResponse, error) {
	request := &calendar.FreeBusyRequest{
		TimeMin: timeMin.Format(time.RFC3339),
		TimeMax: timeMax.Format(time.RFC3339),
		Items:   make([]*calendar.FreeBusyRequestItem, len(calendarIDs)),
	}

	for i, calendarID := range calendarIDs {
		request.Items[i] = &calendar.FreeBusyRequestItem{Id: calendarID}
	}

	response, err := client.Service.Freebusy.Query(request).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve free/busy information: %v", err)
	}

	return response, nil
}

// GetEventColors はカレンダーイベントの色を取得します
// client: CalendarClientのインスタンス
func (client *CalendarClient) GetEventColors() (map[string]*calendar.ColorDefinition, error) {
	colors, err := client.Service.Colors.Get().Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve event colors: %v", err)
	}

	eventColors := make(map[string]*calendar.ColorDefinition)
	for key, value := range colors.Event {
		v := value
		eventColors[key] = &v
	}

	return eventColors, nil
}

// SetEventColor は指定されたイベントに色を設定します
// calendarID: カレンダーのID
// eventID: イベントのID
// colorID: 設定する色のID
func (client *CalendarClient) SetEventColor(calendarID string, eventID string, colorID string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	event.ColorId = colorID

	updatedEvent, err := client.UpdateEvent(calendarID, eventID, event)
	if err != nil {
		return nil, err
	}

	fmt.Println("Event color set successfully.")
	return updatedEvent, nil
}

// GetCalendarACL は指定されたカレンダーのアクセス制御リストを取得します
// calendarID: カレンダーのID
func (client *CalendarClient) GetCalendarACL(calendarID string) ([]*calendar.AclRule, error) {
	acl, err := client.Service.Acl.List(calendarID).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve calendar ACL: %v", err)
	}

	return acl.Items, nil
}

// AddCalendarACL は指定されたカレンダーにアクセス制御ルールを追加します
// calendarID: カレンダーのID
// rule: 追加するアクセス制御ルール
func (client *CalendarClient) AddCalendarACL(calendarID string, rule *calendar.AclRule) (*calendar.AclRule, error) {
	createdRule, err := client.Service.Acl.Insert(calendarID, rule).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to add calendar ACL: %v", err)
	}

	fmt.Println("Calendar ACL added successfully.")
	return createdRule, nil
}

// DeleteCalendarACL は指定されたカレンダーからアクセス制御ルールを削除します
// calendarID: カレンダーのID
// ruleID: 削除するアクセス制御ルールのID
func (client *CalendarClient) DeleteCalendarACL(calendarID string, ruleID string) error {
	err := client.Service.Acl.Delete(calendarID, ruleID).Do()
	if err != nil {
		return fmt.Errorf("unable to delete calendar ACL: %v", err)
	}

	fmt.Println("Calendar ACL deleted successfully.")
	return nil
}

// UpdateCalendarACL は指定されたカレンダーのアクセス制御ルールを更新します
// calendarID: カレンダーのID
// ruleID: 更新するアクセス制御ルールのID
// rule: 更新するアクセス制御ルール
func (client *CalendarClient) UpdateCalendarACL(calendarID string, ruleID string, rule *calendar.AclRule) (*calendar.AclRule, error) {
	updatedRule, err := client.Service.Acl.Update(calendarID, ruleID, rule).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to update calendar ACL: %v", err)
	}

	fmt.Println("Calendar ACL updated successfully.")
	return updatedRule, nil
}

// GetEventAttachments は指定されたイベントの添付ファイルを取得します
// calendarID: カレンダーのID
// eventID: イベントのID
func (client *CalendarClient) GetEventAttachments(calendarID string, eventID string) ([]*calendar.EventAttachment, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	return event.Attachments, nil
}

// AddEventAttachment は指定されたイベントに添付ファイルを追加します
// calendarID: カレンダーのID
// eventID: イベントのID
// attachment: 追加する添付ファイル
func (client *CalendarClient) AddEventAttachment(calendarID string, eventID string, attachment *calendar.EventAttachment) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	event.Attachments = append(event.Attachments, attachment)

	updatedEvent, err := client.UpdateEvent(calendarID, eventID, event)
	if err != nil {
		return nil, err
	}

	fmt.Println("Attachment added successfully.")
	return updatedEvent, nil
}

// RemoveEventAttachment は指定されたイベントから添付ファイルを削除します
// calendarID: カレンダーのID
// eventID: イベントのID
// fileID: 削除する添付ファイルのID
func (client *CalendarClient) RemoveEventAttachment(calendarID string, eventID string, fileID string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	for i, attachment := range event.Attachments {
		if attachment.FileId == fileID {
			event.Attachments = append(event.Attachments[:i], event.Attachments[i+1:]...)
			break
		}
	}

	updatedEvent, err := client.UpdateEvent(calendarID, eventID, event)
	if err != nil {
		return nil, err
	}

	fmt.Println("Attachment removed successfully.")
	return updatedEvent, nil
}

// GetEventAttendees は指定されたイベントの参加者リストを取得します
// calendarID: カレンダーのID
// eventID: イベントのID
func (client *CalendarClient) GetEventAttendees(calendarID string, eventID string) ([]*calendar.EventAttendee, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	return event.Attendees, nil
}

// UpdateEventAttendee は指定されたイベントの参加者情報を更新します
// calendarID: カレンダーのID
// eventID: イベントのID
// attendee: 更新する参加者情報
func (client *CalendarClient) UpdateEventAttendee(calendarID string, eventID string, attendee *calendar.EventAttendee) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	for i, a := range event.Attendees {
		if a.Email == attendee.Email {
			event.Attendees[i] = attendee
			break
		}
	}

	updatedEvent, err := client.UpdateEvent(calendarID, eventID, event)
	if err != nil {
		return nil, err
	}

	fmt.Println("Attendee updated successfully.")
	return updatedEvent, nil
}

// GetEventBySummary は指定されたカレンダー内のイベントを概要で検索します
// calendarID: カレンダーのID
// summary: 検索するイベントの概要
func (client *CalendarClient) GetEventBySummary(calendarID string, summary string) ([]*calendar.Event, error) {
	events, err := client.Service.Events.List(calendarID).Q(summary).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve events by summary: %v", err)
	}

	return events.Items, nil
}

// UpdateEventLocation は指定されたイベントの場所を更新します
// calendarID: カレンダーのID
// eventID: イベントのID
// location: 新しい場所
func (client *CalendarClient) UpdateEventLocation(calendarID string, eventID string, location string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	event.Location = location

	updatedEvent, err := client.UpdateEvent(calendarID, eventID, event)
	if err != nil {
		return nil, err
	}

	fmt.Println("Event location updated successfully.")
	return updatedEvent, nil
}

// GetEventByDateRange は指定された日付範囲内のイベントを取得します
// calendarID: カレンダーのID
// startDate: 開始日
// endDate: 終了日
func (client *CalendarClient) GetEventByDateRange(calendarID string, startDate time.Time, endDate time.Time) ([]*calendar.Event, error) {
	events, err := client.Service.Events.List(calendarID).TimeMin(startDate.Format(time.RFC3339)).TimeMax(endDate.Format(time.RFC3339)).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve events by date range: %v", err)
	}

	return events.Items, nil
}

// UpdateEventDescription は指定されたイベントの説明を更新します
// calendarID: カレンダーのID
// eventID: イベントのID
// description: 新しい説明
func (client *CalendarClient) UpdateEventDescription(calendarID string, eventID string, description string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	event.Description = description

	updatedEvent, err := client.UpdateEvent(calendarID, eventID, event)
	if err != nil {
		return nil, err
	}

	fmt.Println("Event description updated successfully.")
	return updatedEvent, nil
}

// GetEventByLocation は指定されたカレンダー内のイベントを場所で検索します
// calendarID: カレンダーのID
// location: 検索するイベントの場所
func (client *CalendarClient) GetEventByLocation(calendarID string, location string) ([]*calendar.Event, error) {
	events, err := client.Service.Events.List(calendarID).Q(location).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve events by location: %v", err)
	}

	return events.Items, nil
}

// GetEventByAttendee は指定されたカレンダー内のイベントを参加者で検索します
// calendarID: カレンダーのID
// attendeeEmail: 検索する参加者のメールアドレス
func (client *CalendarClient) GetEventByAttendee(calendarID string, attendeeEmail string) ([]*calendar.Event, error) {
	events, err := client.Service.Events.List(calendarID).Q(attendeeEmail).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve events by attendee: %v", err)
	}

	return events.Items, nil
}

// GetEventByIDAndSummary は指定されたイベントIDと概要でイベントを検索します
// calendarID: カレンダーのID
// eventID: イベントのID
// summary: 検索するイベントの概要
func (client *CalendarClient) GetEventByIDAndSummary(calendarID string, eventID string, summary string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	if event.Summary != summary {
		return nil, fmt.Errorf("event summary does not match")
	}

	return event, nil
}

// GetEventByIDAndLocation は指定されたイベントIDと場所でイベントを検索します
// calendarID: カレンダーのID
// eventID: イベントのID
// location: 検索するイベントの場所
func (client *CalendarClient) GetEventByIDAndLocation(calendarID string, eventID string, location string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	if event.Location != location {
		return nil, fmt.Errorf("event location does not match")
	}

	return event, nil
}

// GetEventByIDAndAttendee は指定されたイベントIDと参加者でイベントを検索します
// calendarID: カレンダーのID
// eventID: イベントのID
// attendeeEmail: 検索する参加者のメールアドレス
func (client *CalendarClient) GetEventByIDAndAttendee(calendarID string, eventID string, attendeeEmail string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	for _, attendee := range event.Attendees {
		if attendee.Email == attendeeEmail {
			return event, nil
		}
	}

	return nil, fmt.Errorf("attendee not found in event")
}

// GetEventByIDAndDateRange は指定されたイベントIDと日付範囲でイベントを検索します
// calendarID: カレンダーのID
// eventID: イベントのID
// startDate: 開始日
// endDate: 終了日
func (client *CalendarClient) GetEventByIDAndDateRange(calendarID string, eventID string, startDate time.Time, endDate time.Time) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	eventStart, err := time.Parse(time.RFC3339, event.Start.DateTime)
	if err != nil {
		return nil, err
	}

	eventEnd, err := time.Parse(time.RFC3339, event.End.DateTime)
	if err != nil {
		return nil, err
	}

	if eventStart.After(startDate) && eventEnd.Before(endDate) {
		return event, nil
	}

	return nil, fmt.Errorf("event does not fall within the specified date range")
}

// GetEventByIDAndDescription は指定されたイベントIDと説明でイベントを検索します
// calendarID: カレンダーのID
// eventID: イベントのID
// description: 検索するイベントの説明
func (client *CalendarClient) GetEventByIDAndDescription(calendarID string, eventID string, description string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	if event.Description != description {
		return nil, fmt.Errorf("event description does not match")
	}

	return event, nil
}

// GetEventByIDAndColor は指定されたイベントIDと色でイベントを検索します
// calendarID: カレンダーのID
// eventID: イベントのID
// colorID: 検索するイベントの色ID
func (client *CalendarClient) GetEventByIDAndColor(calendarID string, eventID string, colorID string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	if event.ColorId != colorID {
		return nil, fmt.Errorf("event color does not match")
	}

	return event, nil
}

// GetEventByIDAndReminder は指定されたイベントIDとリマインダーでイベントを検索します
// calendarID: カレンダーのID
// eventID: イベントのID
// method: 検索するリマインダーの方法
// minutes: 検索するリマインダーの時間（分単位）
func (client *CalendarClient) GetEventByIDAndReminder(calendarID string, eventID string, method string, minutes int64) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	for _, reminder := range event.Reminders.Overrides {
		if reminder.Method == method && reminder.Minutes == minutes {
			return event, nil
		}
	}

	return nil, fmt.Errorf("reminder not found in event")
}

// UpdateEventSummary は指定されたイベントの概要を更新します
// calendarID: カレンダーのID
// eventID: イベントのID
// summary: 新しい概要
func (client *CalendarClient) UpdateEventSummary(calendarID string, eventID string, summary string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	event.Summary = summary

	updatedEvent, err := client.UpdateEvent(calendarID, eventID, event)
	if err != nil {
		return nil, err
	}

	fmt.Println("Event summary updated successfully.")
	return updatedEvent, nil
}

// GetEventByIDAndOrganizer は指定されたイベントIDとオーガナイザーでイベントを検索します
// calendarID: カレンダーのID
// eventID: イベントのID
// organizerEmail: 検索するオーガナイザーのメールアドレス
func (client *CalendarClient) GetEventByIDAndOrganizer(calendarID string, eventID string, organizerEmail string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	if event.Organizer.Email != organizerEmail {
		return nil, fmt.Errorf("event organizer does not match")
	}

	return event, nil
}

// GetEventByIDAndStatus は指定されたイベントIDとステータスでイベントを検索します
// calendarID: カレンダーのID
// eventID: イベントのID
// status: 検索するイベントのステータス
func (client *CalendarClient) GetEventByIDAndStatus(calendarID string, eventID string, status string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	if event.Status != status {
		return nil, fmt.Errorf("event status does not match")
	}

	return event, nil
}

// GetEventByIDAndVisibility は指定されたイベントIDと可視性でイベントを検索します
// calendarID: カレンダーのID
// eventID: イベントのID
// visibility: 検索するイベントの可視性
func (client *CalendarClient) GetEventByIDAndVisibility(calendarID string, eventID string, visibility string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	if event.Visibility != visibility {
		return nil, fmt.Errorf("event visibility does not match")
	}

	return event, nil
}

// GetEventByIDAndRecurrence は指定されたイベントIDと繰り返し設定でイベントを検索します
// calendarID: カレンダーのID
// eventID: イベントのID
// recurrence: 検索するイベントの繰り返し設定
func (client *CalendarClient) GetEventByIDAndRecurrence(calendarID string, eventID string, recurrence string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	for _, r := range event.Recurrence {
		if r == recurrence {
			return event, nil
		}
	}

	return nil, fmt.Errorf("recurrence not found in event")
}

// GetEventByIDAndHangoutLink は指定されたイベントIDとハングアウトリンクでイベントを検索します
// calendarID: カレンダーのID
// eventID: イベントのID
// hangoutLink: 検索するイベントのハングアウトリンク
func (client *CalendarClient) GetEventByIDAndHangoutLink(calendarID string, eventID string, hangoutLink string) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	if event.HangoutLink != hangoutLink {
		return nil, fmt.Errorf("event hangout link does not match")
	}

	return event, nil
}

// GetEventByIDAndConferenceData は指定されたイベントIDと会議データでイベントを検索します
// calendarID: カレンダーのID
// eventID: イベントのID
// conferenceData: 検索するイベントの会議データ
func (client *CalendarClient) GetEventByIDAndConferenceData(calendarID string, eventID string, conferenceData *calendar.ConferenceData) (*calendar.Event, error) {
	event, err := client.GetEvent(calendarID, eventID)
	if err != nil {
		return nil, err
	}

	if event.ConferenceData == nil || event.ConferenceData.ConferenceId != conferenceData.ConferenceId {
		return nil, fmt.Errorf("event conference data does not match")
	}

	return event, nil
}

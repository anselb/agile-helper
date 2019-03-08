package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"
//
// 	"golang.org/x/net/context"
// 	"golang.org/x/oauth2"
// 	"golang.org/x/oauth2/google"
// 	"google.golang.org/api/calendar/v3"
// )
//
// // Retrieve a token, saves the token, then returns the generated client.
// func getClient(config *oauth2.Config) *http.Client {
// 	// The file token.json stores the user's access and refresh tokens, and is
// 	// created automatically when the authorization flow completes for the first
// 	// time.
// 	tokFile := "token.json"
// 	tok, err := tokenFromFile(tokFile)
// 	if err != nil {
// 		tok = getTokenFromWeb(config)
// 		saveToken(tokFile, tok)
// 	}
// 	return config.Client(context.Background(), tok)
// }
//
// // Request a token from the web, then returns the retrieved token.
// func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
// 	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
// 	fmt.Printf("Go to the following link in your browser then type the "+
// 		"authorization code: \n%v\n", authURL)
//
// 	var authCode string
// 	if _, err := fmt.Scan(&authCode); err != nil {
// 		log.Fatalf("Unable to read authorization code: %v", err)
// 	}
//
// 	tok, err := config.Exchange(context.TODO(), authCode)
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve token from web: %v", err)
// 	}
// 	return tok
// }
//
// // Retrieves a token from a local file.
// func tokenFromFile(file string) (*oauth2.Token, error) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()
// 	tok := &oauth2.Token{}
// 	err = json.NewDecoder(f).Decode(tok)
// 	return tok, err
// }
//
// // Saves a token to a file path.
// func saveToken(path string, token *oauth2.Token) {
// 	fmt.Printf("Saving credential file to: %s\n", path)
// 	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
// 	if err != nil {
// 		log.Fatalf("Unable to cache oauth token: %v", err)
// 	}
// 	defer f.Close()
// 	json.NewEncoder(f).Encode(token)
// }
//
// // Creates a new event
// func newEvent(srv *calendar.Service, summary, desc, startTime, endTime string) {
// 	event := &calendar.Event{
// 		Summary:     summary,
// 		Description: desc,
// 		Start: &calendar.EventDateTime{
// 			DateTime: startTime,
// 			TimeZone: "UTC",
// 		},
// 		End: &calendar.EventDateTime{
// 			DateTime: endTime,
// 			TimeZone: "UTC",
// 		},
// 	}
//
// 	calendarID := "primary"
// 	event, err := srv.Events.Insert(calendarID, event).Do()
// 	if err != nil {
// 		log.Fatalf("Unable to create event. %v\n", err)
// 	}
// 	fmt.Printf("Event created: %s\n", event.HtmlLink)
// }
//
// // PutEventsOnCalendar adds items from the trello board to the calendar
// func PutEventsOnCalendar() {
// 	// Get credentials
// 	credentials, err := ioutil.ReadFile("credentials.json")
// 	if err != nil {
// 		log.Fatalf("Unable to read client secret file: %v", err)
// 	}
//
// 	// If modifying these scopes, delete your previously saved token.json.
// 	config, err := google.ConfigFromJSON(credentials, calendar.CalendarScope)
// 	if err != nil {
// 		log.Fatalf("Unable to parse client secret file to config: %v", err)
// 	}
// 	client := getClient(config)
//
// 	srv, err := calendar.New(client)
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve Calendar client: %v", err)
// 	}
//
// 	startDate := time.Now().Format(time.RFC3339)
// 	endDate := time.Now().Add(time.Hour * 24 * 7).Format(time.RFC3339)
// 	workDayStart :=
// 	workDayEnd :=
//
// 	// Thanks to @wholien and his repo "freetime" for serving
// 	// as an example for how to use the freebusy calendar
// 	calendarID := "ansel.bridgewater@students.makeschool.com"
// 	calendarIDRequestItem := &calendar.FreeBusyRequestItem{Id: calendarID}
// 	calendarIDRequestItemArr := []*calendar.FreeBusyRequestItem{calendarIDRequestItem}
// 	freeBusyResp, err := srv.Freebusy.Query(&calendar.FreeBusyRequest{Items: calendarIDRequestItemArr, TimeMin: startDate, TimeMax: endDate}).Do()
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve freebusy. %v", err)
// 	}
// 	now := time.Now().UTC()
// 	fmt.Printf("Time %s", now)
// 	for _, v := range freeBusyResp.Calendars {
// 		for _, tp := range v.Busy {
// 			fmt.Println(tp.Start, tp.End)
// 			newEvent(srv, "summary", "description", "2019-03-08T09:00:00-07:00", "2019-03-08T17:00:00-07:00")
//
// 		}
// 	}
// }
//
// func main() {
// 	b, err := ioutil.ReadFile("credentials.json")
// 	if err != nil {
// 		log.Fatalf("Unable to read client secret file: %v", err)
// 	}
//
// 	// If modifying these scopes, delete your previously saved token.json.
// 	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
// 	if err != nil {
// 		log.Fatalf("Unable to parse client secret file to config: %v", err)
// 	}
// 	client := getClient(config)
//
// 	srv, err := calendar.New(client)
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve Calendar client: %v", err)
// 	}
//
// 	startDate := time.Now().Format(time.RFC3339)
// 	endDate := time.Now().Add(time.Hour * 24 * 7).Format(time.RFC3339)
// 	events, err := srv.Events.List("primary").ShowDeleted(false).
// 		SingleEvents(true).TimeMin(startDate).TimeMax(endDate).MaxResults(100).OrderBy("startTime").Do()
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
// 	}
//
// 	newEvent(srv, "summary", "description", "2019-03-08T09:00:00-07:00", "2019-03-08T17:00:00-07:00")
//
// 	// Thanks to wholien and his repo freetime for serving as an example for
// 	// how to use the freebusy calendar
// 	calendarID := "ansel.bridgewater@students.makeschool.com"
// 	fbri := &calendar.FreeBusyRequestItem{Id: calendarID}
// 	fbriarr := []*calendar.FreeBusyRequestItem{fbri}
// 	freeBusyResp, err := srv.Freebusy.Query(&calendar.FreeBusyRequest{Items: fbriarr, TimeMin: startDate, TimeMax: endDate}).Do()
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve freebusy. %v", err)
// 	}
// 	now := time.Now().UTC()
// 	fmt.Printf("Time %s", now)
// 	for _, v := range freeBusyResp.Calendars {
// 		for _, tp := range v.Busy {
// 			fmt.Println(tp.Start, tp.End)
// 		}
// 	}
//
// 	fmt.Println("Upcoming events:")
// 	if len(events.Items) == 0 {
// 		fmt.Println("No upcoming events found.")
// 	} else {
// 		for _, item := range events.Items {
// 			date := item.Start.DateTime
// 			if date == "" {
// 				date = item.Start.Date
// 			}
// 			fmt.Printf("%v (%v)\n", item.Summary, date)
// 		}
// 	}
// }

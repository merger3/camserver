package notifications

import (
	"fmt"
	"log"
	"strings"
	"time"

	chat "github.com/gempir/go-twitch-irc/v4"
	"github.com/gregdel/pushover"
	twitch "github.com/merger3/camserver/managers/twitch"
	//lint:file-ignore ST1001 I want to use dot imports
)

var (
	activeMessageThreshold int = 3
	activeMessageTimeframe int = 9

	inactivityNotifications = []Notification{
		{
			9,
			pushover.NewMessage("Cams are inactive"),
		},
		{
			20,
			pushover.NewMessage("Cams have been inactive for 30 minutes"),
		},
		{
			30,
			pushover.NewMessage("Cams have been inactive for an hour"),
		},
	}
)

type ActiveUser struct {
	Name                     string
	Count                    int
	Timers                   []*time.Timer
	ActiveUsers              chan<- string
	InactiveUsers            chan<- string
	SkipDeletionNotification bool
}

func (au ActiveUser) IsActive() bool {
	if au.Count >= activeMessageThreshold {
		return true
	} else {
		return false
	}
}

func (au *ActiveUser) Delete(skipNotifying bool) {
	fmt.Printf("Deleting %s\n", au.Name)
	for _, t := range au.Timers {
		t.Stop()
	}
	au.Count = 0
	au.SkipDeletionNotification = skipNotifying
	au.InactiveUsers <- au.Name
}

func (au *ActiveUser) Decrement() {
	newTimer := time.NewTimer(time.Duration(activeMessageTimeframe) * time.Minute)
	au.Timers = append(au.Timers, newTimer)
	<-newTimer.C
	au.Count--
	fmt.Printf("%s: %d\n", au.Name, au.Count)
	if au.Count == activeMessageThreshold-1 {
		au.Delete(au.SkipDeletionNotification)
	}
}

func (au *ActiveUser) Increment() {
	au.Count++
	fmt.Printf("%s: %d\n", au.Name, au.Count)
	if au.Count == activeMessageThreshold {
		au.ActiveUsers <- au.Name
	}
	go au.Decrement()
}

type NotificationTimer struct {
	Timer         *time.Timer
	Notifications []Notification
	Index         int
}

type Notification struct {
	Delay   float64
	Message *pushover.Message
}

func NewNotificationTimer(notifications []Notification) *NotificationTimer {
	nt := &NotificationTimer{
		Notifications: notifications,
		Index:         0,
	}
	if len(notifications) != 0 {
		nt.Timer = time.NewTimer(time.Duration(notifications[0].Delay) * time.Minute)
	}
	return nt
}

type NotificationsManager struct {
	Channels          []string
	Pushover          *pushover.Pushover
	Recipient         *pushover.Recipient
	Twitch            *twitch.TwitchManager
	User              *twitch.User
	History           map[string]*ActiveUser
	NotificationTimer *NotificationTimer
	InactiveUsers     chan string
	ActiveUsers       chan string
}

func NewNotificationsManager(twitchManager *twitch.TwitchManager, user *twitch.User, pushbulletKey, pushbulletDevice string, channels ...string) *NotificationsManager {
	nm := NotificationsManager{Twitch: twitchManager, User: user, History: make(map[string]*ActiveUser), InactiveUsers: make(chan string), ActiveUsers: make(chan string), Channels: channels}
	nm.registerListeners()

	listeners := []twitch.Listener{nm.Twitch.Listeners["activity"], nm.Twitch.Listeners["left"], nm.Twitch.Listeners["scenechange"]}
	user.ActiveListeners = append(listeners, user.ActiveListeners...)

	for _, c := range nm.Channels {
		nm.User.Client.Join(c)
	}

	nm.Pushover = pushover.New(pushbulletKey)
	nm.Recipient = pushover.NewRecipient(pushbulletDevice)

	nm.NotificationTimer = NewNotificationTimer(inactivityNotifications)

	go nm.manageActiveUsers()
	go nm.startTimer()
	return &nm
}

func (nm *NotificationsManager) getActiveUsersWithExemption(exemption string) []string {
	var activeUsers []string
	activeUsers = make([]string, 0)
	for _, u := range nm.History {
		if u.IsActive() && u.Name != exemption {
			activeUsers = append(activeUsers, u.Name)
		}
	}
	return activeUsers
}

type NotificationAction int

const (
	addedUser NotificationAction = iota
	removedUser
)

func (nm *NotificationsManager) buildNotification(action NotificationAction, username string) string {
	var msg string
	term := "also"
	if action == addedUser {
		msg += fmt.Sprintf("%s is now on cams", username)
	} else if action == removedUser {
		msg += fmt.Sprintf("%s is off cams", username)
		term = "still"
	}

	activeUsers := nm.getActiveUsersWithExemption(username)
	fmt.Printf("Active users:\n%+v\n", activeUsers)
	switch len(activeUsers) {
	case 1:
		msg += fmt.Sprintf(". %s is %s on cams", activeUsers[0], term)
	case 2:
		msg += fmt.Sprintf(". %s and %s are %s on cams", activeUsers[0], activeUsers[1], term)
	default:
		if len(activeUsers) >= 3 {
			var nameList string
			for i, n := range activeUsers {
				if i == 0 {
					nameList += fmt.Sprintf(". %s, ", n)
				} else if i != len(activeUsers)-1 {
					nameList += n + ", "
				} else {
					nameList += fmt.Sprintf("and %s are %s on cams", n, term)

				}
			}
			msg += nameList
		}
	}
	fmt.Println(msg)
	return msg
}

func (nm *NotificationsManager) manageActiveUsers() {
	for {
		select {
		case u := <-nm.ActiveUsers:
			if u != "merger31" {
				if err := nm.SendWithTimestamp(pushover.NewMessage(nm.buildNotification(addedUser, u)), nm.Recipient); err != nil {
					log.Panic(err)
				}
			}

		case u := <-nm.InactiveUsers:
			if !nm.History[u].SkipDeletionNotification {
				if err := nm.SendWithTimestamp(pushover.NewMessage(nm.buildNotification(removedUser, u)), nm.Recipient); err != nil {
					log.Panic(err)
				}
			}
			delete(nm.History, u)
		}
	}
}

// func (nm *NotificationsManager) craftNotification() string {

// }

func (nm *NotificationsManager) startTimer() {
	t := nm.NotificationTimer
	for {
		<-t.Timer.C

		fmt.Println("inactivity timer triggered")

		if err := nm.SendWithTimestamp(t.Notifications[t.Index].Message, nm.Recipient); err != nil {
			log.Panic(err)
		}

		t.Index++
		if t.Index < len(t.Notifications)-1 {
			newDuration := time.Duration(t.Notifications[t.Index].Delay) * time.Minute
			t.Timer.Reset(time.Duration(newDuration))
		}
	}

}

func (nm *NotificationsManager) registerListeners() {
	nm.Twitch.Listeners["activity"] = nm.activityMonitor
	nm.Twitch.Listeners["left"] = nm.leavingMonitor
	nm.Twitch.Listeners["scenechange"] = nm.sceneChange
}

func cleanName(s string) string {
	msgArray := strings.Split(s, " ")
	if len(msgArray) == 0 {
		return "unknown"
	}
	name := strings.TrimSuffix(msgArray[0], ":")

	return name
}

func (nm *NotificationsManager) activityMonitor(message chat.PrivateMessage) {
	if message.Channel == "alveusgg" && (strings.HasPrefix(message.Message, "!") && !(strings.HasPrefix(message.Message, "!getvolume") || strings.HasPrefix(message.Message, "!setvolume"))) {
		user := message.User.Name
		userHistory, ok := nm.History[user]
		if !ok {
			nm.History[user] = &ActiveUser{Name: user, InactiveUsers: nm.InactiveUsers, ActiveUsers: nm.ActiveUsers}
			userHistory = nm.History[user]
		}

		userHistory.Increment()
		if userHistory.IsActive() {
			t := nm.NotificationTimer
			t.Notifications = inactivityNotifications
			t.Index = 0
			t.Timer.Reset(time.Duration(t.Notifications[0].Delay) * time.Minute)
		}
	}
}

func (nm *NotificationsManager) ApiActivityMonitor(message chat.PrivateMessage) {
	if message.Channel == "alveusgg" && message.User.Name == "alveussanctuary" {
		user := cleanName(message.Message)
		if user == "Scene" || user == "PTZ" || user == "Volumes" || user == "!nightcams" || user == "Feeder" || user == "Clicked" || strings.HasPrefix(message.Message, "{") {
			return
		}
		userHistory, ok := nm.History[user]
		if !ok {
			nm.History[user] = &ActiveUser{Name: user, InactiveUsers: nm.InactiveUsers, ActiveUsers: nm.ActiveUsers}
			userHistory = nm.History[user]
		}

		userHistory.Increment()
		if userHistory.IsActive() {
			t := nm.NotificationTimer
			t.Notifications = inactivityNotifications
			t.Index = 0
			t.Timer.Reset(time.Duration(t.Notifications[0].Delay) * time.Minute)
		}
	}
}

func (nm *NotificationsManager) leavingMonitor(message chat.PrivateMessage) {
	if message.Channel == "alveusgg" && strings.Contains(message.Message, "ppPoof") {
		userHistory, ok := nm.History[message.User.Name]
		if ok && userHistory.IsActive() {
			userHistory.Delete(true)
		}

		t := nm.NotificationTimer
		t.Notifications = append(
			[]Notification{
				{
					3,
					pushover.NewMessage(fmt.Sprintf("%s left and hasn't been replaced", message.User.Name)),
				},
			},
			inactivityNotifications...,
		)
		t.Index = 0
		t.Timer.Reset(time.Duration(t.Notifications[0].Delay) * time.Minute)
	}
}

func (nm *NotificationsManager) sceneChange(message chat.PrivateMessage) {
	if (message.Channel == "alveusgg" || message.Channel == "alveussanctuary") && (strings.HasPrefix(message.Message, "!nightcams") || strings.HasPrefix(message.Message, "!livecams")) {
		t := nm.NotificationTimer
		t.Notifications = append(
			[]Notification{
				{
					2,
					pushover.NewMessage("Returned to livecams"),
				},
			},
			inactivityNotifications...,
		)
		t.Index = 0
		t.Timer.Reset(time.Duration(t.Notifications[0].Delay) * time.Minute)
	}
}

func (nm NotificationsManager) SendWithTimestamp(message *pushover.Message, recipient *pushover.Recipient) error {
	message.Timestamp = time.Now().Unix()
	_, err := nm.Pushover.SendMessage(message, recipient)
	if err != nil {
		return err
	}
	return nil
}

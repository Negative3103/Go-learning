package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	users := generateUsers(1000)
	wg := &sync.WaitGroup{}
	for _, user := range users {
		wg.Add(1)
		go saveUerInfs(user, wg)
	}

	wg.Wait()
	fmt.Println("Time elapsed:", time.Since(time.Now()).String())
}

var actions = []string{
	"log",
	"sign",
	"delete",
	"create",
	"change",
}

type logItem struct {
	actions   string
	timestamp time.Time
}

type Users struct {
	id    int
	email string
	logs  []logItem
}

func (users Users) getActivityInfo() string {
	out := fmt.Sprintf("ID: %d | Email: %s\nActivity Log:\n", users.id, users.email)
	for index, item := range users.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", index+1, item.actions, item.timestamp)
	}
	return out
}

func generateUsers(count int) []Users {
	users := make([]Users, count)
	for index := 0; index < count; index++ {
		users[index] = Users{
			id:    index + 1,
			email: fmt.Sprintf("user%d@ninja.go", index+1),
			logs:  generateLogs(rand.Intn(1000)),
		}
	}
	return users
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)
	for index := 0; index < count; index++ {
		logs[index] = logItem{
			timestamp: time.Now(),
			actions:   actions[rand.Intn(len(actions)-1)],
		}
	}
	return logs
}

func saveUerInfs(user Users, wg *sync.WaitGroup) error {
	fmt.Printf("Writing file for user id: %d\n", user.id)
	filename := fmt.Sprintf("logs/uid_%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	_, err = file.WriteString(user.getActivityInfo())
	if err != nil {
		return err
	}

	wg.Done()
	return nil
}

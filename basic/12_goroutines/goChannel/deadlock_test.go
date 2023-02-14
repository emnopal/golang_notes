package go_channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// deadlock
// ini ketika kita salah implementasi goroutine
// deadlock adalah proses goroutine yg saling tunggu sehingga tidak ada satupun goroutine yg jalan
// simulasi deadlock (contoh pakai mutex lock unlock mechanism)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(sender *UserBalance, receiver *UserBalance, amount int) {
	sender.Lock()
	fmt.Println("Lock", sender.Name)
	sender.Change(-amount)

	time.Sleep(5 * time.Second)

	receiver.Lock()
	fmt.Println("Lock", receiver.Name)
	receiver.Change(amount)

	sender.Unlock()
	receiver.Unlock()
}

func TestDeadLock(t *testing.T) {
	sender := &UserBalance{
		Name:    "Sender",
		Balance: 3000,
	}
	receiver := &UserBalance{
		Name:    "Receiver",
		Balance: 2000,
	}

	amount := 1000

	go Transfer(sender, receiver, amount)

	// result
	// Lock Sender
	// Sender 2000
	// Receiver 2000

	// explain: jadi hasilnya cuma ngeluarin yg sender aja, yg receiver enggak, karena si receiver
	// ini harus nunggu sampe si sender ini di unlock, karena kelamaan nunggu makanya harus di terminate sm golang nya
	// jadi data yg berubah cuma si sender aja, si receiver ini gapernah nerima transferan apapun

	time.Sleep(3 * time.Second)

	fmt.Println(sender.Name, sender.Balance)
	fmt.Println(receiver.Name, receiver.Balance)

}

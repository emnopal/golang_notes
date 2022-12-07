package go_channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.Mutex -> Mutex is Mutual Exclusion
/*
Untuk mengatasi masalah race condition diatas (di function TestRaceCondition) maka gunakan sync.Mutex
Mutex bisa digunakan untuk locking unlocking mutex, ketika kita lock mutex nya, maka tidak ada yg bisa melakukan locking
lagi kecuali kalau kita udah unlock (jadi cuma bisa 1x locking goroutine, seakan2 buat antrian buat lock-unlock goroutine)

*/

func TestMutex(t *testing.T) {
	var mutex sync.Mutex

	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()   // jika ada 1 goroutine yg di lock, maka 999 goroutine menunggu, jadi hasilnya bakalan jadi kaya synchronous padahal asynchronous
				x += 1         // lebih lambat, tapi gak terlalu kelihatan, tapi hasilnya gak race condition
				mutex.Unlock() // ini unlock, maka gantian goroutine yg ke 2 yg di lock
			}
		}()
	}

	time.Sleep(5 * time.Second) // sleep to make concurrent work

	fmt.Println("Counter:", x) // the result must be 100_000; but actually less or more than 100_000 and not equal 100_000
}

// sync.RWMutex -> mutual exclusion buat lock write dan read goroutine
// sebenernya pake mutex doang cukup, cuma rebutan nanti di proses send dan receive nya

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

// write goroutine
func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount // lock write process
	account.RWMutex.Unlock()
}

func (account *BankAccount) AddBalanceNoLock(amount int) {
	account.Balance += amount
}

// read goroutine
// buat safety reason aja locking di read goroutine ini ada; better jika pake ini juga ketika buat locking
func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance // lock read process
	account.RWMutex.RUnlock()
	return balance
}

func (account *BankAccount) GetBalanceNoLock() int {
	return account.Balance
}

func TestRWMutex(t *testing.T) {
	account := &BankAccount{}

	for i := 1; i <= 10; i++ {
		go func() {
			for j := 1; j <= 10; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance: ", account.GetBalance()) // Final Balance: 1000
}

func TestWMutex(t *testing.T) {
	account := &BankAccount{}

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 10; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalanceNoLock())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance: ", account.GetBalanceNoLock()) // Final Balance: 1000
}

func TestRMutex(t *testing.T) {
	account := &BankAccount{}

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 10; j++ {
				account.AddBalanceNoLock(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance: ", account.GetBalance()) // Final Balance: under 1000; Race Condition
}

func TestNoRWMutex(t *testing.T) {
	account := &BankAccount{}

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 10; j++ {
				account.AddBalanceNoLock(1)
				fmt.Println(account.GetBalanceNoLock())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance: ", account.GetBalanceNoLock()) // Final Balance: under 1000; Race Condition
}

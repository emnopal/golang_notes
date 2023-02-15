package goContext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// function buat implement interface dari context
// gunanya ini buat inisialisasi context dan implement interfacenya

// context.Backgroud() <- context kosong; inisialisasi context
// context kosong yang gapernah dibatalkan, timeout, value == nil.
// digunakan di main function, dalam testing dan awal proses request terjadi

// context.TODO()
// sama aja kaya context.Background()
// jika ada kasus, kita gatau mau implement konteks seperti apa
// best practice nya, kalau kita udah tau mau bikin context apa, segera ganti context.TODO() dengan context.Background()

func TestInitGoContext(t *testing.T) {
	// context.Background()
	ctxBg := context.Background()
	fmt.Println("Background -> ", ctxBg) // akan print context.Backgroud

	// context.TODO()
	ctxT := context.TODO()
	fmt.Println("Todo -> ", ctxT) // akan print context.TODO
}

// jadi ketika kita buat context, kita juga bisa bikin child dari context yg udah kita buat (inherit)
// satu context (parent) bisa inherit banyak child, tapi child hanya bisa punya 1 parent
// dan dari child tadi bisa inherit lagi jadi child child yang baru

// contoh
// ctx (parent) => (inherit) ctx.A, ctx.B, ctx.C
// ctx.A (child of ctx) => (inherit) ctx.A.0, ctx.A.1, ctx.A.2
// ctx.A.0 (child of ctx.A who child of ctx) => ctx.A.0.AA, ctx.A.0.AB

// note:
// parent and child akan selalu terhubung
// jadi ketika parent nya punya context maka yang lain akan terpengaruh
// contoh ketika ada pembatalan context dsb di parent, maka child nya aka terpengaruh
// tapi ketika ada child nya yg membatalkan context (sbg contoh), maka yg terpengaruh hanya child dari child nya
// untuk parent nya tidak akan berpengaruh
// (ini berlaku di segala aspek yg berhubungan dengan context, tidak terbatas hanya pada pembatalan context)

// karena context ini immutable, sehingga ketika kita tambahkan parameter/variable ke dalam context, dia bukan menambah parameter/variable
// ke dalam context tersebut, melainkan membuat context baru (berupa child context dari parent context)

// implementasi parent-child (inheritance) di context

// menambahkan value/data ke dalam context dengan withValue(parent, key, value)
// untuk key, silahkan buat type sendiri (contoh ada dibawah) buat menghindari type collision
// note: karena context immutable, jadi dari context.Backgroud() itu kita menurunkan context.Background() bukan merubah value nya

// use this to avoid type collision
type ctxStr string

const (
	ctxAStr ctxStr = "ctxA"
	ctxBStr ctxStr = "ctxB"
	ctxCStr ctxStr = "ctxC"
	ctxDStr ctxStr = "ctxD"
)

func TestGoContextWithValue(t *testing.T) {

	// parent context
	ctxParent := context.Background()

	// child context from ctxParent;
	ctxA := context.WithValue(ctxParent, ctxAStr, "a")
	ctxB := context.WithValue(ctxParent, ctxBStr, "b")

	// child context from ctxParent.ctxA
	ctxC := context.WithValue(ctxA, ctxCStr, "c")

	// child context from ctxParent.ctxA
	ctxD := context.WithValue(ctxB, ctxDStr, "d")

	fmt.Println(ctxParent) // context.Background
	fmt.Println(ctxA)      // context.Background.WithValue(type string, val a)
	fmt.Println(ctxB)      // context.Background.WithValue(type string, val b)
	fmt.Println(ctxC)      // context.Background.WithValue(type string, val a).WithValue(type string, val c)
	fmt.Println(ctxD)      // context.Background.WithValue(type string, val b).WithValue(type string, val d)

	// get value from key
	fmt.Println(ctxC.Value(ctxAStr)) // ctxC has value ctxA since it's inheritance from ctxParent.ctxA, so the result is `a`
	fmt.Println(ctxC.Value(ctxBStr)) // ctxC doesn't have value ctxB since it isn't inheritance from ctxParent.ctxB, so the result is `nil`
	fmt.Println(ctxD.Value(ctxAStr)) // ctxD doesn't have value ctxA since it isn't inheritance from ctxParent.ctxA, so the result is `nil`
	fmt.Println(ctxD.Value(ctxBStr)) // ctxD has value ctxB since it's inheritance from ctxParent.ctxB, so the result is `b`
}

// mengirimkan sinyal cancel ke context dengan context.WithCancel(parent)
// biasanya ini terjadi di goroutine dan goroutine nya biasanya beda
// in case: ketika kita ingin membatalkan sebuah proses goroutine
// note: kondisi goroutine spt ini harus di cek, supaya goroutine bisa benar2 membatalkan proses nya

// contoh kasus: di goroutine leak
// kasus dimana si goroutine ini jalan terus menerus gabisa di stop

// this function will always run since there is no logic that will stop it
// it's called goroutine leak
func counterLeak() chan int {
	dest := make(chan int)
	go func() {
		defer close(dest)
		counter := 1
		for {
			dest <- counter
			counter++
		}
	}()
	return dest
}

// consume goroutine leak
func TestCounterLeak(t *testing.T) {
	fmt.Println(runtime.NumGoroutine()) // print how much goroutine is created; result: 2

	dest := counterLeak()               // call the goroutine leak
	fmt.Println(runtime.NumGoroutine()) // print how much goroutine is created; if success the result must be 3; result: 3

	for n := range dest {
		fmt.Println("counter", n)
		if n == 10 { // this will break if counter is 10
			break
		}
	}

	fmt.Println(runtime.NumGoroutine()) // print how much goroutine is created; but in here the result is not return to 2; result: 3
	// (it's supposed to be 2, since counter goroutine have been terminated by break, but it isn't, so it's called goroutine leak)

	// goroutine leak is dangerous which can lead to memory leak
}

// fix goroutine leak
func counterNoLeak(ctx context.Context) chan int {
	dest := make(chan int)
	go func() {
		defer close(dest)
		counter := 1
		for {
			select {
			case <-ctx.Done(): // this will stop goroutine leak
				return
			default:
				dest <- counter
				counter++
			}
		}
	}()
	return dest
}

func TestCounterNoLeak(t *testing.T) {
	fmt.Println(runtime.NumGoroutine()) // print how much goroutine is created; result: 2

	ctxParent := context.Background()
	ctx, cancel := context.WithCancel(ctxParent)

	dest := counterNoLeak(ctx)          // panggil function nya
	fmt.Println(runtime.NumGoroutine()) // print how much goroutine is created; if success the result must be 3; result: 3

	for n := range dest {
		fmt.Println("counter", n)
		if n == 10 { // this will break if counter is 10
			break
		}
	}
	// kalo udah selesai, panggil cancel, supaya context nya kirim sinyal ke goroutine nya
	cancel()

	// karena dia asynchronous, jadi ketika selesai goroutine nya, jadi dia langsung ke akhir, abis itu baru cancel()
	// nah buat tau hasilnya itu kita bisa sleep aja function nya, supaya buat dia jalanin cancel abis itu ke fmt.print nya

	time.Sleep(2 * time.Second)

	fmt.Println(runtime.NumGoroutine()) // print how much goroutine is created; result: 2

	// goroutine leak problem solved
}

// mengirimkan sinyal timeout ke context dengan context.WithTimeout(parent, timeOut)
// sama kaya cancel, bedanya ini otomatis
// jadi ketika kita set timeout nya berapa, ketika sudah waktunya (timeout) maka akan di cancel

func counterSlow(ctx context.Context) chan int {
	dest := make(chan int)
	go func() {
		defer close(dest)
		counter := 1
		for {
			select {
			case <-ctx.Done(): // this will stop goroutine leak
				return
			default:
				dest <- counter
				counter++
				time.Sleep(1 * time.Second) // make the function is slow
			}
		}
	}()
	return dest
}

func TestCounterNoLeakTimeout(t *testing.T) {
	fmt.Println(runtime.NumGoroutine()) // print how much goroutine is created; result: 2

	ctxParent := context.Background()
	ctx, cancel := context.WithTimeout(ctxParent, 10*time.Second)

	// kalo udah selesai, panggil cancel, supaya context nya kirim sinyal ke goroutine nya
	// tapi literally ini udah otomatis, jadi defer buat mastiin aja klo goroutine nya bener2 udah di terminate
	defer cancel()

	dest := counterSlow(ctx)            // panggil function nya
	fmt.Println(runtime.NumGoroutine()) // print how much goroutine is created; if success the result must be 3; result: 3

	for n := range dest {
		fmt.Println("counter", n) // disini gaada break nya, karena kita mau manfaatin timeout nya
	}

	// karena dia asynchronous, jadi ketika selesai goroutine nya, jadi dia langsung ke akhir, abis itu baru cancel()
	// nah buat tau hasilnya itu kita bisa sleep aja function nya, supaya buat dia jalanin cancel abis itu ke fmt.print nya

	time.Sleep(2 * time.Second)

	fmt.Println(runtime.NumGoroutine()) // print how much goroutine is created; result: 2

	// goroutine leak problem solved
}

// selain dengan timeout, buat cancel otomatis juga bisa menggunakan deadline
// kalo timeout dengan time second, nah klo deadline dengan waktu yang real (misal besok, 2 hari lagi, dst)
// deadline ini timeout nya lebih jelas dan spesifik
// context.Deadline(parent, time)
func TestCounterNoLeakDeadline(t *testing.T) {
	fmt.Println(runtime.NumGoroutine()) // print how much goroutine is created; result: 2

	ctxParent := context.Background()
	ctx, cancel := context.WithDeadline(ctxParent, time.Now().Add(10*time.Second)) // more specific to date

	// kalo udah selesai, panggil cancel, supaya context nya kirim sinyal ke goroutine nya
	// tapi literally ini udah otomatis, jadi defer buat mastiin aja klo goroutine nya bener2 udah di terminate
	defer cancel()

	dest := counterSlow(ctx)            // panggil function nya
	fmt.Println(runtime.NumGoroutine()) // print how much goroutine is created; if success the result must be 3; result: 3

	for n := range dest {
		fmt.Println("counter", n) // disini gaada break nya, karena kita mau manfaatin timeout nya
	}

	// karena dia asynchronous, jadi ketika selesai goroutine nya, jadi dia langsung ke akhir, abis itu baru cancel()
	// nah buat tau hasilnya itu kita bisa sleep aja function nya, supaya buat dia jalanin cancel abis itu ke fmt.print nya

	time.Sleep(2 * time.Second)

	fmt.Println(runtime.NumGoroutine()) // print how much goroutine is created; result: 2

	// goroutine leak problem solved
}

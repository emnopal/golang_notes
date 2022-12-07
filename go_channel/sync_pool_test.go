package go_channel

import (
	"fmt"
	"sync"
	"testing"
)

// sync.Pool
/*
implementasi dari design pattern dari object pool pattern.
singkatnya pool ini digunakan untuk simpan data, kita juga bisa ambil datanya untuk digunakan dari pool, dan setelah dipakai
kita bisa balikan lagi datanya ke pool
pool di golang ini udah aman dari race condition, jadi gausah pusing2 handle race condition

implementasi real life nya:
- sering buat manage koneksi ke database
 karena bikin koneksi ke database itu lumayan mahal,
 dibanding jika kita mau konek ke database harus pakai koneksi baru
 lebih baik pakai pool saja, jadi kita hanya ambil dari pool jika dibutuhkan
 kalau tidak dibutuhkan maka koneksinya akan kesimpan di pool
*/

func TestPool(t *testing.T) {
	var pool sync.Pool
	var group sync.WaitGroup

	// memasukan data ke pool
	for i := 0; i <= 10; i++ {
		pool.Put(fmt.Sprint("Data", i)) // masukin 100 data ke pool
	}

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)
			data := pool.Get() // data yg di get bakalan hilang di pool
			fmt.Println(data)
			pool.Put(data)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Finish")
}

func TestPoolNoPutBack(t *testing.T) {

	// disini demonstrasi dari gimana kalau pool nya cuma dipanggil get aja tanpa di putback lagi ke pool nya
	// jadi disini bakalan gak kesimpen, karena data yg ada di goroutine ini bakalan dihapus kalau udah dipanggil pakai get

	var pool sync.Pool
	var group sync.WaitGroup

	// memasukan data ke pool
	for i := 0; i <= 10; i++ {
		pool.Put(fmt.Sprint("Data", i)) // masukin 100 data ke pool
	}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			data := pool.Get() // data yg di get bakalan hilang di pool
			fmt.Println(data)  // bakalan banyak data yg nil, karena pool nya habis
			// jadi make sure kalau data yg ada di pool kalau udah di get di balikin lagi ke pool pake put
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Finish")

}

func TestPoolNoPutBackDefaultValue(t *testing.T) {

	// disini demonstrasi dari gimana kalau pool nya cuma dipanggil get aja tanpa di putback lagi ke pool nya
	// jadi disini bakalan gak kesimpen, karena data yg ada di goroutine ini bakalan dihapus kalau udah dipanggil pakai get

	// nah karena hilang maka datanya jadi nil, nah buat ganti nil tersebut jadi apapun, kita bisa override attribute new
	// yg ada di struct pool nya

	var pool sync.Pool
	var group sync.WaitGroup

	// override nya disini, jadi dia bakalan return Empty Pool instead of nil
	pool = sync.Pool{
		New: func() interface{} {
			return "Empty Pool"
		},
	}

	// memasukan data ke pool
	for i := 0; i <= 10; i++ {
		pool.Put(fmt.Sprint("Data", i)) // masukin 100 data ke pool
	}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			data := pool.Get() // data yg di get bakalan hilang di pool
			fmt.Println(data)  // bakalan banyak data yg nil, karena pool nya habis
			// jadi make sure kalau data yg ada di pool kalau udah di get di balikin lagi ke pool pake put
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Finish")

}

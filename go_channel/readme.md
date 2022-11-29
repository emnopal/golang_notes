# Channel

## Apa itu?
1. Channel adalah tempat komunikasi secara synchronous yg bisa dilakukan dari goroutine (jadi konsepnya, si channel ini mau bikin goroutine synchronous, karena klo goroutine asynchronous, dia gabisa ngambil return value dari goroutine function)
2. Di channel itu ada sender dan receiver, biasanya sender dan receiver itu beda goroutine
3. Saat send data ke channel, goroutine akan ter-block sampe ada yg receive channel tsb (jadi jika ada goroutine function, dia harus await/tunggu sampe ada yg receive, mirip sm await yg ada di bahasa lain)
4. So, channel dikategorikan sbg alat komunikasi synchronous (krn ada blocking)
5. Channel itu konsep nya mirip sama async await di bahasa pemrograman lain, (dan ternyata goroutines itu agak lain sm async await)

## Flow nya

```go
message := make(chan string, 1) // goroutine sender
message <- 'hello from sender' // assign to goroutine sender
msg := <- message // goroutine receiver receive goroutine sender and pass to variable
fmt.Println(msg)
```


- Kasus Normal: 

    Goroutine sender kirim data -> Channel (tampung semua goroutine) -> Goroutine receiver ambil data dari goroutine sender -> Selesai

- Kasus Lain (receiver gak ngambil data):

    Goroutine sender kirim data -> Channel (tampung semua goroutine) (Goroutine nunggu dan gak jalan) -!> Goroutine receiver belum/gak ambil data dari goroutine sender -> Tidak Selesai

- Kasus Lain (sender gak kirim data):

    Goroutine sender belum kirim data -> Channel (tampung semua goroutine) (Goroutine nunggu dan gak jalan) -!> Goroutine receiver belum ambil data dari goroutine sender krn datanya ga ada (nunggu) -> Tidak Selesai

Note:
1. Channel hanya nampung 1 data, klo mau channel nya nerima data lain, harus di tunggu sampai proses nya selesai
2. Channel hanya menerima 1 jenis data, misal int ya int semua, string ya string semua dll
3. Channel harus di close, atau bisa sebabkan memory leak
# Go routines

## Introduction

Parallelism and Concurrency in Golang. **Parallelism and Concurrency is different!**

### Parallelism

Parallelism or Parallel programming is one of the way to solve complex problem with separating the problem into small part and running all of separated problem in the same time in one or multiple resources like CPUs, GPUs, Memories and another resources

#### Example of Parallelism

- Running apps simultaneously or multitasking (like running office, editor, browser, etc)
- Server operating CRUD in same time
- Machine Learning processing lots of data into model in same time with one or multiple GPUs or TPUs
- One or multiple CPUs or GPUs solving math equation simultaneously in the same time

### Process vs Thread

Di parallel programming, itu ada istilah process dan thread, lalu apa bedanya?

|Process|Thread|
|-|-|
|Process adalah sebuah ekseskusi program|Thread adalah segmen/bagian dari process|
|Process mengkonsumsi memory yg besar|Thread menggunakan memory yg kecil|
|Process saling terisolasi dengan process lain|Thread bisa saling berhubungan jika dalam process yg sama|
|Process lama untuk dijalankan dan dihentikan|Thread cepat untuk dijalankan dan dihentikan|

Contoh gampang nya:

- Process itu satu aplikasi penuh (misal: Google Chrome)
- Thread itu bagian dari aplikasi yg berjalan (misal: Extension yg depends ke Google Chrome, Tabs dari Google Chrome)

### Parallel vs Concurrency

Apa bedanya?

|Parallel|Concurrency|
|-|-|
|Parallel menjalankan aplikasi secara bersamaan|Concurrency menjalankan aplikasi secara bergantian|
|Parallel membutuhkan banyak Thread|Concurrency hanya butuh sedikit Thread|
|Parallel tidak bisa gonta-ganti Thread, karena semua Thread harus dijalankan dalam waktu yang sama|Concurrency bisa gonta-ganti Thread, karena Threadnya bisa dijalankan secara bergantian dalam waktu yang sama|

Contoh gampangnya:

- Parallel itu sekaligus ngerjain sesuatu (misal: Buka Browser)
- Concurrency itu mengerjakan dan menyelesaikan pekerjaan/task secara bergantian (misal: Website di Internet gak harus nunggu semua page ke load untuk menampilkan isi website (walaupun tidak utuh))

### CPU-bound dan IO-bound

- CPU-bound itu ketergantungan aplikasi terhadap sumber daya yg membutuhkan processing unit seperti CPU, GPU atau TPU dan biasanya CPU-bound depends banget sama speed atau jumlah core/tensor/cuda atau bahkan jumlah dari processing unit tersebut. Biasanya aplikasi CPU-bound gak pakai konsep Parallelism karena konsep kerjanya harus barengan
- Contoh CPU-bound: OS, Machine Learning, AI, IoT, Game
- Golang secara default dia pake conccurency jadi ga masuk ke CPU-bound, tapi bisa aja, tapi conccurency nya jadi gak kepake

- I/O bound itu ketergantungan aplikasi terhadap kecepatan input dan output device yg digunakan, oleh karena itu kebanyakan dari I/O bound itu butuh harddisk/ssd/ram/memori yang besar
- Contoh I/O bound: accessing file, website, server CRUD, I/O Database, scrapping

## Goroutines

### Introduction Goroutines

- Goroutine istilahnya "thread ringan" yang dikelola oleh Go Runtime (gaperlu pusing ngelola goroutine)
- Ukutan goroutine di memori sangat kecil, hanya sekitar 2kb, jauh lebih ringan dari Thread biasa yg ukurannya 500x lipat (1000kb)
- Tidak seperti thread biasa yg jalannya parallel, goroutine berjalan secara conccurent
- Goroutine ini sangat ringan, jadi kita bisa buat sebanyak-banyak nya goroutine tanpa takut memory leak

### Cara kerja

- Goroutine dijalankan oleh Go scheduler di dalam thread, dimana jumlah thread nya sebanyak GOMAXPROCS (ini go env nya nilai default nya sesuai sama core CPU)
- Goroutines itu sebenarnya bukan Thread, tapi kaya process (process disini bukan maksud dari Process diatas melainkan seperti running sebuah code, biar lebih gampang pake istilah process aja (dengan p kecil)) yg berjalan di dalam Thread, makanya disebut thread ringan, sehingga goroutines ini bukan pengganti Thread
- Di golang semua sudah otomatis, gak perlu diatur lagi management thread nya secara manual, krn sudah diatur di Go Scheduler

### Terminology

- G: Goroutine
- M: Thread (Machine)
- P: Processor

### Note

Goroutine ini native golang, jadi gaperlu ngelakuin import apapun disini (kalo di bahasa lain kan perlu async await atau promise), cukup tambahin keyword go setelah manggil function

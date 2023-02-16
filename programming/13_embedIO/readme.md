# Golang Embed IO

Package embed merupakan fitur yang mempermudah baca isi file ketika compile time, secara otomatis isi file dimasukan ke dalam variable

Cara embed nya:

``` golang
// Tambahkan komentar seperti ini di variable yg mau kita embed

//go:embed `namafile.extension`
embed := functionEmbed(...)

// note: variable yg mau kita embed harus diluar function, gabisa didalam function
```

Note:
    - Golang Embed tersedia di package embed dan hanya tersedia di golang versi 1.16 keatas
    - Golang Embed akan menyimpan file ke dalam binary golang ketika di compile, sehingga, ketika
      Program Golang telah di compile, maka golang tidak akan membaca file dari luar lagi
      Kelebihannya, kita gaperlu takut atau repot ketika file yg ada di project kita hilang, berubah atau corrupt
      Kekurangannya, gak real time, kalau ada perubahan, kalian harus re-compile lagi

# Go Module

## Buat Module Baru

Buat module baru: `go mod init <nama_module>`

## Tambah dependencies

Tambah module dependencies: `go get <nama_module>`

## Upgrade dependencies

Upgrade module dependencies: ganti tags di dalam go.mod ke versi terbaru atau versi yang lebih tinggi lalu lakukan `go get`

## Downgrade dependencies

Downgrade module dependencies: ganti tags di dalam go.mod ke versi lama atau versi yang lebih rendah lalu lakukan `go get`

## Menggunakan versi lain (bukan yang terbaru)

Caranya sama seperti upgrade module, yaitu tinggal ganti tags di go.mod sesuai yang kita mau, lalu lakukan `go get`

## Release Module

Golang terintegrasi dengan git, untuk release module, kita hanya perlu buat tag di git:

`git tag v<version>`

kemudian push ke git dengan command:

`git push origin <tag_name>`

Note: versioning in golang is using semver (semantic versioning) system

Contoh: bisa dilihat [disini](https://github.com/emnopal/private-go-module)

## Upgrade Module

Buat upgrade module tinggal tambahkan tag baru ke git

## Go Package

Saya buat simple go package sendiri, bisa dilihat [disini](https://github.com/emnopal/private-go-module)

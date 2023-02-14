# Go Module

## Buat Module Baru

Buat module baru: `go mod init <nama_module>`

## Tambah dependencies

Tambah module dependencies: `go get <nama_module>`

## Upgrade dependencies (Minor)

Upgrade module dependencies: ganti tags di dalam go.mod ke versi terbaru atau versi yang lebih tinggi lalu lakukan `go get`

## Downgrade dependencies (Minor)

Downgrade module dependencies: ganti tags di dalam go.mod ke versi lama atau versi yang lebih rendah lalu lakukan `go get`

## Menggunakan versi lain (bukan yang terbaru) (Minor)

Caranya sama seperti upgrade module, yaitu tinggal ganti tags di go.mod sesuai yang kita mau, lalu lakukan `go get`

**Note**: Minor disini maksudnya update nya tidak menambahkan versi di bagian pertama, contoh dari v1.0.0 ke v.1.0.1 atau v1.1.1

## Upgrade dependencies (Major)

Upgrade module dependencies: hapus require yang versi lama, contoh, jika ada: `require github.com/nama_repo/nama_module` (note: ini jika tidak ada prefix nya) atau `require github.com/nama_repo/nama_module/v1` silahkan dihapus, lalu lakukan `go get require github.com/nama_repo/nama_module/v2`

## Downgrade dependencies (Major)

Upgrade module dependencies: hapus require yang versi lama, contoh, jika ada: `require github.com/nama_repo/nama_module/v2` silahkan dihapus, lalu lakukan `go get require github.com/nama_repo/nama_module/v1` atau `go get require github.com/nama_repo/nama_module` (note: ini jika tidak ada prefix nya)

## Menggunakan versi lain (bukan yang terbaru) (Major)

Upgrade module dependencies: hapus require yang versi lama, contoh, jika ada: `require github.com/nama_repo/nama_module/v<previous>` silahkan dihapus, lalu lakukan `go get require github.com/nama_repo/nama_module/v<other>`

**Note**: Major disini maksudnya update nya menambahkan versi di bagian pertama, contoh dari v1.0.0 ke v.2.0.0

## Release Module

Golang terintegrasi dengan git, untuk release module, kita hanya perlu buat tag di git:

`git tag v<version>`

kemudian push ke git dengan command:

`git push origin <tag_name>`

Note: versioning in golang is using semver (semantic versioning) system

Contoh: bisa dilihat [disini](https://github.com/emnopal/private-go-module)

## Upgrade Module

Buat upgrade module tinggal tambahkan tag baru ke git

## Major Upgrade Module

Ganti nama module dengan menambahkan prefix v<new_version_numbering>, contoh: `module github.com/nama_repo/nama_module/v2`, lalu buat tag nya dengan: `git tag v2.0.0` lalu push ke git.

## Go Package

Saya buat simple go package sendiri, bisa dilihat [disini](https://github.com/emnopal/private-go-module)

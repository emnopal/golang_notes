# Go Import

1. By default, file golang hanya bisa akses file golang lainnya yg berada di package yg sama
2. Kalau di package yg beda, gunakan import
3. naming package nya harus simple, gaada campuran huruf besar kecil dan campuran _ atau -, misal:
   1. namanya harus `package somepackage` tidak direkomendasikan `package SomePackage` atau `package somePackage` atau `package some_package`
   2. karena nama package harus mengikuti nama folder, jadi rekomendasi ini juga diterapin ke nama folder tempat library tersebut berada
4. di golang naming function nya harus diawali pake huruf besar misal `func Function()`, kalo engga misal `func function()` gabisa diimport (ini terkait access modifier)
   1. klo awalnya huruf besar `func Function()` ini menandakan kalau access modifier nya public
   2. klo awalnya huruf kecil `func function()` ini menandakan kalau access modifier nya private
   3. Ini berlaku juga buat struct, interface, variable dsb
5. kalau mau bikin main function `func main()` harus di package yg namanya `package main`
6. kalau function/struct dsb nya sama tapi package nya sama itu gabisa, tapi klo beda bisa

Note: disable go module sebelum run code disini (go module itu nanti dulu): `go env -w GO111MODULE=off`

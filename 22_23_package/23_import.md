# Go Import

1. By default, file golang hanya bisa akses file golang lainnya yg berada di package yg sama
2. Kalau di package yg beda, gunakan import
3. naming package nya harus simple, gaada campuran huruf besar kecil dan campuran _ atau -, misal:
   1. namanya harus `package somepackage` gaboleh `package SomePackage` atau `package somePackage` atau `package some_package`
   2. karena nama package harus mengikuti nama folder, jadi aturan ini juga wajib diterapin ke nama folder tempat library tersebut berada
4. di golang naming function nya harus diawali pake huruf besar misal `func Function()`, kalo engga misal `func function()` gabisa diimport (ini terkait access modifier)
   1. klo awalnya huruf besar `func Function()` ini menandakan kalau access modifier nya public
   2. klo awalnya huruf kecil `func function()` ini menandakan kalau access modifier nya private
5. package import rules: `parent_directory/<lokasi_package>/nama_package(atau sebutan lain: namespace)`. Contoh:
   ```
   -- package_parent
        -- module1
            -> module10.go
            -> module11.go
        -- module2
            -> module2.go
        -> main.go

    keterangan:
        -- = folder
        -> = file
    ```
    misal, kita mau import module10, module11 dan module2 di main.go, maka lakukan import seperti ini:
    ```
    package main

    import (
        "package_parent/module1"
        "package_parent/module2
    )

    func main() {
        module1.Module10()
        module1.Module11()
        module2.Module2()
    }
    ```

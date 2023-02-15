# Context

## Apa itu context?

- Context adalah data yang membawa:
  1. value
  2. signal cancel
  3. signal timeout
  4. signal deadline
    yang dibuat per request (misal lewat HTTP request)
- Context mempermudah developer meneruskan value dan signal antar proses
- Context dapat mengirim data request dan signal ke proses lain
- Karena context, kita bisa batch batalin semua proses dengan hanya kirim signal ke context

- Contoh Context
    1. Ada proses goroutine 1,2,3 yang 1 ada context, yg 2 dan 3 gaada context:
    `goroutineProc.1 (context) -> goroutineProc.2 (!context) -> goroutineProc.3 (!context)`
    2. Kita kirim semua context ke proc 2 dan 3
    `goroutineProc.1 (context) -> goroutineProc.2 (context) -> goroutineProc.3 (context)`
    3. Kita kirim signal (contohnya cancel disini) ke proc 2 dan proc 3
    `goroutineProc.1 (context: KILL) -> goroutineProc.2 (context: KILL) -> goroutineProc.3 (context: KILL)`
    4. Ketika ada sinyal cancel di proc1 maka proc2 dan 3 otomatis ikut cancel juga, karena dia menerima sinyal dari proc1
    `!goroutineProc.1 (context: KILL) -> !goroutineProc.2 (context: KILL) -> !goroutineProc.3 (context: KILL)`
    5. Kalo gak kita kirim context ke process lain, maka process tidak akan di cancel
    `!goroutineProc.1 (context: KILL) -> goroutineProc.2 (!context) -> goroutineProc.3 (!context)`

Jadi literally context itu cuma data aja, jadi jika ada context yang dikirim ke process lain, si proses itu tinggal ngebaca aja

Note: Context itu immutable;  Context ada di interface context di dalam package context

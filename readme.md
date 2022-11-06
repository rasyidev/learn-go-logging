# Go-Lang Logging
## Pengenalan logging
- File berisi informasi kejadian dalam sistem
- Biasanya ditulis dalam file, lalu dikirim ke database jika dibutuhkan
- Logging: Aksi menambah informasi ke log file
- Logging sudah menjadi standar industri untuk menampilkan informasi yang terjadi di dalam aplikasi
- Logging dapat digunakan pula untuk proses debugging

## Logging Library
- Package `log` bawaan golang fiturnya terbatas sehingga terdapat library tambahan untuk implementasi logging.
- Package Logging:
  - Logrus  : https://github.com/sirupsen/logrus
  - Zap     : https://github.com/uber-go/zap
  - Zerolog : https://github.com/rs/zerolog
  - dll.

## `Logger`
- Struct utama pada Logrus untuk melakukan Logging
- `Logger.New()`

## Level
- Penentuan prioritas atau jenis dari sebuah kejadian
- Hal paling penting dalam Logging (secara general)
- Level dimulai dari level rendah ke tinggi
- Logrus mendukung banyak level:
  - Trace: `logger.Trace()`
  - Debug: `logger.Debug()`
  - Info: `logger.Info()`
  - Warn: `logger.Warn()`
  - Error: `logger.Error()` 
  - Fatal: `logger.Fatal()` -> Memanggil `os.Exit(1)` setelah logging
  - Panic: `logger.Panic()` -> Memanggil `panic()` setelah logging
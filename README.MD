# Golang RESTful API dengan Arsitektur Domain-Driven Design (DDD)

![Go](https://img.shields.io/badge/Go-1.21%2B-blue.svg) ![Fiber](https://img.shields.io/badge/Fiber-v3-cyan.svg) ![GORM](https://img.shields.io/badge/GORM-v1.25%2B-brightgreen.svg) ![MySQL](https://img.shields.io/badge/MySQL-8.0-orange.svg)

Ini adalah proyek RESTful API yang dibangun dengan **Golang**, menerapkan arsitektur **Domain-Driven Design (DDD)** yang bersih dan modular. Proyek ini merupakan hasil konversi dan pengembangan dari repositori NestJS [belajar-nestjs-restful-api](https://github.com/ProgrammerZamanNow/belajar-nestjs-restful-api).

Tujuan proyek ini adalah menyediakan _boilerplate_ yang kokoh dan siap produksi (_production-ready_) untuk membangun layanan _backend_ yang _scalable_ dan mudah dikelola.

---

## ✨ Fitur & Teknologi

- **Framework**: **Fiber v3** - Web framework yang ekspresif dan super cepat, terinspirasi dari Express.js.
- **Arsitektur**: **Domain-Driven Design (DDD)** - Memisahkan logika bisnis inti dari detail infrastruktur, menghasilkan kode yang lebih bersih dan mudah diuji.
- **Database**: **MySQL** dengan **GORM** sebagai ORM yang intuitif.
- **Konfigurasi**: **Viper** - Manajemen konfigurasi yang mendukung file YAML, variabel lingkungan, dan lainnya.
- **Validasi**: **go-playground/validator** - Validasi _struct_ dan _field_ yang andal.
- **Migrasi DB**: **golang-migrate/migrate** - Mengelola skema database dengan file `.sql` yang terversi.
- **Logging**: **Logrus** - _Logging_ terstruktur dengan implementasi **Goroutine + Channel** untuk _logging_ asinkron ke file tanpa memblokir _request_.
- **Otentikasi**: **JWT** (_JSON Web Tokens_) untuk mengamankan _endpoint_.
- **Fitur Lanjutan**:
  - Respon API standar menggunakan **Generics** Go.
  - _Custom error handler_ yang konsisten.
  - _Dependency Injection_ manual yang jelas di `main.go`.

---

## 📂 Arsitektur Proyek

Proyek ini mengikuti prinsip _Clean Architecture_ dengan pemisahan _layer_ yang ketat.

```plaintext
/
├── cmd/server/main.go      # Entry point & Dependency Injection
├── config/                 # File konfigurasi .yaml
├── internal/
│   ├── application/        # Layer Service/Use Case (Logika Bisnis)
│   ├── domain/             # Layer Domain (Inti Bisnis: Model & Interface)
│   ├── infrastructure/     # Layer Infrastruktur (Implementasi: DB, Logger)
│   └── interfaces/         # Layer Presentasi (Handler, Router, Middleware)
├── migrations/             # File migrasi SQL
├── pkg/util/               # Utilitas umum (password, token, validator)
└── test/                   # Unit & Integration Test

domain: Jantung dari aplikasi. Berisi entity (model) dan interface (kontrak) untuk repository dan service. Layer ini tidak bergantung pada detail eksternal.

application: Berisi implementasi logika bisnis (use case) yang mengorkestrasi data dari repository.

infrastructure: Implementasi konkret dari interface yang ada di domain. Misalnya, koneksi database GORM, repository GORM, dan setup logger.

interfaces: Pintu masuk ke aplikasi. Berisi handler HTTP (Fiber), router, dan middleware. Bertugas menerjemahkan request HTTP menjadi panggilan ke layer application.

🚀 Panduan Memulai
Prasyarat

Go versi 1.21 atau lebih baru.

Docker & Docker Compose (dianjurkan untuk database).

golang-migrate/migrate CLI. Instal dengan perintah:

Bash
go install -tags 'mysql' [github.com/golang-migrate/migrate/v4/cmd/migrate@latest](https://github.com/golang-migrate/migrate/v4/cmd/migrate@latest)
1. Clone Repositori

Bash
git clone [https://github.com/dimasafifudin11/golang-ddd-pattern-restful-api.git](https://github.com/dimasafifudin11/golang-ddd-pattern-restful-api.git)
cd golang-ddd-pattern-restful-api
2. Instal Dependensi

Salin dan jalankan perintah go get berikut di terminal Anda untuk mengunduh semua package yang dibutuhkan.

Bash
# Framework & Tools Wajib
go get [github.com/gofiber/fiber/v3](https://github.com/gofiber/fiber/v3)
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get [github.com/spf13/viper](https://github.com/spf13/viper)
go get [github.com/go-playground/validator/v10](https://github.com/go-playground/validator/v10)
go get [github.com/golang-migrate/migrate/v4/database/mysql](https://github.com/golang-migrate/migrate/v4/database/mysql)
go get [github.com/golang-migrate/migrate/v4/source/file](https://github.com/golang-migrate/migrate/v4/source/file)
go get [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)

# Lainnya
go get [github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt/v5)
go get golang.org/x/crypto/bcrypt
go get [github.com/stretchr/testify](https://github.com/stretchr/testify)
Atau, cukup jalankan:

Bash
go mod tidy
3. Setup Konfigurasi

Salin file contoh konfigurasi. File config.yaml akan diabaikan oleh Git.

Bash
cp config/config.example.yaml config/config.yaml
Buka config/config.yaml dan sesuaikan dengan konfigurasi database dan JWT Anda.

4. Setup Database (via Docker)

Pastikan Docker sudah berjalan. Jalankan perintah berikut untuk membuat container MySQL.

Bash
docker run --name mysql-ddd -p 3306:3306 -e MYSQL_ROOT_PASSWORD=rahasia -e MYSQL_DATABASE=golang_ddd -d mysql:8.0
Penting: Sesuaikan MYSQL_ROOT_PASSWORD dan MYSQL_DATABASE agar cocok dengan yang ada di config.yaml.

5. Jalankan Migrasi Database

Gunakan migrate CLI untuk membuat tabel yang dibutuhkan.

Bash
# Format: migrate -database "mysql://<user>:<password>@tcp(<host>:<port>)/<dbname>" -path <migrations_folder> <command>
migrate -database "mysql://root:rahasia@tcp(localhost:3306)/golang_ddd" -path migrations up
6. Jalankan Aplikasi

Bash
go run cmd/server/main.go
Server akan berjalan di http://localhost:8080 (atau port lain sesuai konfigurasi).

🧪 Menjalankan Unit Test
Untuk menjalankan semua unit test dan melihat coverage:

Bash
go test ./... -v -cover
📖 Dokumentasi Endpoint API
Anda dapat menggunakan file api_tests.http dengan ekstensi REST Client di VS Code untuk menguji API.

Authentication

Method	Endpoint	Deskripsi	Body / Payload
POST	/api/register	Mendaftarkan user baru.	{"name":"", "email":"", "password":""}
POST	/api/login	Login untuk mendapatkan token JWT.	{"email":"", "password":""}
Users (Protected)

Method	Endpoint	Deskripsi	Header Wajib
GET	/api/users/current	Mendapatkan profil user saat ini.	Authorization: Bearer <token>
```

# API Service Proyek Transparansi Pemerintah

Selamat datang di dokumentasi API untuk proyek transparansi pengadaan pemerintah. Service ini menyediakan backend untuk platform yang menghubungkan pemerintah, vendor, dan masyarakat, dengan tujuan meningkatkan akuntabilitas dan efisiensi melalui teknologi.

API ini dibangun menggunakan **Golang** dengan **MongoDB** sebagai database, dan menerapkan prinsip **Clean Architecture** untuk memastikan kode yang bersih, terstruktur, dan mudah dikelola.

## ✨ Fitur Utama

-   **Manajemen Profil**: Registrasi dan pengelolaan profil untuk `GOVERNMENT`, `VENDOR`, dan `CITIZEN` menggunakan wallet address sebagai identitas utama.
-   **Manajemen Proyek**: Pemerintah dapat membuat, melihat, dan mengelola tender proyek.
-   **Manajemen Proposal**: Vendor dapat mengajukan proposal untuk proyek yang tersedia.
-   **Arsitektur Bersih**: Pemisahan yang jelas antara logika bisnis, data, dan framework.

## 🛠️ Teknologi yang Digunakan

-   **Bahasa**: Golang
-   **Framework**: Gin Gonic
-   **Database**: MongoDB
-   **Konfigurasi**: godotenv
-   **UUID**: google/uuid

## 📋 Prasyarat

Sebelum memulai, pastikan Anda telah menginstal:

-   [Go](https://golang.org/doc/install) (versi 1.18 atau lebih baru)
-   [MongoDB](https://www.mongodb.com/try/download/community)
-   [Git](https://git-scm.com/)

## 🚀 Instalasi & Konfigurasi

1.  **Clone repository ini:**
    ```sh
    git clone [https://github.com/yourusername/golang-govtech-api.git](https://github.com/yourusername/golang-govtech-api.git)
    cd golang-govtech-api
    ```

2.  **Buat file `.env`** di root direktori dan isi dengan konfigurasi berikut:
    ```env
    # Server Configuration
    SERVER_PORT=8080

    # MongoDB Configuration
    MONGO_URI=mongodb://localhost:27017
    MONGO_DB_NAME=govtech
    ```

3.  **Install dependensi Go:**
    ```sh
    go mod tidy
    ```

## ▶️ Menjalankan Proyek

Untuk menjalankan API service, gunakan perintah berikut dari direktori root:

```sh
go run ./cmd/api/main.go
```

Server akan berjalan di `http://localhost:8080`.

---

## 📚 Dokumentasi API

Semua endpoint berada di bawah base URL: `/api/v1`

### Endpoint Health Check

-   **`GET /health`**
    -   **Deskripsi**: Memeriksa apakah service berjalan dengan baik.
    -   **Response Sukses (200 OK)**:
        ```json
        {
          "status": "UP"
        }
        ```

---

### 👤 Manajemen Profil (`/profiles`)

#### 1. Membuat Profil Baru

-   **`POST /profiles`**
    -   **Deskripsi**: Mendaftarkan wallet address baru sebagai salah satu dari tiga peran (`GOVERNMENT`, `VENDOR`, `CITIZEN`).
    -   **Request Body**:
        ```json
        {
          "wallet_address": "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B",
          "role": "VENDOR",
          "details": {
            "company_name": "PT. Teknologi Maju",
            "nib": "1234567890123",
            "npwp": "0123456789012345",
            "office_address": "Jl. Raya Inovasi No. 10",
            "domicile_address": "Jl. Raya Inovasi No. 10",
            "contact_email": "kontak@teknologimaju.com",
            "contact_number": "08123456789"
          }
        }
        ```
    -   **Response Sukses (201 Created)**:
        ```json
        {
          "_id": "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B",
          "role": "VENDOR",
          "is_active": true,
          "details": {
            "company_name": "PT. Teknologi Maju",
            "nib": "1234567890123",
            "npwp": "0123456789012345",
            "office_address": "Jl. Raya Inovasi No. 10",
            "domicile_address": "Jl. Raya Inovasi No. 10",
            "contact_email": "kontak@teknologimaju.com",
            "contact_number": "08123456789"
          },
          "created_at": "2025-07-24T21:45:00.123Z",
          "updated_at": "2025-07-24T21:45:00.123Z",
          "deleted_at": null
        }
        ```
    -   **Response Error (409 Conflict)**:
        ```json
        {
          "error": "profile with this wallet address already exists"
        }
        ```

#### 2. Mendapatkan Profil Berdasarkan Wallet

-   **`GET /profiles/:walletAddress`**
    -   **Deskripsi**: Mengambil detail profil berdasarkan `walletAddress`.
    -   **URL Params**:
        -   `walletAddress` (string, required): Alamat wallet yang terdaftar.
    -   **Response Sukses (200 OK)**:
        *(Sama seperti response sukses saat membuat profil)*
    -   **Response Error (404 Not Found)**:
        ```json
        {
          "error": "profile not found"
        }
        ```

---

### 🏗️ Manajemen Proyek (`/projects`)

#### 1. Membuat Proyek Baru

-   **`POST /projects`**
    -   **Deskripsi**: Membuat tender proyek baru. Hanya dapat dilakukan oleh profil dengan peran `GOVERNMENT`.
    -   **Request Body**:
        ```json
        {
          "government_wallet": "0x4E9a23b816a04b1625983A7E3f52CEaD8b7b25A5",
          "project_name": "Pembangunan Sistem E-Arsip Digital",
          "description": "Membangun sistem pengarsipan digital terpusat untuk pemerintah kota.",
          "images": ["[https://example.com/image.png](https://example.com/image.png)"],
          "budget_wei": "50000000000000000000",
          "smart_contract_address": "0x...",
          "proposal_deadline": "2025-08-31T23:59:59Z",
          "voting_deadline": "2025-09-15T23:59:59Z"
        }
        ```
    -   **Response Sukses (201 Created)**:
        ```json
        {
          "_id": "a1b2c3d4-e5f6-7890-1234-567890abcdef",
          "government_wallet": "0x4E9a23b816a04b1625983A7E3f52CEaD8b7b25A5",
          "project_name": "Pembangunan Sistem E-Arsip Digital",
          "description": "Membangun sistem pengarsipan digital terpusat untuk pemerintah kota.",
          "images": ["[https://example.com/image.png](https://example.com/image.png)"],
          "budget_wei": { "$numberDecimal": "50000000000000000000" },
          "status": "OPEN_FOR_PROPOSAL",
          "smart_contract_address": "0x...",
          "winning_proposal_id": null,
          "proposal_deadline": "2025-08-31T23:59:59Z",
          "voting_deadline": "2025-09-15T23:59:59Z",
          "created_at": "2025-07-24T21:50:00.123Z",
          "updated_at": "2025-07-24T21:50:00.123Z",
          "deleted_at": null
        }
        ```

#### 2. Mendapatkan Semua Proyek

-   **`GET /projects`**
    -   **Deskripsi**: Mengambil daftar semua proyek yang aktif.
    -   **Response Sukses (200 OK)**:
        ```json
        [
          {
            "_id": "a1b2c3d4-e5f6-7890-1234-567890abcdef",
            "project_name": "Pembangunan Sistem E-Arsip Digital",
            "status": "OPEN_FOR_PROPOSAL",
            "...": "..."
          }
        ]
        ```

#### 3. Mendapatkan Detail Proyek Berdasarkan ID

-   **`GET /projects/:id`**
    -   **Deskripsi**: Mengambil detail satu proyek berdasarkan ID uniknya.
    -   **URL Params**:
        -   `id` (string, required): UUID dari proyek.
    -   **Response Sukses (200 OK)**:
        *(Sama seperti response sukses saat membuat proyek)*
    -   **Response Error (404 Not Found)**:
        ```json
        {
          "error": "project not found"
        }
        ```

---

### 📄 Manajemen Proposal (`/proposals`)

#### 1. Membuat Proposal Baru

-   **`POST /proposals`**
    -   **Deskripsi**: Mengajukan proposal baru untuk sebuah proyek. Hanya dapat dilakukan oleh profil dengan peran `VENDOR`.
    -   **Request Body**:
        ```json
        {
          "project_id": "a1b2c3d4-e5f6-7890-1234-567890abcdef",
          "vendor_wallet": "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B",
          "proposal_name": "Proposal Sistem E-Arsip Cepat & Aman",
          "description": "Penawaran sistem e-arsip dengan teknologi terbaru...",
          "images": [],
          "requested_budget_wei": "48000000000000000000",
          "onchain_payload": {
            "targets": [],
            "values": [],
            "calldatas": []
          }
        }
        ```
    -   **Response Sukses (201 Created)**:
        ```json
        {
          "_id": "f0e9d8c7-b6a5-4321-fedc-ba9876543210",
          "project_id": "a1b2c3d4-e5f6-7890-1234-567890abcdef",
          "vendor_wallet": "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B",
          "proposal_name": "Proposal Sistem E-Arsip Cepat & Aman",
          "description": "Penawaran sistem e-arsip dengan teknologi terbaru...",
          "images": [],
          "requested_budget_wei": { "$numberDecimal": "48000000000000000000" },
          "ai_summary": "",
          "onchain_payload": {
            "targets": [],
            "values": [],
            "calldatas": []
          },
          "status": "SUBMITTED",
          "created_at": "2025-07-24T21:55:00.123Z",
          "updated_at": "2025-07-24T21:55:00.123Z",
          "deleted_at": null
        }
        ```

#### 2. Mendapatkan Semua Proposal untuk Sebuah Proyek

-   **`GET /projects/:id/proposals`**
    -   **Deskripsi**: Mengambil semua proposal yang telah diajukan untuk proyek dengan ID tertentu.
    -   **URL Params**:
        -   `id` (string, required): UUID dari proyek.
    -   **Response Sukses (200 OK)**:
        ```json
        [
          {
            "_id": "f0e9d8c7-b6a5-4321-fedc-ba9876543210",
            "proposal_name": "Proposal Sistem E-Arsip Cepat & Aman",
            "vendor_wallet": "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B",
            "status": "SUBMITTED",
            "...": "..."
          }
        ]
        ```

---

## 🏛️ Struktur Proyek (Clean Architecture)

-   **/cmd/api**: Titik masuk aplikasi (`main.go`).
-   **/internal/config**: Memuat konfigurasi dari `.env`.
-   **/internal/domain**: Berisi *struct* entitas inti (model data). Tidak memiliki dependensi ke lapisan lain.
-   **/internal/platform**: Berisi implementasi teknis seperti koneksi database (`database`) dan setup server HTTP (`server`).
-   **/internal/[fitur]**: Setiap fitur utama (misal: `profile`, `project`) memiliki modulnya sendiri, yang berisi:
    -   `handler.go`: Menangani request & response HTTP.
    -   `service.go`: Berisi logika bisnis (use case).
    -   `repository.go`: Berisi implementasi konkret dari interaksi dengan database.

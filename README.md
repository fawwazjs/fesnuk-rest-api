# fesnuk-rest-api
Penugasan Club Webdev x Hactiv8 Assignment Membangun REST API Social Media (User, Post, dan Like)   
Ahmad Wildan Fawwaz   
5027241001

# **Fesnuk REST API**
Fesnuk REST API adalah simulasi sederhana dari sebuah aplikasi media sosial, dibangun menggunakan Golang dan framework Gin. Data disimpan sepenuhnya di memori (RAM), sehingga semua data akan hilang setiap kali aplikasi dimatikan. API ini menyediakan endpoint untuk mengelola pengguna (User), postingan (Post), suka (Like), komentar (Comment), dan fitur follower/following.  

---

## **Daftar Isi**
1.  [Tujuan Proyek](#1-tujuan-penugasan)
2.  [Fitur Utama](#2-fitur-utama)
3.  [Struktur Proyek](#3-struktur-proyek)
4.  [Cara Menjalankan](#4-cara-menjalankan)
5.  [Daftar Endpoint & Contoh Penggunaan](#5-daftar-endpoint--contoh-penggunaan)
    * [User Endpoints](#user-endpoints)
    * [Post Endpoints](#post-endpoints)
    * [Like Endpoints](#like-endpoints)
    * [Comment Endpoints (Bonus)](#comment-endpoints)
    * [Follower/Following Endpoints (Bonus)](#followerfollowing-endpoints)
6.  [Struktur Response JSON](#6-struktur-response-json)
7.  [Screenshot Hasil Testing](#7-screenshot-hasil-testing)

---

## **1. Tujuan Penugasan**
* Membangun REST API sederhana menggunakan **Golang** dan *framework* **Gin**.
* Mensimulasikan aplikasi media sosial dengan fitur dasar seperti User, Post, dan Like.
* Menerapkan **validasi input** untuk memastikan integritas data.
* Menyediakan **format respons JSON yang konsisten**.
* Menggunakan **kode status HTTP yang sesuai** (misal: 200 OK, 201 Created, 400 Bad Request, 404 Not Found).
* Mengimplementasikan fitur **bonus** seperti komentar dan fitur *follower/following*.

---

## **2. Fitur Utama**
* **User Management:** Fungsionalitas lengkap untuk registrasi, melihat profil, memperbarui informasi, dan menghapus pengguna.
* **Post Management:** Membuat postingan baru, melihat semua postingan, melihat detail postingan, melihat postingan yang dibuat oleh user tertentu, dan menghapus postingan.
* **Like Functionality:** User dapat menyukai postingan, melihat daftar user yang menyukai post tertentu, dan melihat semua post yang disukai oleh seorang user.
* **Comment Functionality (Bonus):** Menambahkan komentar pada postingan dan mengambil semua komentar untuk post tertentu.
* **Follower/Following (Bonus):** Mengikuti dan berhenti mengikuti user lain, melihat daftar *follower* (pengikut), dan melihat daftar *following* (yang diikuti).
* **Post Filtering (Bonus):** Memfilter postingan berdasarkan kata kunci yang ada pada konten post.
* **Data in Memory:** Seluruh data disimpan sementara di RAM.
* **Consistent JSON Response:** Setiap respons API mengikuti struktur JSON yang sesuai.
* **Appropriate HTTP Status Codes:** Penggunaan kode status HTTP yang standar untuk mengindikasikan keberhasilan atau kegagalan operasi.
* **Basic Logging:** Middleware logger sederhana untuk memantau *request* masuk dan respons.

---

## **3. Struktur Proyek**
```
fesnuk-api/
├── main.go               # File utama untuk menjalankan aplikasi dan mendefinisikan rute
├── go.mod                # Mengelola dependensi proyek Go
├── go.sum                # Mengunci versi dependensi
├── models/               # Definisi struktur data (struct) untuk entitas API
│   ├── user.go           # Struktur User
│   ├── post.go           # Struktur Post
│   ├── like.go           # Struktur Like
│   ├── comment.go        # Struktur Comment
│   └── follower.go       # Struktur Follower
└── handlers/             # Logika(handler) untuk setiap endpoint API
    ├── common.go         # Fungsi helper untuk respons JSON yang konsisten
    ├── user_handler.go   # Handler untuk endpoint User
    ├── post_handler.go   # Handler untuk endpoint Post
    ├── like_handler.go   # Handler untuk endpoint Like
    ├── comment_handler.go # Handler untuk endpoint Comment
    └── follower_handler.go # Handler untuk endpoint Follower/Following
```

---

## **4. Cara Menjalankan**
1.  **Pastikan Go Terinstal:**
    Buka terminal atau Command Prompt dan ketik `go version`. Jika Go belum terinstal, unduh dan ikuti instruksi instalasi dari [go.dev/dl](https://go.dev/dl/).

2.  **Siapkan Direktori Proyek:**
    Jika belum punya, buat folder untuk proyek Anda dan masuk ke dalamnya:
    Jika mengunduh proyek ini dari Git, navigasi saja ke folder `fesnuk-api`.
    ```bash
    cd fesnuk-api
    ```
    

4.  **Inisialisasi Go Modules & Instal Dependensi:**
    Dari dalam direktori `fesnuk-api`, jalankan perintah berikut untuk menginisialisasi Go Modules dan mengunduh *library* yang diperlukan:
    ```bash
    go mod init fesnuk-api # Hanya jika membuat proyek dari awal
    go get [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
    go get [github.com/google/uuid](https://github.com/google/uuid)
    ```

5.  **Jalankan Aplikasi:**
    Di terminal yang sama, pastikan berada di direktori `fesnuk-api`, lalu jalankan perintah:
    ```bash
    go run .
    ```
    Aplikasi akan mulai berjalan dan *listening* pada port `8080`. Lalu output seperti:
    ```
    Fesnuk API Server running on port 8080
    [GIN-debug] Listening and serving HTTP on :8080
    ```
    **Biarkan terminal tetap terbuka selama melakukan pengujian API.**

---

## **5. Daftar Endpoint & Contoh Penggunaan**

Semua *endpoint* diakses melalui URL dasar: `http://localhost:8080`.

**Penting:**
* Karena data disimpan di memori, setiap kali me-restart server (`go run .`), **semua data akan hilang**. Maka perlu membuat ulang data (user, post, like, comment, follow) dari awal untuk pengujian.
* Saat melakukan *request* `POST` yang membuat data baru (user, post, like, comment), *response* akan menyertakan `id` unik. **Sangat penting untuk menyalin ID ini** karena akan menggunakannya untuk *request* `GET`, `PUT`, dan `DELETE` selanjutnya.

### **User Endpoints**
| Method | Endpoint              | Deskripsi                                    |
| :----- | :-------------------- | :------------------------------------------- |
| `POST` | `/users`              | Registrasi user baru                         |
| `GET`  | `/users`              | Ambil semua user                             |
| `GET`  | `/users/:id`          | Ambil profil user berdasarkan ID             |
| `PUT`  | `/users/:id`          | Update profil user berdasarkan ID            |
| `DELETE` | `/users/:id`        | Hapus user berdasarkan ID                    |

#### `POST /users` – Registrasi user baru
* **Deskripsi:** Membuat akun pengguna baru di Fesnuk.
* **Request Body (JSON):**
    ```json
    {
      "username": "asaka101",
      "email": "asaka101@fesnuk.com",
      "bio": "Pengguna Fesnuk yang aktif!"
    }
    ```
* **Response (201 Created):**
    ```json
    {
      "message": "User created successfully",
      "data": {
        "id": "a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d",
        "username": "asaka101",
        "email": "asaka101@fesnuk.com",
        "bio": "Pengguna Fesnuk yang aktif!"
      },
      "error": null
    }
    ```
* **Response (400 Bad Request - Validasi):**
    ```json
    {
      "message": "Username and email cannot be empty",
      "data": null,
      "error": "Validation error"
    }
    ```
* **Screenshot Postman:**
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/79c43dd0-f5ea-409f-a691-1a34230b2b6c" />
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/576be38d-5c16-45d6-b7af-f33183507b56" />
<img width="1269" height="731" alt="Image" src="https://github.com/user-attachments/assets/b3315dd4-2d06-4690-bc23-86f44fdeec76" />
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/0b9721aa-fa57-4f27-8983-de4022ed7dfa" />
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/07685237-655c-41f3-a3b1-a4e9e10e9043" />
<img width="1280" height="730" alt="Image" src="https://github.com/user-attachments/assets/64f1dd83-2ca0-4f97-a17a-d0b622b39972" />

#### `GET /users` – Ambil semua user
* **Deskripsi:** Mengambil daftar semua pengguna terdaftar di Fesnuk.
* **Request:** Tidak ada *body*.
* **Response (200 OK):**
    ```json
    {
      "message": "Users retrieved successfully",
      "data": [
        {
          "id": "a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d",
          "username": "asaka101",
          "email": "asaka101@fesnuk.com",
          "bio": "Pengguna Fesnuk yang aktif!"
        }
        // ... user lainnya
      ],
      "error": null
    }
    ```
* **Screenshot Postman:**
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/08bdb0d7-2936-4e9d-9ea5-213993c84592" />
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/ba9c915a-dbbc-46b9-a4fa-2e37998740f3" />

#### `GET /users/:id` – Ambil profil user
* **Deskripsi:** Mengambil detail profil pengguna tertentu berdasarkan ID.
* **URL Contoh:** `/users/a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d`
* **Response (200 OK):**
    ```json
    {
      "message": "User retrieved successfully",
      "data": {
        "id": "a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d",
        "username": "asaka101",
        "email": "asaka101@fesnuk.com",
        "bio": "Pengguna Fesnuk yang aktif!"
      },
      "error": null
    }
    ```
* **Response (404 Not Found):**
    ```json
    {
      "message": "User not found",
      "data": null,
      "error": "Not Found"
    }
    ```
* **Screenshot Postman:**
<img width="1284" height="725" alt="Image" src="https://github.com/user-attachments/assets/ed438a2f-478b-4e37-8152-117490b9e24d" />
<img width="1274" height="734" alt="Image" src="https://github.com/user-attachments/assets/bb17b8ce-5b5b-430f-8979-da9d8d85fdba" />

#### `PUT /users/:id` – Update profil user
* **Deskripsi:** Memperbarui informasi profil pengguna yang sudah ada.
* **URL Contoh:** `/users/a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d`
* **Request Body (JSON):**
    ```json
    {
      "username": "asaka115",
      "email": "asaka210@fesnuk.com",
      "bio": "Pengguna Fesnuk yang sangat aktif!"
    }
    ```
* **Response (200 OK):**
    ```json
    {
      "message": "User updated successfully",
      "data": {
        "id": "a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d",
        "username": "asaka115",
        "email": "asaka210@fesnuk.com",
        "bio": "Pengguna Fesnuk yang sangat aktif!"
      },
      "error": null
    }
    ```
* **Screenshot Postman:**
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/2c8c125d-f8bb-48bb-bd70-13e084060f78" />
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/fbaca9ef-31d5-47c1-bbd3-76e3284c8c70" />

#### `DELETE /users/:id` – Hapus user
* **Deskripsi:** Menghapus akun pengguna dari Fesnuk. Tindakan ini juga akan menghapus semua post, like, comment, dan relasi follow yang terkait dengan user ini.
* **URL Contoh:** `/users/a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d`
* **Response (200 OK):**
    ```json
    {
      "message": "User deleted successfully",
      "data": null,
      "error": null
    }
    ```
* **Screenshot Postman:**
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/f11d806a-b255-464e-96fe-48b4ede24ba0" />
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/b6f348fc-35a3-4bd0-ae77-193da9bf4621" />
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/90c7aa70-2906-4956-9cd0-58a112dffeff" />

---

### **Post Endpoints**

| Method | Endpoint              | Deskripsi                                    |
| :----- | :-------------------- | :------------------------------------------- |
| `POST` | `/posts`              | Buat postingan baru                          |
| `GET`  | `/posts`              | Ambil semua post (mendukung filtering keyword) |
| `GET`  | `/posts/:id`          | Ambil detail post berdasarkan ID             |
| `GET`  | `/users/:id/posts`    | Ambil semua post dari satu user              |
| `DELETE` | `/posts/:id`        | Hapus post berdasarkan ID                    |

#### `POST /posts` – Buat postingan baru
* **Deskripsi:** Membuat postingan baru oleh pengguna tertentu.
* **Request Body (JSON):**
    ```json
    {
      "user_id": "a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d", // ID Asaka (contoh)
      "content": "Senangnya Fesnuk ada fitur baru! Siapa setuju?"
    }
    ```
* **Response (201 Created):**
    ```json
    {
      "message": "Post created successfully",
      "data": {
        "id": "e6f7g8h9-i0j1-2k3l-4m5n-6o7p8q9r0s1t",
        "user_id": "a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d",
        "content": "Senangnya Fesnuk ada fitur baru! Siapa setuju?",
        "created_at": "2025-07-14T09:30:00Z"
      },
      "error": null
    }
    ```
* **Response (400 Bad Request - Validasi):**
    ```json
    {
      "message": "Post content cannot be empty",
      "data": null,
      "error": "Validation error"
    }
    ```
* **Screenshot Postman:**
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/a093b614-0d08-4b4d-9302-cd7a1a660379" />
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/95abddc7-5c62-4ada-995d-58f2e52a5b06" />
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/e56a54d2-d112-4df9-b62e-01ae41658949" />

#### `GET /posts` – Ambil semua post (mendukung filtering keyword)
* **Deskripsi:** Mengambil daftar semua postingan. Dapat difilter berdasarkan kata kunci pada konten post menggunakan *query parameter* `keyword`.
* **Request (Tanpa Keyword):** `http://localhost:8080/posts`
* **Request (Dengan Keyword):** `http://localhost:8080/posts?keyword=fesnuk`
* **Response (200 OK - Semua Post):**
    ```json
    {
      "message": "Posts retrieved successfully",
      "data": [
        {
          "id": "e6f7g8h9-i0j1-2k3l-4m5n-6o7p8q9r0s1t",
          "user_id": "a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d",
          "content": "Senangnya Fesnuk ada fitur baru! Siapa setuju?",
          "created_at": "2025-07-14T09:30:00Z"
        }
        // ... post lainnya
      ],
      "error": null
    }
    ```
* **Response (200 OK - Filtered Post):**
    ```json
    {
      "message": "Posts retrieved successfully",
      "data": [
        {
          "id": "e6f7g8h9-i0j1-2k3l-4m5n-6o7p8q9r0s1t",
          "user_id": "a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d",
          "content": "Senangnya Fesnuk ada fitur baru! Siapa setuju?",
          "created_at": "2025-07-14T09:30:00Z"
        }
      ],
      "error": null
    }
    ```
* **Screenshot Postman:**
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/6afb99e6-c81e-441f-ad6b-1e06a1772b1f" />
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/e83fc103-7a6d-4fa1-b684-c57155f1708b" />
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/f833b9a4-888a-4ba7-a5f9-12c4e62d330a" />

#### `GET /posts/:id` – Ambil detail post
* **Deskripsi:** Mengambil detail postingan tertentu berdasarkan ID.
* **URL Contoh:** `/posts/e6f7g8h9-i0j1-2k3l-4m5n-6o7p8q9r0s1t`
* **Response (200 OK):**
    ```json
    {
      "message": "Post retrieved successfully",
      "data": {
        "id": "e6f7g8h9-i0j1-2k3l-4m5n-6o7p8q9r0s1t",
        "user_id": "a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d",
        "content": "Senangnya Fesnuk ada fitur baru! Siapa setuju?",
        "created_at": "2025-07-14T09:30:00Z"
      },
      "error": null
    }
    ```
* **Screenshot Postman:**
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/67b67f47-6f83-4f50-b4ce-1635bc89dc8a" />

#### `GET /users/:id/posts` – Ambil semua post dari satu user
* **Deskripsi:** Mengambil semua postingan yang dibuat oleh pengguna dengan ID tertentu.
* **URL Contoh:** `/users/a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d/posts`
* **Response (200 OK):**
    ```json
    {
      "message": "User posts retrieved successfully",
      "data": [
        {
          "id": "e6f7g8h9-i0j1-2k3l-4m5n-6o7p8q9r0s1t",
          "user_id": "a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d",
          "content": "Senangnya Fesnuk ada fitur baru! Siapa setuju?",
          "created_at": "2025-07-14T09:30:00Z"
        }
        // ... post lainnya dari user yang sama
      ],
      "error": null
    }
    ```
* **Screenshot Postman:**
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/11b2e743-a3ce-4433-9635-0bea5117536c" />

#### `DELETE /posts/:id` – Hapus post
* **Deskripsi:** Menghapus postingan tertentu. Tindakan ini juga akan menghapus semua like dan komentar yang terkait dengan post tersebut.
* **URL Contoh:** `/posts/e6f7g8h9-i0j1-2k3l-4m5n-6o7p8q9r0s1t`
* **Response (200 OK):**
    ```json
    {
      "message": "Post deleted successfully",
      "data": null,
      "error": null
    }
    ```
* **Screenshot Postman:**
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/d0f4c687-fb39-4a5d-85a7-a953531f6d54" />
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/c5192121-02a6-4cd0-a1eb-8d3e5fc8657f" />
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/c48f5313-721d-423f-adae-5a0a79246acc" />

---

### **Like Endpoints**

| Method | Endpoint              | Deskripsi                                    |
| :----- | :-------------------- | :------------------------------------------- |
| `POST` | `/likes`              | User menyukai post                           |
| `GET`  | `/posts/:id/likes`    | Lihat siapa saja yang like post tertentu     |
| `GET`  | `/users/:id/likes`    | Lihat semua like dari seorang user           |

#### `POST /likes` – User menyukai post
* **Deskripsi:** Mencatat bahwa seorang user menyukai suatu postingan.
* **Request Body (JSON):**
    ```json
    {
      "user_id": "a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d", // ID Asaka (contoh)
      "post_id": "e6f7g8h9-i0j1-2k3l-4m5n-6o7p8q9r0s1t"  // ID  Post Yuzuru (contoh)
    }
    ```
* **Response (201 Created):**
    ```json
    {
      "message": "Like created successfully",
      "data": {
        "id": "f8g9h0i1-j2k3-4l5m-6n7o-8p9q0r1s2t3u",
        "user_id": "a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d",
        "post_id": "e6f7g8h9-i0j1-2k3l-4m5n-6o7p8q9r0s1t"
      },
      "error": null
    }
    ```
* **Response (409 Conflict - Sudah Like):**
    ```json
    {
      "message": "User has already liked this post",
      "data": null,
      "error": "Validation error"
    }
    ```
* **Screenshot Postman:**
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/f77d1fe0-a292-4cb4-982d-e943f47af172" />
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/a54f3978-d0ae-4439-83ba-5da39edbb620" />

#### `GET /posts/:id/likes` – Lihat siapa saja yang like post tertentu
* **Deskripsi:** Mengambil daftar semua like yang diterima oleh postingan tertentu.
* **URL Contoh:** `/posts/e6f7g8h9-i0j1-2k3l-4m5n-6o7p8q9r0s1t/likes`
* **Response (200 OK):**
    ```json
    {
      "message": "Likes for post retrieved successfully",
      "data": [
        {
          "id": "f8g9h0i1-j2k3-4l5m-6n7o-8p9q0r1s2t3u",
          "user_id": "a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d",
          "post_id": "e6f7g8h9-i0j1-2k3l-4m5n-6o7p8q9r0s1t"
        }
        // ... like lainnya untuk post yang sama
      ],
      "error": null
    }
    ```
* **Screenshot Postman:**
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/9eb95407-3c48-436a-85b7-8774018fe2f1" />

#### `GET /users/:id/likes` – Lihat semua like dari seorang user
* **Deskripsi:** Mengambil daftar semua like yang dibuat oleh user tertentu.
* **URL Contoh:** `/users/a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d/likes`
* **Response (200 OK):**
    ```json
    {
      "message": "Likes by user retrieved successfully",
      "data": [
        {
          "id": "f8g9h0i1-j2k3-4l5m-6n7o-8p9q0r1s2t3u",
          "user_id": "a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d",
          "post_id": "e6f7g8h9-i0j1-2k3l-4m5n-6o7p8q9r0s1t"
        }
        // ... like lainnya dari user yang sama
      ],
      "error": null
    }
    ```
* **Screenshot Postman:**
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/2ffc00b5-1cd6-43e3-8e34-873be57a158b" />
<img width="1366" height="768" alt="Image" src="https://github.com/user-attachments/assets/8d70ede2-b04c-4f86-9577-57b70b76bf54" />

---

### **Comment Endpoints**

| Method | Endpoint              | Deskripsi                                    |
| :----- | :-------------------- | :------------------------------------------- |
| `POST` | `/comments`           | Buat komentar baru untuk post                |
| `GET`  | `/posts/:id/comments` | Lihat semua komentar pada post tertentu      |

#### `POST /comments` – Buat komentar baru untuk post
* **Deskripsi:** Menambahkan komentar ke postingan tertentu oleh user.
* **Request Body (JSON):**
    ```json
    {
      "user_id": "b1c2d3e4-f5g6-7h8i-9j0k-1l2m3n4o5p6q", // ID Yuzuru (contoh)
      "post_id": "e6f7g8h9-i0j1-2k3l-4m5n-6o7p8q9r0s1t", // ID Post Asaka (contoh)
      "content": "Gwehj!!😎☝"
    }
    ```
* **Response (201 Created):**
    ```json
    {
      "message": "Comment created successfully",
      "data": {
        "id": "c1d2e3f4-g5h6-7i8j-9k0l-1m2n3o4p5q6r",
        "user_id": "b1c2d3e4-f5g6-7h8i-9j0k-1l2m3n4o5p6q",
        "post_id": "e6f7g8h9-i0j1-2k3l-4m5n-6o7p8q9r0s1t",
        "content": "Gwehj😎☝!!",
        "created_at": "2025-07-14T10:00:00Z"
      },
      "error": null
    }
    ```
* **Screenshot Postman:**
    **

#### `GET /posts/:id/comments` – Lihat semua komentar pada post tertentu
* **Deskripsi:** Mengambil semua komentar yang terkait dengan postingan tertentu.
* **URL Contoh:** `/posts/e6f7g8h9-i0j1-2k3l-4m5n-6o7p8q9r0s1t/comments`
* **Response (200 OK):**
    ```json
    {
      "message": "Comments for post retrieved successfully",
      "data": [
        {
          "id": "c1d2e3f4-g5h6-7i8j-9k0l-1m2n3o4p5q6r",
          "user_id": "b1c2d3e4-f5g6-7h8i-9j0k-1l2m3n4o5p6q",
          "post_id": "e6f7g8h9-i0j1-2k3l-4m5n-6o7p8q9r0s1t",
          "content": "Gwehj😎☝!!",
          "created_at": "2025-07-14T10:00:00Z"
        }
        // ... komentar lainnya untuk post yang sama
      ],
      "error": null
    }
    ```
* **Screenshot Postman:**
    **

---

### **Follower/Following Endpoints (Bonus)**

| Method | Endpoint                  | Deskripsi                                    |
| :----- | :------------------------ | :------------------------------------------- |
| `POST` | `/users/:id/follow`       | User melakukan follow ke user lain           |
| `DELETE` | `/users/:id/unfollow`     | User berhenti follow user lain               |
| `GET`  | `/users/:id/followers`    | Lihat daftar follower dari user tertentu     |
| `GET`  | `/users/:id/following`    | Lihat daftar user yang diikuti oleh user tertentu |

#### `POST /users/:id/follow` – User melakukan follow ke user lain
* **Deskripsi:** Mencatat bahwa seorang `follower_id` mulai mengikuti `user_id` yang ditentukan di URL.
* **URL Contoh:** `/users/a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d/follow` (Asaka yang difollow)
* **Request Body (JSON):**
    ```json
    {
      "follower_id": "b1c2d3e4-f5g6-7h8i-9j0k-1l2m3n4o5p6q" // Yuzuru yang follow
    }
    ```
* **Response (201 Created):**
    ```json
    {
      "message": "User followed successfully",
      "data": {
        "id": "g1h2i3j4-k5l6-7m8n-9o0p-1q2r3s4t5u6v",
        "follower_id": "b1c2d3e4-f5g6-7h8i-9j0k-1l2m3n4o5p6q",
        "following_id": "a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d"
      },
      "error": null
    }
    ```
* **Response (400 Bad Request - Follow Diri Sendiri):**
    ```json
    {
      "message": "User cannot follow themselves",
      "data": null,
      "error": "Validation error"
    }
    ```
* **Response (409 Conflict - Sudah Follow):**
    ```json
    {
      "message": "User already follows this user",
      "data": null,
      "error": "Validation error"
    }
    ```
* **Screenshot Postman:**
    **

#### `DELETE /users/:id/unfollow` – User berhenti follow user lain
* **Deskripsi:** Mencatat bahwa seorang `follower_id` berhenti mengikuti `user_id` yang ditentukan di URL.
* **URL Contoh:** `/users/a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d/unfollow` (Alice yang di-unfollow)
* **Request Body (JSON):**
    ```json
    {
      "follower_id": "b1c2d3e4-f5g6-7h8i-9j0k-1l2m3n4o5p6q" // Bob yang unfollow
    }
    ```
* **Response (200 OK):**
    ```json
    {
      "message": "User unfollowed successfully",
      "data": null,
      "error": null
    }
    ```
* **Screenshot Postman:**
    **

#### `GET /users/:id/followers` – Lihat daftar follower dari user tertentu
* **Deskripsi:** Mengambil daftar pengguna yang mengikuti user dengan ID tertentu.
* **URL Contoh:** `/users/a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d/followers` (Lihat follower Asaka)
* **Response (200 OK):**
    ```json
    {
      "message": "Followers retrieved successfully",
      "data": [
        {
          "id": "b1c2d3e4-f5g6-7h8i-9j0k-1l2m3n4o5p6q",
          "username": "yuzuru13",
          "email": "yuzuru13@fesnuk.com",
          "bio": "Koboi di Fesnuk."
        }
        // ... follower lainnya
      ],
      "error": null
    }
    ```
* **Screenshot Postman:**
    **

#### `GET /users/:id/following` – Lihat daftar user yang diikuti oleh user tertentu
* **Deskripsi:** Mengambil daftar pengguna yang diikuti oleh user dengan ID tertentu.
* **URL Contoh:** `/users/b1c2d3e4-f5g6-7h8i-9j0k-1l2m3n4o5p6q/following` (Lihat siapa yang diikuti Yuzuru)
* **Response (200 OK):**
    ```json
    {
      "message": "Following retrieved successfully",
      "data": [
        {
          "id": "a1b2c3d4-e5f6-4f7f-8c9d-0e1f2a3b4c5d",
          "username": "asaka115",
          "email": "asaka210@fesnuk.com",
          "bio": "Pengguna Fesnuk yang aktif!"
        }
        // ... user lain yang diikuti
      ],
      "error": null
    }
    ```
* **Screenshot Postman:**
    **

---

## **6. Struktur Response JSON**

Semua respons dari Fesnuk API akan mengikuti format JSON yang konsisten untuk memudahkan parsing dan penanganan:

```json
{
  "message": "Deskripsi singkat hasil operasi.",
  "data": { /* Objek atau array data yang relevan dengan request, bisa null jika tidak ada */ },
  "error": "Pesan error jika ada, bisa null jika tidak ada error"
}
```
* message: Sebuah string yang memberikan deskripsi singkat dan mudah dimengerti tentang hasil operasi (misal: "User created successfully", "User not found").
* data: Dapat berupa objek JSON tunggal (untuk detail satu entitas), array JSON (untuk daftar entitas), atau null. Bagian ini berisi data yang diminta atau data hasil dari operasi (misal: objek user yang baru dibuat, daftar postingan).
* error: Sebuah string yang berisi pesan error yang lebih spesifik jika terjadi kesalahan. Jika operasi berhasil, nilai ini akan menjadi null.

**Contoh Penggunaan Status Code HTTP:**
* 200 OK: Menunjukkan bahwa request berhasil diproses dan respons dikembalikan (digunakan untuk operasi GET, PUT, dan DELETE yang sukses).
* 201 Created: Menunjukkan bahwa request berhasil dan satu atau lebih sumber daya baru telah dibuat (digunakan untuk operasi POST yang sukses).
* 400 Bad Request: Menunjukkan bahwa request tidak valid karena masalah pada input yang diberikan (misal: validasi data gagal, body request salah format).
* 404 Not Found: Menunjukkan bahwa sumber daya yang diminta tidak ditemukan di server.
* 409 Conflict: Menunjukkan adanya konflik dalam request karena sumber daya yang mencoba dibuat sudah ada (misal: mencoba mengikuti user yang sudah diikuti).
* 500 Internal Server Error: Menunjukkan bahwa terjadi kesalahan tak terduga di sisi server yang mencegah request diproses.

## **7. Screenshot Hasil Testing**


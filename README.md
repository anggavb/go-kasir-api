# Kasir API

Ini adalah API sederhana untuk aplikasi kasir yang dibuat dengan Go.

## Deskripsi

API ini menyediakan endpoint untuk mengelola produk dan kategori.

## Cara Menjalankan

1.  Clone repositori ini.
2.  Jalankan perintah `go run main.go` untuk memulai server.
3.  Server akan berjalan di `http://localhost:8080`.

## Endpoint API

### Produk

-   **GET /api/product**
    -   Mendapatkan semua produk.

-   **GET /api/product/{id}**
    -   Mendapatkan produk berdasarkan ID.

-   **POST /api/product**
    -   Membuat produk baru.
    -   Contoh body:
        ```json
        {
          "nama": "Nama Produk",
          "harga": 10000,
          "stock": 100
        }
        ```

-   **PUT /api/product/{id}**
    -   Memperbarui produk berdasarkan ID.
    -   Contoh body:
        ```json
        {
          "nama": "Nama Produk Baru",
          "harga": 12000,
          "stock": 50
        }
        ```

-   **DELETE /api/product/{id}**
    -   Menghapus produk berdasarkan ID.

### Kategori

-   **GET /api/categories**
    -   Mendapatkan semua kategori.

-   **POST /api/categories**
    -   Membuat kategori baru.
    -   Contoh body:
        ```json
        {
          "name": "Nama Kategori"
        }
        ```

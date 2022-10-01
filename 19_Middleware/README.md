# (19) Middleware

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Middleware

### 1. Pengertian Middleware

Middleware adalah entitas yang menghubungkan proses request atau response sebuah server. Singkatnya, middleware merupakan blok kode yang dieksekusi oleh server antara setelah menerima request dari client dan sebelum mengembalikan response ke client.

### 2. Kegunaan Middleware

Middleware mengizinkan kita untuk mengimplementasikan berbagai macam fungsionalitas seperti yang didefinisikan, yaitu diantara HTTP sesudah request dan sebelum response. Adapun contoh kegunaan dari middleware yaitu, melakukan logging aktivitas yang terjadi pada sistem, pengecekan autentikasi user, handle CORS, handle trailing slash pada URI, dan masih banyak kegunaan lainnya.

### 3. Middleware pada Echo

Echo menyediakan banyak middleware, seperti semua yang disebutkan pada poin kedua. Namun di echo dibagi lagi menjadi dua bagian, yaitu middleware yang dijalankan sebelum router melakukan proses (`Echo#Pre`) dan middleware yang dijalankan sesudah router melakukan proses dan juga memiliki akses penuh pada API context echo (`Echo#Use`). Contoh Echo#Pre yaitu meng-override HTTP method dan menambahkan trailing slash. Kemudian contoh Echo#Use yaitu, membatasi ukuran dari request body dan logger.

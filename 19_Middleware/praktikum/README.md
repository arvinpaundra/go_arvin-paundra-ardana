# Middleware

#### Arvin Paundra Ardana - Golang A

### Mengimplementasikan Logging dan Autentikasi Lewat Middleware

Link [**solusi**](https://github.com/arvinpaundra/go_arvin-paundra-ardana/tree/master/19_Middleware/praktikum).

#### 1. Autentikasi menggunakan JWT

Untuk melakukan autentikasi menggunakan JWT untuk memproteksi route yang hanya diakses oleh user yang terdaftar pada sistem. Pada skema ini, saya melakukan pengecekan token yang tardapat pada Header `Authorization`. Sebelum itu, untuk mendapatkan token diperlukan login terlebih dahulu karena nantinya akan di-generate kan oleh server sebuah token JWT. Setelah mendapatkan token, token tersebut diletakkan pada Header Authorization. Untuk route yang terproteksi atau yang memerlukan token untuk memvalidasi bahwa user tersebut telah login, saya menggunakan middleware dari echo yaitu `middleware.JWT` dengan parameter berupa bentuk byte dari secret key untuk generate JWT yang nantinya akan digunakan untuk melakukan pengecekan valid tidaknya (jika ada) token tersebut. Jika ada dan valid, maka akan diteruskan ke route yang dituju.

Screenshot hasil :

- Sebelum login

![img1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/19_Middleware/screenshots/Screenshot_9.png)

- Melakukan login untuk mendapatkan token

![img2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/19_Middleware/screenshots/Screenshot_8.png)

- Setelah login dan melampirkan token pada header Authorization

![img3](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/19_Middleware/screenshots/Screenshot_10.png)

#### 2. Logging

Logging digunakan untuk menampilkan aktivitas seperti waktu akses, route tujuan dsb. Pada echo terdapat sebuah middleware yang dapat melakukan logging yaitu `middleware.LoggingWithConfig` yang menerima parameter config logging yang bertipe struct. Hasil logging ini nantinya akan ditampilkan pada terminal.

Screenshot hasil :

![img4](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/19_Middleware/screenshots/Screenshot_11.png)

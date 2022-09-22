# (16) Intro RESTful API

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Intro RESTful API

### 1. Application Programming Interface (API)

Application Programming Interface atau API adalah kumpulan dari fungsi dan prosedur yang menizinkan pembuatan aplikasi yang mengakses fitur atau data dari sebuah sistem operasi, aplikasi, atau service lainnya. Singkatnya API adalah sebuah sekumpulan fungsi, prosedur, dan sejenisnya yang menjadi penghubung atau jembatan antara program internal dengan program eksternal.

Pada kasus ketika membuat sebuah sistem yang berbasis dengan web, API sangat sering digunakan sebagai media komunikasi antara user web browser dengan server penyedia sistem tersebut untuk saling bertukar data. Contohnya ketika kita ingin membuat sebuah sistem yang terintegrasi dengan pembayaran online dan kita tidak ingin membuat sistem pembayaran online sendiri maka kita bisa menggunakan layanan penyedia pembayaran online, seperti contohnya yang terkenal adalah Midtrans. Untuk menghubungkan sistem internal kita dengan sistem milik Midtrans, kita dapat memanfaatkan yang namanya sebuah API.

### 2. REpresentational State Transfer (REST API)

REpresentational State Transfer atau disingkat dengan REST API adalah sebuah arsitektur pertukaran atau komunikasi data yang mana memanfaatkan protokol HTTP, yaitu request dan response. Request itu sendiri merupakan permintaan yang dikirimkan oleh user melalui web browser untuk nantinya diproses oleh server dan hasilnya akan dikembalikan sebagai response atau jawaban. Bentuk format response yaitu dapat berupa format JSON ataupun XML. Namun seiring berjalannya waktu, JSON lebih sering digunakan karena ringan dan mudah untuk dibaca oleh manusia dari pada XML.

Pada REST API terdapat sebuah HTTP method yang dimana itu merupakan sebuah bentuk representasi dalam melakukan mengaksesan ataupun manipulasi resource data. Contohnya, method GET digunakan untuk menampilkan atau memanggil data, POST untuk mengirimkan data ke server dan biasanya terdapat payload untuk menyimpan data untuk dikirim ke server, PUT untuk mengubah data, DELETE untuk menghapus data, dan masih banyak lagi yang lainnya.

RESTful API atau salah satu implementasi dari REST API juga terdapat HTTP Status Code atau Response Code yang dimana menunjukkan sebuah informasi (dinotasikan dengan angka) dari hasil kembalian yang dikirimkan oleh server ke user. Adapun untuk code-nya seperti 100-199 merupakan response yang bersifat informasional. Lalu, 200-299 berisi response sukses, seperti 200 untuk OK atau request berhasil mengakses sebuah resource. Selanjutnya, 300-399 berisi pesan redirect, seperti 304 yaitu Not Modified atau response tidak berubah dengan response sebelumnya saat mengakses resource yang sama. Kemudian, 400-499 dan 500-599 yang berisi tentang eror baik pada sisi user atau eror pada sisi server, seperti 404 Not Found yaitu user mengakses sebuah resource yang datanya tidak ada pada server dan 500 Internal Server Error yang berarti server tersebut mengalami eror dengan kondisi eror yang belum diketahui penyebabnya.

Selain Browser, ada beberapa tools yang bisa digunakan untuk melakukan request ke server dan biasanya sering digunakan pada saat proses testing atau debugging pada waktu development. Contohnya seperti Postman dan Insomnia.

### 3. REST API Design Best Practices

Untuk membuat sebuah REST API yang baik dan mudah untuk di-develop dikemudian hari maka diperlukan sebuah design yang baik dalam merancang sebuah REST API. Berikut adalah beberapa diantaranya.

a. Gunakan kata benda daripada kata kerja

Menggunakan kata benda lebih disarankan daripada menggunakan kata kerja aksi, karena menurut saya lebih menjelaskan bahwa kata benda ini lebih merepresentasikan bahwa ini dipergunakan untuk menyimpan sebuah resource data dan karena juga nantinya terdapat sebuah HTTP method yang dimana juga nantinya akan semakin menjelaskan untuk aksi-aksinya dalam proses manipulasi data tersebut.

Contoh :

```
Good :
GET /users
GET /users/123

Bad :
GET /getAllUsers
GET /getDetailUser/123
```

b. Gunakan kata jamak

Gunakan kata jamak daripada kata tunggal, karena nantinya di dialamnya pasti berisi banyak data dan bukan hanya berisi satu data.

Contoh :

```
Good:
GET /chats
GET /chats/123

Bad :
GET /chat
GET /chat/123
```

c. Gunakan Resource Nesting atau resource bersarang untuk menunjukkan relasi atau hirarki

Menggunakan Resource Nesting juga memperjelas bahwa terdapat beberapa entitas yang saling berhubungan atau berelasi yang membentuk sebuah kebersinambungan antara entitas tersebut.

Contoh :

```
GET /users
GET /user/123
GET /users/123/repositories/321
```

d. Filtering, Sorting, dan Paging

Best practice untuk melakukan filtering, sorting, dan paging sebuah data, kita dapat memanfaatkan dengan menggunakan query parameter. Query parameter sendiri diletakkan setelah nama endpoint dari API tersebut dan dipisahkan dengan tanda tanya (?) lalu diikuti dengan nama variable query parameter-nya.

Contoh :

```
Filtering :
GET /news?category=sport
GET /stories?lang=id

Sorting :
GET /users?sort=name:asc
GET /users?sort=name:desc

Paging :
GET /menus?limit=10&page=3
```

e. Versioning

Sesuai dengan namanya, agar kualitas API tersebut menjadi lebih baik misal dalam segi performance, keefisienan, keefektifan, maka dari pengembangan itulah diperlukannya sebuah versioning API. Salah satu tujuan versioning selain digunakan untuk menunjukkan keterbaharuan versi pengembangan API, versioning juga biasanya digunakan untuk tetap memberikan akses kepada pengguna API sebelumnya. Ini dilakukan karena biasanya pengguna atau service tersebut tidak kompatibel ketika menggunakan API versi terbaru dan harus dengan terpaksa menggunakan versi terdahulunya maka disitulah versioning API dapat membantu permasalahan tersebut.

Contoh :

```
https://developer.github.com/v3
```

f. API Documentation

API Documentation ini sangat krusial, karena disini merupakan berupa catatan-catatan pengembang dalam melakukan development REST API tersebut, mulai dari apa saja endpoint-nya, kemudian field apa yang dibutuhkan ketika mengakses tersebut, bagian mana saja yang perlu diberikan akses autentikasi dan otorisasi, dan sejenisnya.

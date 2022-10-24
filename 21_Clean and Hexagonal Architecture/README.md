# (21) Clean and Hexagonal Architecture

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada materi Clean and Hexagonal Architecture

### 1. Kendala Sebelum Mendesign Clean Architecture

Berikut ini merupakan kendala yang bisa diselesaikan dengan menggunakan Clean Architecture :

- Independent of Framework

Pada Clean Architecture mengijinkan kita untuk menggunakan seperti contohnya framework sebagai sebuah tools.

- Testable

Business logic bisa ditesting tanpa harus menggunakan UI, database, web server, atau elemen eksternal lainnya.

- Independent of UI

User Interface atau UI bisa diubah dengan mudah, tanpa harus mengubah sistem lainnya.

- Independent of Database

Business logic tidak terikat dengan database, jadi kita bisa mengubah atau swap database dengan mudah.

- Independent of Any External

Pada kenyataannya, business logic tidak perlu tahu tentang apa yang ada diluar lingkaran dari diagram Clean Architecture, seperti contohnya tidak perlu tahu framework apa yang digunakan atau arsitektur komunikasi pertukaran data apa yang akan digunakan.

### 2. Benefit dari Clean Architecture

Berikut merupakan keuntungan dari Clean Architecture :

- Clean Architecture merupakan sebuah struktur standar.

- Development lebih cepat dalam jangka panjang.

- Mocking dependency menjadi lebih mudah dalam melakukan unit test.

- Mudah mengganti dari prototype ke solusi yang sesuai.

### 3. Struktur dari Clean Architecture

Berikut merupakan struktur dari Clean Architecture :

- Entities atau Repository layer

Business object yang mencerminkan konsep dari aplikasi yang dikelola.

- Use case atau Domain atau Service layer

Berisikan business logic dari aplikasi yang didevelop.

- Controller atau Presentation layer

Mempresentasikan data ke sebuah screen dan menangani interaksi pengguna.

- Drivers atau Data layer

Mengatur data aplikasi, contohnya menerima data dari network atau mengatur data cache.

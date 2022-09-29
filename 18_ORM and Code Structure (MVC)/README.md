# (18) ORM and Code Structure MVC

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi ORM and Code Structure

### 1. Definisi ORM

Object Relational Mapping atau ORM merupakan sebuah teknik pemrograman dengan mengkonversi data diantara sistem yang tidak kompatibel menggunakan paradigma pemrograman OOP. Salah satu library ORM golang yang terkenal yaitu adalah `gorm`. Pada gorm sangat banyak fitur yang mendukung proses development, seperti database relationship, join, migration, seeder, dan masih banyak lagi. Gorm mendukung database migration yang dimana merupakan sebuah mekanisme untuk meng-update versi database untuk memenuhi perubahan versi aplikasi. Update ini sendiri beisa mencakup ke versi terbaru atau melakukan rollback untuk ke versi sebelumnya. Database migration ini sangat berguna karena dapat melakukan update dan rollback secara mudah, melacak struktur pada database, menuliskan riwayat struktur database pada sebuah kode, dan akan selalu kompatibel dengan perubahan aplikasi.

### 2. Kelebihan dan Kekurangn Menggunakan ORM

Penggunaan ORM sendiri memiliki beberapa kelebihan dan kekurangan, seperti sebagai berikut.

a. Kelebihan ORM

- Mengurangi penulisan query yang berulang.

- Secara otomatis melakukan konversi data menjadi bentuk object.

- Dapat melakukan screening data atau pengecekan data dengan mudah.

- Beberapa ORM mendukung caching pada query-nya.

b. Kekurangan ORM

- Menambahkan baris kode baru yang menyebabkan proses berlebih.

- Memuat data relasi yang tidak diperlukan.

- Tidak efisien jika digunakan untuk melakukan query yang kompleks, seperti join lebih dari 10 tabel.

- Tidak mendukung beberapa fungsi SQL bawaan, seperti Force Index pada MySQL.

### 3. Model - View - Controller (MVC)

MVC merupakan sebuah cara yang populer untuk melakukan pengorganisiran folder atau kode pada sebuah projek. Ide dibalik MVC yitu masing-masing bagian dari kode tersebut memiliki tujuan dan masing-masing tujuan tersebut berbeda antara satu dengan lainnya. Tujuan diperlukannya strukturisasi yaitu untuk membuat aplikasi menjadi modular dengan mengimplementasikan memisahkan fokus urusan per masing-masing bagian dan juga mengurangi konflik ketika versioning.

Adapun fungsi masing-masing dari tiga bagian MVC :

a. Model

Model merupakan bagian yang bertugas untuk menyiapkan, mengatur, memanipulasi, dan mengorganisasikan data yang ada di database.

b. View

View merupakan bagian yang bertugas untuk menampilkan informasi dalam bentuk Graphical User Interface (GUI).

c. Controller

Controller merupakan bagian yang bertugas untuk menghubungkan serta mengatur model dan view agar dapat saling terhubung.

# (12) Database - DDL - DML

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Database - DDL - DML

### 1. Pengertian Database

Database merupakan sekumpulan data yang disimpan secara terorganisir. Database dibedakan menjadi SQL yang mana direpresentasikan ke dalam bentuk tabel dan NoSQL direpresentasikan ke dalam sebuah document yang terdiri dari sebuah key dan value. Contoh dari SQL yaitu MySQL dan PostgreSQL, sedangkan NoSQL yaitu MongoDB dan Redis.

Pada database tardapat hubungan atau disebut sebagai `Database Relationship`. Berikut macam-macam database relationship.

#### a. One-to-one relationship

One-to-one relationship merupakan hubungan dimana data dalam satu tabel dapat berhubungan hanya dengan satu data pada tabel lain. Contoh implementasinya yaitu, relasi antara produk dengan deskripsi produk, dimana satu produk hanya memiliki satu deskripsi dan begitu pula sebaliknya.

#### b. One-to-many relationship

One-to-many relationship merupakan hubungan dimana sebuah data pada satu tabel berhubungan dengan nol, satu, atau lebih data dari tabel lain. Contoh implementasinya yaitu, relasi antara tabel customer dengan tabel riwayat belanja, dimana satu customer dapat memiliki lebih dari satu riwayat pembelian.

#### c. Many-to-many relationship

Many-to-many relationship merupakan hubungan dimana beberapa data pada sebuah tabel berhubungan dengan beberapa data pada tabel lain. Contoh implementasinya yaitu, relasi antara tabel mahasiswa dengan tabel matakuliah, diamana satu mahasiswa dapat memiliki banyak matakuliah dan 1 matakuliah dapat dimiliki oleh banyak mahasiswa.

### 2. Database Definition Language (DDL)

Database Definition Language atau DDL merupakan perintah yang digunakan oleh database untuk membuat atau memodifikasi terkait struktur pada database, seperti tabel, index, dsb.

Contoh perintah DDL :

#### a. Perintah `CREATE`

Sesuai dengan namanya, create merupakan salah satu perintah DDL yang dimana bertujuan untuk membuat sebuah database, tabel, atau yang lainnya.

Contoh :

```sql
-- Membuat database
CREATE DATABASE altera_online_shop;

-- Menggunakan database
USE altera_online_shop;

-- Membuat tabel
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    address TEXT
);
```

#### b. Perintah `ALTER`

Perintah alter biasanya digunakan untuk mengubah struktur dari sebuah tabel, seperti mengubah nama dan kolom, menghapus sebuah kolom, mengganti tipe data dari sebuah kolom, dan sebagainya.

Contoh :

```sql
-- Menambahkan sebuah kolom ke dalam tabel users
ALTER TABLE user ADD join_date DATETIME;
```

#### c. Perintah `DROP`

Perintah drop digunakan untuk mengahapus sebuah database atau menghapus sebuah tabel atau bisa juga untuk menghapus sebuah kolom pada sebuah tabel.

Contoh :

```sql
-- Membuat database
CREATE DATABASE twitter;

-- Menghapus database
DROP DATABASE twitter;
```

#### d. Perintah `RENAME`

Perintah rename biasa digunakan untuk mengganti nama sebuah tabel atau index.

Contoh :

```sql
--Mengubah nama tabel users
RENAME TABLE users TO customers;
```

### 3. Database Manipulation Language (DML)

Database Manipulation Language atau DML merupakan perintah untuk mengubah atau memanipulasi data yang ada pada sebuah database.

- Contoh DML Statement Operation :

#### a. Perintah `INSERT`

Perintah insert digunakan untuk menambahkan data pada sebuah tabel.

Contoh :

```sql
-- Menambahkan data pada tabel users
INSERT INTO users (id, name, address) VALUES(1, "Arvin Paundra", "Cilacap");
```

#### b. Perintah `SELECT`

Perintah select digunakan untuk mengambil atau menampilkan data dari satu atau lebih tabel. Pada saat mengambil data pada tabel bisa menggunakan kondisi sesuai dengan kondisi yang diperlukan.

Contoh :

```sql
-- Menampilkan semua data pada tabel
SELECT * FROM users;

-- Menampilkan kolom tertentu pada tabel dengan ditambahkan kondisi
SELECT id, name FROM user WHERE id = 1;
```

#### c. Perintah `UPDATE`

Perintah update digunakan untuk mengubah data pada tabel. Pada pengubahannya harus ditambahkan sebuah kondisi jika hanya ingin mengubah data tertentu.

Contoh :

```sql
-- Mengubah data nama user dengan id = 1
UPDATE users SET name = "Arvin Alterra" WHERE id = 1;
```

#### d. Perintah `DELETE`

Perintah delete digunakan untuk menghapus data pada sebuah tabel. Sama seperti update, jika ingin menghapus data tertentu maka gunakan sebuah kondisi untuk menghapusnya.

Contoh :

```sql
-- Menghapus data pada tabel users dengan id = 10
DELETE FROM users WHERE id = 10;
```

- Contoh DML Statement :

#### a. Perintah `LIKE`

Perintah like digunakan untuk melakukan pencarian data pada saat menggunakan perintah select.

Contoh :

```sql
-- Mencari data pada tabel users dengan nama yang terdapat arvin di dalamnya
SELECT * FROM users WHERE name LIKE "%arvin%";
```

#### b. Perintah `BETWEEN`

Perintah between digunakan untuk menyeleksi data sesuai dengan range atau jangkauan yang telah ditentukan.

Contoh :

```sql
-- Menampilkan data users dari id dari 1 sampai 10
SELECT * FROM users WHERE id BETWEEN 1 AND 10;
```

#### c. Perintah `ORDER BY`

Perintah order by digunakan untuk mengurutkan data yang telah diseleksi secara menaik atau menurun sesuai dengan kolom yang telah ditentukan

Contoh :

```sql
-- Menampilkan data users dengan urutan name dari bawah
SELECT * FROM users ORDER BY name DESC;
```

#### d. Perintah `LIMIT`

Perintah limit digunakan untuk membatasi data yang akan ditampilkan sesuai dengan yang telah ditentukan. Biasanya digunakan pada pagination.

Contoh :

```sql
-- Menampilkan data users dengan batas data yang ditampilkan yaitu 10 baris
SELECT * FROM users LIMIT 10;
```

# (13) Join - Union - Agregasi - Subquery - Function (DBMS)

#### Arvin Paundra Ardana - Golang A

## Poin Penting pada Materi Join - Union - Agregasi - Subquery - Function

### 1. Join dan Union Tabel SQL

Untuk menggabungkan beberapa tabel pada SQL dapat meggunakan perintah `JOIN` atau `UNION`. Perbedaan antara join dengan union yaitu, join menggunakan klausa untuk mengkombinasikan record dari dua atau lebih tabel, sedangkan union mengkombinasikan kumpulan hasil dari dua atau lebih perintah select (tanpa klausa).

a. Jenis-jenis perintah pada join :

- Inner Join

Perintah ini digunakan untuk melakukan join tabel dengan menyeleksi dan return record yang sama pada kedua tabel atau mengembalikan record dari dua tabel atau lebih yang memenuhi syarat.

Contoh :

```sql
> SELECT u.*, r.role FROM users u
> INNER JOIN roles r ON u.role_id = r.id;
```

- Left Join

Mengembalikan semua record dari tabel sebelah kiri dan record yang sama pada tabel sebelah kanan.

Contoh :

```sql
> SELECT p.*, pt.name FROM products p
> LEFT JOIN product_types pt ON p.product_type_id = pt.id;
```

- Right Join

Mengembalikan semua record dari tabel sebelah kanan dan record yang sama pada tabel sebelah kiri.

Contoh :

```sql
> SELECT td.*, p.name AS product_name FROM transaction_details td
> RIGHT JOIN products p ON p.id = td.product_id;
```

b. Implementasi dari union :

```sql
> SELECT * FROM transactions WHERE id = 1
> UNION
> SELECT * FROM transactions WHERE id = 2;
```

### 2. Fungsi Agregasi pada SQL

Fungsi agregasi adalah fungsi yang digunakan untuk melakukan perhitungan pada beberapa nilai dan mengembalikan hasilnya dalam satu nilai seperti rata-rata semua nilai, jumlah semua nilai, dan nilai maksimum & minimum di antara kelompok nilai tertentu.

Macam-macam fungsi agregat SQL :

#### a. Fungsi MIN()

Fungsi MIN digunakan untuk mengembalikan nilai terkecil dari sebuah kolom.

Contoh mengambil saldo karyawan yang terkecil pada tabel employees :

```sql
> SELECT MIN(balance) FROM employees;
```

#### b. Fungsi MAX()

Fungsi MAX digunakan untuk mengembalikan nilai terbesar dari sebuah kolom.

Contoh mengambil saldo karyawan yang terbesar pada tabel employees :

```sql
> SELECT MAX(balance) FROM employees;
```

#### c. Fungsi SUM()

Fungsi SUM digunakan untuk menjumlahkan semua data pada tabel tertentu.

Contoh menjumlahkan seluruh saldo karyawan pada tabel employees :

```sql
> SELECT SUM(balance) FROM employees;
```

#### d. Fungsi AVG()

Fungsi AVG digunakan untuk mencari nilai rata-rata atau mean pada tabel tertentu.

Contoh mengitung rata-rata saldo karyawan pada tabel employees :

```sql
> SELECT AVG(balance) FROM employees;
```

#### e. Fungsi COUNT()

Fungsi COUNT digunakan untuk mengembalikan jumlah baris dari dari proses seleksi data.

Contoh jumlah karyawan pada tabel employees :

```sql
> SELECT COUNT(name) FROM employees;
```

#### f. Fungsi HAVING()

Fungsi HAVING digunakan untuk menyeleksi data berdasarkan kriteria tertentu, dimana kriteria berupa fungsi agregat.

Contoh menyeleksi karyawan yang memiliki id lebih dari 10 :

```sql
> SELECT * FROM employees
> GROUP BY name
> HAVING COUNT(id) > 10;
```

### 3. Fungsi pada SQL

Fungsi merupakan sebuah kumpulan statement yang akan mengembalikan sebuah nilai balik pada pemanggilnya. Trigger merupakan salah satu jenis dari fungsi yang ada pada SQL.

Contoh fungsi trigger secara otomatis menghapus data transaction_details apabila id dari transactions yang dihapus itu ada di tabel transaction_details :

```sql
> DELIMITER $$
> CREATE TRIGGER delete_transaction_details
> BEFORE DELETE ON transactions FOR EACH ROW
> BEGIN
> DECLARE v_transaction_id INT;
> SET v_transaction_id = OLD.id;
> DELETE FROM transaction_details WHERE transaction_id = v_transaction_id;
> END; $$
```

# Database - DDL - DML

#### Arvin Paundra Ardana - Golang A

### Part 1 - Schema Database

Mengimplementasikan schema database dengan tema Digital Outlet Pulsa. Link [ERD](https://drive.google.com/file/d/1I2WmJsIjkvqtxOCTqwItHP_mHj97F_OS/view?usp=sharing).

Schema ERD yang saya susun sebagai berikut.

![ERD](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/ERD-outlet-pulsa.png)

Penjelasan :

- Tabel product_description dengan tabel products memliki relasi **one-to-one**, dimana satu product hanya boleh memiliki satu deskripsi, dan begitu juga sebaliknya.

- Tabel product_types dengan tabel products memiliki relasi **one-to-many**, dimana satu tipe produk bisa dimiliki oleh banyak produk dan produk hanya boleh memiliki satu tipe produk.

- Tabel operators dengan tabel products memiliki relasi **one-to-many**, dimana operator bisa dimiliki oleh banyak produk dan produk hanya boleh memiliki satu operator.

- Tabel products dengan tabel transactions memiliki relasi **one-to-many**, dimana satu produk dapat digunakan pada banyak transaksi dan transaksi hanya dapat membeli satu tipe produk.

- Tabel transactions dengan tabel transaction_detail memiliki relasi **one-to-one**, dimana transaksi hanya memiliki satu detail transaksi dan detail transaksi hanya berhubungan dengan satu transaksi.

- Tabel transaction_detail dengan tabel payment_methods memiliki relasi **many-to-one**, dimana detail transaksi hanya bisa bertransaksi dengan satu metode pembayaran dan methode pembayaran bisa digunakan oleh banyak transaksi.

- Tabel transactions dengan tabel users memiliki relasi **one-to-many**, dimana transaksi hanya bisa dibuat oleh masing-masing pengguna dan pengguna bisa membuat banyak transaksi.

- Tabel histories dengan tabel users memiliki relasi **one-to-many**, dimana riwayat belanja hanya bisa dibuat oleh masing-masing pengguna dan pengguna dapat memiliki banyak riwayat belanja.

- Tabel histories dengan tabel transactions memiliki relasi **one-to-one**, dimana riwayat belanja hanya dapat memiliki satu relasi dengan transaksi dan begitu juga sebaliknya.

- Tabel histories dengan tabel transaction_detail memiliki relasi **one-to-one**, dimana riwayat belanja hanya dapat memiliki satu relasi dengan detail transaksi dan begitu juga sebaliknya.

### Part 2 - Data Definition Language

#### 1. Membuat database `alta_online_shop`

Untuk membuat database baru dapat menggunakan perintah `CREATE` dan untuk menggunakan database tersebut dapat menggunakan perintah `USE`.

Implementasi :

![CREATE DB](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_2.png)

#### 2. Implementasi ERD ke dalam bentuk DDL

- Membuat tabel users

Membuat tabel users dengan kolom-kolomnya yaitu antara lain id (PK), name, address, birth_date, status_user, gender, email, password, created_at, updated_at.

Implementasi :

![tabel users](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_3.png)

- Membuat tabel product_types

Membuat tabel product_types dengan kolom-kolomnya yaitu antara lain id (PK), type, created_at, updated_at.

Implementasi :

![tabel product_types](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_4.png)

- Membuat tabel operators

Membuat tabel operators dengan kolom-kolomya yaitu antara lain id (PK), operator_name, created_at, updated_at.

Implementasi :

![tabel operators](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_5.png)

- Membuat tabel products

Membuat tabel products dengan kolom-kolomnya yaitu antara lain id (PK), product_type_id (FK), operator_id (FK), product_name, price, created_at, updated_at.

Implementasi :

![tabel products](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_7.png)

- Membuat tabel product_description

Membuat tabel product_description dengan kolom-kolomnya yaitu antara lain id (PK), product_id (FK), description, created_at, updated_at.

Implementasi :

![tabel product_description](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_8.png)

- Membuat tabel payment_methods

Membuat tabel payment_methods dengan kolom-kolomnya yaitu antara lain id (PK), method_name, created_at, updated_at.

Implementasi :

![payment_methods](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_9.png)

- Membuat tabel transactions

Membuat tabel transactions dengan kolom-kolomnya yaitu antara lain id (PK), user_id (FK), product_id (FK), invoice.

Implementasi :

![tabel transactions](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_11.png)

- Membuat tabel transaction_detail

Membuat tabel transacion_detail dengan kolom-kolomnya yaitu antara lain id (PK), transaction_id (FK), payment_method_id (FK), total_price, created_at, updated_at.

Implementasi :

![tabel transaction_detail](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_12.png)

- Membuat tabel histories

Membuat tabel histories dengan kolom-kolomnya yaitu antara lain id (PK), user_id (FK), product_id (FK), transaction_id (FK), transaction_detail (FK), status, created_at, updated_at.

Implementasi :

![tabel histories](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_16.png)

#### 3. Membuat tabel `kurir`

Menambahkan tabel kurir dengan kolom-kolomnya yaitu antara lain id (PK), name, created_at, updated_at. Untuk menambahkan tabel pada sebuah database dapat menggunakan perintah `CREATE TABLE`.

Implementasi :

![tabel couriers](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_17.png)

#### 4. Menambahkan kolom `ongkos_dasar` pada tabel kurir

Menambahkan satu kolom baru yaitu ongkos_dasar pada tabel kurir. Untuk menambahkan kolom pada sebuah tabel dapat dilakukan menggunakan perintah `ALTER TABEL`.

Implementasi :

![kolom ongkos_dasar](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_18.png)

#### 5. Mengubah nama tabel kurir menjadi `shipping`

Mengubah nama tabel kurir menjadi shipping dapat dilakukan menggunakan perintah `RENAME TABLE`.

Implementasi :

![tabel shipping](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_19.png)

#### 6. Menghapus tabel shipping

Menghapus tabel shipping karena dirasa tidak diperlukam dapat dilakukan dengan menggunakan perintah `DROP TABLE`.

Implementasi :

![drop table shipping](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_20.png)

#### 7. Menambahkan jenis relasi baru

Menambahkan relasi baru antara lain sebagai berikut.

#### a. Menambahkan relasi one-to-one

Menambahkan relasi one-to-one yaitu antara tabel payment_methods dengan tabel payment_method_description. Yaitu dimana satu metode pembayaran hanya dapat berelasi dengan satu deskripsi metode pembayaran dan begitu juga sebaliknya.

Implementasi :

![one-to-one](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_21.png)

#### b. Menambahkan relasi one-to-many

Menambahkan relasi one-to-many yaitu antara tabel users dengan tabel addresses. Yaitu dimana satu pengguna bisa memiliki banyak alamat dan alamat tersebut hanya memiliki masing-masing satu pengguna. Sebelum menambahkan tabel addresses, hapus dahulu kolom address pada tabel users.

Implementasi :

![one-to-many](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_22.png)

#### c. Menambahkan relasi many-to-many

Menambahkan relasi many-to-many yaitu antara tabel users dengan tabel user_payment_method_detail. Yaitu dimana pengguna memiliki banyak metode pambayaran seperti via bank, e-wallet dsb yang ada pada tabel user_payment_method_detail dan tabel user_payment_method_detail dapat berelasi atau digunakan oleh banyak pengguna.

Implementasi :

![many-to-many](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/12_Database%20-%20DDL%20-DML/screenshots/Screenshot_23.png)

# Join - Union - Agregasi - Subquery - Function (DBMS)

#### Arvin Paundra Ardana - Golang A

Link [File SQL](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/praktikum/queries.sql>).

## Insert - Select - Update - Delete

### 1. Perintah Insert

#### a. Insert 5 operators ada table operators.

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_6.png>)

#### b. Insert 3 product type

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_7.png>)

#### c. Insert 2 product dengan product type id = 1, dan operators id = 3

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_8.png>)

#### d. Insert 3 product dengan product type id = 2, dan operators id = 1

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_9.png>)

#### e. Insert 3 product dengan product type id = 3, dan operators id = 4

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_10.png>)

#### f. Insert product description pada setiap product

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_11.png>)

#### g. Insert 3 payment methods

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_12.png>)

#### h. Insert 5 user pada tabel users

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_13.png>)

#### i. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_14.png>)

#### j. Insert 3 product di masing-masing transaksi

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_15.png>)

### 2. Perintah Select

#### a. Tampilkan nama user / pelanggan dengan gender Laki-laki atau M

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_16.png>)

#### b. Tampilkan product dengan id = 3

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_17.png>)

#### c. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung huruf 'a'

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_18.png>)

#### d. Hitung jumlah user / pelanggan dengan status gender Perempuan

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_19.png>)

#### e. Tampilkan data pelanggan dengan urutan sesuai nama abjad

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_20.png>)

#### f. Tampilkan 5 data pada data product

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_21.png>)

### 3. Perintah Update

#### a. Ubah data product id 1 dengan nama 'product dummy'

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_22.png>)

#### b. Update qty = 3 pada transaction detail dengan product id = 1

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_23.png>)

### 4. Perintah Delete

#### a. Delete pada data table product dengan id = 1

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_24.png>)

#### b. Delete pada tabel product dengan product type id = 1

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_25.png>)

## Join - Union - Subquery - Function

#### 1. Gabungkan data transaksi dari user id 1 dan user 2

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_26.png>)

#### 2. Tampilkan jumlah harga transaksi user id 1

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_27.png>)

#### 3. Tampilkan total transaksi dengan product type 2

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_28.png>)

#### 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_29.png>)

#### 5. Tampilkan semua field table transaction, field name table product dan field name table user

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_30.png>)

#### 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_31.png>)

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_32.png>)

Transaksi detail dengan id 2 terhapus ketika transaksi id 2 dihapus

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_33.png>)

#### 7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_34.png>)

Menghapus transaksi detail dengan id 15

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_35.png>)

Transaksi detail dengan id 15 telah terhapus

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_36.png>)

Total qty pada transaksi dengan id 15 berkurang menjadi 0

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_37.png>)

#### 7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus

![img](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/13_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Screenshot_38.png>)

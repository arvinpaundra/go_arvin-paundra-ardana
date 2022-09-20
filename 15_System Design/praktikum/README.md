# System Design

#### Arvin Paundra Ardana - Golang A

### Problem 1 - Diagram

#### a. Entity Relationship Diagram

Link [ERD](https://drive.google.com/file/d/1dqpgFFH3SPWFUMOZ3k0UeD6vaYhlwZ3r/view?usp=sharing)

Membuat ERD pada sistem yang saya namai dengan nama Expenses System. Pada ERD ini, saya menentukan bahwa terdapat 3 entitas yaitu antara lain, users, expense_types, dan expenses.

Penjelasan masing-masing entitas :

- Entitas `users`

Pada entitas ini, saya membuat field yang terdiri dari id, fullname, role, gender, balance, email, password, created_at, dan updated_at. Entitas users bertujuan untuk menyimpan data-data dari user dan juga termasuk balance atau saldo terbaru dari user tersebut.

- Entitas `expense_types`

Pada entitas ini, saya membuat field yang terdiri dari id, name, created_at, dan updated_at. Entitas expense_types bertujuan untuk menyimpan tipe pengeluaran tersebut, contoh untuk keperluan konsumsi, transportasi, dsb.

- Entitas `expenses`

Pada entitas ini, saya membuat field yang terdiri dari id, user_id, expense_type_id, item_name, qty, total_expenses, created_at, dan updated_at. Entitas expenses bertujuan untuk mencatat pengeluaran dari user yang dimana direlasikan antara users.id dengan expenses.user_id untuk foreign key nya. Lalu, tipe dari pengeluaran tersebut yang direlasikan dengan expense_type.id dengan expenses.expense_type_id sebagai foreign key. Selanjutnya, terdapat item_name untuk nama barangnya, qty untuk jumlah barangnya, dan total_expenses untuk total biaya pengeluarannya.

Adapun untuk database relationship-nya yaitu antara lain :

- Relasi antara `users` dengan `expenses` yaitu **one-to-many**, karena satu user bisa memiliki banyak data pengeluaran.

- Relasi antara `expense_types` dengan `expenses` yaitu **one-to-many**, karena expense_type atau tipe pengeluaran dapat digunakan oleh banyak data expenses atau pengeluaran.

Hasil ERD :

![ERD](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/15_System%20Design/screenshots/ERD-expenses-system.png)

#### b. Use Case Diagram

Untuk pada use case, saya menentukan bahwa terdapat dua aktor, yaitu user dan admin. Beriktu penjabarannya.

#### 1. User

- User dapat Login ke sistem.

- Pada user dapat menambahkan data pengeluaran dengan mengisikan field yang sudah ditentukan pada tahap ERD.

- User dapat melihat semua data pengeluaran dari awal input sampai input terbaru.

- User dapat melakukan pencarian data pengeluaran, baik berdasarkan tipe pengeluaran ataupun tanggal pengeluaran. Pada use case ini include untuk menampilkan detail data yang dicari.

#### 2. Admin

- Admin dapat Login ke sistem.

- Admin dapat menambahkan data tipe pengeluaran untuk nantinya digunakan oleh user untuk menambahkan pengeluaran sesuai dengan pengeluarannya tersebut.

Hasil Use Case :

![use-case](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/15_System%20Design/screenshots/Use-case-expenses-system.png)

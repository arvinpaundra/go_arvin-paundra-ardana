# ORM and Code Structure

#### Arvin Paundra Ardana - Golang A

### Problem 1 - API CRUD User Using Database

Link solusi [**Problem 1**](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/tree/master/18_ORM%20and%20Code%20Structure%20(MVC)/praktikum/Problem%201>).

Membuat sebuah REST API sederhana dengan menggunakan framework `echo` dengan terkoneksi dengan ORM `gorm` untuk melakukan manipulasi data pada database MySQL.

Berikut merupakan langkah yang saya terapkan untuk CRUD :

a. Menambahkan user

Untuk menambahkan user meggunakan gorm, dapat menggunakan method `Save` yang dimana menerima sebuah parameter yang mereferensikan ke alamat memori data yang berisi payload yang sesuai dengan struct User.

Screenshot hasil :

![img1](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/18_ORM%20and%20Code%20Structure%20(MVC)/screenshots/Screenshot_9.png>)

b. Menampilkan seluruh data users

Untuk menampilkan semua data users pada database dengan menggunakan gorm, dapat menggunakan method `Find` yang dimana menerima parameter bertipe struct User.

Screenshot hasil :

![img2](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/18_ORM%20and%20Code%20Structure%20(MVC)/screenshots/Screenshot_10.png>)

c. Menampilkan user berdasarkan id

Untuk menampilkan user berdasarkan id yang diinginkan, dapat menggunakan method `Where` yang menerima parameter kolom yang dijadikan kondisi dan value-nya, kemudian di chain dengan method `Take` yang dimana menerima parameter bertipe struct user dan nantinya akan mengembalikan data berdasarkan kondisinya dan tanpa pengurutan secara spesifik.

![img3](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/18_ORM%20and%20Code%20Structure%20(MVC)/screenshots/Screenshot_11.png>)

d. Melakukan update data user berdasarkan id

Untuk melakukan update menggunakan gorm, dapat menggunakan method `Where` untuk kondisi berdasarkan apa data tersebut diubah dan dilanjutkan menggunakan `Updates` yang mana menerima parameter berupa struct atau map user.

Screenshot hasil :

- mengubah data user dengan id = 2

![img4](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/18_ORM%20and%20Code%20Structure%20(MVC)/screenshots/Screenshot_12.png>)

- Data user dengan id = 2 berubah

![img5](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/18_ORM%20and%20Code%20Structure%20(MVC)/screenshots/Screenshot_13.png>)

e. Menghapus data user berdasarkan id

Untuk menghapus data menggunakan gorm, dapat menggunakan method `Where` untuk kondisinya dan ditambah menggunakan method `Delete`. Untuk default pada gorm sendiri pada saat delete yaitu akan melakukan soft delete atau hanaya meng-update kolom deleted_at yang nantinya akan dihilangkan visibilitasnya atau tidak akan ditampilkan lagi. Untuk melakukan hard delete dapat menambahkan method `Unscoped`.

Screenshot hasil :

- Menghapus data user dengan id = 1

![img6](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/18_ORM%20and%20Code%20Structure%20(MVC)/screenshots/Screenshot_14.png>)

- Data user dengan id = 1 terhapus

![img7](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/18_ORM%20and%20Code%20Structure%20(MVC)/screenshots/Screenshot_15.png>)

### Problem 2 - Structuring Project with Layered Architecture

Link solusi [**Problem 2**]().

Melakukan restrukturisasi projek Problem 1 dengan menggunakan architectural pattern Model, View, Controller (MVC). Pada problem ini saya membuat struktur folder sebagi berikut.

a. Config

Pada folder config bertanggung jawab untuk melakukan konfigurasi terkait database, seperti hostname, port database, username, password, dan nama database yang digunakan. Di folder ini juga digunakan untuk melakukan migrasi dari bentuk struktur object struct golang ke bentuk tabel database.

b. Controllers

Pada folder controllers bertanggunajawab untuk meng-handle logic sistem, request dan response dari client, perantara untuk komunikasi ke view ataupun model, dsb.

c. Lib->Database

Pada folder database yang ada di folder lib bertanggunajawab untuk melakukan query menggunakan ORM.

d. Models

Pada folder models bertanggunajawab untuk definisi bentuk dari struktur folder yang akan dimigrasikan atau definisi dari tabel yang akan dibuat, atau bisa dapat juga sebagai konfigurasi seperti penamaan pada format json.

e. Routes

Pada folder routes bertanggunajawab untuk mendefinisikan route atau endpoint API yang sedang dibangun

Berikut merupakan struktur folder dibuat yang menggunakan MVC.

![img8](<https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/18_ORM%20and%20Code%20Structure%20(MVC)/screenshots/Screenshot_16.png>)

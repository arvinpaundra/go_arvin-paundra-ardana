# Intro Echo Golang

#### Arvin Paundra Ardana - Golang A

## Static API CRUD User

Membuat static API yang menerapkan CRUD user dengan menggunakan framework `echo`.

Berikut link [**repository**](https://github.com/arvinpaundra/rest-echo/blob/master/main.go).

Berikut penyelesaian masalah untuk membuat sebuah static API menggunakan echo.

a. Menambahkan user

Untuk menambahkan user, buat sebuah fungsi controller yang bernama CreateUserController yang dimana menerima parameter yang berupa echo.Context dan return interface error. Untuk langkah pertama yang perlu dilakukan yaitu membuat instance dari struct User yang nantinya digunakan untuk menampung request body yang berisi data user baru. Untuk melakukan binding data atau menyimpan data dari request body yang dikirim oleh client, dapat menggunakan method `Bind` yang berasal dari context dengan passing parameternya berupa alamat memori instance User. Kemudian, untuk men-generate id secara otomatis yaitu menggunakan for loop, dimana jika panjang dari slice users itu 0, maka id user akan dimulai dari satu, namun jika sudah terisi maka akan mengambil id pada slice index terakhir lalu ditambahkan satu untuk id user terbaru. Setelah itu, gabungkan instance user dengan slice users menggunakan method `append`. Terakhir, return response dengan format json menggunakan method `JSON` yang ada pada context, yang mana pada parameter pertama menerima http status code dan yang kedua berupa repsonse nya yang berupa map dengan key string dan value interface kosong.

Menambahkan user pertama :

![img1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/17_Intro%20Echo%20Golang/screenshots/Screenshot_27.png)

Menambahkan user kedua :

![img2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/17_Intro%20Echo%20Golang/screenshots/Screenshot_28.png)

b. Menampilkan semua user

Untuk menampilkan semua user, buat sebuah fungsi controller yang bernama GetUsersController yang dimana menerima parameter yang berupa echo.Context dan return interface error. Didalam controller ini bisa untuk langsung me-returnkan response menggunakan method `JSON` dari echo, dengan parameter pertama yaitu status OK (200) dan data seluruh users yang dibungkus dengan map[string]interface{}.

Menampilkan semua users :

![img3](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/17_Intro%20Echo%20Golang/screenshots/Screenshot_29.png)

c. Menampilkan user berdasarkan id

Untuk menampilkan user berdasarkan id, buat sebuah fungsi controller yang bernama GetUserController yang dimana menerima parameter yang berupa echo.Context dan return interface error. Hampir sama dengan menampilkan semua users, yang berbeda yaitu client passing sebuah id user yang dimana diambil dari method `Param` dari echo dan sekaligus ubah ke bentuk integer karena semua request params pasti bertipe string. Langkah selanjutnya, loop slice user untuk mencari id yang sesuai id yang di passing oleh client. Jika id tersebut ada pada slice users, maka akan me-returnkan data user tersebut yang sesuai dengan id yang dimaksud. Namun jika gagal maka akan me-returnkan imformasi bahwa user dengan id yang dimaksud tidak ada.

Menampilkan user berdasarkan id :

![img4](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/17_Intro%20Echo%20Golang/screenshots/Screenshot_30.png)

d. Mengubah user berdasarkan id

Untuk mengubah user berdasarkan id, buat sebuah fungsi controller yang bernama UpdateUserController yang dimana menerima parameter yang berupa echo.Context dan return interface error. Untuk mengubah data user hampir sama seperti menambahkan user, yaitu sama-sama menerima request body yang berisi data user. Data user tersebut disimpan pada instance User melalui method `Bind` dari echo. Setelah itu, looping slice users untuk melakukan pencarian data user yang memiliki id yang sama dengan yang di user passing oleh client. Jika id ditemukan sama, maka ubah isi data user tersebut dengan data user terbaru. Namun, jika id tersebut tidak sesuai atau tidak ada dengan id yang ada pada users, maka akan me-returnkan informasi bahwa user tersebut tidak ditemukan.

Mengubah data user dengan id = 2 :

![img5](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/17_Intro%20Echo%20Golang/screenshots/Screenshot_31.png)

Hasil ketika menampilkan user dengan id = 2 :

![img6](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/17_Intro%20Echo%20Golang/screenshots/Screenshot_32.png)

e. Menghapus user berdasarkan id :

Untuk menghapus user berdasarkan id, buat sebuaj fungsi controller yang bernama DeleteUserController yang dimana menerima parameter yang berupa echo.Context dan return interface error. Untuk menghapus data caranya mirip seperti menampilkan detail user, seperti menamngkap terlebih dahulu id yang di passing oleh client, lalu melakukan looping untuk mencocokan id yang di passig oleh client dengan id user yang ada pada slice users. Yang beda yaitu, kita menghapus data tersebut jika id sama dengan id yang ada pada slice users. Untuk menghapus element pada slice, yaitu dapat dengan cara memisahkan element slice seperti cara dibawah ini.

```go
users = append(users[:i], users[i+1:]...)
```

Menghapus data user dengan id = 1 :

![img7](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/17_Intro%20Echo%20Golang/screenshots/Screenshot_33.png)

Menampilkan data user dengan id = 1 :

![img8](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/17_Intro%20Echo%20Golang/screenshots/Screenshot_34.png)

# String - Advance Function - Pointer - Method - Struct and Interface

#### Arvin Paundra Ardana - Golang A

#### 1. Problem 1 - Compare String

Link [Problem 1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/10_String%20-%20Advance%20Function%20-%20Pointer%20-%20Method%20-%20Struct%20and%20Interface/praktikum/Problem%201/main.go)

Problem 1, mencari substring pada sebuah string.

Untuk mencari sebuah substring pada sebuah string, disini saya membagi dua buah fungsi. Fungsi pertama yaitu fungsi ContainSubtr untuk mengecek apakah substring tersebut ada pada sebuah string yang telah ditentukan. Pada fungsi ini menerima dua buah parameter yaitu input dan juga substring dengan return value-nya yaitu string. Di dalam fungsi tersebut dilakukan sebuah pengecekan menggunakan bantuan package `strings.Contains`, jika substring ada pada input string maka akan mengahsilkan nilai true dan akan mengembalikan nilai substring tersebut, namun jika false maka akan mengembalikan sebuah string yang memberitahu bahwa substring tersebut tidak cocok. Kemudian pada fungsi kedua, yaitu fungsi Compare yaitu sama menerima dua parameter dengan input string dan substring dengan return value-nya string. Di dalam fungsi ini hanya melakukan pengecekan string yang paling panjang, jika string b lebih panjang dari a, maka akan mengembalikan fungsi ContainSubtr dengan parameter b didepan a dan jika a lebih panjang dari b maka akan mengembalikan fungsi ContainSubtr dengan parameter a terlebih dahulu dan baru b.

Screenshot :

![Problem 1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/10_String%20-%20Advance%20Function%20-%20Pointer%20-%20Method%20-%20Struct%20and%20Interface/screenshots/Screenshot_36.png)

#### 2. Problem 2 - Caesar Cipher

Link [Problem 2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/10_String%20-%20Advance%20Function%20-%20Pointer%20-%20Method%20-%20Struct%20and%20Interface/praktikum/Problem%202/main.go)

Problem 2, singkatnya melakukan penggeseran sebuah huruf sesuah dengan jumlah offsetnya.

Untuk memulainya, saya membuat sebuah slice penampung dengan nama `substitute` untuk menyimpan masing masing karakter yang telah digeser. Setelah itu, lakukan loop parameter input. Pada proses loop, buat variabel `shift` untuk menyimpan pergeseran nilai dari ascii masing-masing karakter dengan algoritmanya yaitu `(ascii karakter + offset - 97) % 26 + 97` dan dari ini akan mendapatkan ascii karakter yang telah digeser. Kemudian, ubah variabel shift tersebut menjadi bentuk rune dan simpan pada variabel `asciiToString`. Terakhir pada proses loop, append dengan slice substitute tadi untuk masing masing karakter yang telah digeser. Setelah loop selesai, gabungkan karakter-karakter pada slice substitute dengan bantuan package `strings.Join()`.

Screenshot :

![Problem 2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/10_String%20-%20Advance%20Function%20-%20Pointer%20-%20Method%20-%20Struct%20and%20Interface/screenshots/Screenshot_37.png)

#### 3. Problem 3 - Swap Two Number Using Pointer

Link [Problem 3](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/10_String%20-%20Advance%20Function%20-%20Pointer%20-%20Method%20-%20Struct%20and%20Interface/praktikum/Problem%203/main.go)

Problem 3, melakukan swap antara nilai a dan b dengan menggunakan pointer.

Untuk melakukan swap nilai menggunakan pointer, saya membuat sebuah variabel `temp` untuk menyimpan nilai *a. Kemudian, ubah nilai *a dengan nilai *b dan nilai *b dengan nilai temp. Perlu diperhatikan bahwa ( \* ) berarti melakukan deferencing atau mengakses nilai pada sebuah variabel pointer.

Screenshot :

![Problem 3](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/10_String%20-%20Advance%20Function%20-%20Pointer%20-%20Method%20-%20Struct%20and%20Interface/screenshots/Screenshot_38.png)

#### 4. Problem 4 - Min and Max Using Pointer

Link [Problem 4](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/10_String%20-%20Advance%20Function%20-%20Pointer%20-%20Method%20-%20Struct%20and%20Interface/praktikum/Problem%204/main.go)

Problem 4, menemukan nilai maksimum dan minimum menggunakan pointer.

Untuk mencari nilai maksimum dan minimum menggunakan pointer, inisialisasi return name max dan min dengan nilai awal yaitu index 0 dari slice varargs *numbers. Setelah itu lakukan loop pada numbers. Di dalam loop, dilakukan pengecekan apabila *value lebih besar dari max maka tukar max dengan *value. Setelah itu cek juga apabila jika *value lebih kecil dari min, tukar nilai min dengan \*value. Setelah loop selesai, return min dan max.

Screenshot :

![Problem 4](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/10_String%20-%20Advance%20Function%20-%20Pointer%20-%20Method%20-%20Struct%20and%20Interface/screenshots/Screenshot_39.png)

#### 5. Problem 5 - Student Score

Link [Problem 5](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/10_String%20-%20Advance%20Function%20-%20Pointer%20-%20Method%20-%20Struct%20and%20Interface/praktikum/Problem%205/main.go)

Problem 5, mencari nilai rata-rata, siswa dengan nilai maksimun dan minimum dari 5 siswa.

Method pertama, yaitu Average digunakan untuk menghitung rata-rata nilai 5 siswa. Pada method tersebut dibuat variabel `average` dengan tipe float untuk menyimpan nilai desimal. Setelah itu lakukan loop pada field score yang merupakan sebuah slice yang ada pada struct Student. pada masing-masing iterasinya, jumlahkan average dengan value yang ada pada field score. Setelah loop selesai, bagi dengan jumlah siswa yang ada pada struct Student. Terakhir return variabel average.

Method kedua, yaitu Min digunakan untuk mencari nilai terkecil pada field score beserta nama siswanya. Method tersebut mengembalikan nama variabel min dan name. Inisialisasi min dengan `s.score[0]` dan name dengan `s.name[0]`. Setelah itu, loop field score dan di dalamnya lakukan sebuah pengecekan, apabila value lebih kecil dari min maka tukar nilai min dengan value dan tukar nilai name dengan `s.name[i]`. Setelah loop selesai, return min dan name.

Method ketiga, yaitu Max digunakan untuk mencari nilai terbesar pada field score beserta nama siswanya. Method tersebut mengembalikan nama variabel max dan name. Inisialisasi max dengan `s.score[0]` dan name dengan `s.name[0]`. Setelah itu, loop field score dan di dalamnya lakukan sebuah pengecekan, apabila value lebih besar dari max maka tukar nilai max dengan value dan tukar nilai name dengan `s.name[i]`. Setelah loop selesai, return max dan name.

Screenshot :

![Problem 5](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/10_String%20-%20Advance%20Function%20-%20Pointer%20-%20Method%20-%20Struct%20and%20Interface/screenshots/Screenshot_40.png)

#### 6. Substitution Cipher

Link [Problem 6](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/10_String%20-%20Advance%20Function%20-%20Pointer%20-%20Method%20-%20Struct%20and%20Interface/praktikum/Problem%206/main.go)

Problem 6, singkatnya yaitu membuat sebuah algoritma enkripsi beserta dekripsi yang sederhana menggunakan algoritma Caesar Cipher.

Method Encode, digunakan untuk melakukan enkripsi atau pengubahan string ke bentuk string yang telah digeser. buat variabel string nameEncode dan slice string temp. Setelah itu, lakukan loop field `s.Name` dan di dalamnya buat variabel `shift` untuk menyimpan pergeseran nilai dari ascii masing-masing karakter dengan algoritmanya yaitu `(ascii karakter + 17 /* saya mengasumsikan bahwa digeser 17 karena tidak ada keterangan khusus untuk jumlah geser string */ - 97) % 26 + 97` dan dari ini akan mendapatkan ascii karakter yang telah digeser. Kemudian, ubah variabel shift tersebut menjadi bentuk rune dan simpan pada variabel `asciiToString`. Terakhir pada proses loop, append dengan slice temp tadi untuk masing masing karakter yang telah digeser. Setelah loop selesai, gabungkan karakter-karakter pada slice temp dengan bantuan package `strings.Join()`.

Method Decode, digunakan untuk melakukan dekripsi atau pengubahan string yang telah digeser ke bentuk string semula. buat variabel string nameDecode dan slice string temp. Setelah itu, lakukan loop field `s.Name` dan di dalamnya buat variabel `shift` untuk menyimpan pergeseran nilai dari ascii masing-masing karakter dengan algoritmanya yaitu `(ascii karakter - 17 - 97) % 26 + 97` dan dari ini akan mendapatkan ascii karakter yang telah digeser ke bentuk semula. Kemudian, ubah variabel shift tersebut menjadi bentuk rune dan simpan pada variabel `asciiToString`. Terakhir pada proses loop, append dengan slice temp tadi untuk masing masing karakter yang telah digeser. Setelah loop selesai, gabungkan karakter-karakter pada slice temp dengan bantuan package `strings.Join()`.

Screenshot :

![Problem 6](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/10_String%20-%20Advance%20Function%20-%20Pointer%20-%20Method%20-%20Struct%20and%20Interface/screenshots/Screenshot_41.png)

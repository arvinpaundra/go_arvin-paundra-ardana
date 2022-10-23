# (20) Unit Testing

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Unit Testing

### 1. Pengertian Software Testing

Software testing merupakan sebauh proses dari menganalisis sebuah software item untuk mendeteksi perbedaan antara kondisi yang aktual dengan kondisi yang diekspetasikan dan untuk evaluasi fitur dari sebuah software item.

Tujuan adanya testing antara lain :

- Menghindari regresi

- Kepercayaan pada level **Refactoring**

- Meningkatkan **Code Design**

- Kode yang terdokumentasi

- Scaling tim

Adapun level dari sebuah testing yaitu antara lain dari yang terkecil :

- Unit test

Menguji bagian paling kecil atau operasi logika tunggal dari sebuah aplikasi melalui method.

- Integration test

Menguji sebuah modul atau sub-sistem melalui API.

- UI atau End to End

Menguji interaksi antara keseluruhan sistem melalui user interface atau antarmuka.

Framework untuk melakukan unit testing di golang yaitu ada testify. Tujuan digunakannya framework pada saat melakukan unit test yaitu karena framework menyediakan tools dan struktu yang diperlukan untuk menjalankan unit testing secara efisien dan juga agar lebih terfokus pada tahap pengembangan sistem.

### 2. Struktur pada Unit Test

Struktur atau pola yang biasa digunakan dalam melakukan sebuah unit testing ada dua, antara lain :

- Centralize, yang berarti semua file test tersebut berada pada sebuah test folder.

- Menyimpan test bersamaan dengan production, yang berarti file test tersebut berada pada folder yang sama dengan tempat dimana fungsi yang akan ditest itu ada.

Jika breakdown kembali, test file berisi koleksi dari test suites. Test suites kemudian berisi koleksi dari test cases dan test fixtures.

### 3. Runner, Mocking, dan Coverage

a. Runner

- Tool yang menjalankan test file.

- menggunakan watch mode yang mana sangat membantu karena apabila terdapat sebuah perubahan pada kode, secara otomatis akan menjalankan test kembali yang mana membuat TDD atau Test Driven Development menjadi lebih efisien.

b. Mocking

Merupakan sebuah objek palsu yang mensimulasikan behaviour dari objek asli. Mock digunakan ketika akan mengetest modul third party, database, dsb.

c. Coverage

Digunakan untuk memastikan apakah mereka telah membuat test yang cukup untuk keseluruhan sistem tersebut. Coverage tool menganalisis kode aplikasi ketika test dijalankan. Format report coverage dapat berupa CLI, XML, HTML, LCOV.

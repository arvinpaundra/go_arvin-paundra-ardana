# (9) Clean Code

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Clean Code

#### 1. Pengertian dan Kegunaan Clean Code

Clean code merupakan istilah untuk kode yang mudah dibaca, dipahami, dan diubah oleh programmer. Kegunaan dari clean code antara lain.

- Memudahkan developer dalam bekerja dengan berkolaborasi.

- Memudahkan dalam melakukan sebuah improvement, khususnya dalam membuat sebuah fitur.

- Mempercepat proses development.

#### 2. Karakteristik pada Clean Code

Kode dapat dikatakan clean code jika memenuhi karaketristik dibawah ini.

a. Mudah dipahami

Kode harus mudah dipahami dengan tidak hanya dipahami oleh komputer saja, namun juga harus dapat dipahami oleh sesama manusia. Seperti penamaan sebuah variabel yang sesuai dengan representasinya. Contoh `Users` merupakan representasi dari banyak user.

b. Mudah dieja dan dicari

Penamaan variabel harus jelas dan mudah ketika dicari. Contoh `PHI` dideklarasikan secara global dan konstan, diletakan pada awal baris, dan menggunakan kapital karena nilainya konstan.

c. Singkat namun mendeskripsikan konteks

Pemberian penamaan singkat dan menunjukkan sesuai konteks yang ada. Contoh `fullName`, singkat dan menunjukkan konteks bahwa variabel tersebut berisi nama lengkap seseorang.

d. Konsisten

Harus konsisten, tidak boleh menyimpang dari gaya awal. Contoh pada penamaan class menggunakan PascalCase, maka semua penamaan class harus menggunakan PascalCase.

e. Hindari penambahan konteks yang tidak perlu

Hindari penambahan yang sudah jelas merupakan sebuah properti dari objek tersebut. Contoh, terdapat sebuah struct dengan nama User dan pada fieldnya tidak perlu UserID, UserEmail, UserPassword. Cukup hilangkan saja kata user pada masing masing field, karena itu sudah menjelaskan bahwa field tersebut milik struct User.

f. Komentar

Beri komentar yang menjelaskan bagaimana sebuah blok kode tersebut diekseksi tanpa harus memberi komentar pada setiap baris blok kode tersebut.

g. Good Function

Fungsi yang baik tidak boleh memiliki banyak parameter dan fungsi yang baik harus menghindari side effect.

h. Gunakan konvensi

Usahakan menggunakan style guide penulisan kode yang telah terbukti efektif dan dipakai oleh perusahaan-perusahaan besar. Contoh pada golang terdapat style guide milik [CockroachDB](https://wiki.crdb.io/wiki/spaces/CRDB/pages/181371303/Go+Golang+coding+guidelines) atau dari [go.dev](https://go.dev/doc/effective_go) langsung.

i. Formatting

Untuk melakukan formatting yaitu dapat memperhatikan indentasi, lebar baris karakter 80-120 karakter, satu class terdiri dari 300-500 baris, dsb. Atau untuk memudahkan dapat dengan menggunakan formatter yang builtin pada sebuah IDE.

#### 3. Prinsip pada Clean Code

Adapun beberapa prinsip pada clean code, yaitu :

- KISS (Keep It Simple, Stupid)

Hindari membuat sebuah fungsi yang dapat melakukan banyak tugas. Fungsi harus single responsibility atau hanya memiliki satu tugas. Jangan gunakan terlalu banyak argumen atau parameter pada sebuah fungsi.

- DRY (Don't Repeat Yourself)

Buat sebuah fungsi yang bisa digunakan secara berulang-ulang sesuai kebutuhan untuk mengurangi redundansi.

- Refactoring

Refactoring merupakan sebuah proses restrukturisasi sebuah kode yang telah dibuat tanpa harus mengubah inti logika dari kode tersebut. Refactoring biasanya dilakukan dengan cara membuat sebuah abstraksi, memecah fungsi atau class menjadi bagian yang lebih sederhana, memperbaiki penamaan kode, dan menghilangkan duplikasi dengan mengganti menjadi sebuah reusable function.

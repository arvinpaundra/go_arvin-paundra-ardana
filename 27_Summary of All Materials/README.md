# (3) Version Control and Branch Management (Git)

#### Arvin Paundra Ardana - Golang A

## 3 Point Penting pada Materi Git

#### 1. Useful Tools

Git sangat membantu para developer untuk memanajemen pekerjaan projek mereka mulai dari siapa saja yang mengubah kode, kapan kode itu diubah, siapa saja yang berkontribusi, apa saja yang diperbarui pada kode tersebut, dan sebagainya.

Dengan git juga akan membuat sebuah kebiasaan baru untuk seorang developer agar selalu mendokumentasikan setiap perubahan yang telah mereka buat. Hal ini sangat penting mengingat jika yang dikerjakan merupakan sebuah projek yang skalanya besar yang akan menyulitkan proses pengembangan apabila tidak ada sebuah tools untuk melakukan pelacakan pada setiap perubahan kode.

Git akan sangat membantu ketika digunakan dengan berkolaborasi dengan orang banyak. Karena akan lebih terlihat bahwa git itu akan merekam dan menampilkan semua commit perubahan yang mereka buat sehingga akan lebih mudah untuk dipantau dan dilacak pada setiap perubahannya.

#### 2. Bagian-bagian pada Git

- Local repository
  Local repository merupakan sebuah komputer yang digunakan oleh developer dimana itu merupakan tempat melakukan sebuah pengubahan dan penyimpanan sebuah kode secara lokal (pada komputer masing-masing developer).
- Stagging Area
  Stagging area merupakan sebuah area dimana developer melakukan stagging atau menandai suatu kode yang telah ia ubah atau buat yang kemudian akan dilakukan sebuah commit pada kode tersebut yang nantinya akan didstribusikan ke remote repository.
- Remote Repository
  Remote repository merupakan sebuah komputer server yang digunakan untuk menyimpan seluruh kode-kode yang di commit oleh developer atau tempat central untuk menyimpan setiap perubahan kode.
- Stash Area
  Stash area merupakan sebuah penyimpanan sementara perubahan saat ini pada kode tanpa harus melakukan commit ke remote repository.

#### 3. Best Practices Penggunaan Git

Best practices pada git yaitu dapat dengan menggunakan git workflow untuk projek yang berskala besar.

- Git workflow hanya memperbolehkan branch development untuk melakukan checkout branch master agar branch master tidak terganggu dengan branch lain.
- Hindari untuk mengedit langsung pada branch development, lebih baik buat satu branch baru yang men-checkout branch development jika memang diperlukan.
- Jika ingin membuat sebuah fitur baru maka checkout dan merge pada branch development.
- Jika akan melakukan merge fitur baru ke branch master, hanya gunakan branch development pada saat melakukan merging tersebut.

Opsi lain yang dapat digunakan yaitu dapat menggunakan trunk-based development yang biasanya dipakai pada projek yang tidak terlalu besar. Opsi ini menggunakan satu buah branch untuk berkolaborasi, misalnya branch master dapat di checkout langsung oleh beberapa branch.

# (4) Introduction to Algorithm and Golang

#### Arvin Paundra Ardana - Golang A

## Poin Penting pada Materi Algoritma dan Golang

### 1. Pengertian Algoritma

Algoritma adalah sebuah prosedur komputasi yang didefinisikan dengan baik yang biasanya mengambil nilai sebagai input dan menghasilkan beberapa nilai sebagai output. Contoh penerapan algoritma yaitu seperti pengecekan bilangan prima, sorting, searching, dsb.

### 2. Karakteristik Algoritma

- Algoritma memiliki batas awal dan akhir.
- Instruksi terdefinisikan dengan baik.
- Efektif dan efisien.

### 3. Tiga Konsep Algoritma

- Sequential (Berurutan)
- Branching (Percabangan)
- Looping (Perulangan)

### 4. Beberapa Cara Merepresentasikan Sebuah Algoritma

##### 1. Pseudocode

Pseudocode adalah bagian dari algoritma yang bertujuan untuk memahami alur logika dari suatu program.

##### 2. Flowchart

Flowchart adalah suatu bagan dengan simbol tertentu yang menggambarkan urutan dan hubungan antar proses secara mendetail.

# (5) Basic Programming

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Basic Programming

### 1. Variabel pada Golang

Variabel biasanya digunakan untuk menyimpan nilai dan pemberian nama data secara deskriptif.

Terdapat dua keyword dalam melakukan deklarasi variabel, yaitu :

#### a. Keyword **var**

Keyword var digunakan ketika data yang disimpan nantinya dapat berubah ubah nilainya. Pada var memiliki address yang mana merupakan sebuah tempat dimana nilai tersebut disimpan dalam memori.

Beberapa cara melakukan deklarasi variabel menggunakan keyword var.

- Long

```go
...

var name string = "Arvin Paundra Ardana"

...
```

- Short

```go
...

name := "Arvin Paundra Ardana"

...
```

\*Golang akan secara otomatis mendeteksi bahwa variabel name bertipe data string.

#### b. Keyword **const**

Keyword const digunakan ketika hendak menyimpan sebuah nilai yang konstan atau tidak akan diubah untuk kedepannya. Pada const tidak akan bisa untuk melakukan pengubahan sebuah nilai (immutable). Berbeda dengan var, const tidak memiliki address.

Berikut cara untuk melakukan deklarasi variabel menggunakan keyword const.

```go
...

const phi float64 = 3.14

...
```

### 2. Tipe Data pada Golang

Golang memiliki 3 buah tipe data dasar yaitu string, boolean, dan numeric.

#### a. String

Sama seperti pada bahasa pemrograman lainnya, string merupakan sekumpulan karakter yang ditandai dengan tanda kutip diawal dan diakhirnya. Pada golang, string direpresentasi seperti potongan byte.

#### b. Boolean

Golang terdapat tipe data boolean yang hanya berisi true atau false.

#### c. Numeric

Tipe data numerik di golang memiliki sebuah range yang mana merupakan batas seberapa panjang data dapat disimpan pada variabel, contohnya `int8` yang mampu menyimpan dengan range -128 sampai 127. Pada golang terdapat tiga jenis tipe data numerik sebagai berikut.

- Integer atau `int`

Integer merupakan sebuah tipe data numerik yang berupa bilangan bulat bukan pecahan. terdapat dua jenis tipe integer, yaitu `int` dan `uint`.

- `float`

Tipe data numerik ini dapat digunakan untuk menyimpan angka pecahan atau desimal.

- `complex`

Tipe data numerik terakhir ini biasa digunakan untuk menghitung perhitungan matematika tingkat tinggi seperti bilangan imajiner dan sejenisnya.

### 3. Control Structure Branching and Looping pada Golang

#### a. Branching

Pada golang untuk menentukan sebuah kontrol alur logika program disediakan menggunakan dua cara, yaitu menggunakan IF atau menggunakan SWITCH.

- **IF** statement

Penggunaan statement if pada golang sebagai berikut.

```go
...

if kondisi_satu {
    // do something
} else if kondisi_dua {
    // do something
} else {
    // do something
}

...
```

- **SWITCH** statement

Penggunaan statement switch pada golang sebagai berikut.

```go
...

switch ekspresi {
case case_1:
    // do something
case case_2:
    // do something
default:
    // do something
}

...
```

\*Switch pada golang tidak perlu menambahkan keyword `break` pada setiap akhir case, karena golang sudah secara otomatis akan menghentikan proses switch ketika kondisi terpenuhi.

#### b. Looping

Pada golang secara garis besar untuk melakukan sebuah proses perulangan dapat menggunakan beberapa cara sebagai berikut.

- Cara pertama

```go
...

for inisialisasi; kondisi; pengubah {
    // do something
}

...
```

- Cara kedua

```go
...

for index, value := range array {
    // do something
}

...
```

# (6) Data Structure

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Data Structure

### 1. Big O Notation

Big O Notation merupakan sebuah cara untuk memastikan sebuah keefisienan sebuah algoritma.

#### a. Time Complexity

Time Complexity merupakan sebuah kompleksitas perhitungan yang menunjukkan total waktu seberapa lama untuk menjalankan sebuah algoritma.

Macam-macam time complexity :

- O(1) - Constant time

Constant time berarti banyaknya sebuah inputan tidak mempengaruhi waktu proses dari algoritma tersebut.

- O(log n) - Logarithmic time

Logarithmic time yaitu jika memberikan inputan sebesar **n** terhadap sebuah fungsi, jumlah tahapan yang dilakukan oleh fungsi tersebut berkurang berdasarkan suatu faktor.

- O(n) - Linear time

Linear time adalah ketika proses runtime dari fungsi berbanding lurus dengan jumlah input yang diberikan.

- O(n^2) - Quadratic time

Quadratic time adalah ketika proses runtime dari sebuah fungsi adalah sebesar **n^2**, dimana **n** adalah sebuah inputan dari fungsi tersebut.

#### b. Space Complexity

Space Complexity adalah jumlah memori yang digunakan oleh sebuah algoritma atau program.

### 2. Array, Slice, dan Map pada Golang

Pada bahasa golang terdapat tiga buah struktur data, dimana ketiga itu ialah array, slice, dan map.

#### a. Array

Array merupakan kumpulan dari elemen atau nilai dengan tipe data yang sama. Array memiliki panjang yang tetap atau statik (fixed allocation size). Nilai pada array ditandai dengan index yang mana itu berupa integer yang biasanya digunakan untuk memanipulasi nilai pada array. Index pada array tersebut dimulai dari 0.

Contoh struktur data array :

```go
...

// deklarasi array perlu menentukan panjang array
var mahasiswa [5]string{"Arvin", "Paundra", "Ardana"}

// menambahkan nilai pada array
mahasiswa[3] = "Pesulap"
mahasiswa[4] = "Merah" // dimulai dari indeks 0 maka berakhir di indeks ke 4

// mencetak nilai pada array
fmt.Println(mahasiswa)
fmt.Println(mahasiswa[1])

...
```

#### b. Slice

Slice pada dasarnya merupakan sebuah array, yaitu sama-sama kumpulan nilai atau elemen dengan tipe data yang sama. Perbedaannya adalah slice merupakan referensi tiap elemen tersebut dan juga slice memiliki panjang nilai yang dinamis (dynamic allocation size). Yang unik pada slice yaitu, ketika ada slice baru yang terbentuk dari slice atau array lama, maka slice baru tersebut memiliki alamat memori yang sama dengan slice lama yang mengakibatkan jika terdapat perubahan pada slice baru akan berdampak juga pada slice lama, begitu juga sebaliknya.

Contoh struktur data slice :

```go
...

// mendeklarasikan array
var mahasiswa = [5]string{"Arvin", "Paundra", "Ardana"}

fmt.Println(mahasiswa)

// copy nilai array ke dalam slice
var copied = mahasiswa[0:3]

// mengubah nilai slice akan berdampak pada array
copied = append(copied, "Pesulap", "Merah")

// cetak array dan slice
fmt.Println(mahasiswa)
fmt.Println(copied)

...
```

#### c. Map

Struktur data yang terakhir adalah map, dimana map sama seperti array yang merupakan kumpulan nilai atau elemen, bedanya yaitu map terdiri dari key dan value. Key dari map memiliki tipe data bisa selain integer (contoh string) dan juga key tersebut bersifat **unik** dari key lainnya.

```go
...

// deklarasi map
var mahasiswa = map[string]string{"nama": "Arvin", "domisili": "Cilacap"}

// menambahkan key pada map
mahasiswa["gender"] = "pria"
mahasiswa["hobi"] = "Main Game FF"

// menghapus data pada map berdasarkan key
delete(mahasiswa, "hobi")

// cetak data map
fmt.Println(mahasiswa)

...
```

### 3. Fungsi pada Golang

Fungsi merupakan sekumpulan blok kode yang dibungkus dengan nama tertentu. Tujuan fungsi biasanya digunakan untuk mengurangi pengulangan baris kode dan membuat kode menjadi lebih rapih dan mudah untuk dipahami.

#### a. Fungsi dan paramater

Parameter merupakan nilai yang dikirimkan ke sebuah blok fungsi yang nantinya akan dilakukan sebuah operasi. Parameter bersifat opsional tergantung dengan kebutuhan. Untuk membuat fungsi pada golang dimulai dengan menggunakan keyword `func` dan diikuti dengan nama fungsinya.

```go
...

// deklarasi sebuah fungsi tanpa parameter
func greetings() {
    fmt.Println("Halo, Alterra Academy!")
}

// deklarasi sebuah fungsi dengan parameter
func job(name string, jobType string) {
    fmt.Printf("%s adalah seorang %s", name, jobType)
}

func main() {
    // memanggil fungsi dalam program main
    greetings()
    job("Arvin", "Backend Developer")
}

...
```

#### b. Fungsi dengan return nilai

Return nilai merupakan nilai yang dikembalikan oleh sebuah fungsi dari hasil operasi pada fungsi tersebut. Nilai yang di-return pada fungsi harus sama dengan tipe yang telah didefinisikan. Untuk menandakan bahwa fungsi tersebut memerlukan pengembalian sebuah nilai, yaitu tambahkan sebuah tipe sebelum kurung kurawal.

```go
...

// deklarasi fungsi dengan single return value
func exponentiation(base int, power int) int {
    var result = base

    if power == 0 {
        result = 1

        return result
    } else {
            for i:=0; i < power-1; i++ {
            result = result * base
        }
    }

    return result
}

// deklarasi fungsi dengan multiple return value
func getFullname(first, last, string) (string, string) {
    return first, last
}

func main() {
    // mengambil nilai return dari fungsi
    var result = exponentiation(2, 3)
    fmt.Println(result)

    var first, last = getFullname("Arvin", "Ardana")
    fmt.Println(first)
    fmt.Println(last)
}

...
```

#### c. Fungsi dengan return nama variabel

Fungsi yang me-return nama variabel berarti bahwa fungsi tersebut data mengembalikan nilai yang ada pada sebuah variabel di dalam fungsi tersebut. Sama seperti return nilai lainnya pada saat deklarasi, yang membedakan hanya menambahkan nama variabel yang akan di-return beserta dengan tipe datanya.

```go
...

func square(side int) (wide int) {
    wide = side * side

    return
}

func main() {
    // pemanggilan fungsi
    var wide = square(4)

    fmt.Println(wide)
}

...
```

# (7) Time Complexity and Space Complexity

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Time Complexity and Space Complexity

#### 1. Time Complexity

Time Complexity merupakan sebuah kompleksitas perhitungan yang menunjukkan total waktu seberapa lama untuk menjalankan sebuah algoritma.

Macam-macam time complexity :

- O(1) - Constant time

Constant time berarti banyaknya sebuah inputan tidak mempengaruhi waktu proses dari algoritma tersebut.

- O(n) - Linear time

Linear time adalah ketika proses runtime dari fungsi berbanding lurus dengan jumlah input yang diberikan.

- O(log n) - Logarithmic time

Logarithmic time yaitu jika memberikan inputan sebesar n terhadap sebuah fungsi, jumlah tahapan yang dilakukan oleh fungsi tersebut berkurang berdasarkan suatu faktor.

- O(n^2) - Quadratic time

Quadratic time adalah ketika proses runtime dari sebuah fungsi adalah sebesar n^2, dimana n adalah sebuah inputan dari fungsi tersebut.

#### 2. Time Limit

Untuk komputer pada masa sekarang ini bisa mengerjakan 10^8 operasi dengan kurang dari satu detik. Batasan waktu yang diset pada pengujian online biasanya satu sampai sepuluh detik.

- Pada input `n <= 1.000.000`, perkiraan time complexity nya adalah O(n) atau O(n log n)
- Pada input `n <= 10.000`, perkiraan time complexity nya adalah O(n^2)
- Pada input `n <= 500`, perkiraan time complexity nya adalah O(n^3)

#### 3. Space Complexity

Space Complexity adalah jumlah memori yang digunakan oleh sebuah algoritma atau program.

# (8) Recursive - Number Theory - Sorting - Searching

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Recursive - Number Theory - Sorting - Searching

#### 1. Recursive Function

Recursive atau rekursif merupakan sebuah fungsi yang memanggil dirinya sendiri untuk dieksekusi atau mengulang baris kode yang ada pada fungsi tersebut dengan cara memanggil dirinya sendiri yang bertujuan untuk mengurangi kompleksitas kode dan meningkatkan keterbacaan kode (readability).

Pada fungsi rekursif terdapat beberapa hal yang perlu diperhatikan, antara lain seperti :

#### a. Batas fungsi rekursif

Tidak seperti perulangan menggunakan loop yang bisa tak terbatas (endless loop), rekursif memiliki sebuah batasan terhadap pemanggilannya. Batas itu sendiri sampai adanya sebuah error warning dengan pesannya `stack overflow` yang menandakan sebuah batas dari fungsi rekursif.

#### b. Break point pada rekursif

Fungsi rekursif harus memiliki sebuah break point atau titik penghenti agar fungsi tersebut memiliki sebuah batasan harus sampai kapan fungsi rekursif tersebut harus dieksekusi. Jika tidak memberikan break point maka secara terus menerus fungsi rekursif tersebut akan selalu dieksekusi dan akan mengakibatkan error pada eksekusi.

#### c. Harus memiliki return value

Fungsi rekursif harus memiliki sebuah return value yang mana merupakan salah satu cara break point pada rekursif. Fungsi dari return value lainnya yaitu supaya kembalian dari fungsi rekursif bisa dieksekusi kembali secara berulang.

Pada penggunaannya fungsi rekursif digunakan pada contoh kasus seperti mencari nilai faktorial dari sebuah bilangan atau mencari deret fibonacci.

#### 2. Number Theory

Number theory merupakan cabang dari ilmu matematika yang memperlajari tentang integer (bilangan bulat). Contoh dari penerapakan number theory yaitu KPK, FPB, faktorial, faktor prima, dll.

#### 3. Sorting dan Searching

#### a. Sorting

Sorting merupakan sebuah proses dari mengurutkan data pada urutan tertentu. Pada implementasinya, sorting biasa digunakan untuk mengurutkan sebuah angka, huruf, dan data lainnya.

Contoh algoritma yang digunakan untuk melakukan sorting :

- Selection Sort - O(n^2)

Konsep dari selection sort yaitu mencari nilai terkecil dan menukarnya dengan elemen pertama pada sebuah array. Selanjutnya, sama hanya tinggal melanjutkan proses sebelumnya dengan tidak mengikutkan elemen pertama pada array karena sudah ditukar pada proses sebelumnya. Time complexity pada selection sort yaitu O(n^2), yang mana tidak efisien jika harus mengurutkan data dengan jumlah elemen yang besar.

- Counting Sort - O(n + k)

Konsep dari counting sort yaitu yang pertama mencari jumlah dari array tersebut, lalu yang kedua tinggal melakukan iterasi pada array tersebut dengan urutan yang di increment. Time complexity pada counting sort yaitu O(n + k).

- Merge Sort - O(log n)

Konsep merge sort yaitu membagi array menjadi dua bagian dan secara terus menerus dibagi sampai menjadi sebuah individual elemen. Kemudian elemen tersebut digabungkan kembali dengan elemen pertama yang terkecil. Time complexity pada merge sort yaitu O(log n).

#### b. Searching

Searching merupakan sebuah proses mencari letak nilai yang yang ada pada sebuah kumpulan nilai.

Contoh algoritma yang digunakan untuk melakukan searching :

- Linear Search

Linear search bekerja dengan melakukan perulangan melalui array sampai query ditemukan, yang mana membentuk sebuah algoritma linear. Time complexity pada linear search yaitu O(n).

- Binary Search - O(log n)

Binary search bekerja dengan mengurangi separuh setiap ruang pencarian di setiap iterasi setelah membandingkan nilai target ke nilai tengah dari ruang pencarian. Time complexity pada binary search yaitu O(log n).

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

# (10) String - Advance Function - Pointer - Method - Struct and Interface

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi String - Advance Function - Pointer - Method - Struct and Interface

### 1. Advance Function dan Method pada Golang

#### a. Advance Function

Pada golang terdapat beberapa advance function atau fungsi tingkat lanjut yang dapat membantu dalam pengembangan menggunakan golang. Berikut macam-macam advance function pada golang.

- Variadic Function

Variadic function merupakan fungsi yang memiliki varargs (variable argument) dan pada akhir parameternya ditandai titik tiga. Pada parameter terakhir tersebut sebenarnya merupakan sebuah slice sementara. Variadic function digunakan ketika kita tidak tahu seberapa banyak input yang akan di pass ke fungsi tersebut dan kemudian meningkatkan juga keterbacaan (readability).

Contoh variadic function :

```go
...

// Daripada melakukan ini
func notVariadic(target int, numbers []int) bool {
    // awesome code

    return true
}

// Lebih baik gunakan variadic function
func myVariadic(target int, numbers ...int) bool {
    // awesome code

    return true
}

...

```

- Anonymous Function

Seperti namanya, anonymous function merupakan fungsi tanpa nama. Kegunaan fungsi anonymous ini biasanya digunakan untuk langsung menjalankan fungsi tersebut tanpa harus memanggilnya atau membuat sebuah inline function.

Contoh anonymous function :

```go
...

function main() {
    // Anonymous function ini langsung dijalankan tanpa harus dipanggil
    func() {
        fmt.Println("Hallo, nama saya Arvin Paundra Ardana")
    }()

    // Anonymous function bisa digunakan sebagai value (function as value)
    anonFunc := func(value string) {
        fmt.Println(value + " adalah Hengker Pro")
    }

    // Memanggil anonymous function yang di set sebagai value
    anonFunc("Arvin")
}

...
```

- Closure

Closure merupakan sebuah special type dari anonymous function yang mereferensikan variabel yang dideklarasikan diluar dari fungsi tersebut. Closure bisa digunakan untuk menghindari side effect, seperti contohnya mengubah nilai variabel yang sudah dideklarasikan diluar fungsi tersebut (data isolation).

Contoh closure :

```go
...

func main() {
    var counter = 0

    increment := func() {
        /*
         *  Untuk menghindari pengubahan nilai pada variabel global counter,
         *  maka deklarasikan ulang variabel counter didalam block scope fungsi ini
         */
        var counter = 0
        counter++
        fmt.Println(counter)
    }

    increment()
    increment()
    increment()

    // counter tidak akan berubah nilainya
    fmt.Println(counter)
}

...
```

- Defer Function

Defer function merupakan fungsi yang dijalankan ketika parent function selesai dieksekusi. Fungsi defer bekerja layaknya stack, yaitu LIFO atau Last In First Out yang mana membuat fungsi ini akan dijalankan seolah-olah naik keatas setelah parent function selesai dieksekusi. Yang menarik pada fungsi defer yaitu, fungsi defer akan tetap dijalankan meskipun terdapat panic atau error pada parent function-nya.

Contoh defer function :

```go
...

func runApplication(value int) {
    defer func() {
        fmt.Println("Fungsi selesai dijalankan")
    }()

    var result = value / 10

    fmt.Println(result)
}

func main() {
    runApplication(100)
}

...
```

#### b. Method

Method pada golang merupakan sebuah fungsi yang menempel pada sebuah type, contohnya menempel pada sebuah type struct. Jadi pada struct tersebut seolah olah memiliki sebuah fungsi yang hanya milik type struct tersebut. Kegunaan method biasanya untuk menghindari konflik penamaan sebuah fungsi, membantu menulis gaya object-oriented pada golang yang karena pada dasarnya golang sendiri tidak mendukung object-oriented, dan meningkatkan keterbacaan (readability).

Contoh method :

```go
...

type Organism struct {
    Kingdom string
    Class string
    Order string
    Genus string
}

// deklarsi method pada struct User
func (o Organism) GetKingdom() string {
    return o.Kingdom
}

func (o Organism) GetGenus() string {
    return o.Genus
}

func main() {
    var cat = Organism{
        Kingdom: "Animalia",
        Class: "Mammals",
        Order: "Carnivore",
        Genus: "Felis",
    }

    // Pemanggilan method
    fmt.Println(cat.GetKingdom())
    fmt.Println(cat.GetGenus())
}

...
```

### 2. Tipe Data Pointer

Pointer merupakan tipe data spesial pada golang. Pointer merupakan variabel yang menyimpan address atau alamat nilai dari variabel lain. Pointer memiliki kemampuan untuk mengubah nilai yang ada pada address tersebut yang mana membuat variabel yang nilainya disimpan pada address tersebut akan ikut berubah.

Dua operator penting pada pointer :

a. Referencing (&)

Referencing merupakan kemampuan untuk mengembalikan address dari sebuah variabel dan mengakses address dari sebuah variabel ke variabel pointer.

b. Deferencing (\*)

Deferencing merupakan kemampuan untuk mengakses nilai pada sebuah address dan digunakan untuk mendeklarasikan variabel pointer.

Contoh pointer :

```go
...

func main() {
    // Deklarasi variabel biasa
    var a = 0

    // Deklarasi variabel pointer dimana variabel ini mereference ke variabel a
    var b *int = &a

    fmt.Println("Nilai A :", a)
    fmt.Println("Address A :", &a)
    fmt.Println("Nilai B :", *b)
    fmt.Println("Address dari value B :", b)

    // Mendeklarasikan variabel dengan built-in new
    var c = new(int)
    *c = 1000
    fmt.Println("Nilai C :", *c)
    fmt.Println("Address C :", c)
}

...
```

### 3. Struct dan Interface

#### a. Struct

Struct merupakan tipe user-defined yang berisi collection dari nama field/properti atau fungsi/method. Struct juga merupakan template data yang digunakan untuk menggabungkan satu atau lebih tipe data.

Cara penggunaan struct :

```go
...

// deklarasi type struct
type Customer struct {
    Name string,
    City string,
    Age int
}

func main() {
    // penggunaan struct pertama
    var cust1 Customer

    cust1.Name = "Bill Gates"
    cust1.City = "Jakarta"
    cust1.Age = 56

    fmt.Println(cust1)

    // penggunaan struct kedua (direkomendasikan)
    var cust2 = Customer{
        Name: "Messi",
        City: "Purwodadi",
        Age: 34
    }

    fmt.Println(cust2)

    // penggunaan struct ketiga (tidak direkomedasikan karena harus sesuai dengan urutan field nya)
    var cust3 = Customer{"Ronaldo", "Nganjuk", 36}

    fmt.Println(cust3)
}

...
```

#### b. Interface

Interface merupakan tipe data abstrak dan tidak memiliki implementasi langsung. Interface hanya berisi method-method dan hanya digunakan untuk kontrak perjanjian method. Interface bisa digunakan secara berulang oleh struct yang berbeda. Semua struct yang memiliki fungsi yang ada pada interface, maka berhak untuk menginplementasikan interface tersebut.

Contoh penggunaan interface :

```go
...

type HasName interface {
    GetName() string
}

type User struct {
    Name string
    Age int
}

// mengimplementasikan interface
func (u User) GetName() string {
    return u.Name
}

func main() {
    var user HasName

    user = User{
        Name: "Samsudin",
        Age: 45,
    }

    fmt.Println("Hello,", user.GetName())
}

...
```

Pada golang juga terdapat interface kosong. Interface kosong merupakan interface tanpa deklarasi method yang membuat secara otomatis semua tipe data akan menjadi implementasinya.

Contoh penggunaan interface kosong :

```go
...

func main() {
    // mengimplementasi semua tipe data
    var i interface{}

    i = 10
    fmt.Println(i)

    i = "John Doe"
    fmt.Println(i)
}

..
```

Type assertion merupakan konversi tipe data dari implementasi interface kosong ke tipe data yang diinginkan. Untuk menlakukan konversinya pun harus diperhatikan untuk selalu sesuai dengan tipe data aslinya agar tidak terjadi panic.

Contoh penggunaan type assertion :

```go
...

func main() {
    var random interface{}

    random = 1000

    var result = random.(int) * 50
    fmt.Println("Result:", result)
}

...
```

## Bonus Materi

### 1. Package

Package merupakan sebuah collection dari fungsi dan data. Package digunakan untuk mengorganisir kode program yang ada pada golang.

Contoh penggunaan package :

```go
package main

// rest of code
...
```

### 2. Error Handling

Tidak seperti bahasa pemrograman lainnya, golang tidak memiliki yang namanya try catch, melainkan pada golang dinamakan `panic` dan `recover`. Panic merupakan sebuah fungsi yang memberhentikan jalannya program dan panic biasanya dipanggil jika terdapat error pada saat program berjalan. Sedangkan recover adalah sebuah fungsi yang digunakan untuk menangkap data panic. Dengan adanya recover, panic akan berhenti dan program akan tetap bisa lanjut untuk berjalan.

Contoh penggunaan panic dan recover :

```go
...

func endApp() {
    message := recover()
    if message != nil {
        fmt.Println("Terjadi error message:", message)
    }

    fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
    defer endApp()

    if error {
        panic("APLIKASI ERROR")
    }

    fmt.Println("Aplikasi berjalan")
}

func main() {
    runApp(false)
}

...
```

# (11) Concurrent Programming

#### Arvin Paundra Ardana - Golang A

## 4 Poin Penting pada Materi Concurrent Programming

### 1. Sequential vs Parallel vs Concurrent

- Sequential Program

Sequential program merupakan sebuah proses yang dimana harus menunggu task sebelumnya selesai terlebih dahulu sebelum task selanjutnya dijalankan.

- Parallel Program

Parallel program merupakan proses dimana menjalankan dua atau lebih task secara bersamaan atau simultan. Untuk menjalankan parallel program ini diperlukan lebih dari satu core pada parangkat.

- Concurrent Program

Pada concurrent program, task bisa dikerjakan secara independen dan mungkin muncul secara bersamaan atau simultan. Pada golang juga mendukung concurrent program secara parallel.

### 2. Goroutines

Concurrency pada golang disebut dengan `goroutines`. Goroutines merupakan fungsi atau method yang berjalan secara concurrent (independen) dengan fungsi atau method lainnya. Jika dibandingkan dengan thread, goroutine itu lebih kecil. Fungsi `main` pada golang merupakan salah satu implementasi dari goroutine. Pada golang juga terdapat configurasi untuk membatasi thread sistem operasi untuk mengeksekusi goroutine.

Single goroutine :

```go
package main

import (
    "fmt"
    "time"

    func hello(s string) {
        for i := 0; i < 5; i++ {
            time.Sleep(1 * time.Second)
            fmt.Println(s)
        }
    }

    func main() {
        go func() {
            hello("hello") // menggunakan goroutine
        }()

        hello("world") // menggunakan thread
    }
)
```

Multiple goroutine :

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    go func() {
        fmt.Println("Aku adalah goroutine")
    }()

    go func() {
        fmt.Println("Aku adalah goroutine juga")
    }()

    fmt.Println("Aku bukan goroutine")

    time.Sleep(1 * time.Second)
}
```

Cara membatasi penggunaan thread dalam mengeksekusi goroutine :

```
// tambahkan perintah berikut pada saat akan mengeksekusi program go

$ time GOMAXPROCS=1 go run main.go
```

### 3. Channel dan Select

#### a. Channel

Channel merupakan sebuah objek komunikasi yang digunakan oleh goroutine agar bisa saling berkomunikasi antara goroutine satu dengan goroutine lainnya pada concurrent program. Channel bekerja layaknya stack, yaitu FIFO (First In First Out) dimana nilai pertama yang masuk akan menjadi nilai yang pertama kali keluar. Pada channel terdapat buffered dan unbuffered channel.

- Buffered channel

Buffered channel merupakan channel yang telah ditentukan kapasitasnya (diinisialisasi kapasitasnya) untuk menyimpan messages. Buffered channel akan memblok goroutine jika channel kosong dan akan menunggu sampai terisi. Buffered channel juga akan memblok apabila channel sudah terisi penuh dengan kapasitasnya namun terdapat perintah untuk mengisi channel yang sudah dalam keadaan penuh tersebut.

- Unbuffered channel

Unbuffered channel merupakan channel yang tidak memiliki kapasitas pada waktu inisialisasi. Unbuffered channel akan memblok goroutine jika channel tersebut kosong dan akan menunggu sampai channel tersebut terisi.

Penggunaan channel :

```go
...

func main() {
    // deklarasi channel
    ch := make(chan string)

    go func() {
        ch <- "test" // mengisi channel dengan value ping
    }

    message := <- ch // mengambil nilai yang ada pada channel

    fm.Println("Message pada channel : ", message)
}

...
```

Buffered channel :

```go
...

func main() {
    // deklarasi buffered channel
    ch := make(chan string, 3)

    // rest of code
}
...
```

Unbuffered channel :

```go
...

func main() {
    // deklarasi unbuffered channel
    ch := make(chan string)

    // rest of code
}

...
```

#### b. Select

Select digunakan untuk mempermudah pengontrolan komunikasi data lewat satu ataupun banyak channel. Cara penggunaan select untuk kontrol channel sama seperti menggunakan switch.

Contoh penggunaan select :

```go
package main

import (
    "fmt"
    "runtime"
)

func getAverage(numbers []int, ch chan float64) {
    var sum = 0

    for _, e := range numbers {
        sum += e
    }
    ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
    var max = numbers[0]

    for _, e := range numbers {
        if max < e {
            max = e
        }
    }

    ch <- max
}

func main() {
    runtime.GOMAXPROCS(2)

    var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 3, 4, 6, 3}
    fmt.Println("numbers :", numbers)

    var ch1 = make(chan float64)
    go func(numbers []int, ch chan float64) {
        getAverage(numbers, ch)
    }(numbers, ch1)

    var ch2 = make(chan int)
    go func(numbers []int, ch chan int) {
        getMax(numbers, ch)
    }(numbers, ch2)

    for i := 0; i < 2; i++ {
        select {
            case avg := <- ch1:
                fmt.Printf("Avg \t: %.2f \n", avg)
            case max := <- ch2:
                fmt.Printf("Max \t: %d \n", max)
        }
    }
}
```

### 4. Race Condition

Race condition merupakan dimana lebih dari satu goroutine mengakses data yang sama pada waktu yang sama yang membuat nilai dari data tersebut menjadi kacau.

Menangani data race pada golang dapat dilakukan dengan cara sebagai berikut.

- Blocking dengan WaitGroup

Cara yang paling mudah untuk menangani race condition pada golang, yaitu menggunakan cara block akeses read sampai operasi write itu selesai.

Contoh :

```go
package main

import (
    "fmt"
    "sync"
    "runtime"
)

func doPrint(wg *sync.WaitGroup, message string) {
    defer wg.Done()
    fmt.Println(message)
}

func main() {
    runtime.GOMAXPROCS(2)

    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        var data = fmt.Sprintf("data %d", i)

        wg.Add(1)

        go func() {
            doPrint(&wg, data)
        }()
    }

    wg.Wait()
}
```

- Channel Blocking

Channel blocking mengijinkan goroutine untuk tersinkronisasi dengan goroutine lainnya untuk beberapa saat.

Contoh :

```go
package main

import "fmt"

func getNumber() int {
    var i int

    done := make(chan struct{})

    go func() {
        i = 5

        done <- struct{}{}
    }()

    <- done

    return i
}

func main() {
    fmt.Println(getNumber())
}
```

- Mutex

Mutex atau mutual exclusion merupakan sebuah mekanisme yang digunakan untuk membuka dan mengunci sebuah data.

Contoh :

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type SafeNumber struct {
    val int
    m sync.Mutex
}

func (i *SafeNumber) Get() int {
    i.m.Lock()

    defer i.m.Unlock()

    return i
}

func (i *SafeNumber) Set(val int) {
    i.m.Lock()

    defer i.m.Unlock()

    i.val = val
}

func getNumber() int {
    i := &SafeNumber{}

    go func() {
        i.Set(5)
    }()

    time.Sleep(1 * time.Second)
    return i.Get()
}

func main() {
    fmt.Println(getNumber())
}
```

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

### 2. Data Definition Language (DDL)

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

### 3. Data Manipulation Language (DML)

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

# (14) Configuration Management and CLI

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Configuration Management and CLI

### 1. Command Line Interface

Command Line Interface atau CLI merupakan sebuah interface yang berbasis teks yang sangat cepat dan powerful yang digunakan oleh developer untuk bekerja secara efisien dan efektif untuk berkomunikasi dengan komputer agar mampu menyelesaikan berbagai pekerjaan.

Kelebihan dari CLI antara lain :

a. Kontrol granular dari OS atau aplikasi.

b. Manajemen lebih cepat dari sejumlah sistem informasi.

c. Kemampuan untuk menyimpan script untuk mengotomatisasi pekerjaan umum.

Pada device yang berbasis UNIX menggunakan shell sebagai CLI nya. CLI juga diterapkan pada sejumlah aplikasi, seperti MySQL CLI Client, redis-cli, dsb.

### 2. Unix Shell

Unix Shell merupakan salah satu dari jenis CLI. Pada shell terdapat dua pembagian umum pengguna, yaitu Root User dimana mampu untuk memanipulasi semua direktori yang ada pada device tersebut dan Normal User yang hanya memiliki akses terbatas seperti hanya mampu memanipulasi direktori /home/{username} saja. Ada tambahan satu, yaitu User + Sudoers yang bertindak sebagai Root User secara sementara.

Contoh perintah populer yang sering digunakan pada UNIX Shell :

a. Perintah untuk direktori

- `pwd` untuk menampilkan lokasi ada pada direktori mana user saat ini.

- `ls` untuk menampilkan list berupa nama baik file ataupun folder pada sebuah direktori.

- `mkdir` untuk membuat folder baru pada sebuah direktori. Perhatikan bahwa dalam membuat folder dengan mkdir jika ingin folder tersebut terdapat spasi, maka gunakan backslash (`\`).

- `cd` untuk berpindah direktori satu ke direktori lainnya. Pro tips, untuk berpindah keluar dari direktori saat ini, maka gunakan titik-titk (`../`).

- `rm` untuk menghapus sebuah file. Untuk menghapus sebuah direktori maka bisa menggunakan `rmdir`.

- `cp` untuk menduplikasi sebuah file atau folder ke direktori yang dituju.

- `mv` untuk memindahkan file atau folder ke direktori yang dituju. Perintah mv bisa digunakan untuk mengubah nama jika memindahkan ke dorektori yang sama dan dengan memasukkan nama yang berbeda pada saat melakukan proses pemindahan.

b. Perintah untuk file

- `touch` untuk membuat sebuah file baru dengan nama yang diinginkan.

- `cat` untuk menampilkan isi dari sebuah file (read-only).

- `nano` untuk membuat (jika belum ada) dan mengedit secara langsung file tersebut layaknya sebuah text editor namun bawaan dari CLI.

- `chmod` untuk mengubah permission atau izin dari sebuah file, misal mengubah siapa saja yang dapat melihat, mengubah isi, dan mengeksekusi file tersebut.

c. Perintah untuk network

- `ping` untuk melakukan pengujian koneksi internet dengan mengirimkan "paket" ke host tujuan, misal google.com.

- `ssh` untuk melakukan transfer data antara dua host secara aman karena menggunakan sebuah key.

- `netstat` untuk menampilkan data statistik koneksi jaringan dari dan ke komputer yang sedang digunakan.

- `ifconfig` untuk mengkonfigurasikan atau menampilkan informasi konfigurasi network interface perangkat saat ini.

- `wget` untuk mendownload file dari internet (web).

- `curl` untuk transfer data ke atau dari server menggunakan protokol yang didukung, seperti HTTP, FTP, dll.

d. Perintah untuk utility

- `man` untuk menampilkan petunjuk dari sebuah perintah pada CLI.

- `env` untuk menampilkan seluruh konfigurasi environment variable pada perangkat tersebut.

- `echo` untuk menampilkan text/string yang di pass sebagai argument.

- `date` untuk menampilkan waktu dan tanggal saat ini sesuai dengan sistem.

- `which` untuk mencari file yang dapat dieksekusi atau lokasi program dari sistem file.

- `sudo` untuk memberikan akses sementara Normal User untuk bertindak sebagai Root User.

Mengubah akses permission atau izin dari sebuah file agar dapat dibaca, diubah, atau dieksekusi oleh pengguna. Contoh file dengan permission `-rwxr--r--` jika dibedah menjadi (`- | rwx | r-- | r--`) :

- `-` di paling awal berarti menandakan bahwa itu merupakan sebuah file, jika `d` maka direktori.

- `rwx` awal berarti pemilik atau owner dapat melakukan baca, ubah, dan eksekusi file tersebut.

- `r--` kedua berarti group owner hanya dapat melakukan baca.

- `r--` ketiga berarti pengguna lain hanya dapat melakukan baca.

### 3. Shell Script

Shell merupakan program yang berfungsi sebagai jembatan penghubung antara user dengan kernel. Sedangkan, Shell Script merupakan sebuah bahasa pemrograman yang di-compile oleh perintah shell. File Shell script berekstensi `.sh` dan kemudian diawal baris harus dimulai dengan `#!/bin/sh` yang menandakan keberadaan kompiler/transpiler dari bahasa shell.

Contoh dari Shell Script dengan nama file `test.sh` :

```sh
#!/bin/sh

# echo perintah untuk menampilkan format string
# $1 menandakan menerima satu parameter
echo "Hai, $1"

# menjalankan perintah "date" pada shell script
$(date)
```

Cara mengeksekusi Shell Script diatas :

```
> ./test.sh arvin
```

# (15) System Design

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi System Design

### 1. System Design

#### a. Diagram design

Diagram merupakan sebuah informasi yang direpresentasikan dengan simbol-simbol menggunakan teknik visualisasi.

Software design yang sering digunakan :

- Flowchart atau diagram alir adalah diagram yang menggambarkan secara rinci langkah-langkah atau alur dari proses pemrograman atau algoritma.

- Use case diagram adalah diagram yang meringkaskan detail dari pengguna sistem atau yang dikenal sebagai aktor dan beserta interaksinya dengan sistem tersebut.

- Entity relationship diagram merupakan tipe dari flowchart yang mengilustrasikan bagaimana entitas-entitas, seperti orang, objek, atau konsep yang berhubungan satu dengan yang lainnya dalam sebuah sistem.

- High Level Architecture menjelaskan arsitektur yang akan digunakan untuk mengembangkan sebuah sistem (design sistem secara keseluruhan`) tanpa harus tahu detail implementasi nantinya seperti apa.

#### b. Karakteristik dari sistem terdistribusi

- Scalability

Scalability atau skalabilitas merupakan kemampuan dari sebuah sistem, proses, atau jaringan untuk berkembang dan mengatur permintaan yang meningkat. Semua sistem bisa untuk berkembang karena banyak alasannya, seperti perkembangan volume data atau bertambahnya jumlah pekerjaan, seperti jumlah transaksi.

- Reliability

Reliability merupakan kemungkinan sebuah sistem akan mengalami kegagalan untuk jangka waktu tertentu. Secara sederhana, sebuah sistem terdistribusi dianggap reliable atau dapat diandalkan apabila tetap mampu memberikan service atau layanan bahkan ketika satu atau beberapa hardware atau software-nya mengalami kegagalan saat sedang berjalan.

- Availability

Availability merupakan waktu sistem tetap beroperasi untuk menjalankan fungsi yang diperlukan pada waktu tertentu. Itu merupakan ukuran sederhana dari persentase waktu bahwa suatu sistem, service, atau mesin tetap beroperasional diluar kondisi normal.

- Efficiency

Efficiency merupakan cara bagaimana memaksimalkan kerja dari sebuah sistem agar dapat bekerja secara seefisien mungkin, namun dengan penggunaan resource atau sumber daya yang rendah.

- Serviceability or Manageability

Serviceability or Manageability merupakan kemudahan dan kecepatan dengan yang mana sebuah sistem bisa diperbaiki atau dimaintain atau dirawat. Jika waktu untuk memperbaiki kegagalan sebuah sistem terus meningkat, maka akan berimbas dengan menurunkannya availability atau keberadaan sistem.

#### c. Horizontal scaling dan Vertical scaling

- Horizontal scaling merupakan meningkatkan server ke arah samping atau mengingkatkan kuantitas jumlah server untuk meningkatkan ability atau kemampuan server, seperti contohnya untuk memproses jumlah request yang lebih banyak dibanding sebelumnya.

- Vertical scaling atau meningkatkan server ke atas yaitu meningkatkan resource dari server tersebut, misalnya meningkatkan kapasitas SSD, RAM, atau jumlah core CPU yang dimana cost atau biayanya lebih murah dibanding horizontal scaling.

#### d. Job/Work Queue

Job/Work Queue adalah struktur data yang dikelola oleh perangkat lunak job scheduler yang berisi pekerjaan untuk dijalankan. Pengguna mengirimkan program mereka yang ingin mereka jalankan, "job", ke antrian untuk pemrosesan batch. Scheduler mengurus queue atau antrian sebagai kumpulan dari job yang tersedia untuk dijalankan.

#### e. Load Balancing

Load Balancer merupakan salah satu komponen penting dalam sistem terdistribusi. Load Balancer membantu untuk menyebarkan atau membagi traffic antara cluster dari server untuk meningkatkan responsifitas dan ketersediaan dari aplikasi, website, atau database. Load Balancer juga melacak status dari semua resource ketika sedang mendistribusikan request. Salah satu kegunaan lain dari Load Balancer yaitu dapat mencegah serangan siber seperti DDoS.

### 2. Monolith vs Microservices

- Monolith atau Monolitik adalah aplikasi yang memiliki single code base dengan banyak module didalamnya. Module dibagi menjadi fitur bisnis atau fitur teknis. Monolitik memiliki sistem single build yang mana mem-build keseluruhan aplikasi beserta dependensinya. Monolitik juga hanya memiliki single executable atau deployable binary.

- Microservices merupakan service yang di-deploy secara independen yang telah dimodelkan sesuai dengan domain bisnisnya. Microservice berkomunikasi antar service-nya melalui network dan sebagai sebuah arsitektur pilihan yang menawarkan banyak opsi penyelesaian masalah yang dihadapi.

### 3. Database dan Optimasi

#### a. SQL dan NoSQL

- SQL atau relational database

SQL atau relasional database merupakan database yang terstruktur dan memiliki skema yang telah didefinisikan pada awal pembuatannya. Kelebihan dari SQL yaitu, dirancang untuk segala keperluan, memiliki standar yang jelas, terdapat banyak tools pembantu.

Kelebihan lainnya yaitu mandukung `ACID` :

`Atomicity` : Transaksi terjadi semua atau tidak sama sekali.

`Consistency` : Data tertulis merupakan data valid yang ditentukan berdasarkan aturan tertentu.

`Isolation` : Pada saat terjadi request yang terjadi secara bersamaan (concurrent), memastikan bahwa transaksi dieksekusi seperti dijalankan secara sekuensial.

`Durability` : Jaminan bahwa transaksi yang telah tersimpan itu tetap tersimpan.

- NoSQL atau non-relational database

Tidak seperti SQL, NoSQL tidak terstruktur dan memiliki skema yang dinamis. NoSQL merupakan DBMS yang menyediakan mekanisme yang lebih fleksibel dibandingkan dengan model RDBMS (sifat ACID). NoSQL cocok untuk digunakan ketika membutuhkan skema yang fleksibel, tidak memerlukan ACID, data logging, caching, dsb.

#### b. Caching

Cache digunakan untuk menyimpan data yang baru saja di request agar tidak melakukan request ulang ke tujuan yang sama untuk sementara waktu. Cache merupakan memori jangka pendek yang memiliki jumlah kapasitas yang terbatas namun lebih cepat daripada sumber data dan berisi items yang paling sering diakses.

#### c. Database replication

Redundansi merupakan komponen atau fungsi kritikal dari duplikasi dari duplikasi dari sebuah sistem dengan maksud meningkatkan reliability atau keandalan dari sistem, biasa pada bentuk dari sebuah backup atau untuk meningkatkan kinerja sistem yang sebenarnya.

#### d. Database Indexing

Indexing merupakan sebuah cara untuk mengoptimasi performa dari sebuah database dengan meminimalisir penggunaan akses memori yang diperlukan ketika sedang menjalankan sebuah query database. Indexig merupakan sebuah teknik struktur data yang mana digunakan untuk secara cepat melokasikan dan mengakses data pada sebuah database.

# (16) Intro RESTful API

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Intro RESTful API

### 1. Application Programming Interface (API)

Application Programming Interface atau API adalah kumpulan dari fungsi dan prosedur yang menizinkan pembuatan aplikasi yang mengakses fitur atau data dari sebuah sistem operasi, aplikasi, atau service lainnya. Singkatnya API adalah sebuah sekumpulan fungsi, prosedur, dan sejenisnya yang menjadi penghubung atau jembatan antara program internal dengan program eksternal.

Pada kasus ketika membuat sebuah sistem yang berbasis dengan web, API sangat sering digunakan sebagai media komunikasi antara user web browser dengan server penyedia sistem tersebut untuk saling bertukar data. Contohnya ketika kita ingin membuat sebuah sistem yang terintegrasi dengan pembayaran online dan kita tidak ingin membuat sistem pembayaran online sendiri maka kita bisa menggunakan layanan penyedia pembayaran online, seperti contohnya yang terkenal adalah Midtrans. Untuk menghubungkan sistem internal kita dengan sistem milik Midtrans, kita dapat memanfaatkan yang namanya sebuah API.

### 2. REpresentational State Transfer (REST API)

REpresentational State Transfer atau disingkat dengan REST API adalah sebuah arsitektur pertukaran atau komunikasi data yang mana memanfaatkan protokol HTTP, yaitu request dan response. Request itu sendiri merupakan permintaan yang dikirimkan oleh user melalui web browser untuk nantinya diproses oleh server dan hasilnya akan dikembalikan sebagai response atau jawaban. Bentuk format response yaitu dapat berupa format JSON ataupun XML. Namun seiring berjalannya waktu, JSON lebih sering digunakan karena ringan dan mudah untuk dibaca oleh manusia dari pada XML.

Pada REST API terdapat sebuah HTTP method yang dimana itu merupakan sebuah bentuk representasi dalam melakukan mengaksesan ataupun manipulasi resource data. Contohnya, method GET digunakan untuk menampilkan atau memanggil data, POST untuk mengirimkan data ke server dan biasanya terdapat payload untuk menyimpan data untuk dikirim ke server, PUT untuk mengubah data, DELETE untuk menghapus data, dan masih banyak lagi yang lainnya.

RESTful API atau salah satu implementasi dari REST API juga terdapat HTTP Status Code atau Response Code yang dimana menunjukkan sebuah informasi (dinotasikan dengan angka) dari hasil kembalian yang dikirimkan oleh server ke user. Adapun untuk code-nya seperti 100-199 merupakan response yang bersifat informasional. Lalu, 200-299 berisi response sukses, seperti 200 untuk OK atau request berhasil mengakses sebuah resource. Selanjutnya, 300-399 berisi pesan redirect, seperti 304 yaitu Not Modified atau response tidak berubah dengan response sebelumnya saat mengakses resource yang sama. Kemudian, 400-499 dan 500-599 yang berisi tentang eror baik pada sisi user atau eror pada sisi server, seperti 404 Not Found yaitu user mengakses sebuah resource yang datanya tidak ada pada server dan 500 Internal Server Error yang berarti server tersebut mengalami eror dengan kondisi eror yang belum diketahui penyebabnya.

Selain Browser, ada beberapa tools yang bisa digunakan untuk melakukan request ke server dan biasanya sering digunakan pada saat proses testing atau debugging pada waktu development. Contohnya seperti Postman dan Insomnia.

### 3. REST API Design Best Practices

Untuk membuat sebuah REST API yang baik dan mudah untuk di-develop dikemudian hari maka diperlukan sebuah design yang baik dalam merancang sebuah REST API. Berikut adalah beberapa diantaranya.

a. Gunakan kata benda daripada kata kerja

Menggunakan kata benda lebih disarankan daripada menggunakan kata kerja aksi, karena menurut saya lebih menjelaskan bahwa kata benda ini lebih merepresentasikan bahwa ini dipergunakan untuk menyimpan sebuah resource data dan karena juga nantinya terdapat sebuah HTTP method yang dimana juga nantinya akan semakin menjelaskan untuk aksi-aksinya dalam proses manipulasi data tersebut.

Contoh :

```
Good :
GET /users
GET /users/123

Bad :
GET /getAllUsers
GET /getDetailUser/123
```

b. Gunakan kata jamak

Gunakan kata jamak daripada kata tunggal, karena nantinya di dialamnya pasti berisi banyak data dan bukan hanya berisi satu data.

Contoh :

```
Good:
GET /chats
GET /chats/123

Bad :
GET /chat
GET /chat/123
```

c. Gunakan Resource Nesting atau resource bersarang untuk menunjukkan relasi atau hirarki

Menggunakan Resource Nesting juga memperjelas bahwa terdapat beberapa entitas yang saling berhubungan atau berelasi yang membentuk sebuah kebersinambungan antara entitas tersebut.

Contoh :

```
GET /users
GET /user/123
GET /users/123/repositories/321
```

d. Filtering, Sorting, dan Paging

Best practice untuk melakukan filtering, sorting, dan paging sebuah data, kita dapat memanfaatkan dengan menggunakan query parameter. Query parameter sendiri diletakkan setelah nama endpoint dari API tersebut dan dipisahkan dengan tanda tanya (?) lalu diikuti dengan nama variable query parameter-nya.

Contoh :

```
Filtering :
GET /news?category=sport
GET /stories?lang=id

Sorting :
GET /users?sort=name:asc
GET /users?sort=name:desc

Paging :
GET /menus?limit=10&page=3
```

e. Versioning

Sesuai dengan namanya, agar kualitas API tersebut menjadi lebih baik misal dalam segi performance, keefisienan, keefektifan, maka dari pengembangan itulah diperlukannya sebuah versioning API. Salah satu tujuan versioning selain digunakan untuk menunjukkan keterbaharuan versi pengembangan API, versioning juga biasanya digunakan untuk tetap memberikan akses kepada pengguna API sebelumnya. Ini dilakukan karena biasanya pengguna atau service tersebut tidak kompatibel ketika menggunakan API versi terbaru dan harus dengan terpaksa menggunakan versi terdahulunya maka disitulah versioning API dapat membantu permasalahan tersebut.

Contoh :

```
https://developer.github.com/v3
```

f. API Documentation

API Documentation ini sangat krusial, karena disini merupakan berupa catatan-catatan pengembang dalam melakukan development REST API tersebut, mulai dari apa saja endpoint-nya, kemudian field apa yang dibutuhkan ketika mengakses tersebut, bagian mana saja yang perlu diberikan akses autentikasi dan otorisasi, dan sejenisnya.

# (17) Intro Echo Golang

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Intro Echo Golang

### 1. Echo Golang

Echo merupakan sebuah micro-framework milik golang. Echo menawarkan performa yang tinggi, mudah untuk dipadupadankan dengan library lain (extensible), dan juga minimalis. Kemudian tambahan untuk kemudahan yang ditawarkan oleh echo yaitu, router yang teroptimisasi, data rendering, data binding, middleware, mudah dalam mengatur skalabilitasnya. Echo dikatakan sebuah minimalis framework karena pada framework ini tidak mengusung banyak library atau third party lain seperti contohnya driver untuk database yang mana berarti harus di get secara terpisah.

### 2. Routing dan Controller pada Echo

a. Routing

Handling routing pada echo sangat mudah, dimana yang pertama yaitu buat terlebih dahulu instance dari echo dengan cara :

```go
e := echo.New()
```

Setelah melakukan instasiasi, maka tinggal pilih saja http method yang akan digunakan. Contoh akan menampilkan data users, maka gunakan method `GET`, kemudian isi parameter pertama dengan path atau endpoint-nya, parameter kedua yaitu berupa fungsi handler nya atau controller yang akan melakukan pemrosesan request response, dan parameter ketiga yaitu middleware yang bersifat opsional.

```go
e.GET("/users", GetUsersController)
```

Jika pada path atau endpoint tersebut ingin mengirimkan sebuah parameter yang sifatnya dinamis, seperti data id, maka dapat ditulis seperti berikut.

```go
e.GET("/users/:id", GetUserController)
```

b. Controller

Setelah melakukan membuat sebuah route atau endpoint nya, maka langkah selanjutnya yaitu membuat sebuah controller yang dimana berisi terkait logic dari sistem yang akan dibangun. Fungsi controller pada echo sendiri memerlukan sebuah parameter yaitu `echo.Context` dan return value nya berupa interface error. Berikut contoh sederhana controller untuk menampilkan data users.

```go
...

func GetUsersController(c echo.Context) error {
    // return berupa data JSON
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success get all users",
        "users": users,
    })
}

...
```

### 3. Handling Data pada Echo

a. Data Rendering

Pada echo dapat mengembalikan berbagai macam jenis HTTP response, seperti JSON, XML, HTML, dan masih banyak yang lainnya. Contoh melakukan render data atau response dengan bentuk JSON.

```go
...

func GetProductController(c echo.Context) error {
    product := map[string]interface{} {
        "name": "Samsung J2 Prime",
        "category": "phone",
        "price": 320,
    }

    // return data product sebagai JSON
	return c.JSON(http.StatusOK, products)
}

...
```

b. Retrieve Data

- Request Params

Seperti yang telah disebutkan pada saat membuat endpoint, biasanya pada endpoint tersebut memerlukan sebuah data yang sifatnya dinamis, seperti contohnya yaitu menangkap data id yang dikirimkan oleh client. Contoh :

```go
...

func main() {
    ...

    e.GET("/products/:id", GetIdProduct)

    ...
}

func GetIdProduct(c echo.Context) error {
    // tangkap id menggunakan method Param milik echo.Context
    productId, _ := strconv.Atoi(c.Param(id))

    return c.JSON(http.StatusOK, map[string]interface{}{
        "product_id": productId,
    })
}

...
```

- Query Parameter

Sama seperti request params, query parameter juga merupakan data yang di passing lewat url endpoint, yang membedakan adalah query parameter biasanya bersifat opsional. Contoh :

```go
...

func main() {
    ...

    e.GET("/products", GetProductsController)

    ...
}

func GetProductsController(c echo.Context) error {
    // tangkap query parameter menggunakan method QueryParam milik echo.Context
    search := c.QueryParam("search")
    limit := c.QueryParam("limit")
    page := c.QueryParam("page")

    return c.JSON(http.StatusOK, map[string]interface{}{
        "search": search,
        "limit": limit,
        "page": page,
    })
}

...
```

- Form Value

Form Value digunakan untuk menangkap nilai sesuai antara nama field pada form dan struct melalui bentuk JSON. Contoh :

```go
...

type Product struct {
    Name        string  `json:"name" form:"name"`
    Category    string  `json:"category" form:"category"`
    Price       int     `json:"price" form:"price"`
}

func main() {
    ...

    e.POST("/products", CreateProductController)

    ...
}

func CreateProductController(c echo.Context) error {
    // tangkap value dari form
    name := c.FormValue("name")
    category := c.FormValue("category")
    price := c.FormValue("price")

    product := Product{
        Name: name,
        Category: category,
        Price: price,
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success create product",
        "product": product
    })
}

...
```

c. Binding Data

Data binding mekanismenya mirip seperti menggunakan FormValue, akan tetapi lebih mudah dan lebih singkat dalam penulisannya. Contoh :

```go
...

type Product struct {
    Name        string  `json:"name" form:"name"`
    Category    string  `json:"category" form:"category"`
    Price       int     `json:"price" form:"price"`
}

func main() {
    ...

    e.POST("/products", CreateProductController)

    ...
}

func CreateProductController(c echo.Context) error {
    // buat instance Product sebagai penampung payload
    product := Product{}
    c.Bind(&product)

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success create product",
        "product": product
    })
}

...
```

# (18) ORM and Code Structure MVC

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi ORM and Code Structure

### 1. Definisi ORM

Object Relational Mapping atau ORM merupakan sebuah teknik pemrograman dengan mengkonversi data diantara sistem yang tidak kompatibel menggunakan paradigma pemrograman OOP. Salah satu library ORM golang yang terkenal yaitu adalah `gorm`. Pada gorm sangat banyak fitur yang mendukung proses development, seperti database relationship, join, migration, seeder, dan masih banyak lagi. Gorm mendukung database migration yang dimana merupakan sebuah mekanisme untuk meng-update versi database untuk memenuhi perubahan versi aplikasi. Update ini sendiri beisa mencakup ke versi terbaru atau melakukan rollback untuk ke versi sebelumnya. Database migration ini sangat berguna karena dapat melakukan update dan rollback secara mudah, melacak struktur pada database, menuliskan riwayat struktur database pada sebuah kode, dan akan selalu kompatibel dengan perubahan aplikasi.

### 2. Kelebihan dan Kekurangn Menggunakan ORM

Penggunaan ORM sendiri memiliki beberapa kelebihan dan kekurangan, seperti sebagai berikut.

a. Kelebihan ORM

- Mengurangi penulisan query yang berulang.

- Secara otomatis melakukan konversi data menjadi bentuk object.

- Dapat melakukan screening data atau pengecekan data dengan mudah.

- Beberapa ORM mendukung caching pada query-nya.

b. Kekurangan ORM

- Menambahkan baris kode baru yang menyebabkan proses berlebih.

- Memuat data relasi yang tidak diperlukan.

- Tidak efisien jika digunakan untuk melakukan query yang kompleks, seperti join lebih dari 10 tabel.

- Tidak mendukung beberapa fungsi SQL bawaan, seperti Force Index pada MySQL.

### 3. Model - View - Controller (MVC)

MVC merupakan sebuah cara yang populer untuk melakukan pengorganisiran folder atau kode pada sebuah projek. Ide dibalik MVC yitu masing-masing bagian dari kode tersebut memiliki tujuan dan masing-masing tujuan tersebut berbeda antara satu dengan lainnya. Tujuan diperlukannya strukturisasi yaitu untuk membuat aplikasi menjadi modular dengan mengimplementasikan memisahkan fokus urusan per masing-masing bagian dan juga mengurangi konflik ketika versioning.

Adapun fungsi masing-masing dari tiga bagian MVC :

a. Model

Model merupakan bagian yang bertugas untuk menyiapkan, mengatur, memanipulasi, dan mengorganisasikan data yang ada di database.

b. View

View merupakan bagian yang bertugas untuk menampilkan informasi dalam bentuk Graphical User Interface (GUI).

c. Controller

Controller merupakan bagian yang bertugas untuk menghubungkan serta mengatur model dan view agar dapat saling terhubung.

# (19) Middleware

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Middleware

### 1. Pengertian Middleware

Middleware adalah entitas yang menghubungkan proses request atau response sebuah server. Singkatnya, middleware merupakan blok kode yang dieksekusi oleh server antara setelah menerima request dari client dan sebelum mengembalikan response ke client.

### 2. Kegunaan Middleware

Middleware mengizinkan kita untuk mengimplementasikan berbagai macam fungsionalitas seperti yang didefinisikan, yaitu diantara HTTP sesudah request dan sebelum response. Adapun contoh kegunaan dari middleware yaitu, melakukan logging aktivitas yang terjadi pada sistem, pengecekan autentikasi user, handle CORS, handle trailing slash pada URI, dan masih banyak kegunaan lainnya.

### 3. Middleware pada Echo

Echo menyediakan banyak middleware, seperti semua yang disebutkan pada poin kedua. Namun di echo dibagi lagi menjadi dua bagian, yaitu middleware yang dijalankan sebelum router melakukan proses (`Echo#Pre`) dan middleware yang dijalankan sesudah router melakukan proses dan juga memiliki akses penuh pada API context echo (`Echo#Use`). Contoh Echo#Pre yaitu meng-override HTTP method dan menambahkan trailing slash. Kemudian contoh Echo#Use yaitu, membatasi ukuran dari request body dan logger.

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

# (21) Clean and Hexagonal Architecture

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada materi Clean and Hexagonal Architecture

### 1. Kendala Sebelum Mendesign Clean Architecture

Berikut ini merupakan kendala yang bisa diselesaikan dengan menggunakan Clean Architecture :

- Independent of Framework

Pada Clean Architecture mengijinkan kita untuk menggunakan seperti contohnya framework sebagai sebuah tools.

- Testable

Business logic bisa ditesting tanpa harus menggunakan UI, database, web server, atau elemen eksternal lainnya.

- Independent of UI

User Interface atau UI bisa diubah dengan mudah, tanpa harus mengubah sistem lainnya.

- Independent of Database

Business logic tidak terikat dengan database, jadi kita bisa mengubah atau swap database dengan mudah.

- Independent of Any External

Pada kenyataannya, business logic tidak perlu tahu tentang apa yang ada diluar lingkaran dari diagram Clean Architecture, seperti contohnya tidak perlu tahu framework apa yang digunakan atau arsitektur komunikasi pertukaran data apa yang akan digunakan.

### 2. Benefit dari Clean Architecture

Berikut merupakan keuntungan dari Clean Architecture :

- Clean Architecture merupakan sebuah struktur standar.

- Development lebih cepat dalam jangka panjang.

- Mocking dependency menjadi lebih mudah dalam melakukan unit test.

- Mudah mengganti dari prototype ke solusi yang sesuai.

### 3. Struktur dari Clean Architecture

Berikut merupakan struktur dari Clean Architecture :

- Entities atau Repository layer

Business object yang mencerminkan konsep dari aplikasi yang dikelola.

- Use case atau Domain atau Service layer

Berisikan business logic dari aplikasi yang didevelop.

- Controller atau Presentation layer

Mempresentasikan data ke sebuah screen dan menangani interaksi pengguna.

- Drivers atau Data layer

Mengatur data aplikasi, contohnya menerima data dari network atau mengatur data cache.

# (22) Docker

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Docker

### 1. Gambaran Tentang Docker

Docker merupakan sebuah open platform untuk developing, shipping, dan running aplikasi. Docker memungkinkan kita untuk memisahkan aplikasi dengan infrastruktur, jadi kita bisa mendeliver software secara cepat. Dengan docker, kita bisa mengelola infrastruktur dengan cara yang sama dengan mengelola aplikasi yang sedang kita kembangkan. Dengan mengambil keuntungan dari metodologi dari docker untuk shipping, testing, dan deploy kode secara cepat, kita bisa secara signifikan mengurangi delay waktu antara menulis kode dengan menjalankannya di production.

### 2. Docker Volume

Docker volume bisa dianggap sebagai storage ata tempat penyimpanan data di container. Karena ketika container mati, secara default data yang ada pada container tersebuta akan ikut terhapus. Untuk itu kita perlu memanfaatkan volume pada docker. Berikut merupakan perintah yang sering digunakan.

- Menampilkan list volume

```
> docker volume ls
```

- Membuat volume

```
> docker volume create <nama_volume>
```

- Menghapus volume

```
> docker volume rm <nama_volume>
```

### 3. Docker Network

Defaultnya container pada docker akan saling terisolasi satu sama lain. Misalnya, kita tidak dapat melakukan request API dari container satu ke container lain. Untuk itu kita harus membuat dan mendaftarkan container pada network yang sama. Berikut merupakan perintah yang sering digunakan.

- Menampilkan list network

```
> docker network ls
```

- Membuat network baru

```
> docker network create <nama_network>
```

- Menghapus network

```
> docker network rm <nama_network>
```

# (23) Compute Service

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Compute Service

### 1. Definisi Deployment

Deployment adalah kegiatan yang bertujuan untuk menyebarkan aplikasi/produk yang telah dikerjakan oleh para pengembang seringkali untuk mengubah dari status sementara ke permanen. Penyebarannya dapat melalui beragam cara tergantung dari jenis aplikasinya, aplikasi web dan API ke server sedangkan aplikasi mobile ke Playstore/Appstore.

### 2. Strategi Deployment

Berikut merupakan strategi dalam melakukan deplyment.

a. Big-Bang/Replacement Deployment Strategy

Kelebihan :

- Mudah diimplementasikan. Cara klasik yaitu tinggal replace projek.

- Perubahan kepada sistem langsung 100% secara instan.

Kekurangan :

- Terlalu beresiko, rata-rata downtime cukup lama.

b. Rollout Deployment Strategy

Kelebihan :

- Lebih aman dan less downtime dari versi sebelumnya.

Kekurangan :

- Akan ada 2 versi aplikasi yang berjalan secara bersamaan sampai semua server terdeploy dan bisa membuat bingung.

- Karena sifatnya perlahan satu persatu, untuk deployment dan rollback lebih lama dari yang Bigbang, karena prosesnya perlahan-lahan sampai semua server terkena efeknya.

- Tidak ada kontrol request. Server yang baru terdeploy dengan aplikasi versi baru langsung mendapat request yang sama banyaknya dengan server yang lain.

c. Blue/Green Deployment Strategy

Kelebihan :

- Perubahan sangat cepat, sekali switch service langsug berubah 100%.

- Tidak ada issue beda versi pada service seperti yang terjadi pada Rollout Deployment.

Kekurangan :

- Resource yang dibutuhkan lebih banyak. Karena untuk setiap deployment kita harus menyediakan service yang serupa environmentnya dengan yang sedang berjalan di production.

- Testing harus benar-benar sangat diprioritaskan sebelum di switch, aplikasi harus kita pastikan aman dari request yang tiba-tiba banyak.

d. Canary Deployment Strategy

Kelebihan :

- Cukup aman

- Mudah untuk rollback jika terjadi error/bug, tanpa berimbas kesemua server.

Kekurangan :

- Untuk mencapai 100% cukup lama dibanding dengan Blue/Green Deployment.

### 3. Proses Deployment

Proses deployment dapat memiliki satu atau banyak langkah, langkah-langkah dapat berjalan secara berurutan atau paralel, selain berbagai langkah penerapan, dapat juga menyertakan langkah-langkah intervensi manual untuk keluar sebelum deployment, langkah-langkah pemberitahuan email untuk memberi tahu semua orang tentang proses, atau bahkan melewatkan langkah dalam keadaan yang berbeda.

# (24) CI/CD

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi CI/CD

### 1. Continuous Integration (CI)

Continuous Integration merupakan sebuah otomasi proses. Pada CI secara otomatis melakukan proses seperti build, test, dan mengintegrasikan code pada sebuah remote repository contohnya Github.

### 2. Continuous Deployment/Delivery (CD)

Pada CD ini terdapat dua versi, yaitu deployment dan delivery. Perbedaan yang menonjol diantara keduanya yaitu untuk Continuous Deployment merupakan otomasi untuk mendeploy perubahan code langsung ke server dan langsung dapat digunakan oleh user. Sedangkan Continuous Delivery merupakan otomasi deliver perubahan code ke server environment production dan harus menunggu adanya approval agar dapat dilanjutkan ke proses deployment dan proses deploy ini dilakukan secara manual.

### 3. Benefit dari CI/CD

Berikut ini merupakan benefit atau keuntungan dari CI/CD.

- Development velocity

Feedback yang berkelanjutan memungkinkan developer untuk melakukan perubahan kecil lebih sering, dibandingkan menunggu satu rilis penuh.

- Stability and reliability

Testing otomatis dan berkelanjutan memastikan bahwa codebase tetap stabil dan siap dirilis kapan saja.

- Business growth

Dibebaskan dari tugas manual, jadi kita dapat memfokuskan sumber daya pada inovasi, kepuasan pelanggan, dan fokus hal lainnya.

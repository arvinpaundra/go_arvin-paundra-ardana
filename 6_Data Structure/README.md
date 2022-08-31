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

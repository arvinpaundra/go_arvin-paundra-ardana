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

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

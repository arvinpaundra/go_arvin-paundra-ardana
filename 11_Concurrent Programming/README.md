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

# Recursive - Number Theory - Sorting - Searching

#### Arvin Paundra Ardana - Golang A

### Problem 1 - Prima ke X

Link [Problem 1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/8_Recursive%20%E2%80%93%20Number%20Theory%20%E2%80%93%20Sorting%20%E2%80%93%20Searching/praktikum/Problem%201/main.go)

Problem 1, membuat sebuah fungsi yang menampilkan bilangan prima sesuai dengan deret urutannya.

Untuk menyelesaikannya bisa dengan membuat dua buah variabel, yang pertama merupakan variabel slice primes untuk menyimpan data bilangan prima yang digenerate dan index untuk inisialisasi perulangan. Setelah itu lakukan pengecekan bila inputan bernilai satu maka langsung me-return bilangan dua, dan jika lebih dari dua maka akan masuk ke sebuah perulangan. Sejatinya pada perulangan tersebut merupakan sebuah endless loop, namun nanti di dalamnya diberi sebuah kondisi sebagai break point agar perulangan tersebut tidak berhenti. Pada perulangan tersebut diberikan kondisi, yang pertama yaitu mengecek bila variabel index bernilai dua dan tiga maka akan dimasukkan ke dalam variabel primes untuk menyimpan bilangan prima. Kemudian sama jika index bernilai lima maka akan dimasukkan ke slice primes. Setelah itu, jika tidak sesuai pada kondisi keempat di atas maka, akan dicek pada increment index bila bernilai genap atau dapat dibagi tiga atau dapat dibagi lima maka bilangan tersebut bukan prima dan akan dimasukkan ke slice primes. Jika panjang slice primes sama dengan inputan maka perulangan akan break atau berhenti. Setelah itu, akan mereturn kan bilangan prima pada index ke `number-1`.

Screenshot :

![Problem 1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/8_Recursive%20%E2%80%93%20Number%20Theory%20%E2%80%93%20Sorting%20%E2%80%93%20Searching/screenshots/Screenshot_22.png)

### Problem 2 - Fibonacci (Recursive)

Link [Problem 2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/8_Recursive%20%E2%80%93%20Number%20Theory%20%E2%80%93%20Sorting%20%E2%80%93%20Searching/praktikum/Problem%202/main.go)

Problem 2, menampilkan deret fibonacci ke-n menggunakan rekursif.

Untuk menampilkan deret fibonacci ke-n menggunakan rekursif yaitu yang pertama mengecek bila inputan number bernilai kurang dari sama dengan satu maka akan akan me-return satu. Kemudian, jika input number bernilai lebih dari satu, maka akan me-return dirinya sendiri (rekursif) dengan parameternya dikurangi satu kemudian ditambahkan dengan dirinya sendiri (rekursif) dengan parameternya dikurangi dua.

Screenshot :

![Problem 2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/8_Recursive%20%E2%80%93%20Number%20Theory%20%E2%80%93%20Sorting%20%E2%80%93%20Searching/screenshots/Screenshot_23.png)

### Problem 3 - Prima Segi Empat

Link [Problem 3](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/8_Recursive%20%E2%80%93%20Number%20Theory%20%E2%80%93%20Sorting%20%E2%80%93%20Searching/praktikum/Problem%203/main.go)

Problem 3, membuat sebuah segiempat dengan high x wide yang berisikan bilangan prima setelah start, dan kemudian hitung jumlahnya.

Untuk menyelesaikan masalah tersebut dapat dilakukan dengan cara, yang pertama dengan membuat empat buah variabel antara lain, total untuk seberapa banyak bilangan prima yang perlu ditampilkan, primes untuk menyimpan data bilangan prima setelah bilangan start, index untuk inisialisasi perulangan yang dimulai dari `start+1`, dan sum untuk menyimpan jumlah penambahan dari bilangan prima yang ada pada slice primes. Setelah itu lakukan perulangan dengan index lebih dari 0. Pada perulangan ini sama seperti pada perulangan problem satu dengan break point ada di dalam loop tersebut. Pada perulangan diberikan kondisi, yaitu mengecek apakah index sama dengan dua atau tiga atau lima, jika true maka akan dimasukkan ke dalam slice primes. Jika false, maka akan masuk ke dalam else dimana didalam else terdapat kondisi untuk mengecek bila index tidak dapat dibagi habis oleh dua, tiga dan lima, maka bilangan tersebut bilangan prima dan akan dimasukkan ke dalam slice primes. Setelah itu dicek apakah panjang dari slice primes sama dengan nilai pada variabel total, jika true maka perulangan tersebut akan break atau terhenti dan jika false maka akan terus berlanjut perulangan tersebut. Setelah itu loop slice primes tersebut untuk dicetak nilainya dan simpan penjumlahan bilangan prima tersebut pada variabel sum. Terakhir tampilkan total penjumlahan bilangan prima tersebut pada variabel sum.

Screenshot :

![Problem 3](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/8_Recursive%20%E2%80%93%20Number%20Theory%20%E2%80%93%20Sorting%20%E2%80%93%20Searching/screenshots/Screenshot_26.png)

### Problem 4 - Total Maksimum dari Deret Bilangan

Link [Problem 4](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/8_Recursive%20%E2%80%93%20Number%20Theory%20%E2%80%93%20Sorting%20%E2%80%93%20Searching/praktikum/Problem%204/main.go)

Problem 4, menemukan total nilai maksimum jumlah bilangan dari deret sebuah list integer secara berurutan.

Ide untuk menyelesaikan problem 4 yaitu dapat dengan menggunakan Kadane's Algorithm. Pertama yaitu membuat variabel inisialisasi nilai minimum dari sebuah integer disini dengan nama MinInt, kemudian variabel currentMax untuk menyimpan nilai terbesar saat ini dengan nilai inisialisasi 0, lalu variabel max dengan nilai yang di assign dengan variabel MinInt. Setelah itu lakukan perulangan, dengan index dari 0, kurang dari panjang array input - 1, dan index increment. Di dalam perulangan, assign ulang variabel currentMax dengan nilai array pada index saat ini. Lakukan pengecekan bila value dari currentMax lebih besar dari max, maka assign ulang nilai max dengan nilai currentMax. Jika currentMax kurang dari 0 maka currentMax di assign kembali menjadi 0. Setelah perulangan tersebut selesai, return nilai max.

Screenshot :

![Problem 4](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/8_Recursive%20%E2%80%93%20Number%20Theory%20%E2%80%93%20Sorting%20%E2%80%93%20Searching/screenshots/Screenshot_27.png)

### Problem 5 - Find Min and Max Number

Link [Problem 5](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/8_Recursive%20%E2%80%93%20Number%20Theory%20%E2%80%93%20Sorting%20%E2%80%93%20Searching/praktikum/Problem%205/main.go)

Problem 5, membuat sebuah program untuk mencari bilangan minimum dan maksimum dari sebuah array dan munculkan juga index dari bilangan tersebut.

Untuk menyelesaikan problem 5 yaitu dapat dilakukan dengan membuat variabel result untuk menyimpan output yang telah diformat, menginisialisasi variabel min untuk menyimpan nilai terkecil dengan inisialisasi nilai input array pada index ke 0, max untuk menyimpan nilai terbesar dengan inisialisasi nilai input array pada index ke 0, minIdx dan maxIdx yang masing-masing digunakan untuk menyimpan index dari bilangan. Setelah itu lakukan perulangan pada array input. Di dalam perulangan tersebut dilakukan pengecekan jika value lebih besar dari nilai max maka assign ulang variabel max dengan value tersebut dan assign ulang juga index-nya pada variabel maxIdx. Lalu jika variabel min kurang dari nilai min maka assign ulang nilai variabel min dengan value tersebut dan assign ulang juga index-nya pada variabel minIdx. Setelah itu, concate sesuai dengan output yang ditentukan dengan bantuan package `strings` dan simpan pada variabel result. Dan terakhir, return result.

Screenshot :

![Problem 5](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/8_Recursive%20%E2%80%93%20Number%20Theory%20%E2%80%93%20Sorting%20%E2%80%93%20Searching/screenshots/Screenshot_28.png)

### Problem 6 - Maximum Buy Product

Link [Problem 6](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/8_Recursive%20%E2%80%93%20Number%20Theory%20%E2%80%93%20Sorting%20%E2%80%93%20Searching/praktikum/Problem%206/main.go)

Problem 6, menampilkan jumlah barang yang dapat dibeli dengan berdasarkan uang yang diberikan.

Untuk menyelesaikan problem ini, buat variabel total untuk menyimpan jumlah barang yang dapat dibeli. Setelah itu sebelum melakukan perulangan, urutkan terlebih dahulu array inputan. Disini menggunakan algoritma `Bubble Sort` untuk melakukan pengurutan harga produk. Pada konsep bubble sort yaitu melakukan perbandingan pada setiap elemen pada array, jika nilai tersebut lebih besar atau lebih kecil dari elemen pembandingnya, maka akan dilakukan swap elemen. Setelah selesai melakukan pengurutan, loop array input tersebut yang telah diurutkan. Di dalam perulangan parameter money di assign ulang dengan money - productPrice index ke i. Lalu dilakukan pengecekan jika money lebih besar sama dengan 0, maka variabel total akan ditambah 1. Setelah perulangan selesai, cetak variabel total tersebut untuk menampilkan total produk yang telah dibeli.

Screenshot :

![Problem 6](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/8_Recursive%20%E2%80%93%20Number%20Theory%20%E2%80%93%20Sorting%20%E2%80%93%20Searching/screenshots/Screenshot_29.png)

### Problem 7 - Playing Domino

Link [Problem 7](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/8_Recursive%20%E2%80%93%20Number%20Theory%20%E2%80%93%20Sorting%20%E2%80%93%20Searching/praktikum/Problem%207/main.go)

Problem 7, membuat program permainan domino.

Untuk menyelesaikan problem 7, buat variabel temp dan max dengan inisialisasi slice kosong. Lakukan loop untuk parameter cards, dimana di dalamnya meng assign ulang variabel temp dengan nilai `cards[i]` dan max dengan nilai `cards[i+1]`. Setelah itu lakukan pengecekan apabila jumlah pada cards[i] lebih sama dengan jumlah deck[i] dan salah satu dari cards sama dengan salah deck. Kemudian di dalam pengecekan tersebut terdapat pengecekan lagi, apabila jumlah dari max lebih besar dari temp, maka yang akan return adalah max, jika tidak maka yang direturn adalah temp.

Screenshot :

![Problem 7](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/8_Recursive%20%E2%80%93%20Number%20Theory%20%E2%80%93%20Sorting%20%E2%80%93%20Searching/screenshots/Screenshot_30.png)

### Problem 8 - Most Appear Item

Link [Problem 8](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/8_Recursive%20%E2%80%93%20Number%20Theory%20%E2%80%93%20Sorting%20%E2%80%93%20Searching/praktikum/Problem%208/main.go)

Problem 8, menampilkan item yang paling sering muncul yang diurutkan berdasarkan jumlah kemunculannya.

Untuk menyelesaikannya, buat variabel map temp untuk menyimpan nilai pada array items dan variabel slice result untuk menyimpan hasil akhirnya. Selenjutnya loop array items untuk diubah menjadi map, hal ini dilakukan karena map memiliki key yang unique yang membuat array item tidak bisa terduplikasi, dan dalam knoversi tersebut akan menambahkan value jika terdapat key yang sama. Setelah itu loop map tersebut untuk diubah menjadi bentuk array dari pair. Terakhir return result.

Screenshot :

![Problem 8](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/8_Recursive%20%E2%80%93%20Number%20Theory%20%E2%80%93%20Sorting%20%E2%80%93%20Searching/screenshots/Screenshot_31.png)

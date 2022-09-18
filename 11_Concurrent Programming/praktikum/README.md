# Concurrent Programming

#### Arvin Paudra Ardanaa - Golang A

### Letter Frequency

Link [Penyelesaian](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/11_Concurrent%20Programming/praktikum/Problem%201/main.go).

Menghitung frekuensi huruf pada teks menggunakan perhitungan parallel.

Untuk menghitung frekuensi huruf secara paralel pada sebuah teks yaitu dapat dicapai menggunakan map dan juga memanfaatkan fungsi `goroutine` pada golang. Disini saya menggunakan `mutex` package built-in dari golang untuk menghandle apabila terjadi race condition. Berikut penyelesaiannya.

- Pertama definisikan terlebih dahulu variabel mutex, words, dan map frequency yang dimana nanti key nya berupa huruf yang ada pada word dan value nya merupakan frekuensi jumlahnya.

- Selanjutnya yaitu membuat fungsi goroutine dengan anonymous function yang menerima 2 buah parameter, yang pertama untuk menerima words dan kedua pointer dari mutex. Di dalam fungsi ini pada awal line kunci terlebih dahulu fungsi ini menggunakan `mutex.Lock()` agar goroutine lain tidak langsung mengakses nilai dari map frequency, karena nantinya pada fungsi ini akan dilakukan konversi huruf yang ada pada words ke dalam map frequency. Setelah lock goroutine ini, selanjutnya terdapat defer `mutex.Unlock()` yang tujuannya ketika fungsi ini selesai maka akan men-defer fungsi unlock dari mutex. Beralih dari mutex, proses berlanjut untuk mengubah semua karakter menjadi huruf kecil (lower case) dan disimpan pada variabel lowerCase. Setelah diubah menjadi lower case, Ubah menjadi slice of string dengan bantuan method Split dari package strings dan simpan pada variabel splittedWords. Loop slice tersebut dan masukkan elemennya ke map frequency (menjadi key). Jika pada proses loop tersebut pada slice nya terdapat lebih dari satu maka nantinya value dari key nya bertambah (menjadi frekuensi).

- Setelah sudah dapat frekuensinya yang disimpan pada map frequency, hapus key yang bukan alpabet a-z. Gunakan fungsi goroutine dengan anonymous function yang menerima parameter pointer mutex. Pada awal line fungsi sama seperti fungsi sebelumnya yaitu me-lock terlebih dahulu dan me-unlock ketika goroutine selesai dijalankan. Untuk mengahapus selain aplabet dapat melakukan loop pada map frequency, dimana didalamnya terdapat kondisi jika key tersebut kurang dari "a" atau lebih dari "z" maka hapus key tersebut. Setelah itu cetak map frequency untuk menampilkan jumlah frekuensi huruf pada teks di variabel words yang yang hanya terdiri dari a-z saja.

- Terakhir, karena proses eksekusi goroutine sangat cepat dan agar tetap bisa menampilkan hasil dari proses perhitungan frekuensi huruf, set fungsi goroutine main untuk selesai dijalankan setelah satu detik kemudian dengan menggunakan package time.

Screenshot hasil :

![Letter Frequency](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/11_Concurrent%20Programming/screenshots/Screenshot_6.png)

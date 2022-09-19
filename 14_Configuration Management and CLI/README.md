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

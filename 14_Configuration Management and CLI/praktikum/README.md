# Configuration Management and CLI

#### Arvin Paundra Ardana - Golang A

### Problem 1

Screenshot hasil instal oh-my-zsh :

![oh-my-zsh](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/14_Configuration%20Management%20and%20CLI/screenshots/Screenshot_7.png)

### Problem 2

Membuat automation script menggunakan shell script.

Link [Automation.sh](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/14_Configuration%20Management%20and%20CLI/praktikum/Problem%202/automate.sh).

Link hasil generate [automation](https://github.com/arvinpaundra/go_arvin-paundra-ardana/tree/master/14_Configuration%20Management%20and%20CLI/praktikum/Problem%202/arvin%20at%20Mon%20Sep%2019%2015.04.05%20WIB%202022).

Untuk membuat sebuah automation script menggunakan shell script sesuai dengan instruksi yang telah diberikan, saya melakukan langkah-langkah berikut.

a. Pertama, buat sebuah variabel date yang bernama timestamp yang berguna untuk menyimpan nilai timestamp atau waktu saat file tersebut dibuat, dimana menggunakan perintah `date` dan diikuti dengan %H yaitu jam dalam format 24 jam, %M yaitu menit, dan %S yaitu detik.

b. Setelah itu membuat directory root menggunakan perintah `mkdir` yang bernama sesuai dengan parameter pertama atau `$1` yang dimasukkan oleh author. Kemudian, untuk menambahkan waktu pada pembuatan folder, saya menggunakan perintah date untuk hari, bulan, tahun dan beserta timestamp-nya.

c. Kemudian, untuk membuat child directory nya sama seperti membuat root directory dan hanya tinggal menambahkan nama folder sesuai dengan perintah yang tertera.

d. Untuk membuat file nya, saya menggunakan perintah `touch` dan kemudian diteruskan dengan nama folder dan diakhiri dengan nama file dan format file nya.

e. Jika direktori dan file tersebut dibuat, maka untuk memasukkan nama sosial medianya yang ada pada parameter kedua `$2` untuk facebook dan parameter ketiga `$3` atau linkedin pada directory about_me, saya menggunakan perintah `echo` untuk menuliskan sebuah teks yang nantinya akan dimasukkan kedalam file sesuai dengan yang ditujunya dan kemudian ditambahkan parameternya yang sesuai pada tempatnya, misalnya facebook untuk parameter kedua dan linkedin untuk parameter ketiga.

f. Setelah itu untuk melakukan download data friendlist yang ada directory my_friends pada file txt, saya menggunakan perintah `curl` untuk mendownload data tersebut dan menggunakan perintah `echo` untuk menuliskan data tersebut ke dalam file list_of_my_friends.txt.

g. Selanjutnya yaitu memasukkan informasi device dari author ke dalam file pada directory my_system_info lalu file about_this_laptop, saya menggunakan yaitu perintah `uname` untuk menampilkan username, `hostname` untuk menampilkan hostname, `uname -a` untuk menampilkan data terkait user pada device tersebut. Kemudian untuk menuliskannya, saya menggunakan perintah `echo` dan agar tidak mereplace file tersebut, saya menambahkan ``>>``` supaya lanjut ke baris selanjutnya.

h. Terakhir, untuk menuliskan koneksi internet saat ini pada file internet_connection.txt pada directory my_system_info, saya menggunakan perintah `ping` dan diikuti dengan target name nya, misalnya google.com. Sama seperti lainnya, untuk menuliskan kedalam file tersebut, saya menggunakan perintah `echo` dan nantinya akan secara otomatis dimasukkan ke dalam file tersebut hasil dari informasi koneksi internet tadi.

Screenshot hasil akhir :

![automate.sh](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/14_Configuration%20Management%20and%20CLI/screenshots/Screenshot_8.png)

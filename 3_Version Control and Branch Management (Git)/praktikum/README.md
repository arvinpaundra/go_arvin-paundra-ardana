
# Version Control and Branch Management (Git)

#### Arvin Paundra Ardana - Golang A

### 1 - Penggunaan Branch pada Git
- Branch master (default)
Secara default, git menggunakan nama branch master pada awal atau inisialisasi pembuatan repository.

- Branch development
Untuk membuat sebuah branch baru  gunakan perintah :
```
$ git branch development
```
maka nanti akan dibuatkan oleh git sebuah branch bernama development. Setelah membuat branch, untuk berpindah ke branch lain (contohnya ke branch master) itu dapat menggunakan perintah ```git checkout master```.

- Branch featureA
Terdapat shortcut untuk membuat lalu berpindah ke branch baru (contohnya pindah ke branch baru featureA), yaitu menggunakan perintah :
```
$ git checkout -b featureA
```
maka secara otomatis akan membuat branch featureA dan langsung pindah ke branch tersebut.

- Branch featureB
Pilih cara diantara kedua tadi untuk membuat branch featureB.
```
$ git branch featureB
$ git checkout featureB

# atau

$ git checkout -b featureB
```

### 2 - Instruksi Perintah Git
- Perintah push

contoh perintah :
```
$ git push origin master
```

Penjelasan :

Perintah push digunakan untuk mengirimkan perubahan kode yang sudah ada di stagging area menuju ke remote repository pada branch master, contohnya github.

- Perintah pull

contoh perintah :
```
$ git pull origin master
```

Penjelasan :

Perintah pull pada git digunakan untuk mengambil commit terbaru beserta perubahan kode pada sebuah branch (contohnya master) dari remote repository ke local repository atau komputer kita.

- Perintah stash

contoh implementasi :
```
$ git stash
$ git stash pop
```

Penjelasan :

Perintah stash biasanya digunakan untuk memindahkan perubahan saat ini (contohnya file terbaru atau perubahan yang belum di stagging) menuju ke stash area atau penyimpanan sementara milik stash. Setelah itu, kode secara otomatis akan kembali seperti waktu terakhir kali melakukan commit. Untuk mengembalikan perubahan yang ada pada stash area yaitu dapat dengan melakukan perintah ```git stash pop```.

- Perintah merge

contoh implementasi :
```
# saat ini ada di branch development.
$ git merge featureA
```

Penjelasan :

Perintah merge digunakan untuk pelakukan penggabungan (merging) antar branch (contohnya pada branch development akan melakukan penggabungan dengan branch featureA).

### 3 - Implementasi Penanganan Sebuah Conflict pada Git
Untuk melakukan penanganan conflict pada git yaitu pada saat melakukan merge dapat dilakukan seperti dibawah ini.

- Awal conflict terjadi
![Gambar 1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/3_Version%20Control%20and%20Branch%20Management%20(Git)/screenshots/Screenshot_16.png)
Pada gambar tersebut terdapat conflict atau eror pada saat merging, karena pada branch development sudah melakukan merge dengan branch featureA yang mengubah line code ke 2 dan juga mengakibatkan pada branch development terdapat 1 commit lebih maju dari sebelumnya. Itulah hal yang menyebabkan conflict pada saat branch development melakukan merge dengan branch featureB.

- Fix conflict
![Gambar 2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/3_Version%20Control%20and%20Branch%20Management%20(Git)/screenshots/Screenshot_17.png)
Untuk menyelesaikan conflict tersebut bisa untuk me-resolve atau menyelesaikan baris code yang terjadi conflict, bisa untuk memilih salah satu perubahan kode atau memilih keduanya.

- Commit perubahan
![Gambar 3](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/3_Version%20Control%20and%20Branch%20Management%20(Git)/screenshots/Screenshot_18.png)
Setelah me-resolve conflict, maka bisa melakukan commit perubahan.

- Push commit
![Gambar 4](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/3_Version%20Control%20and%20Branch%20Management%20(Git)/screenshots/Screenshot_19.png)
Push dan disimpan commit pada remote repository.

- Hasil Akhir
![Gambar 5](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/3_Version%20Control%20and%20Branch%20Management%20(Git)/screenshots/Screenshot_20.png)
Commit perubahan berhasil dipush ke remote repository.
# Data Structure

#### Arvin Paundra Ardana - Golang A

### Problem 1 - Array Merge

Link [Problem 1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/6_Data%20Structure/praktikum/Problem%201/main.go)

Pada problem satu, menggabungkan dua array dan mehilangkan duplikasi nama jika terdapat pada kedua array. Ide pada proses ini yaitu dengan melakukan pengecekan atau perbandingan pada setiap indeks pada array. Untuk mengeceknya yaitu menggunakan dua buah perulangan dimana perulangan yang pertama bertujuan sebagai array yang dicek, sedangkan perulangan kedua berfungsi sebagai array pengecek. Jika pada array yang dicek terdapat nama yang sama dengan array pengecek, maka program akan berhenti dan langsung ke proses selanjutnya. Sedangkan jika tidak sama maka masuk ke dalam sebuah slice string atau variabel penampung. Setelah semua pengecekan selesai atau proses kedua perulangan telah selesai, maka variabel penampung tersebut akan digabungkan (append) dengan slice arrayA, lalu di-return kan.

Screenshot :

![Problem 1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/6_Data%20Structure/screenshots/Screenshot_16.png)

### Problem 2 - Angka Muncul Sekali

Link [Problem 2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/6_Data%20Structure/praktikum/Problem%202/main.go)

Pada problem dua, mencari karakter yang unik atau tidak ada yang double pada sebuah string. Ide pada proses ini yaitu pada setiap karakter dilakukan pengecekan terhadap karakter lain, apakah ada nilai yang sama pada string tersebut. Agar tidak dibaca sama pada waktu pengecekan, pada perulangan kedua dicek bila pada perulangan pertama dengan perulangan kedua berada di indeks yang sama maka program akan tetap dilanjutkan sampai menemukan karakter yang sama jika ada atau sampai akhir perulangan. Karena pada return nya harus sebuah slice integer, maka jika terdapat karakter yang unik sebelum digabungkan (append) perlu untuk diubah kebentuk integer terlebih dahulu menggunakan bantuan package `strconv`, kemudian digabungkan ke dalam sebuah variabel penampung. Setelah proses pengecekan atau kedua perulangan selesai, maka akan me-return variabel penampung atau sebuah slice integer yang berisi karakter unik dari string yang diberikan.

![Problem 2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/6_Data%20Structure/screenshots/Screenshot_17.png)

### Problem 3 - Pair with Target Sum

Link [Problem 3](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/6_Data%20Structure/praktikum/Problem%203/main.go)

Pada problem tiga, Menjumlahkan dua buah bilangan yang ada pada array sesuai dengan besar angka target, dengan catatan yang dikembalikan merupakan indeks dari kedua bilangan yang dijumlahkan dan penjumlahan pertama kali saja yang sesuai dengan nilai target. Ide pada proses ini yaitu membuat dua buah perulangan yang mana setiap perulangan itu dicek pada perulangan kedua, apakah jika dijumlahkan akan sama dengan besar nilai target. Jika hasilnya sama maka akan dimasukkan ke slice integer atau variabel penampung yang mana berisi kedua indeks dari nilai yang ketika dijumlahkan sama dengan nilai target untuk pertama kali, lalu proses akan break atau berhenti dan me-return nilai indeks yang ada pada variabel penampung.

![Problem 3](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/6_Data%20Structure/screenshots/Screenshot_18.png)

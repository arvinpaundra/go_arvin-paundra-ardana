# Time Complexity and Space Complexity

#### Arvin Paundra Ardana - Golang A

### Problem 1 - Bilangan Prima

Link [Problem 1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/7_Time%20Complexity%20%26%20Space%20Complexity/praktikum/Problem%201/main.go)

Pada problem pertama, melakukan optimasi pengecekan apakah sebuah bilangan merupakan bilangan prima atau bukan dengan time complexity-nya lebih cepat dari O(n) atau O(n/2).

Untuk proses yang awal yaitu melakukan pengecekan apakah apakah nilai input sama dengan satu, jika iya maka akan mengembalikan nilai false. Selanjutnya, melakukan pengecekan apakah nilai input bernilai 2 atau 3 atau 5, jika iya maka akan mengembalikan nilai true karena bilangan tersebut prima. Jika tidak memenuhi pengecekan sebelumnya, maka akan dilakukan sebuah perulangan dimana dimulai dari besar input, dengan kondisi input lebih besar dari satu, kemudian decrement sebesar input dibagi dua. Di dalam perulangan terdapat sebuah pengecekan, dimana jika input tersebut dapat dibagi habis oleh 2, 3, atau 5 maka akan mengembalikan nilai false karena bilangan tersebut dapat dibagi oleh bilangan selain satu dan bilangan itu sendiri atau berarti bukan bilangan prima. Jika sampai selesai perulangan tidak memenuhi kondisi pengecekan, maka dapat dipastikan bahwa input tersebut merupakan bilangan prima dan akan mengembalikan nilai true.

Screenshot :

![Problem 1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/7_Time%20Complexity%20%26%20Space%20Complexity/screenshots/Screenshot_20.png)

### Problem 2 - Fast Exponentiation

Link [Problem 2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/7_Time%20Complexity%20%26%20Space%20Complexity/praktikum/Problem%201/main.go)

Pada problem kedua, melakukan optimasi menentukan nilai hasil pangkat atau eksponensial dengan time complexity-nya lebih cepat dari O(n).

Pada proses ini dapat diselesaikan dengan menggunakan fungsi rekursif. Untuk awal prosesnya yaitu melakukan pengecekan apakah nilai input sama dengan 0, jika iya maka akan mengembalikan nilai satu karena berapapun angka yang dipangkatkan dengan satu maka akan menghasilkan nilai satu. Selanjutnya membuat sebuah variabel yang mana menampung nilai dari fungsi rekursif dengan parameter yang kedua dibagi dengan 2. Selanjutnya lakukan pengecekan jika pangkatnya positif maka akan mengalikan kembalian dari fungsi rekursif dengan kembalian dari fungsi rekursif. Jika pangkat negatif maka akan mengalikan kembalian dari fungsi rekursif akan dikalikan dengan kembalian dari fungsi rekursif dikalikan lagi dengan nilai input. Setelah itu akan mengembalikan nilai tersebut.

Screenshot :

![Problem 2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/7_Time%20Complexity%20%26%20Space%20Complexity/screenshots/Screenshot_21.png)

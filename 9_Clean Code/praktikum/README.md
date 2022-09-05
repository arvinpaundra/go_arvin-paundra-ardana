# Clean Code

#### Arvin Paundra Ardana - Golang A

### Problem 1 - Analysis

Link [Problem 1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/9_Clean%20Code/praktikum/Problem%201/main.go)

Problem 1, menganalisis kode before yang dituliskan tanpa mengikuti kebiasaan-kebiasaan penulisan yang diasarankan.

#### Pertanyaan 1

Berapa banyak kekurangan dalam penulisan kode tersebut?

Jawab :

4 (empat)

---

#### Pertanyaan 2

Bagian mana saja terjadi kekurangan tersebut?

Jawab :

- Penulisan pada saat deklarasi struct `user` dan `userservice`.

- Penulisan field pada struct `user` dan `userservice`.

- Penulisan method (struct function) `getallusers` dan `getuserbyid`.

- Penamaan saat deklarasi field `t` pada struct userservice dan variabel `r` pada loop.

---

#### Pertanyaan 3

Tuliskan alasan dari tiap kekurangan tersebut!

Jawab :

- Penulisan pada saat deklarasi struct sebaiknya menggunakan `PascalCase`, karena struct pada golang direpresentasikan sama seperti class.

- Penulisan field pada struct golang untuk best practice nya menggunakan `PascalCase`.

- Penulisan untuk nama method dengan lowercase bukan pendekatan yang baik, terutama pada clean code jika terdapat lebih dari 1 kata. Pada golang sendiri terdapat sebuah access modifier, jika ingin digunakan pada file lain (exported) maka awali dengan huruf kapital atau gunakan `PascalCase`. Namun, jika tidak ingin dapat digunakan pada file lain (unexported) maka hindari penggunaan huruf kapital pada saat awal deklarasi atau gunakan `camelCase`.

- Penamaan pada field `t` dan variabel `r` tidak mudah untuk dipahami dan tidak mendeskripsikan konteks. Maka dapat diganti dengan `Users` untuk `t` dan `user` untuk `r`.

Screenshot :

#### Before :

![Problem 1 Before](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/9_Clean%20Code/screenshots/Screenshot_32.png)

#### After :

![Problem 1 After](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/9_Clean%20Code/screenshots/Screenshot_33.png)

### Problem 2 - Rewrite

Link [Problem 2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/9_Clean%20Code/praktikum/Problem%202/main.go)

Problem 2, mengubah penulisan kode before menggunakan pendekatan clean code.

Yang perlu diubah pada penulisan kode before yaitu, penamaan nama struct yang masih menggunakan lowercase. Kemudian penulisan nama field pada struct dan method diganti menjadi `PascalCase`. Lalu penulisan parameter pada method tambahkecepatan dan variabel pada fungsi main diubah menjadi `camelCase`.

Screenshots :

#### Before

![Problem 2 Before](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/9_Clean%20Code/screenshots/Screenshot_34.png)

#### After

![Problem 2 After](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/9_Clean%20Code/screenshots/Screenshot_35.png)

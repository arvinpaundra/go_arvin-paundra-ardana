# (7) Time Complexity and Space Complexity

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Time Complexity and Space Complexity

#### 1. Time Complexity

Time Complexity merupakan sebuah kompleksitas perhitungan yang menunjukkan total waktu seberapa lama untuk menjalankan sebuah algoritma.

Macam-macam time complexity :

- O(1) - Constant time

Constant time berarti banyaknya sebuah inputan tidak mempengaruhi waktu proses dari algoritma tersebut.

- O(n) - Linear time

Linear time adalah ketika proses runtime dari fungsi berbanding lurus dengan jumlah input yang diberikan.

- O(log n) - Logarithmic time

Logarithmic time yaitu jika memberikan inputan sebesar n terhadap sebuah fungsi, jumlah tahapan yang dilakukan oleh fungsi tersebut berkurang berdasarkan suatu faktor.

- O(n^2) - Quadratic time

Quadratic time adalah ketika proses runtime dari sebuah fungsi adalah sebesar n^2, dimana n adalah sebuah inputan dari fungsi tersebut.

#### 2. Time Limit

Untuk komputer pada masa sekarang ini bisa mengerjakan 10^8 operasi dengan kurang dari satu detik. Batasan waktu yang diset pada pengujian online biasanya satu sampai sepuluh detik.

- Pada input `n <= 1.000.000`, perkiraan time complexity nya adalah O(n) atau O(n log n)
- Pada input `n <= 10.000`, perkiraan time complexity nya adalah O(n^2)
- Pada input `n <= 500`, perkiraan time complexity nya adalah O(n^3)

#### 3. Space Complexity

Space Complexity adalah jumlah memori yang digunakan oleh sebuah algoritma atau program.

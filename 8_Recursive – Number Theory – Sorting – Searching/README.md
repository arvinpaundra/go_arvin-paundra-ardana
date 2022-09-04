# (8) Recursive - Number Theory - Sorting - Searching

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Recursive - Number Theory - Sorting - Searching

#### 1. Recursive Function

Recursive atau rekursif merupakan sebuah fungsi yang memanggil dirinya sendiri untuk dieksekusi atau mengulang baris kode yang ada pada fungsi tersebut dengan cara memanggil dirinya sendiri yang bertujuan untuk mengurangi kompleksitas kode dan meningkatkan keterbacaan kode (readability).

Pada fungsi rekursif terdapat beberapa hal yang perlu diperhatikan, antara lain seperti :

#### a. Batas fungsi rekursif

Tidak seperti perulangan menggunakan loop yang bisa tak terbatas (endless loop), rekursif memiliki sebuah batasan terhadap pemanggilannya. Batas itu sendiri sampai adanya sebuah error warning dengan pesannya `stack overflow` yang menandakan sebuah batas dari fungsi rekursif.

#### b. Break point pada rekursif

Fungsi rekursif harus memiliki sebuah break point atau titik penghenti agar fungsi tersebut memiliki sebuah batasan harus sampai kapan fungsi rekursif tersebut harus dieksekusi. Jika tidak memberikan break point maka secara terus menerus fungsi rekursif tersebut akan selalu dieksekusi dan akan mengakibatkan error pada eksekusi.

#### c. Harus memiliki return value

Fungsi rekursif harus memiliki sebuah return value yang mana merupakan salah satu cara break point pada rekursif. Fungsi dari return value lainnya yaitu supaya kembalian dari fungsi rekursif bisa dieksekusi kembali secara berulang.

Pada penggunaannya fungsi rekursif digunakan pada contoh kasus seperti mencari nilai faktorial dari sebuah bilangan atau mencari deret fibonacci.

#### 2. Number Theory

Number theory merupakan cabang dari ilmu matematika yang memperlajari tentang integer (bilangan bulat). Contoh dari penerapakan number theory yaitu KPK, FPB, faktorial, faktor prima, dll.

#### 3. Sorting dan Searching

#### a. Sorting

Sorting merupakan sebuah proses dari mengurutkan data pada urutan tertentu. Pada implementasinya, sorting biasa digunakan untuk mengurutkan sebuah angka, huruf, dan data lainnya.

Contoh algoritma yang digunakan untuk melakukan sorting :

- Selection Sort - O(n^2)

Konsep dari selection sort yaitu mencari nilai terkecil dan menukarnya dengan elemen pertama pada sebuah array. Selanjutnya, sama hanya tinggal melanjutkan proses sebelumnya dengan tidak mengikutkan elemen pertama pada array karena sudah ditukar pada proses sebelumnya. Time complexity pada selection sort yaitu O(n^2), yang mana tidak efisien jika harus mengurutkan data dengan jumlah elemen yang besar.

- Counting Sort - O(n + k)

Konsep dari counting sort yaitu yang pertama mencari jumlah dari array tersebut, lalu yang kedua tinggal melakukan iterasi pada array tersebut dengan urutan yang di increment. Time complexity pada counting sort yaitu O(n + k).

- Merge Sort - O(log n)

Konsep merge sort yaitu membagi array menjadi dua bagian dan secara terus menerus dibagi sampai menjadi sebuah individual elemen. Kemudian elemen tersebut digabungkan kembali dengan elemen pertama yang terkecil. Time complexity pada merge sort yaitu O(log n).

#### b. Searching

Searching merupakan sebuah proses mencari letak nilai yang yang ada pada sebuah kumpulan nilai.

Contoh algoritma yang digunakan untuk melakukan searching :

- Linear Search

Linear search bekerja dengan melakukan perulangan melalui array sampai query ditemukan, yang mana membentuk sebuah algoritma linear. Time complexity pada linear search yaitu O(n).

- Binary Search - O(log n)

Binary search bekerja dengan mengurangi separuh setiap ruang pencarian di setiap iterasi setelah membandingkan nilai target ke nilai tengah dari ruang pencarian. Time complexity pada binary search yaitu O(log n).

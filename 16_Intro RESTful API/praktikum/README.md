# Intro RESTful API

#### Arvin Paundra Ardana - Golang A

## Part 1

Membuat satu collection dan memanfaatkan fitur environment dan melakukan request terhadap API yang ada di API Documentation menggunakan Postman dari tiga target API berikut.

- Sebelumnya menyiapkan environment variable untuk menyimpan tiga Base URL API dibawah.

![setup environment](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_23.png)

#### a. Target API 1 - News API

Berikut 5 request API dengan News API :

- Menyiapkan environment untuk menyimpan apiKey sebagai autentikasi untuk mengakses API dari News API dengan set type-nya ke API Key dan isi key value-nya dan tambahkan ke query params.

![img1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_6.png)

- Mencari berita tentang Liga Premier.

endpoint : **/everything**

q (query params) : **premier+league**

![img2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_7.png)

- Mencari berita tentang Cuaca.

endpoint : **/everything**

q : **weather**

![img3](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_8.png)

- Mencari berita tentang Bahasa Pemrograman Golang.

endpoint : **/everything**

q : **golang+programming**

![img4](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_9.png)

- Mencari berita headlines di Negara Indonesia dengan limit data yang ditampilkan 10 data.

endpoint : **/top-headlines**

country (query params) : **id**

pageSize (query params) : **10**

![img5](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_10.png)

- Mencari penyedia berita yang ada di seluruh dunia.

endpoint : **/top-headlines/sources**

![img6](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_11.png)

#### b. Target API 2 - Star Wars API

Berikut 5 request API dengan SWAPI :

- Menampilkan seluruh karakter yang ada di Star Wars.

endpoint : **/people**

![img7](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_12.png)

- Menampilkan detail karakter berdasarkan id karakter.

endpoint : **/people**

/:id (req params) : **4**

![img8](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_13.png)

- Menampilkan Data Kendaraan berdasarkan nama yang dimasukkan.

endpoint : **/vehicles**

name (query params) : **sand+crawler**

![img9](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_14.png)

- Menampilkan semua data film Star Wars.

endpoint : **/films**

![img10](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_15.png)

- Menampilkan detail karakter berdasarkan id film.

endpoint : **/films**

/:id : **4**

![img11](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_16.png)

#### c. Target API 3 - Rent a Book API

Berikut menerapkan 4 HTTP Method dengan Rent a Book API :

- Mendapatkan token untuk autentikasi ketika menggunakan API ini dengan menerapkan method GET.

endpoint : **/token**

method : **GET**

![img12](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_17.png)

- Agar dapat digunakan pada seluruh collection yang menggunakan API ini, maka simpan token dengan type Bearer Token dan pastekan token yang telah di GET tadi.

![img13](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_18.png)

- Menambahkan data user dengan method POST.

endpoint : **/user**

method : **POST**

req body :

```json
{
    "id": int,
    "name": string,
    "age": int,
    "sex": string,
    "client_id": int
}
```

![img14](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_19.png)

- Menampilkan detail user berdasarkan id user.

endpoint : **/user**

/:id : **9876**

method : **GET**

![img15](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_20.png)

- Mengubah data user berdasarkan id dengan menggunakan method PUT.

endpoint : **/user**

/:id : **9876**

method : **PUT**

req body :

```json
{
  "id": 9876,
  "name": "Ardana Arvin Paundra",
  "age": 18,
  "sex": "male",
  "client_id": 1
}
```

![img16](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_21.png)

- Menghapus data user berdasarkan id user dengan menggunakan method DELETE.

endpoint : **/user**

/:id : **9876**

method : **DELETE**

![img17](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/16_Intro%20RESTful%20API/screenshots/Screenshot_22.png)

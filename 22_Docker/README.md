# (22) Docker

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Docker

### 1. Gambaran Tentang Docker

Docker merupakan sebuah open platform untuk developing, shipping, dan running aplikasi. Docker memungkinkan kita untuk memisahkan aplikasi dengan infrastruktur, jadi kita bisa mendeliver software secara cepat. Dengan docker, kita bisa mengelola infrastruktur dengan cara yang sama dengan mengelola aplikasi yang sedang kita kembangkan. Dengan mengambil keuntungan dari metodologi dari docker untuk shipping, testing, dan deploy kode secara cepat, kita bisa secara signifikan mengurangi delay waktu antara menulis kode dengan menjalankannya di production.

### 2. Docker Volume

Docker volume bisa dianggap sebagai storage ata tempat penyimpanan data di container. Karena ketika container mati, secara default data yang ada pada container tersebuta akan ikut terhapus. Untuk itu kita perlu memanfaatkan volume pada docker. Berikut merupakan perintah yang sering digunakan.

- Menampilkan list volume

```
> docker volume ls
```

- Membuat volume

```
> docker volume create <nama_volume>
```

- Menghapus volume

```
> docker volume rm <nama_volume>
```

### 3. Docker Network

Defaultnya container pada docker akan saling terisolasi satu sama lain. Misalnya, kita tidak dapat melakukan request API dari container satu ke container lain. Untuk itu kita harus membuat dan mendaftarkan container pada network yang sama. Berikut merupakan perintah yang sering digunakan.

- Menampilkan list network

```
> docker network ls
```

- Membuat network baru

```
> docker network create <nama_network>
```

- Menghapus network

```
> docker network rm <nama_network>
```

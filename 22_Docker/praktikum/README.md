# Docker

#### Arvin Paundra Ardana - Golang A

### Dockerize Application

- Membuat Dockerfile

Membuat Dockerfile yang nantinya akan digunakan untuk konfigurasi image yang akan di build.

![img1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/22_Docker/screenshots/Screenshot_36.png)

- Build source code menjadi sebuah image

Melakukan build source code menjadi sebuah image berdasarkan konfigurasi Dockerfile dengan perintah berikut.

```
> docker build -t arvinpaundra/clean-arch:1.2 .
```

Kemudian untuk menampilkan hasil build image dapat menggunakan perintah berikut.

```
> docker image ls
```

![img2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/22_Docker/screenshots/Screenshot_37.png)

- Push ke image registry

Melakukan unggah image yang telah dibuat ke image registry atau image repository yaitu docker hub dengan perintah berikut.

```
> docker push arvinpaundra/clean-arch:1.2
```

![img3](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/22_Docker/screenshots/Screenshot_38.png)

- Running docker-compose

Menjalankan multi containers atau services dapat menggunakan docker compose. Untuk menjalankan multi containers yaitu go app dan mysql dapat dengan perintah berikut.

```
> docker compose start db_mysql rest_api
```

Sebelum memulai docker compose, dapat melakukan build berdasarkan docker-compose dengan perintah berikut.

```
> docker compose up --build
```

![img4](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/22_Docker/screenshots/Screenshot_39.png)

Hasil ketika menjalankan dengan docker compose go app dan mysql.

![img5](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/22_Docker/screenshots/Screenshot_40.png)

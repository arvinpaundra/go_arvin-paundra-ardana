# CI/CD

#### Arvin Paundra Ardana - Golang A

Link [**Repository**](https://github.com/arvinpaundra/learn-clean-architecture).

## 1. Praktikum Continuous Integration

Membuat pipeline CI/CD untuk memudahkan proses integrasi dan deployment berdasarkan projek sebelumnya, yaitu Clean Architecture.

### a. Konfigurasi file ci.yml

Konfigurasi saya ini untuk menjalankan proses CI secara otomatis yaitu dengan berdasarkan event push pada branch master. Pada proses CI ini saya bagi menjadi dua job, yaitu pertama untuk testing unit testing dan linter, dan kedua untuk build golang menjadi binary.

![img1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/24_CI%20CD/screenshots/Screenshot_36.png)

### b. Commit dan Push

Melakukan commit pada perubahan lalu melakukan push agar proses CI akan ter-trigger.

![img2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/24_CI%20CD/screenshots/Screenshot_37.png)

### c. Continuous Integration berhasil

Untuk proses CI telah berhasil, dimana disitu melakukan job testing seperti unit test dan lint atau code smell, serta build aplikasi menjadi binary.

![img3](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/24_CI%20CD/screenshots/Screenshot_38.png)

## 2. Praktikum Continuous Deployment

### a. Konfigurasi file cd.yml

Sama seperti pada CI, untuk proses CD saya samakan yaitu proses CD akan ter-trigger ketika ada sebuah push pada branch master. Pada proses CD ini saya juga melakukan setup untuk melakukan push docker image ke image repository yaitu docker hub dan kemudian secara otomatis juga akan masuk ke instance EC2 melalui SSH, dimana nanti akan melakukan hal-hal seperti pull image dengant tag `latest` dan melakukan build ulang container berdasarkan image terbaru.

![img4](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/24_CI%20CD/screenshots/Screenshot_40.png)

### b. Setup Github Secrets

Karena untuk proses CD ini terdapat akses seperti login ke docker hub yang mana membutuhkan username dan password serta key SSH untuk mengakses instance, maka diperlukan github secrets untuk menyembunyikan data penting tersebut.

![img7](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/24_CI%20CD/screenshots/Screenshot_45.png)

### c. Commit dan Push

Setelah terdapat event push, maka secara otomatis github actions akan langsung menjalankan dua workflows, yaitu CI dan setelah itu CD.

- Proses kedua workflows sedang berjalan

![img5](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/24_CI%20CD/screenshots/Screenshot_41.png)

- Proses detail CD atau automation deployment ke instance EC2 milik saya

![img6](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/24_CI%20CD/screenshots/Screenshot_42.png)

### d. Image berhasil dipush

Image berdasarkan code yang saya push telah berhasil ter-push pada image repository yaitu docker hub melalui proses CD.

- Sebelum

![img8](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/24_CI%20CD/screenshots/Screenshot_39.png)

- Sesudah

Terdapat image baru dengan tag `latest`.

![img9](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/24_CI%20CD/screenshots/Screenshot_43.png)

### e. Menguji keberhasilan proses CD

Berikut ini hasil dari proses CD yang mana saya uji melalui `curl` untuk mengakses instance EC2 saya melalui DNS publik.

![img10](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/24_CI%20CD/screenshots/Screenshot_44.png)

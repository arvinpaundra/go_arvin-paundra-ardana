# Compute Service

#### Arvin Paundra Ardana - Golang A

### Deployment EC2

- Masuk ke service EC2 milik aws, lalu pilih **Launch Instance** untuk membuat instance ec2 baru. Masukkan nama instance yang diinginkan.

![img1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/23_Compute%20Services/screenshots/Screenshot_36.png)

- Pada **Amazom Machine Image** (AMI), pilih OS yang ingin digunakan. Untuk sekarang saya memilih Ubuntu.

![img2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/23_Compute%20Services/screenshots/Screenshot_37.png)

- Untuk **Instace Type** saya memilih t2.micro yang free.

![img3](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/23_Compute%20Services/screenshots/Screenshot_38.png)

- Pada **Key Pair** wajib untuk diisi, karena digunakan untuk melakukan remoting VM ec2 via SSH atau Putty. Jika belum ada, bisa membuat terlebih dahulu. Disini saya memilih untuk remote access menggunakan SSH.

![img4](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/23_Compute%20Services/screenshots/Screenshot_39.png)

- Untuk **Security Group** saya menggunakan konfigurasi yang sebelumnya pernah saya buat.

![img5](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/23_Compute%20Services/screenshots/Screenshot_40.png)

- Setelah itu sisanya saya biarkan secara default. Lalu **Launch Instance**. Tampilan jika instance ec2 sudah berhasil dibuat.

![img6](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/23_Compute%20Services/screenshots/Screenshot_41.png)

- Kemudian **Connect** atau lakukan remote access antara device kita dengan VM ec2 melalui SSH dengan panduan yang sudah tertera. Jika menggunakan SSH maka wajib melampirkan file **.pem** sebagai Key nya. Berikut jika sudah berhasil mengakses VM tersebut.

![img7](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/23_Compute%20Services/screenshots/Screenshot_42.png)

### Deployment RDS

- Masuk ke service RDS, lalu pilih **Create Database**. Pilih database yang diinginkan. Disini saya menggunakan MySQL.

![img8](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/23_Compute%20Services/screenshots/Screenshot_43.png)

- Pada **Templates** saya memilih Free Tier.

![img9](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/23_Compute%20Services/screenshots/Screenshot_44.png)

- Pada bagian **Settings**, isi nama db cluster, username, password masternya.

![img10](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/23_Compute%20Services/screenshots/Screenshot_45.png)

- Setelah itu saya membiarkan sisanya terkonfigurasi secara default. Jika sudah klik **Create Database**. Jika sudah selesai dibuat, maka dapat melakukan akses database tersebut melalui CLI dan dapat juga melakukan migrasi database.

![img11](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/23_Compute%20Services/screenshots/Screenshot_46.png)

### Pointing Domain

Link [service](http://ec2-18-142-242-25.ap-southeast-1.compute.amazonaws.com/)

- Masuk ke VM ec2 via SSH. Lalu setup untuk install docker dengan perintah berikut.

```
> sudo apt-get install docker.io
```

- Jika sudah terinstall, maka pull image dari registry dengan perintah berikut.

```
> sudo docker pull arvinpaundra/clean-arch:1.2
```

- Cek apakah image sudah terdownload dengan perintah berikut.

```
> sudo docker image ls
```

![img12](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/23_Compute%20Services/screenshots/Screenshot_47.png)

- Setelah pull sukses, saya membuat file .env yang berisi konfigurasi untuk aplikasi golang nanti.

![img13](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/23_Compute%20Services/screenshots/Screenshot_48.png)

- Kemudian setelah membuat file untuk environment, build image yang baru saja terdownload menjadi sebuah container dengan perintah berikut.

```
> sudo docker run --name go-app -p 80:8080 -v $PWD/.env:/.env arvinpaundra/clean-arch:1.2
```

- Setelah build container selesai, cek apakah container tersebut berjalan dengan perintah berikut.

```
> sudo docker ps
```

![img14](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/23_Compute%20Services/screenshots/Screenshot_49.png)

- Cek lagi apakah sudah dapat diakses via DNS milik instance ec2 yang telah dibuat sebelumnya.

![img15](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/23_Compute%20Services/screenshots/Screenshot_50.png)

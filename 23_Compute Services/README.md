# (23) Compute Service

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Compute Service

### 1. Definisi Deployment

Deployment adalah kegiatan yang bertujuan untuk menyebarkan aplikasi/produk yang telah dikerjakan oleh para pengembang seringkali untuk mengubah dari status sementara ke permanen. Penyebarannya dapat melalui beragam cara tergantung dari jenis aplikasinya, aplikasi web dan API ke server sedangkan aplikasi mobile ke Playstore/Appstore.

### 2. Strategi Deployment

Berikut merupakan strategi dalam melakukan deplyment.

a. Big-Bang/Replacement Deployment Strategy

Kelebihan :

- Mudah diimplementasikan. Cara klasik yaitu tinggal replace projek.

- Perubahan kepada sistem langsung 100% secara instan.

Kekurangan :

- Terlalu beresiko, rata-rata downtime cukup lama.

b. Rollout Deployment Strategy

Kelebihan :

- Lebih aman dan less downtime dari versi sebelumnya.

Kekurangan :

- Akan ada 2 versi aplikasi yang berjalan secara bersamaan sampai semua server terdeploy dan bisa membuat bingung.

- Karena sifatnya perlahan satu persatu, untuk deployment dan rollback lebih lama dari yang Bigbang, karena prosesnya perlahan-lahan sampai semua server terkena efeknya.

- Tidak ada kontrol request. Server yang baru terdeploy dengan aplikasi versi baru langsung mendapat request yang sama banyaknya dengan server yang lain.

c. Blue/Green Deployment Strategy

Kelebihan :

- Perubahan sangat cepat, sekali switch service langsug berubah 100%.

- Tidak ada issue beda versi pada service seperti yang terjadi pada Rollout Deployment.

Kekurangan :

- Resource yang dibutuhkan lebih banyak. Karena untuk setiap deployment kita harus menyediakan service yang serupa environmentnya dengan yang sedang berjalan di production.

- Testing harus benar-benar sangat diprioritaskan sebelum di switch, aplikasi harus kita pastikan aman dari request yang tiba-tiba banyak.

d. Canary Deployment Strategy

Kelebihan :

- Cukup aman

- Mudah untuk rollback jika terjadi error/bug, tanpa berimbas kesemua server.

Kekurangan :

- Untuk mencapai 100% cukup lama dibanding dengan Blue/Green Deployment.

### 3. Proses Deployment

Proses deployment dapat memiliki satu atau banyak langkah, langkah-langkah dapat berjalan secara berurutan atau paralel, selain berbagai langkah penerapan, dapat juga menyertakan langkah-langkah intervensi manual untuk keluar sebelum deployment, langkah-langkah pemberitahuan email untuk memberi tahu semua orang tentang proses, atau bahkan melewatkan langkah dalam keadaan yang berbeda.

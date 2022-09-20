# (15) System Design

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi System Design

### 1. System Design

#### a. Diagram design

Diagram merupakan sebuah informasi yang direpresentasikan dengan simbol-simbol menggunakan teknik visualisasi.

Software design yang sering digunakan :

- Flowchart atau diagram alir adalah diagram yang menggambarkan secara rinci langkah-langkah atau alur dari proses pemrograman atau algoritma.

- Use case diagram adalah diagram yang meringkaskan detail dari pengguna sistem atau yang dikenal sebagai aktor dan beserta interaksinya dengan sistem tersebut.

- Entity relationship diagram merupakan tipe dari flowchart yang mengilustrasikan bagaimana entitas-entitas, seperti orang, objek, atau konsep yang berhubungan satu dengan yang lainnya dalam sebuah sistem.

- High Level Architecture menjelaskan arsitektur yang akan digunakan untuk mengembangkan sebuah sistem (design sistem secara keseluruhan`) tanpa harus tahu detail implementasi nantinya seperti apa.

#### b. Karakteristik dari sistem terdistribusi

- Scalability

Scalability atau skalabilitas merupakan kemampuan dari sebuah sistem, proses, atau jaringan untuk berkembang dan mengatur permintaan yang meningkat. Semua sistem bisa untuk berkembang karena banyak alasannya, seperti perkembangan volume data atau bertambahnya jumlah pekerjaan, seperti jumlah transaksi.

- Reliability

Reliability merupakan kemungkinan sebuah sistem akan mengalami kegagalan untuk jangka waktu tertentu. Secara sederhana, sebuah sistem terdistribusi dianggap reliable atau dapat diandalkan apabila tetap mampu memberikan service atau layanan bahkan ketika satu atau beberapa hardware atau software-nya mengalami kegagalan saat sedang berjalan.

- Availability

Availability merupakan waktu sistem tetap beroperasi untuk menjalankan fungsi yang diperlukan pada waktu tertentu. Itu merupakan ukuran sederhana dari persentase waktu bahwa suatu sistem, service, atau mesin tetap beroperasional diluar kondisi normal.

- Efficiency

Efficiency merupakan cara bagaimana memaksimalkan kerja dari sebuah sistem agar dapat bekerja secara seefisien mungkin, namun dengan penggunaan resource atau sumber daya yang rendah.

- Serviceability or Manageability

Serviceability or Manageability merupakan kemudahan dan kecepatan dengan yang mana sebuah sistem bisa diperbaiki atau dimaintain atau dirawat. Jika waktu untuk memperbaiki kegagalan sebuah sistem terus meningkat, maka akan berimbas dengan menurunkannya availability atau keberadaan sistem.

#### c. Horizontal scaling dan Vertical scaling

- Horizontal scaling merupakan meningkatkan server ke arah samping atau mengingkatkan kuantitas jumlah server untuk meningkatkan ability atau kemampuan server, seperti contohnya untuk memproses jumlah request yang lebih banyak dibanding sebelumnya.

- Vertical scaling atau meningkatkan server ke atas yaitu meningkatkan resource dari server tersebut, misalnya meningkatkan kapasitas SSD, RAM, atau jumlah core CPU yang dimana cost atau biayanya lebih murah dibanding horizontal scaling.

#### d. Job/Work Queue

Job/Work Queue adalah struktur data yang dikelola oleh perangkat lunak job scheduler yang berisi pekerjaan untuk dijalankan. Pengguna mengirimkan program mereka yang ingin mereka jalankan, "job", ke antrian untuk pemrosesan batch. Scheduler mengurus queue atau antrian sebagai kumpulan dari job yang tersedia untuk dijalankan.

#### e. Load Balancing

Load Balancer merupakan salah satu komponen penting dalam sistem terdistribusi. Load Balancer membantu untuk menyebarkan atau membagi traffic antara cluster dari server untuk meningkatkan responsifitas dan ketersediaan dari aplikasi, website, atau database. Load Balancer juga melacak status dari semua resource ketika sedang mendistribusikan request. Salah satu kegunaan lain dari Load Balancer yaitu dapat mencegah serangan siber seperti DDoS.

### 2. Monolith vs Microservices

- Monolith atau Monolitik adalah aplikasi yang memiliki single code base dengan banyak module didalamnya. Module dibagi menjadi fitur bisnis atau fitur teknis. Monolitik memiliki sistem single build yang mana mem-build keseluruhan aplikasi beserta dependensinya. Monolitik juga hanya memiliki single executable atau deployable binary.

- Microservices merupakan service yang di-deploy secara independen yang telah dimodelkan sesuai dengan domain bisnisnya. Microservice berkomunikasi antar service-nya melalui network dan sebagai sebuah arsitektur pilihan yang menawarkan banyak opsi penyelesaian masalah yang dihadapi.

### 3. Database dan Optimasi

#### a. SQL dan NoSQL

- SQL atau relational database

SQL atau relasional database merupakan database yang terstruktur dan memiliki skema yang telah didefinisikan pada awal pembuatannya. Kelebihan dari SQL yaitu, dirancang untuk segala keperluan, memiliki standar yang jelas, terdapat banyak tools pembantu.

Kelebihan lainnya yaitu mandukung `ACID` :

`Atomicity` : Transaksi terjadi semua atau tidak sama sekali.

`Consistency` : Data tertulis merupakan data valid yang ditentukan berdasarkan aturan tertentu.

`Isolation` : Pada saat terjadi request yang terjadi secara bersamaan (concurrent), memastikan bahwa transaksi dieksekusi seperti dijalankan secara sekuensial.

`Durability` : Jaminan bahwa transaksi yang telah tersimpan itu tetap tersimpan.

- NoSQL atau non-relational database

Tidak seperti SQL, NoSQL tidak terstruktur dan memiliki skema yang dinamis. NoSQL merupakan DBMS yang menyediakan mekanisme yang lebih fleksibel dibandingkan dengan model RDBMS (sifat ACID). NoSQL cocok untuk digunakan ketika membutuhkan skema yang fleksibel, tidak memerlukan ACID, data logging, caching, dsb.

#### b. Caching

Cache digunakan untuk menyimpan data yang baru saja di request agar tidak melakukan request ulang ke tujuan yang sama untuk sementara waktu. Cache merupakan memori jangka pendek yang memiliki jumlah kapasitas yang terbatas namun lebih cepat daripada sumber data dan berisi items yang paling sering diakses.

#### c. Database replication

Redundansi merupakan komponen atau fungsi kritikal dari duplikasi dari duplikasi dari sebuah sistem dengan maksud meningkatkan reliability atau keandalan dari sistem, biasa pada bentuk dari sebuah backup atau untuk meningkatkan kinerja sistem yang sebenarnya.

#### d. Database Indexing

Indexing merupakan sebuah cara untuk mengoptimasi performa dari sebuah database dengan meminimalisir penggunaan akses memori yang diperlukan ketika sedang menjalankan sebuah query database. Indexig merupakan sebuah teknik struktur data yang mana digunakan untuk secara cepat melokasikan dan mengakses data pada sebuah database.

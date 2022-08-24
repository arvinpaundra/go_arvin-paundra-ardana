# (3) Version Control and Branch Management (Git)

#### Arvin Paundra Ardana - Golang A

## 3 Point Penting pada Materi Git

#### 1. Useful Tools

Git sangat membantu para developer untuk memanajemen pekerjaan projek mereka mulai dari siapa saja yang mengubah kode, kapan kode itu diubah, siapa saja yang berkontribusi, apa saja yang diperbarui pada kode tersebut, dan sebagainya.

Dengan git juga akan membuat sebuah kebiasaan baru untuk seorang developer agar selalu mendokumentasikan setiap perubahan yang telah mereka buat. Hal ini sangat penting mengingat jika yang dikerjakan merupakan sebuah projek yang skalanya besar yang akan menyulitkan proses pengembangan apabila tidak ada sebuah tools untuk melakukan pelacakan pada setiap perubahan kode.

Git akan sangat membantu ketika digunakan dengan berkolaborasi dengan orang banyak. Karena akan lebih terlihat bahwa git itu akan merekam dan menampilkan semua commit perubahan yang mereka buat sehingga akan lebih mudah untuk dipantau dan dilacak pada setiap perubahannya.

#### 2. Bagian-bagian pada Git

- Local repository
  Local repository merupakan sebuah komputer yang digunakan oleh developer dimana itu merupakan tempat melakukan sebuah pengubahan dan penyimpanan sebuah kode secara lokal (pada komputer masing-masing developer).
- Stagging Area
  Stagging area merupakan sebuah area dimana developer melakukan stagging atau menandai suatu kode yang telah ia ubah atau buat yang kemudian akan dilakukan sebuah commit pada kode tersebut yang nantinya akan didstribusikan ke remote repository.
- Remote Repository
  Remote repository merupakan sebuah komputer server yang digunakan untuk menyimpan seluruh kode-kode yang di commit oleh developer atau tempat central untuk menyimpan setiap perubahan kode.
- Stash Area
  Stash area merupakan sebuah penyimpanan sementara perubahan saat ini pada kode tanpa harus melakukan commit ke remote repository.

#### 3. Best Practices Penggunaan Git

Best practices pada git yaitu dapat dengan menggunakan git workflow untuk projek yang berskala besar.

- Git workflow hanya memperbolehkan branch development untuk melakukan checkout branch master agar branch master tidak terganggu dengan branch lain.
- Hindari untuk mengedit langsung pada branch development, lebih baik buat satu branch baru yang men-checkout branch development jika memang diperlukan.
- Jika ingin membuat sebuah fitur baru maka checkout dan merge pada branch development.
- Jika akan melakukan merge fitur baru ke branch master, hanya gunakan branch development pada saat melakukan merging tersebut.

Opsi lain yang dapat digunakan yaitu dapat menggunakan trunk-based development yang biasanya dipakai pada projek yang tidak terlalu besar. Opsi ini menggunakan satu buah branch untuk berkolaborasi, misalnya branch master dapat di checkout langsung oleh beberapa branch.

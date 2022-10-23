# Clean and Hexagonal Architecture

#### Arvin Paundra Ardana - Golang A

### Problem 1 - Rewrite Code

Problem pertama yaitu melakukan rewrite code yang awalnya menggunakan MVC, diubah menggunakan Clean Architecture.

- Layer Repository

Pada layer repository dibuat sebuah interface yang berisi fungsi GetAll, Create, Login yang nantinya digunakan sebagai komunikasi atau kontrak oleh layer usecase jika ingin mengimplementasikan layer repository tersebut.

![img1](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/21_Clean%20and%20Hexagonal%20Architecture/screenshots/Screenshot_36.png)

- Layer Usecase

Pada layer usecase dibuat sebuah interface yang berisi fungsi GetAll, Create, Login yang sama fungsinya sebagai komunikasi atau kontrak oleh layer controller jika ingin mengimplementasikan layer usecase ini.

![img2](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/21_Clean%20and%20Hexagonal%20Architecture/screenshots/Screenshot_37.png)

- Layer Controller

Pada layer controller mengimplementasikan semua method yang ada pada layer usecase milik user.

![img3](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/21_Clean%20and%20Hexagonal%20Architecture/screenshots/Screenshot_38.png)

### Problem 2 - Implementasi JWT

Implementasi JWT sebagai autentikasi yang dieksekusi melalui middleware pada saat melakukan definisi endpoint pada API. Di route, semua layer diinstasiasi dengan berdasarkan dependensinya.

![img4](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/21_Clean%20and%20Hexagonal%20Architecture/screenshots/Screenshot_39.png)

### Problem 3 - Unit Test

Melakukan unit testing pada layer controller.

![img5](https://github.com/arvinpaundra/go_arvin-paundra-ardana/blob/master/21_Clean%20and%20Hexagonal%20Architecture/screenshots/Screenshot_40.png)

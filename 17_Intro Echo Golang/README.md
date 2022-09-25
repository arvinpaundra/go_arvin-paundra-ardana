# (17) Intro Echo Golang

#### Arvin Paundra Ardana - Golang A

## 3 Poin Penting pada Materi Intro Echo Golang

### 1. Echo Golang

Echo merupakan sebuah micro-framework milik golang. Echo menawarkan performa yang tinggi, mudah untuk dipadupadankan dengan library lain (extensible), dan juga minimalis. Kemudian tambahan untuk kemudahan yang ditawarkan oleh echo yaitu, router yang teroptimisasi, data rendering, data binding, middleware, mudah dalam mengatur skalabilitasnya. Echo dikatakan sebuah minimalis framework karena pada framework ini tidak mengusung banyak library atau third party lain seperti contohnya driver untuk database yang mana berarti harus di get secara terpisah.

### 2. Routing dan Controller pada Echo

a. Routing

Handling routing pada echo sangat mudah, dimana yang pertama yaitu buat terlebih dahulu instance dari echo dengan cara :

```go
e := echo.New()
```

Setelah melakukan instasiasi, maka tinggal pilih saja http method yang akan digunakan. Contoh akan menampilkan data users, maka gunakan method `GET`, kemudian isi parameter pertama dengan path atau endpoint-nya, parameter kedua yaitu berupa fungsi handler nya atau controller yang akan melakukan pemrosesan request response, dan parameter ketiga yaitu middleware yang bersifat opsional.

```go
e.GET("/users", GetUsersController)
```

Jika pada path atau endpoint tersebut ingin mengirimkan sebuah parameter yang sifatnya dinamis, seperti data id, maka dapat ditulis seperti berikut.

```go
e.GET("/users/:id", GetUserController)
```

b. Controller

Setelah melakukan membuat sebuah route atau endpoint nya, maka langkah selanjutnya yaitu membuat sebuah controller yang dimana berisi terkait logic dari sistem yang akan dibangun. Fungsi controller pada echo sendiri memerlukan sebuah parameter yaitu `echo.Context` dan return value nya berupa interface error. Berikut contoh sederhana controller untuk menampilkan data users.

```go
...

func GetUsersController(c echo.Context) error {
    // return berupa data JSON
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success get all users",
        "users": users,
    })
}

...
```

### 3. Handling Data pada Echo

a. Data Rendering

Pada echo dapat mengembalikan berbagai macam jenis HTTP response, seperti JSON, XML, HTML, dan masih banyak yang lainnya. Contoh melakukan render data atau response dengan bentuk JSON.

```go
...

func GetProductController(c echo.Context) error {
    product := map[string]interface{} {
        "name": "Samsung J2 Prime",
        "category": "phone",
        "price": 320,
    }

    // return data product sebagai JSON
	return c.JSON(http.StatusOK, products)
}

...
```

b. Retrieve Data

- Request Params

Seperti yang telah disebutkan pada saat membuat endpoint, biasanya pada endpoint tersebut memerlukan sebuah data yang sifatnya dinamis, seperti contohnya yaitu menangkap data id yang dikirimkan oleh client. Contoh :

```go
...

func main() {
    ...

    e.GET("/products/:id", GetIdProduct)

    ...
}

func GetIdProduct(c echo.Context) error {
    // tangkap id menggunakan method Param milik echo.Context
    productId, _ := strconv.Atoi(c.Param(id))

    return c.JSON(http.StatusOK, map[string]interface{}{
        "product_id": productId,
    })
}

...
```

- Query Parameter

Sama seperti request params, query parameter juga merupakan data yang di passing lewat url endpoint, yang membedakan adalah query parameter biasanya bersifat opsional. Contoh :

```go
...

func main() {
    ...

    e.GET("/products", GetProductsController)

    ...
}

func GetProductsController(c echo.Context) error {
    // tangkap query parameter menggunakan method QueryParam milik echo.Context
    search := c.QueryParam("search")
    limit := c.QueryParam("limit")
    page := c.QueryParam("page")

    return c.JSON(http.StatusOK, map[string]interface{}{
        "search": search,
        "limit": limit,
        "page": page,
    })
}

...
```

- Form Value

Form Value digunakan untuk menangkap nilai sesuai antara nama field pada form dan struct melalui bentuk JSON. Contoh :

```go
...

type Product struct {
    Name        string  `json:"name" form:"name"`
    Category    string  `json:"category" form:"category"`
    Price       int     `json:"price" form:"price"`
}

func main() {
    ...

    e.POST("/products", CreateProductController)

    ...
}

func CreateProductController(c echo.Context) error {
    // tangkap value dari form
    name := c.FormValue("name")
    category := c.FormValue("category")
    price := c.FormValue("price")

    product := Product{
        Name: name,
        Category: category,
        Price: price,
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success create product",
        "product": product
    })
}

...
```

c. Binding Data

Data binding mekanismenya mirip seperti menggunakan FormValue, akan tetapi lebih mudah dan lebih singkat dalam penulisannya. Contoh :

```go
...

type Product struct {
    Name        string  `json:"name" form:"name"`
    Category    string  `json:"category" form:"category"`
    Price       int     `json:"price" form:"price"`
}

func main() {
    ...

    e.POST("/products", CreateProductController)

    ...
}

func CreateProductController(c echo.Context) error {
    // buat instance Product sebagai penampung payload
    product := Product{}
    c.Bind(&product)

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success create product",
        "product": product
    })
}

...
```

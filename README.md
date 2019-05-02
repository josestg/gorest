
# Pre-Test PRIVY : Full Rest API

## Quick Start
Semua Endpoint pada aplikasi ini diuji menggukan Aplikasi [Postman](https://www.getpostman.com/).

**1. Mengunduh project ke $GOPATH/src**
```bash
# mengunduh project ke GOPATH/src
go get github.com/josestnggng/Pretest-privy-full-rest-api

```

**catatan** :
abaikan jika ada pesan berikut
`josestnggng/Pretest-privy-full-rest-api/make_seed.go:5:6: main redeclared in this block
    previous declaration at josestnggng/Pretest-privy-full-rest-api/main.go:5:6
`

**2. Menjalankan Aplikasi**

#### Linux
```bash
# pindah ke folder GOPATH
cd $GOPATH/github.com/josestnggng/Pretest-privy-full-rest-api

# menbangkitkan table dari seed
make seed
atau
go make_seed

# Memulai aplikasi
make start
atau
go run main.go
```

#### Windows
```bash
# pindah ke folder GOPATH
cd %GOPATH%\src\github.com\josestnggng\Pretest-privy-full-rest-api

# menbangkitkan table dari seed
go run make_seed.go

# Memulai aplikasi
go run main.go
```

### Catatan
Karena untuk menanambah `image` juga meng-upload file ke server maka `body` pada Postman adalah `form-data` dan `header` sebagai berikut :
`key` : `Content-Type` dan `value` : `application/x-www-form-urlencoded`.

Contoh :

![source code.](https://github.com/josestnggng/Pretest-privy-full-rest-api/blob/master/uploads/ss.png)


# Endpoints
- [Products](#Products)
- [Categories](#Categories)
- [Images](#Images)
- [Category Products](#Category-Products)
- [Product Images](#Product-Images)

## Products

| Route | Description | params |
| --- | --- | --- |
| `GET  api/products` | Mengembalikan semua `product` | None |
| `GET  api/products/{id}` | Mengembalikan satu `product` sesuai parameter `id` | Product ID  |
| `POST api/products` | Memasukkan `product` baru ke database | None |
| `PUT api/products/{id}` | Meng-update `product` sesuai `id` | Product ID |
| `DELETE api/products/{id}` | Menghapus satu `product` sesuai `id` | Product ID |

<br>

Contoh Request :
```http
GET     api/products
```
Respond :
```json
[
    {
        "id": 1,
        "name": "Statistika",
        "description": "Buku pengantar statistika dasar",
        "enable": true
    },
    {
        "id": 2,
        "name": "Pemograman Golang",
        "description": "Buku pengantar pemograman",
        "enable": true
    }
]

```




## Categories

| Route | Description | params |
| --- | --- | --- |
| `GET  api/categories` | Mengembalikan semua `category` | None |
| `GET  api/categories/{id}` | Mengembalikan satu `category` sesuai parameter `id` | category ID  |
| `POST api/categories` | Memasukkan `category` baru ke database | None |
| `PUT api/categories/{id}` | Meng-update `category` sesuai `id` | category ID |
| `DELETE api/categories/{id}` | Menghapus satu `category` sesuai `id` | Product ID |

<br>

Contoh Request :
```http
GET     api/categories
```
Respond :
```json
[
    {
        "id": 1,
        "name": "Matematika",
        "enable": true
    },
    {
        "id": 2,
        "name": "Geofisika",
        "enable": true
    }
]

```

## Images

| Route | Description | params |
| --- | --- | --- |
| `GET  api/images` | Mengembalikan semua `image` | None |
| `GET  api/images/{id}` | Mengembalikan satu `image` sesuai parameter `id` | image ID  |
| `POST api/images` | Memasukkan `image` baru ke database dan meng-upload ke file `image` ke sever. | None |
| `PUT api/images/{id}` | Meng-update `image` sesuai `id` | image ID |
| `DELETE api/images/{id}` | Menghapus satu `image` sesuai `id` | Product ID |

<br>

Contoh Request :
```http
GET     api/images
```
Respond :
```json
[
    {
        "id": 1,
        "name": "Image 1",
        "file": "upload-121371.png",
        "enable": false
    },
    {
        "id": 2,
        "name": "Image 2",
        "file": "upload-121362.png",
        "enable": false
    }
]

```

## Category Products

| Route | Description | params |
| --- | --- | --- |
| `GET  api/category-products` | Mengembalikan semua `category-product` | None |
| `GET  api/category-products/{pid}/{cid}` | Mengembalikan satu `category-product` sesuai parameter `pid` dan `cid` | productID (pid), categoryID (cid)  |
| `POST api/category-products` | Memasukkan `category-product` baru ke database | None |
| `PUT api/category-products/{pid}/{cid}` | Meng-update `category-product` sesuai parameter`pid` dan `cid` | productID (pid), categoryID (cid) |
| `DELETE api/category-products/{pid}/{cid}` | Menghapus satu `category-product` sesuai parameter`pid` dan `cid` | productID (pid), categoryID (cid) |

<br>

Contoh Request :
```http
GET     api/category-products
```
Respond :
```json
[
    {
        "id": 1,
        "name": "Statistika",
        "description": "Buku pengantar statistika dasar",
        "enable": true,
        "categories": [
            {
                "id": 1,
                "name": "Matematika",
                "enable": true
            },
            {
                "id": 3,
                "name": "Informatika",
                "enable": true
            },
            {
                "id": 4,
                "name": "Sains",
                "enable": true
            }
        ]
    }
]

```


## Product Images

| Route | Description | params |
| --- | --- | --- |
| `GET  api/product-images` | Mengembalikan semua `product-images` | None |
| `GET  api/product-images/{pid}/{iid}` | Mengembalikan satu `product-images` sesuai parameter `pid` dan `iid` | productID (pid), imageID (iid)  |
| `POST api/product-images` | Memasukkan `product-images` baru ke database | None |
| `PUT api/product-images/{pid}/{iid}` | Meng-update `product-images` sesuai parameter`pid` dan `iid` | productID (pid), imageID (iid) |
| `DELETE api/product-images/{pid}/{iid}` | Menghapus satu `product-images` sesuai parameter`pid` dan `iid` | productID (pid), imageID (iid) |

<br>

Contoh Request :
```http
GET     api/product-images
```
Respond :
```json
[
    {
        "id": 1,
        "name": "Statistika",
        "description": "Buku pengantar statistika dasar",
        "enable": true,
        "images": [
            {
                "id": 1,
                "name": "Image 1",
                "file": "upload-121371.png",
                "enable": false
            },
            {
                "id": 4,
                "name": "Image 4",
                "file": "upload-121344.png",
                "enable": false
            }
        ]
    }
]

```

# Database Management

## Edu Management 1.1

### Implementation technique

Siswa akan melaksanakan sesi live code di 15 menit terakhir dari sesi mentoring dan di awasi secara langsung oleh Mentor. Dengan penjelasan sebagai berikut:

- **Durasi**: 15 menit pengerjaan
- **Submit**: Maximum 10 menit setelah sesi mentoring menggunakan `grader-cli submit`
- **Obligation**: Wajib melakukan _share screen_ di breakout room yang akan dibuatkan oleh Mentor pada saat mengerjakan Live Coding.

### Description

Edu Management API ini digunakan untuk mengakses data dari repositori `student` dan `teacher`. API ini memungkinkan pengguna untuk melakukan berbagai operasi **CR** _(Create, Read)_ pada data `student` dan **CRU** _(Create, Read, Update)_ `teacher`.

API ini memiliki dua endpoint yaitu: `/student`, dan `/teacher`.

- Endpoint `/student` digunakan untuk mengakses data student.
- Endpoint `/teacher` digunakan untuk mengakses data teacher.

Pada setiap endpoint, API ini memiliki beberapa _sub-endpoint_ yang berfungsi untuk melakukan operasi CRUD pada data. Sub-endpoint tersebut meliputi:

- `/student`
  - `/get-all`: untuk mengambil semua data
  - `/get`: untuk mengambil data dengan ID tertentu
  - `/add`: untuk menambahkan data baru

- `/teacher`
  - `/get-all`: untuk mengambil semua data
  - `/get`: untuk mengambil data dengan ID tertentu
  - `/add`: untuk menambahkan data baru
  - `/update`: untuk memperbarui data yang sudah ada

API ini menggunakan `http.ServeMux` sebagai router untuk mengarahkan permintaan ke fungsi yang sesuai. API ini juga mengambil dua parameter yaitu `studentRepo` dan `teacherRepo` yang masing-masing adalah instance dari `repo.StudentRepository` dan `repo.TeacherRepository`. Instance ini digunakan untuk mengakses data dari repositori.

API ini dapat dijalankan dengan memanggil fungsi `Start()`, yang akan menampilkan pesan di console bahwa server sedang berjalan dan menjalankan server pada <http://localhost:8080>.

### Constraints

Pada live code ini, kamu harus melengkapi fungsi dari repository `student` dan `teacher` ini memiliki implementasi function-function berikut:

- `repository/student.go`

  - `FetchAll`: Function ini akan mengambil semua data mahasiswa yang ada di dalam tabel `students` pada database. Selanjutnya, data mahasiswa tersebut akan di-scan dan dimasukkan ke dalam slice `[]model.Student`.
    - Jika proses tersebut berhasil, function akan mengembalikan slice tersebut beserta nilai `nil` sebagai `error`.
    - Namun jika terjadi error pada proses tersebut, function akan mengembalikan `nil` sebagai slice dan `error` yang terjadi.

  - `FetchByID`: Function ini akan mengambil data mahasiswa yang memiliki `id` yang sesuai dengan nilai yang diberikan sebagai argumen. Pertama-tama, function akan mengeksekusi sebuah query untuk mencari data mahasiswa dengan `id` yang sesuai. Hasil dari query tersebut akan di-scan ke dalam variabel `model.Student`.
    - Jika proses tersebut berhasil, function akan mengembalikan pointer `model.Student` beserta nilai `nil` sebagai `error`.
    - Namun jika terjadi error pada proses tersebut, function akan mengembalikan `nil` sebagai pointer dan `error` yang terjadi.

  - `Store`: Function ini akan menyimpan data mahasiswa yang diberikan sebagai argumen ke dalam database. Pertama-tama, function akan mengeksekusi sebuah query `INSERT` untuk memasukkan data mahasiswa baru ke dalam tabel `students`. Query tersebut akan menggunakan nilai dari variabel `model.Student` yang diberikan sebagai argumen.
    - Jika proses tersebut berhasil, function akan mengembalikan `nil` sebagai `error`.
    - Namun jika terjadi `error` pada proses tersebut, function akan mengembalikan `error` yang terjadi.

- `repository/teacher.go`
  - `FetchAll`: Function ini akan mengambil semua data guru yang ada di dalam tabel `teachers` pada database. Selanjutnya, data guru tersebut akan di-scan dan dimasukkan ke dalam slice `[]model.Teacher`.
    - Jika proses tersebut berhasil, function akan mengembalikan slice tersebut beserta nilai `nil` sebagai `error`.
    - Namun jika terjadi error pada proses tersebut, function akan mengembalikan `nil` sebagai slice dan `error` yang terjadi.

  - `FetchByID`: Function ini akan mengambil data guru yang memiliki `id` yang sesuai dengan nilai yang diberikan sebagai argumen. Pertama-tama, function akan mengeksekusi sebuah query untuk mencari data guru dengan `id` yang sesuai. Hasil dari query tersebut akan di-scan ke dalam variabel `model.Teacher`.
    - Jika proses tersebut berhasil, function akan mengembalikan pointer `model.Teacher` beserta nilai `nil` sebagai `error`.
    - Namun jika terjadi error pada proses tersebut, function akan mengembalikan `nil` sebagai pointer dan `error` yang terjadi.

  - `Store`: Function ini akan menyimpan data guru yang diberikan sebagai argumen ke dalam database. Pertama-tama, function akan mengeksekusi sebuah query `INSERT` untuk memasukkan data guru baru ke dalam tabel `teachers`. Query tersebut akan menggunakan nilai dari variabel `model.Teacher` yang diberikan sebagai argumen.
    - Jika proses tersebut berhasil, function akan mengembalikan `nil` sebagai `error`.
    - Namun jika terjadi `error` pada proses tersebut, function akan mengembalikan `error` yang terjadi.

  - `Update`: Function ini akan mengupdate data guru yang memiliki `id` yang sesuai dengan nilai yang diberikan sebagai argumen. Pertama-tama, function akan mengeksekusi sebuah query `UPDATE` untuk mengubah data guru dengan `id` yang sesuai. Query tersebut akan menggunakan nilai dari variabel `model.Teacher` dan `id` yang diberikan sebagai argumen.
    - Jika proses tersebut berhasil, function akan mengembalikan `nil` sebagai `error`.
    - Namun jika terjadi `error` pada proses tersebut, function akan mengembalikan `error` yang terjadi.

### **Perhatian**

Sebelum kalian menjalankan `grader-cli test`, pastikan kalian sudah mengubah database credentials pada file **`main.go`** (line 65) dan **`main_test.go`** (line 17) sesuai dengan database kalian. Kalian cukup mengubah nilai dari  `"username"`, `"password"` dan `"database_name"`saja.

Contoh:

```go
dbCredentials = Credential{
    Host:         "localhost",
    Username:     "postgres", // <- ubah ini
    Password:     "postgres", // <- ubah ini
    DatabaseName: "kampusmerdeka", // <- ubah ini
    Port:         5432,
}
```

### Test Case Examples

#### Test Case 1

**Input**:

Untuk function `Store`, input yang diberikan adalah pointer ke `model.Student` yang berisi data mahasiswa baru yang ingin disimpan.

**Expected Output / Behavior**:

Function akan mengembalikan `error`. `error` akan bernilai `nil` jika tidak terjadi kesalahan, sedangkan jika terjadi kesalahan dalam menjalankan perintah SQL untuk menyimpan data mahasiswa baru, function akan mengembalikan `error` yang sesuai.

**Explanation**:

1. Function `Store` dipanggil dengan pointer ke `model.Student` yang berisi data mahasiswa baru yang ingin disimpan.
2. Function akan mengeksekusi perintah SQL yang akan menyimpan data mahasiswa baru ke tabel `students`.
3. Function ini akan mengembalikan `error` jika terjadi kesalahan saat memasukan data dan mengembalikan `nil` jika insert berhasil.

#### Test Case 2

**Input**:

Untuk function `FetchAll`, tidak memerlukan input apa pun.

**Expected Output / Behavior**:

Function akan mengembalikan slice of `model.Student` dan `error`. `error` akan bernilai `nil` jika tidak terjadi kesalahan, sedangkan slice `model.Student` akan berisi data seluruh mahasiswa yang ada dalam database.

**Explanation**:

1. Function `FetchAll` dipanggil.
2. Function akan mengeksekusi perintah SQL yang akan mengambil data seluruh mahasiswa dari tabel `students`.
3. Jika terjadi kesalahan dalam menjalankan perintah SQL, function akan mengembalikan `nil` dan `error` yang sesuai.
4. Function akan melakukan perulangan untuk setiap baris hasil query yang diperoleh.
5. Setiap baris hasil query yang diperoleh akan dimasukkan ke dalam slice `model.Student` menggunakan `append`.
6. Setelah semua baris hasil query dimasukkan ke dalam slice `model.Student`, function akan mengembalikan slice `model.Student` dan `error`.

#### Test Case 3

**Input**:

Untuk function `Update`, input yang diberikan adalah sebuah bilangan bulat `id` yang merepresentasikan ID guru yang ingin diupdate dan pointer ke `model.Teacher` yang berisi data guru yang baru.

**Expected Output / Behavior**:

Function akan mengembalikan `error`. `error` akan bernilai `nil` jika tidak terjadi kesalahan, sedangkan jika terjadi kesalahan dalam menjalankan perintah SQL, function akan mengembalikan `error` yang sesuai.

**Explanation**:

1. Function `Update` dipanggil dengan ID guru yang ingin diupdate dan data guru yang baru.
2. Function akan mengeksekusi perintah SQL dengan data guru yang diberikan dan ID guru yang ingin diupdate, yang akan mengupdate data guru yang memiliki ID tersebut pada tabel `teachers`.
3. Jika terjadi kesalahan dalam menjalankan perintah SQL, function akan mengembalikan `error` yang sesuai.
4. Jika tidak terjadi kesalahan, function akan mengembalikan `nil`.

### Note

Gunakan perintah curl untuk melakukan pengujian terhadap beberapa endpoint yang ada pada aplikasi di atas, contoh:

1. Untuk menambahkan data siswa baru:

   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"name": "Aditira", "address": "Jakarta", "class": "A"}' http://localhost:8080/student/add
   ```

2. Untuk mengambil semua data siswa:

   ```bash
   curl http://localhost:8080/student/get-all
   ```

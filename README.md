🎓 Sistem Pelaporan Prestasi Mahasiswa

Backend REST API – UAS Pemrograman Backend Lanjut

👤 Identitas Mahasiswa

Nama: Shelyna Riska Amanatullah
NIM: 434231005
Kelas: C-2

📌 Deskripsi Project

Aplikasi ini merupakan Backend REST API yang digunakan untuk mengelola pelaporan prestasi mahasiswa.
Sistem dirancang untuk mendukung:

1. Role Based Access Control (RBAC)
2. Autentikasi JWT
3. Pelaporan prestasi dengan struktur dinamis menggunakan MongoDB
4. Verifikasi prestasi oleh dosen wali
5. Manajemen user, mahasiswa, dan dosen
6. Upload lampiran prestasi
7. Dashboard dan statistik prestasi
8. Integrasi 2 database (PostgreSQL + MongoDB)
9. Mengacu pada dokumen SRS (Software Requirement Specification) Sistem Pelaporan Prestasi Mahasiswa

📚 Arsitektur Sistem
Sistem menggunakan dua database sekaligus, yaitu PostgreSQL untuk data relasional dan MongoDB untuk data yang dinamis.

1. PostgreSQL
   Digunakan untuk data yang memiliki relasi tetap, seperti:
   Users, Roles, Permissions, Role Permissions, Students, Lecturers, dan Achievement References (relasi ke MongoDB).

2. MongoDB
   Digunakan untuk menyimpan detail prestasi yang dapat berbeda-beda setiap mahasiswa. Data seperti jenis prestasi, detail dinamis, lampiran, tags, dan points disimpan sebagai dokumen.

🛡️ Role dan Akses (RBAC)
Setiap role memiliki permission berbeda yang tersimpan di database.

1. Admin
   Admin memiliki seluruh akses tanpa pengecualian.
   Admin dapat mengelola seluruh user, melihat semua prestasi, melihat seluruh data mahasiswa dan dosen, mengatur dosen pembimbing mahasiswa, melihat statistik dan laporan prestasi, serta mengakses seluruh fungsi mahasiswa dan dosen wali.

2. Mahasiswa
   Mahasiswa dapat membuat prestasi, mengedit prestasi yang masih draft, menghapus draft, submit prestasi, melihat prestasi miliknya sendiri, meng-upload file lampiran, melihat daftar dosen, melihat list advisee dosen wali, dan melihat statistik prestasi.

3. Dosen Wali
   Dosen wali dapat melihat prestasi mahasiswa bimbingan, melakukan verifikasi prestasi, menolak prestasi dengan catatan, melihat riwayat prestasi, melihat seluruh data mahasiswa, melihat detail mahasiswa bimbingan, membaca prestasi mahasiswa tertentu, melihat statistik, dan melihat laporan prestasi mahasiswa.

🔄 Alur Sistem

1. Mahasiswa membuat laporan prestasi
   Detail prestasi disimpan ke MongoDB, dan reference-nya disimpan ke PostgreSQL dengan status "draft".

2. Mahasiswa melakukan submit prestasi
   Status pada PostgreSQL berubah dari draft menjadi submitted.

3. Dosen wali melihat daftar prestasi mahasiswa bimbingan
   Sistem mengambil data mahasiswa bimbingan dari PostgreSQL, lalu mengambil detail prestasi dari MongoDB.

4. Dosen melakukan verifikasi atau penolakan
   Dosen mengisi verified_by, verified_at atau rejection_note. Status berubah menjadi verified atau rejected.

5. Admin memantau seluruh prestasi
   Admin dapat melihat seluruh data prestasi untuk keperluan monitoring, statistik, dan evaluasi.

🛠 Teknologi yang Digunakan
   Backend dibangun menggunakan Go dengan framework Fiber. Autentikasi menggunakan JWT, PostgreSQL sebagai database relasional, MongoDB sebagai database NoSQL untuk prestasi, pgx dan mongo-driver sebagai driver database, Swagger (opsional) untuk dokumentasi, serta Fiber Multipart untuk upload file.

🗂 Struktur Database
1. PostgreSQL
a. Users
   Menyimpan data user seperti username, email, password_hash, role_id, dan status aktif.

b. Roles
   Berisi daftar role yaitu Admin, Mahasiswa, dan Dosen Wali.

c. Permissions
   Berisi daftar izin seperti achievement:create, achievement:update_own, user:manage, report:view_statistics, lecturer:read_advisees, dan lainnya.

d. Role Permissions
   Berfungsi sebagai penghubung antara role dan permissions.

e. Students
   Berisi informasi mahasiswa seperti student_id, program_study, academic_year, dan advisor_id.

f. Lecturers
   Berisi data dosen seperti lecturer_id dan department.

g. Achievement References
   Menghubungkan data prestasi di PostgreSQL dengan MongoDB. Terdapat informasi seperti mongo_achievement_id, status, submitted_at, verified_by, verified_at, dan rejection_note.

2. MongoDB – Collection achievements

   Contoh format dokumen prestasi:
  {
    "studentId": "UUID",
    "achievementType": "competition",
    "title": "...",
    "description": "...",
    "details": {},
    "attachments": [],
    "tags": [],
    "points": 0,
    "createdAt": "ISODate",
    "updatedAt": "ISODate"
  }

▶️ Cara Menjalankan Project

1. Clone repository
  git clone https://github.com/cococimoryy1/projectuas_be.git
  cd projectuas_be

2. Install dependencies
   go mod tidy

3. Setup file .env
   APP_PORT=8080
   POSTGRES_DSN=postgres://postgres:YOURPASSWORD@localhost:5432/prestasi_mahasiswa?sslmode=disable
   MONGO_URI=mongodb://localhost:27017
   MONGO_DB=prestasi_mahasiswa
   JWT_SECRET=supersecret

4. Jalankan server
   go run main.go

Jika berhasil, sistem menampilkan:
PostgreSQL connected
MongoDB connected
Server running on port 8080

📡 Daftar Endpoint Utama
Auth

POST /login
GET /profile

Achievements
GET /achievements
GET /achievements/:id
POST /achievements
PUT /achievements/:id
DELETE /achievements/:id
POST /achievements/:id/submit
POST /achievements/:id/verify
POST /achievements/:id/reject
GET /achievements/:id/history
POST /achievements/:id/attachments

Students & Lecturers
GET /students
GET /students/:id
GET /students/:id/achievements
PUT /students/:id/advisor
GET /lecturers
GET /lecturers/:id/advisees

Reports
GET /reports/statistics
GET /reports/student/:id
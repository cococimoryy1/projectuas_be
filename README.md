# 🎓 Sistem Pelaporan Prestasi Mahasiswa  
Backend API – UAS Pemrograman Backend Lanjut

## 👤 Identitas Mahasiswa
- **Nama:** SHELYNA RISKA AMANATULLAH 
- **NIM:** 434231005  
- **Kelas:** C-2

---

# 📌 Deskripsi Project

Project ini adalah aplikasi **Backend REST API** yang digunakan untuk **mengelola pelaporan prestasi mahasiswa**.  
Sistem mendukung:

- Role Based Access Control (RBAC)  
- Autentikasi JWT  
- Pelaporan prestasi dinamis (menggunakan MongoDB)  
- Verifikasi prestasi oleh dosen wali  
- Manajemen pengguna, mahasiswa, dosen  
- Dashboard / statistik dasar  
- Upload lampiran prestasi  

Semua spesifikasi mengacu pada dokumen resmi:

**Software Requirement Specification (SRS) – Sistem Pelaporan Prestasi Mahasiswa**.

---

# 📚 Arsitektur Sistem

Sistem menggunakan **dua database**:

### 🟦 PostgreSQL (Data Relasional)
Digunakan untuk data yang memiliki relasi tetap:
- Users
- Roles
- Permissions
- Role Permissions
- Students
- Lecturers
- Achievement References (relasi ke MongoDB)

### 🟩 MongoDB (Prestasi Dinamis)
Digunakan untuk menyimpan detail prestasi yang **bisa berbeda-beda** setiap mahasiswa:
- Types: competition, publication, organization, certification, academic, other
- Field dinamis berdasarkan tipe prestasi
- Lampiran (attachments)
- Tags
- Points
- Metadata (createdAt, updatedAt)

---

# 🛡️ Role & Akses (RBAC)

### 1️⃣ **Admin**
**Akses:**
- Kelola semua user (CRUD)
- Set role user
- Lihat semua prestasi mahasiswa
- Kelola data dosen & mahasiswa
- Role management

### 2️⃣ **Mahasiswa**
**Akses:**
- Membuat prestasi  
- Meng-edit prestasi berstatus **draft**  
- Menghapus prestasi **draft**  
- Submit untuk verifikasi  
- Melihat prestasi milik sendiri  
- Upload file prestasi  

### 3️⃣ **Dosen Wali**
**Akses:**
- Melihat prestasi mahasiswa bimbingan  
- Memverifikasi prestasi  
- Menolak prestasi & memberi catatan  

---

# 🔄 Alur Sistem Sesuai Modul

### 1. **Mahasiswa membuat laporan prestasi**
- Data prestasi disimpan ke MongoDB  
- Reference (status draft) disimpan ke PostgreSQL

### 2. **Mahasiswa Submit Prestasi**
- Status di PostgreSQL berubah dari **draft → submitted**

### 3. **Dosen Wali melihat daftar prestasi mahasiswa bimbingan**
- Ambil student_id dari tabel `students`
- Ambil referensi prestasi dari PostgreSQL
- Ambil detail prestasi dari MongoDB

### 4. **Dosen memverifikasi / menolak**
- `verified_at`, `verified_by` di PostgreSQL
- Status berubah ke **verified** atau **rejected**

### 5. **Admin bisa melihat semua prestasi**
Untuk keperluan rekap & monitoring.

---

# 🛠️ Teknologi yang Digunakan

| Komponen | Teknologi |
|---------|-----------|
| Backend Framework | Go + Fiber |
| Auth | JWT |
| Database Relasional | PostgreSQL |
| Database NoSQL | MongoDB |
| Documentation | (Opsional) Swagger |
| ORM / Driver | pgx / mongo-driver |
| File Upload | Fiber Multipart |

---

# 🗂️ Struktur Database

## 🟦 PostgreSQL Tables

### 1. Users  
- username  
- email  
- password_hash  
- role_id  
- is_active  

### 2. Roles  
- Admin  
- Mahasiswa  
- Dosen Wali  

### 3. Permissions  
- achievement:create  
- achievement:read  
- achievement:update  
- achievement:delete  
- achievement:verify  
- user:manage  

### 4. Role Permissions  
Mapping role → permission

### 5. Students  
- student_id  
- program_study  
- academic_year  
- advisor_id (dosen)

### 6. Lecturers  
- lecturer_id  
- department  

### 7. Achievement References  
Relasi Postgres → MongoDB  
- mongo_achievement_id  
- status (draft/submitted/verified/rejected)  
- verified_by  
- rejection_note  

---

## 🟩 MongoDB – Collection `achievements`

Format dokumen sesuai modul:

```json
{
  "studentId": "UUID",
  "achievementType": "competition",
  "title": "...",
  "description": "...",
  "details": { ... },
  "attachments": [],
  "tags": [],
  "points": 0,
  "createdAt": ISODate(),
  "updatedAt": ISODate()
}

▶️ Cara Menjalankan Project
1️⃣ Clone repository
git clone https://github.com/cococimoryy1/projectuas_be.git
cd projectuas_be

2️⃣ Install dependencies
go mod tidy

3️⃣ Setup .env
APP_PORT=8080
POSTGRES_DSN=postgres://postgres:YOURPASSWORD@localhost:5432/prestasi_mahasiswa?sslmode=disable
MONGO_URI=mongodb://localhost:27017
MONGO_DB=prestasi_mahasiswa

JWT_SECRET=supersecret

4️⃣ Jalankan server
go run main.go

Jika sukses:
PostgreSQL connected
MongoDB connected
Server running on port 8080
📡 Endpoint Utama (Ringkas)

Auth
POST /login
GET /profile
Achievements
POST /achievements
PUT /achievements/:id
DELETE /achievements/:id
POST /achievements/:id/submit
POST /achievements/:id/verify
POST /achievements/:id/reject
Students / Lecturers
GET /lecturers/:id/advisees
GET /students/:id/achievements

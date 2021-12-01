package hellogo

func SayHello(name string) string {
	return "Hello, " + name
}

/*
====== GIT NOTES =======

1. membuat repo :
git init

2. menambah remote origin dari github :
buat project di github lalu
git remote add origin https://github.com/Faridalim/hellogo.git

3. menambahkan perubahan :
git add . (jika semua)
git add hello.go (jika add file tertentu saja)

4. mengecek status perubahan :
git status

5. membuat commit atas perubahan :
git commit -m "deskripsi commit"

6. mengepush commit ke server github :
git push origin master

7. memberi versi rilis dengan tag :
git tag v1.0.0

8. mengepush rilis ke remote server :
git push origin v1.0.0

Untuk rilis versi baru, tinggal diulang prosesnya dari step 3 s.d step 8
*/

/*
==== Cara meng-upgrade dependency : ======

1. mengubah versi di go.mod
2. run di terminal : go get
*/

/*
===== implementasi major upgrade : ==========

diusahakan backward compatible
jika bener2 terpaksa tidak backward compatible, maka sebaiknya ganti nama module
sehingga tidak merusak aplikasi yang sudah ada yg menggunakan module tersebut

biasanya nama modulnya di go.mod diubah dengan menambah slash :
sebelum : module github.com/Faridalim/hellogo
sesudah : module github.com/Faridalim/hellogo/v2

jadi menambahkan /v2. Jadi reponya sama gapapa, cuma dia harus menambahkan /v2 untuk
memakai yang versi 2
*/

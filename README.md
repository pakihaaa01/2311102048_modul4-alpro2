# 2311102048_modul4-alpro2
laporan praktikum modul 4 algoritma pemrograman 2
1. program no. 1 berfungsi untuk menghitung permutasi dan kombinasi dari dua set bilangan yang di berikan oleh pengguna. terdapat beebrapa fungsi dalam program tersebut:
factorial(n int) int: 
- menghitung faktorial dari bilangan
- menggunakan perulangan untuk mengalikan semua bilangan dari 1 hingga n
permutasi(n, r int) int:
- menghitung permutasi dari n objek yang diambil r sekaligus.
- menggunakan rumus: P(n,r)= n!/(n-r).
kombinasi(n, r) int:
- menghitung kombinasi dari n objek yang diambil r sekaligus.
- menggunakan rumus: C(n,r)= n!/r!(n-r).
func main()
- program meminta pengguna untuk memasukkan empat bilangan bulat: a, b, c, d.
- menghitung permutasi dan kombinasi untuk dua set:
  - set pertama: permutasi dan kombinasi dari a dan c.
  - set kedua: permutasi dan kombinasi dari b dan d.
- menampilkan hasil perhitungan permutasi dan kombinasi untuk kedua set tersebut
2. program no. 2 dibuat untuk menentukan pemenang dari sebuah kompetisi pemrograman berdasarkan jumlah soal yang diselesaikan dan total waktu yang dibutuhkan. fungsi fungsi yang ada di program:
hitungSkor(waktu[8]int, soal int, skor int):
- menghitung jumlah soal yang diselesaikan dan total waktu yang dibutuhkan.
- menggunakan pointer untuk mengembalikan hasil perhitungan.
func main()
- program meminta pengguna untuk memasukkan nama peserta dan waktu penyelesaian untuk 8 soal.
- jika waktu penyelesaian lebih dari 300 menit, soal dianggap tidak selesai.
- program menghitung jumlah soal yang diselesaikan dan total waktu untuk setiap peserta.
- pemenang ditentukan berdasarkan:
  - jumlah soal yang diselesaikan terbanyak.
  - jika jumlah soal sama, pemenang adalah yang memiliki total waktu paling sedikit
- program berakhir ketika pengguna memasukkan "Selesai".
3. program mencetak deret bilangan berdasarkan aturan tertentu, sering dikenal sebagai "Collatz Conjecture" atau "3n + 1 problem". berikut adalah deskripsi dari program tersebut:
cetakDeret(n int)
- mencetak deret bilangan dimulai dari n.
- jika n adalah bilangan genap, bagi n dengan 2.
- jika n adalah bilangan ganjil, ubah n menjadi 3n + 1.
- proses berlanjut hingga n menjadi 1, dan program mencetak 1 sebagai akhir deret.
yang harus diperhatikan adalah pengguna harus memasukkan bilangan bulat positif yang kurang dari 1.000.000 dan jika input tidak valid, program menampilakan pesan kesalahan.
jika input valid maka program memanggil fungsi cetakDeret untuk mencetak deret bilangan sesuai aturan.

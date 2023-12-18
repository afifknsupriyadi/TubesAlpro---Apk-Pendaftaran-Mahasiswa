package main

import (
	"fmt"
)

const NMAX int = 100
const ADMINMAX int = 2
const USERMAX int = 100

type mahasiswa struct {
	nama                     string
	mtk, bindo, bing, fisika int
	prodi                    string
	status                   bool
}
type jurusan struct {
	namaProdi  string
	akreditasi string
	kuota      int
	terisi     int
}

type arrMhs [NMAX]mahasiswa
type arrJurusan [NMAX]jurusan

type admin struct {
	user    string
	pass    int
	isAdmin bool
}
type admins struct {
	info   [ADMINMAX]admin
	nAdmin int
}

var dataAdmin = admins{
	info: [ADMINMAX]admin{
		{user: "admin1", pass: 1234, isAdmin: true},
		{user: "admin2", pass: 5678, isAdmin: true},
	},
	nAdmin: 2,
}

type user struct {
	Username  string
	Password  string
	isStudent bool
}
type arrUser [USERMAX]user

func header() {
	fmt.Println("^+=============================================^+")
	fmt.Println("^|       APLIKASI PENDAFTARAN MAHASISWA        ^|")
	fmt.Println("^|           ALGORITMA PEMROGRAMAN             ^|")
	fmt.Println("^+---------------------------------------------^+")
	fmt.Println("^|          AFIF KURNIAWAN SUPRIYADI           ^|")
	fmt.Println("^|            MOHAMMAD FAIZ HAIKAL             ^|")
	fmt.Println("^+=============================================^+")

}

func menu() {
	var tabProdi arrJurusan
	var tabMhs arrMhs

	var nProdi, nMhs int
	var opsi string
	var opsi2 string

	header()
	for opsi != "3" {
		fmt.Println("\n^+---------------------------------------------^+")
		fmt.Println("^|                 MENU UTAMA                  ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 1. Admin                                    ^|")
		fmt.Println("^| 2. Mahasiswa                                ^|")
		fmt.Println("^| 3. Keluar                                   ^|")
		fmt.Println("^+---------------------------------------------^+")

		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&opsi)

		for opsi != "1" && opsi != "2" && opsi != "3" {
			fmt.Print("Input tidak valid, Pilihan Anda : ")
			fmt.Scan(&opsi)
		}

		if opsi == "1" {
			administrator()
			for opsi2 != "6" {
				fmt.Println("\n^+=============================================^+")
				fmt.Println("^|       INSTITUT TEKNNOLOGI BOJONGSOANG       ^|")
				fmt.Println("^+---------------------------------------------^+")
				fmt.Println("^|                MENU ADMIN                   ^|")
				fmt.Println("^+=============================================^+")
				fmt.Println("^| 1. Add Data Jurusan                         ^|")
				fmt.Println("^| 2. Add Data Mahasiswa                       ^|")
				fmt.Println("^| 3. Edit Data                                ^|")
				fmt.Println("^| 4. Sort Data Mahasiswa                      ^|")
				fmt.Println("^| 5. Pencarian Data Mahasiswa                 ^|")
				fmt.Println("^+---------------------------------------------^+")
				fmt.Println("^| 6. Keluar                                   ^|")
				fmt.Println("^+---------------------------------------------^+")

				fmt.Print("Pilihan Anda: ")
				fmt.Scan(&opsi2)

				for opsi2 != "1" && opsi2 != "2" && opsi2 != "3" && opsi2 != "4" && opsi2 != "5" && opsi2 != "6" {
					fmt.Print("Input tidak valid, Pilihan Anda : ")
					fmt.Scan(&opsi2)
				}
				if opsi2 == "1" {
					addProdi_byAdmin(&tabProdi, &nProdi)
				} else if opsi2 == "2" {
					if nProdi != 0 {
						addMhs_byAdmin(&tabMhs, tabProdi, &nMhs, nProdi)
					} else {
						fmt.Println("Data Jurusan tidak ada.")
					}
				} else if opsi2 == "3" {
					if nProdi != 0 && nMhs != 0 {
						editData_byAdmin(&tabProdi, &tabMhs, &nProdi, &nMhs)
					} else {
						fmt.Println("Data Jurusan atau data mahasiswa tidak ada.")
					}
				} else if opsi2 == "4" {
					if nMhs != 0 {
						sortingMenu(tabMhs, tabProdi, nMhs, nProdi)
					} else {
						fmt.Println("Data Mahasiswa tidak ada.")
					}
				} else if opsi2 == "5" {
					if nProdi != 0 && nMhs != 0 {
						pencarianData(tabMhs, tabProdi, nMhs, nProdi)
					} else {
						fmt.Println("Data Jurusan atau data mahasiswa tidak ada.")
					}
				}
			}

			fmt.Println("^+---------------------------------------------^+")
			fmt.Println("^|       ......Exiting from admin......     	^|")
			fmt.Println("^+---------------------------------------------^+")

		} else if opsi == "2" {
			loginSystem_asMhs(tabProdi, tabMhs, nProdi, nMhs)
		}
	}
	fmt.Println("\n^+---------------------------------------------^+")
	fmt.Println("^|    ......Exiting from application......   	^|")
	fmt.Println("^|           ----Terima Kasih----          	^|")
	fmt.Println("^+---------------------------------------------^+")

}

func loginSystem_asAdmin(A *admins) *admin {
	var username string
	var password int

	fmt.Println("\n^+---------------------------------------------^+")
	fmt.Println("^|        Login System as Administrator        ^|")
	fmt.Println("^+---------------------------------------------^+")

	fmt.Print("Username: ")
	fmt.Scan(&username)

	fmt.Print("Password: ")
	fmt.Scan(&password)

	for i := 0; i < A.nAdmin; i++ {
		if A.info[i].user == username && A.info[i].pass == password {
			return &A.info[i]
		}
	}
	return nil
}

func administrator() {
	cekAdmin := loginSystem_asAdmin(&dataAdmin)

	for cekAdmin == nil {
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^|               Login failed              	^|")
		fmt.Println("^|       Username atau Password salah         	^|")
		fmt.Println("^+---------------------------------------------^+")
		cekAdmin = loginSystem_asAdmin(&dataAdmin)

	}

	if cekAdmin.isAdmin == true {
		fmt.Println("^+---------------------------------------------^+")
		fmt.Printf("^|         Selamat Datang, admin %s!       ^|\n", cekAdmin.user)
		fmt.Println("^+---------------------------------------------^+")
	}
}

func addProdi_byAdmin(P *arrJurusan, nP *int) {
	fmt.Println("\n^+---------------------------------------------^+")
	fmt.Println("^|             INPUT DATA JURUSAN              ^|")
	fmt.Println("^+---------------------------------------------^+")

	fmt.Scan(&P[*nP].namaProdi)
	for P[*nP].namaProdi != "#" {
		fmt.Scan(&P[*nP].akreditasi, &P[*nP].kuota)
		*nP++

		fmt.Scan(&P[*nP].namaProdi)
	}
	fmt.Println("^+---------------------------------------------^+")
}

func addMhs_byAdmin(M *arrMhs, P arrJurusan, nM *int, nP int) {
	var idx_prodi int
	fmt.Println("\n^+---------------------------------------------^+")
	fmt.Println("^|            INPUT DATA MAHASISWA             ^|")
	fmt.Println("^+---------------------------------------------^+")
	fmt.Scan(&M[*nM].nama)
	for M[*nM].nama != "#" {
		fmt.Scan(&M[*nM].mtk, &M[*nM].bindo, &M[*nM].bing, &M[*nM].fisika, &M[*nM].prodi)

		idx_prodi = seqProdi(&P, &nP, M[*nM].prodi)
		for idx_prodi == -1 {
			fmt.Print("Jurusan invalid. Masukkan jurusan: ")
			fmt.Scan(&M[*nM].prodi)

			idx_prodi = seqProdi(&P, &nP, M[*nM].prodi)
		}

		P[idx_prodi].terisi++
		*nM++
		fmt.Scan(&M[*nM].nama)
	}
	fmt.Println("^+---------------------------------------------^+")
}

func seqProdi(P *arrJurusan, nP *int, x string) int {
	var found int = -1
	var i int = 0

	for i < *nP && found == -1 {
		if x == P[i].namaProdi {
			found = i
		}
		i++
	}

	return found
}

func editData_byAdmin(P *arrJurusan, M *arrMhs, nP *int, nM *int) {
	var xProdi, xNama string
	var opsi string

	for opsi != "7" {
		fmt.Println("\n^+---------------------------------------------^+")
		fmt.Println("^|                  EDIT DATA                  ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 1. View Data Jurusan                        ^|")
		fmt.Println("^| 2. Edit Data Jurusan                        ^|")
		fmt.Println("^| 3. Delete Data Jurusan                      ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 4. View Data Mahasiswa                      ^|")
		fmt.Println("^| 5. Edit Data Mahasiswa                      ^|")
		fmt.Println("^| 6. Delete Data Mahasiswa                    ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 7. Keluar                                   ^|")
		fmt.Println("^+---------------------------------------------^+")

		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&opsi)

		for opsi != "1" && opsi != "2" && opsi != "3" && opsi != "4" && opsi != "5" && opsi != "6" && opsi != "7" {
			fmt.Print("Input tidak valid, Pilihan Anda : ")
			fmt.Scan(&opsi)
		}

		if opsi == "1" {
			viewProdi(*P, *nP)
			fmt.Println("^+---------------------------------------------^+")
		} else if opsi == "2" {
			editJurusan_byAdmin(P, nP)
		} else if opsi == "3" {
			viewProdi(*P, *nP)
			fmt.Println("^+---------------------------------------------^+")
			fmt.Print("\nNama Jurusan yg akan dihapus: ")
			fmt.Scan(&xProdi)
			deleteJurusan_byAdmin(P, nP, xProdi)

		} else if opsi == "4" {
			viewMahasiswa(*M, *P, *nM, *nP)
			fmt.Println("^+---------------------------------------------^+")
		} else if opsi == "5" {
			editMahasiswa_byAdmin(M, P, nM, nP)
		} else if opsi == "6" {
			viewMahasiswa(*M, *P, *nM, *nP)
			fmt.Println("^+---------------------------------------------^+")
			fmt.Print("\nNama Mahasiswa yg akan dihapus: ")
			fmt.Scan(&xNama)
			deleteMahasiswa_byAdmin(M, nM, xNama)
		}
	}
	fmt.Println("^+---------------------------------------------^+")
	fmt.Println("^|     ......Exiting from menu edit......     	^|")
	fmt.Println("^+---------------------------------------------^+")
}

// Edit Data Jurusan
func editJurusan_byAdmin(P *arrJurusan, nP *int) {
	var opsi string
	var xProdi string

	for opsi != "5" {
		fmt.Println("\n^+---------------------------------------------^+")
		fmt.Println("^|             HALAMAN EDIT JURUSAN            ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 1. Edit Nama Jurusan                        ^|")
		fmt.Println("^| 2. Edit Akreditasi Jurusan                  ^|")
		fmt.Println("^| 3. Edit Kuota Jurusan                       ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 4. View Data Jurusan                        ^|")
		fmt.Println("^| 5. Keluar                                   ^|")
		fmt.Println("^+---------------------------------------------^+")

		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&opsi)

		for opsi != "1" && opsi != "2" && opsi != "3" && opsi != "4" && opsi != "5" {
			fmt.Print("Input tidak valid, Pilihan Anda : ")
			fmt.Scan(&opsi)
		}

		if opsi == "1" {
			viewProdi(*P, *nP)
			fmt.Println("^+---------------------------------------------^+")
			fmt.Print("Ubah Prodi\t: ")
			fmt.Scan(&xProdi)
			editProdi_byAdmin(P, *nP, xProdi)
		} else if opsi == "2" {
			viewProdi(*P, *nP)
			fmt.Println("^+---------------------------------------------^+")
			fmt.Print("Masukkan Prodi\t\t: ")
			fmt.Scan(&xProdi)
			editAkreditasi_byAdmin(P, *nP, xProdi)
		} else if opsi == "3" {
			viewProdi(*P, *nP)
			fmt.Println("^+---------------------------------------------^+")
			fmt.Print("Masukkan Prodi\t: ")
			fmt.Scan(&xProdi)
			editKuota_byAdmin(P, *nP, xProdi)
		} else if opsi == "4" {
			viewProdi(*P, *nP)
			fmt.Println("^+---------------------------------------------^+")
		}
	}
}
func viewProdi(P arrJurusan, nP int) {
	var i int

	fmt.Println("\n^+---------------------------------------------^+")
	fmt.Println("^|        DAFTAR JURUSAN IT BOJONGSOANG        ^|")
	fmt.Println("^+---------------------------------------------^+")

	for i < nP {
		fmt.Println(P[i].namaProdi, P[i].akreditasi, P[i].kuota)
		i++
	}
}
func editProdi_byAdmin(P *arrJurusan, nP int, X string) {
	var found int
	found = seqProdi(P, &nP, X)

	if found == -1 {
		fmt.Println("Jurusan tidak ditemukan")
	} else {
		fmt.Print("Input Prodi\t: ")
		fmt.Scan(&P[found].namaProdi)
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| Nama jurusan berhasil diubah                ^|")
		fmt.Println("^+---------------------------------------------^+")

	}
}

func editAkreditasi_byAdmin(P *arrJurusan, nP int, X string) {
	var found int
	found = seqProdi(P, &nP, X)

	if found == -1 {
		fmt.Println("Jurusan tidak ditemukan")
	} else {
		fmt.Print("Input Akreditasi\t: ")
		fmt.Scan(&P[found].akreditasi)
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| Data akreditasi berhasil diubah             ^|")
		fmt.Println("^+---------------------------------------------^+")
	}
}
func editKuota_byAdmin(P *arrJurusan, nP int, X string) {
	var found int
	found = seqProdi(P, &nP, X)

	if found == -1 {
		fmt.Println("Jurusan tidak ditemukan")
	} else {
		fmt.Print("Input Kuota\t: ")
		fmt.Scan(&P[found].kuota)
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| Data kuota jurusan berhasil diubah          ^|")
		fmt.Println("^+---------------------------------------------^+")
	}
}

func deleteJurusan_byAdmin(P *arrJurusan, nP *int, X string) {
	var found, i int

	found = seqProdi(P, nP, X)
	if found == -1 {
		fmt.Println("Jurusan tidak ditemukan")
	} else {
		i = found
		for i <= *nP-2 {
			P[i] = P[i+1]
			i = i + 1
		}
		*nP--
	}
	fmt.Println("^+---------------------------------------------^+")
	fmt.Println("^| Data jurusan berhasil dihapus               ^|")
	fmt.Println("^+---------------------------------------------^+")
}

// Edit Data Mahasiswa
func editMahasiswa_byAdmin(M *arrMhs, P *arrJurusan, nM *int, nP *int) {
	var xNama string
	var opsi string

	for opsi != "5" {
		fmt.Println("\n^+---------------------------------------------^+")
		fmt.Println("^|            HALAMAN EDIT MAHASISWA           ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 1. Edit Nama Mahasiswa                      ^|")
		fmt.Println("^| 2. Edit Nilai                               ^|")
		fmt.Println("^| 3. Edit Pilihan Jurusan                     ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 4. View Data Mahasiswa                      ^|")
		fmt.Println("^| 5. Keluar                                   ^|")
		fmt.Println("^+---------------------------------------------^+")

		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&opsi)

		for opsi != "1" && opsi != "2" && opsi != "3" && opsi != "4" && opsi != "5" {
			fmt.Print("Input tidak valid, Pilihan Anda : ")
			fmt.Scan(&opsi)
		}

		if opsi == "1" {
			viewMahasiswa(*M, *P, *nM, *nP)
			fmt.Println("^+---------------------------------------------^+")
			fmt.Print("Nama yg diubah\t: ")
			fmt.Scan(&xNama)

			editNamaMhs_byAdmin(M, *nM, xNama)
		} else if opsi == "2" {
			viewMahasiswa(*M, *P, *nM, *nP)
			fmt.Println("^+---------------------------------------------^+")
			fmt.Print("Nama mahasiswa yg nilainya akan diubah: ")
			fmt.Scan(&xNama)

			dataNilai_byAdmin(*M, *P, *nM, *nP, xNama)
			editNilai_byAdmin(M, *P, *nM, *nP, xNama)
		} else if opsi == "3" {
			viewMahasiswa(*M, *P, *nM, *nP)
			fmt.Println("^+---------------------------------------------^+")
			fmt.Print("Nama mahasiswa yg jurusannya akan diubah: ")
			fmt.Scan(&xNama)

			dataNilai_byAdmin(*M, *P, *nM, *nP, xNama)
			editPilihanJurusan_byAdmin(M, *P, *nM, *nP, xNama)
		} else if opsi == "4" {
			viewMahasiswa(*M, *P, *nM, *nP)
			fmt.Println("^+---------------------------------------------^+")
		}
	}
}

func lulus(M *arrMhs, P *arrJurusan, nM int, nP *int) {
	var idx_prodi int
	var average float64
	const pasingGrade float64 = 90

	for i := 0; i <= nM; i++ {
		idx_prodi = seqProdi(P, nP, M[i].prodi)

		M[i].status = false
		if idx_prodi != -1 {
			average = float64((M[i].mtk + M[i].bindo + M[i].bing + M[i].fisika) / 4)
			if average >= pasingGrade && P[idx_prodi].terisi <= P[idx_prodi].kuota {
				M[i].status = true
			}
		}
	}

}

func viewMahasiswa(M arrMhs, P arrJurusan, nM int, nP int) {
	var idx_prodi int

	fmt.Println("\n^+---------------------------------------------^+")
	fmt.Println("^| DATA CALON MAHASISWA JURUSAN IT BOJONGSOANG ^|")
	fmt.Println("^+---------------------------------------------^+")

	for i := 0; i < nM; i++ {
		idx_prodi = seqProdi(&P, &nP, M[i].prodi)
		if idx_prodi == -1 {
			fmt.Printf("%d. %s %d %d %d %d - \n", i+1, M[i].nama, M[i].mtk, M[i].bindo, M[i].bing, M[i].fisika)
		} else {
			lulus(&M, &P, nM, &nP)
			fmt.Printf("%d. %s %d %d %d %d %s %t\n", i+1, M[i].nama, M[i].mtk, M[i].bindo, M[i].bing, M[i].fisika, M[i].prodi, M[i].status)
		}
	}
}

func editNamaMhs_byAdmin(M *arrMhs, nM int, X string) {
	var found int
	found = seqMhs(*M, nM, X)

	if found == -1 {
		fmt.Println("Nama mahasiswa tidak ditemukan")
	} else {
		fmt.Print("Nama Baru\t: ")
		fmt.Scan(&M[found].nama)
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| Nama mahasiswa berhasil diubah              ^|")
		fmt.Println("^+---------------------------------------------^+")
	}
}

func editNilai_byAdmin(M *arrMhs, P arrJurusan, nM int, nP int, X string) {
	var opsi string
	var found int
	found = seqMhs(*M, nM, X)

	for opsi != "6" {

		fmt.Println("\n^+---------------------------------------------^+")
		fmt.Println("^|              HALAMAN EDIT NILAI             ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 1. Edit Nilai Matematika                    ^|")
		fmt.Println("^| 2. Edit Nilai B.Indonesia                   ^|")
		fmt.Println("^| 3. Edit Nilai B.Inggris                     ^|")
		fmt.Println("^| 4. Edit Nilai Fisika                        ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 5. View Data Nilai                          ^|")
		fmt.Println("^| 6. Keluar                                   ^|")
		fmt.Println("^+---------------------------------------------^+")

		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&opsi)

		for opsi != "1" && opsi != "2" && opsi != "3" && opsi != "4" && opsi != "5" && opsi != "6" {
			fmt.Print("Input tidak valid, Pilihan Anda : ")
			fmt.Scan(&opsi)
		}

		if found == -1 {
			fmt.Println("Nama mahasiswa tidak ditemukan")
		} else {
			if opsi == "1" {
				fmt.Print("\nNilai Matematika baru\t: ")
				fmt.Scan(&M[found].mtk)

				fmt.Println("^+---------------------------------------------^+")
				fmt.Println("^| Nilai Matematika berhasil diubah            ^|")
				fmt.Println("^+---------------------------------------------^+")
			} else if opsi == "2" {
				fmt.Print("\nNilai B.Indonesia baru\t: ")
				fmt.Scan(&M[found].bindo)

				fmt.Println("^+---------------------------------------------^+")
				fmt.Println("^| Nilai B.Indonesia berhasil diubah           ^|")
				fmt.Println("^+---------------------------------------------^+")
			} else if opsi == "3" {
				fmt.Print("\nNilai B.Inggris baru\t: ")
				fmt.Scan(&M[found].bing)

				fmt.Println("^+---------------------------------------------^+")
				fmt.Println("^| Nilai B.Inggris berhasil diubah             ^|")
				fmt.Println("^+---------------------------------------------^+")
			} else if opsi == "4" {
				fmt.Print("\nNilai Fisika baru\t: ")
				fmt.Scan(&M[found].fisika)

				fmt.Println("^+---------------------------------------------^+")
				fmt.Println("^| Nilai Fisika berhasil diubah                ^|")
				fmt.Println("^+---------------------------------------------^+")
			} else if opsi == "5" {
				dataNilai_byAdmin(*M, P, nM, nP, X)
			}
		}
	}
}

func dataNilai_byAdmin(M arrMhs, P arrJurusan, nM int, nP int, xNama string) {
	var idx_mhs int
	var idx_prodi int

	idx_mhs = seqMhs(M, nM, xNama)

	if idx_mhs == -1 {
		fmt.Println("\nData Mahasiswa Tidak Ditemukan.")
	} else {
		idx_prodi = seqProdi(&P, &nP, M[idx_mhs].prodi)
		if idx_prodi == -1 {
			M[idx_mhs].prodi = "-"
		}

		fmt.Println("\n^+---------------------------------------------^+")
		fmt.Println("^|             DATA CALON MAHASISWA            ^|")
		fmt.Println("^+---------------------------------------------^+")

		fmt.Println("Nama\t\t\t:", M[idx_mhs].nama)
		fmt.Println("Nilai Matematika\t:", M[idx_mhs].mtk)
		fmt.Println("Nilai B.Indonesia\t:", M[idx_mhs].bindo)
		fmt.Println("Nilai B.Inggris\t\t:", M[idx_mhs].bing)
		fmt.Println("Nilai Fisika\t\t:", M[idx_mhs].fisika)
		fmt.Println("Jurusan\t\t\t:", M[idx_mhs].prodi)
		fmt.Println("^+---------------------------------------------^+")
	}
}

func editPilihanJurusan_byAdmin(M *arrMhs, P arrJurusan, nM int, nP int, X string) {
	var found int
	var idx_prodi int

	found = seqMhs(*M, nM, X)
	if found == -1 {
		fmt.Println("Nama mahasiswa tidak ditemukan")
	} else {
		viewProdi_mhs(P, nP)
		fmt.Print("Jurusan Baru\t: ")
		fmt.Scan(&M[found].prodi)

		idx_prodi = seqProdi(&P, &nP, M[found].prodi)
		for idx_prodi == -1 {
			fmt.Print("\nJurusan invalid.")
			fmt.Print("\nJurusan Baru\t: ")
			fmt.Scan(&M[found].prodi)

			idx_prodi = seqProdi(&P, &nP, M[found].prodi)
		}

		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| Data jurusan berhasil diubah                ^|")
		fmt.Println("^+---------------------------------------------^+")
	}
}

func deleteMahasiswa_byAdmin(M *arrMhs, nM *int, X string) {
	var found, i int

	found = seqMhs(*M, *nM, X)
	if found == -1 {
		fmt.Println("Jurusan tidak ditemukan")
	} else {
		i = found
		for i <= *nM-2 {
			M[i] = M[i+1]
			i = i + 1
		}
		*nM--
	}
	fmt.Println("^+---------------------------------------------^+")
	fmt.Println("^| Data mahasiswa berhasil dihapus             ^|")
	fmt.Println("^+---------------------------------------------^+")
}

// Sorting Data
func sortingMenu(M arrMhs, P arrJurusan, nM int, nP int) {
	var flagNilai string
	var opsi string

	for opsi != "4" {

		fmt.Println("\n^+---------------------------------------------^+")
		fmt.Println("^|      HALAMAN PENGURUTAN DATA MAHASISWA      ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 1. Pengurutan Nilai                         ^|")
		fmt.Println("^| 2. Pengurutan Jurusan                       ^|")
		fmt.Println("^| 3. Pengurutan Nama                          ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 4. Keluar                                   ^|")
		fmt.Println("^+---------------------------------------------^+")

		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&opsi)

		for opsi != "1" && opsi != "2" && opsi != "3" && opsi != "4" {
			fmt.Print("Input tidak valid, Pilihan Anda : ")
			fmt.Scan(&opsi)
		}

		if opsi == "1" {
			fmt.Println("\n^+---------------------------------------------^+")
			fmt.Println("^| Daftar Mata Pelajaran:                      ^|")
			fmt.Println("^| - Matematika                               ^|")
			fmt.Println("^| - B.Indonesia                              ^|")
			fmt.Println("^| - B.Inggris                                ^|")
			fmt.Println("^| - Fisika                                   ^|")
			fmt.Println("^+---------------------------------------------^+")
			fmt.Print("Mata Pelajaran yg diurutkan: ")

			fmt.Scan(&flagNilai)
			sortNilai(&M, nM, flagNilai)
			fmt.Println("\n^+---------------------------------------------^+")
			fmt.Printf("^| Sorting Data Berdasarkan Nilai %s ^|\n", flagNilai)
			fmt.Print("^+---------------------------------------------^+")
			viewMahasiswa(M, P, nM, nP)
		} else if opsi == "2" {
			sortJurusan(&M, nM)
			fmt.Println("\n^+---------------------------------------------^+")
			fmt.Println("^|       Sorting Data Berdasarkan Jurusan      ^|")
			fmt.Print("^+---------------------------------------------^+")
			viewMahasiswa(M, P, nM, nP)
		} else if opsi == "3" {
			sortNama(&M, nM)
			fmt.Println("\n^+---------------------------------------------^+")
			fmt.Println("^|   Sorting Data Berdasarkan Nama Mahasiswa   ^|")
			fmt.Print("^+---------------------------------------------^+")
			viewMahasiswa(M, P, nM, nP)
		}
	}
}

func sortNilai(M *arrMhs, nM int, xNilai string) {
	var pass, i int
	var temp mahasiswa

	pass = 1
	for pass <= nM-1 {
		i = pass
		temp = M[pass]
		if xNilai == "Matematika" {
			for i > 0 && temp.mtk > M[i-1].mtk {
				M[i] = M[i-1]
				i--
			}
		} else if xNilai == "B.Indonesia" {
			for i > 0 && temp.bindo > M[i-1].bindo {
				M[i] = M[i-1]
				i--
			}
		} else if xNilai == "B.Inggris" {
			for i > 0 && temp.bing > M[i-1].bing {
				M[i] = M[i-1]
				i--
			}
		} else if xNilai == "Fisika" {
			for i > 0 && temp.fisika > M[i-1].fisika {
				M[i] = M[i-1]
				i--
			}
		}
		M[i] = temp
		pass = pass + 1
	}
}

func sortJurusan(M *arrMhs, nM int) {
	var pass, i int
	var temp mahasiswa

	pass = 1
	for pass <= nM-1 {
		i = pass
		temp = M[pass]
		for i > 0 && temp.prodi < M[i-1].prodi {
			M[i] = M[i-1]
			i--
		}
		M[i] = temp
		pass = pass + 1
	}
}

func sortNama(M *arrMhs, nM int) {
	var pass, i int
	var temp mahasiswa

	pass = 1
	for pass <= nM-1 {
		i = pass
		temp = M[pass]
		for i > 0 && temp.nama < M[i-1].nama {
			M[i] = M[i-1]
			i--
		}
		M[i] = temp
		pass = pass + 1
	}
}

// Pencarian Data
func pencarianData(M arrMhs, P arrJurusan, nM int, nP int) {
	var opsi string
	var xNama string
	var xProdi string

	for opsi != "5" {
		fmt.Println("\n^+---------------------------------------------^+")
		fmt.Println("^|       HALAMAN PENCARIAN DATA MAHASISWA      ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 1. Mahasiswa diterima                       ^|")
		fmt.Println("^| 2. Mahasiswa ditolak                        ^|")
		fmt.Println("^| 3. Pencarian by nama                        ^|")
		fmt.Println("^| 4. Pencarian by jurusan                     ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 5. Keluar                                   ^|")
		fmt.Println("^+---------------------------------------------^+")

		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&opsi)

		for opsi != "1" && opsi != "2" && opsi != "3" && opsi != "4" && opsi != "5" {
			fmt.Print("Input tidak valid, Pilihan Anda : ")
			fmt.Scan(&opsi)
		}

		if opsi == "1" {
			mhsLulus(M, P, nM, nP)
		} else if opsi == "2" {
			mhsGagal(M, P, nM, nP)
		} else if opsi == "3" {
			fmt.Print("\nNama mahasiswa yg akan dicari: ")
			fmt.Scan(&xNama)
			dataNilai_byAdmin(M, P, nM, nP, xNama)
		} else if opsi == "4" {
			fmt.Print("\nNama jurusan yg akan dicari: ")
			fmt.Scan(&xProdi)
			searchProdi(M, P, nM, nP, xProdi)
		}
	}
}

func mhsLulus(M arrMhs, P arrJurusan, nM int, nP int) {

	fmt.Println("\n^+---------------------------------------------^+")
	fmt.Println("^|          DAFTAR MAHASISWA DITERIMA          ^|")
	fmt.Println("^|       INSTITUT TEKNOLOGI BOJONGSOANG        ^|")
	fmt.Println("^+---------------------------------------------^+")
	for i := 0; i < nM; i++ {
		lulus(&M, &P, nM, &nP)
		if M[i].nama != "#" && M[i].status == true {
			fmt.Println(M[i].nama, M[i].prodi)
		}
	}
}

func mhsGagal(M arrMhs, P arrJurusan, nM int, nP int) {
	var idx_prodi int
	fmt.Println("\n^+---------------------------------------------^+")
	fmt.Println("^|           DAFTAR MAHASISWA DITOLAK          ^|")
	fmt.Println("^|        INSTITUT TEKNOLOGI BOJONGSOANG       ^|")
	fmt.Println("^+---------------------------------------------^+")

	for i := 0; i < nM; i++ {
		idx_prodi = seqProdi(&P, &nP, M[i].prodi)
		if idx_prodi == -1 {
			M[i].prodi = "-"
		}

		lulus(&M, &P, nM, &nP)
		if M[i].nama != "#" && M[i].status == false {
			fmt.Println(M[i].nama, M[i].prodi)
		}
	}
}

func searchProdi(M arrMhs, P arrJurusan, nM int, nP int, x string) {

	var idx_prodi = seqProdi(&P, &nP, x)

	fmt.Println("\n^+---------------------------------------------^+")
	fmt.Println("^| DATA CALON MAHASISWA                        ^|")
	fmt.Printf("^| PROGRAM STUDI %s    ^|\n", x)
	fmt.Println("^| INSTITUT TEKNOLOGI BOJONGSOANG              ^|")
	fmt.Println("^+---------------------------------------------^+")

	if idx_prodi == -1 {
		fmt.Println("\nData Prodi Tidak Ditemukan.")
	} else {
		for i := 0; i < nM; i++ {
			if M[i].prodi == x {
				fmt.Println(M[i].nama, M[i].prodi, M[i].status)
			}
		}
	}
}

// Mahasiswa
func signUp(U *arrUser, nU *int) {
	var cekUsername bool

	fmt.Println("\n^+---------------------------------------------^+")
	fmt.Println("^|                   Sign-Up                   ^|")
	fmt.Println("^+---------------------------------------------^+")

	var username, password string
	fmt.Print("Create Username: ")
	fmt.Scan(&username)

	fmt.Print("Create Password: ")
	fmt.Scan(&password)

	cekUsername = usernameExist(*U, *nU, username)
	if cekUsername == true {
		fmt.Println("Username telah ada. Gunakan username lain. ")
		return
	}

	U[*nU].Username = username
	U[*nU].Password = password
	U[*nU].isStudent = true
	*nU++

	fmt.Println("Sign-up berhasil. Silahkan sign-in untuk melanjutkan")

}

func signIn(U arrUser, nU int) int {
	// Algoritma Sequential Search
	var found int = -1

	fmt.Println("\n^+---------------------------------------------^+")
	fmt.Println("^|                   Sign-In                   ^|")
	fmt.Println("^+---------------------------------------------^+")

	var username, password string
	fmt.Print("Username: ")
	fmt.Scan(&username)

	fmt.Print("Password: ")
	fmt.Scan(&password)

	for i := 0; i < nU; i++ {
		if U[i].Username == username && U[i].Password == password {
			found = i
		}
	}
	return found

}

func usernameExist(U arrUser, nU int, xUser string) bool {
	for i := 0; i < nU; i++ {
		if U[i].Username == xUser {
			return true
		}
	}
	return false
}

func loginSystem_asMhs(P arrJurusan, M arrMhs, nP int, nM int) {
	var opsi string
	var U arrUser
	var cekUSer, nU int

	for opsi != "3" {

		fmt.Println("\n^+---------------------------------------------^+")
		fmt.Println("^|                MENU MAHASISWA               ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 1. Sign-Up                                  ^|")
		fmt.Println("^| 2. Sign-In                                  ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 3. Keluar                                   ^|")
		fmt.Println("^+---------------------------------------------^+")

		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&opsi)

		for opsi != "1" && opsi != "2" && opsi != "3" {
			fmt.Print("Input tidak valid, Pilihan Anda : ")
			fmt.Scan(&opsi)
		}

		if opsi == "1" {
			signUp(&U, &nU)
		} else if opsi == "2" {
			cekUSer = signIn(U, nU)
			if cekUSer != -1 {
				fmt.Println(U[nU].Username)
				fmt.Printf("Sign-In berhasil. Selamat Datang %s!\n", U[cekUSer].Username)

				menuMahasiswa(P, M, nP, nM)

			} else {
				fmt.Println("Username atau Password salah. Silahkan coba lagi atau buat akun baru pada menu Sign-Up. ")
			}
		}

	}
	fmt.Println("Keluar dari Mahasiswa")
}

func menuMahasiswa(P arrJurusan, M arrMhs, nP int, nM int) {
	var nMhs int
	var opsi string

	fmt.Println("\n^+---------------------------------------------^+")
	fmt.Println("^|           SELAMAT DATANG DI PORTAL          ^|")
	fmt.Println("^|          PENDAFTARAN MAHASISWA BARU         ^|")
	fmt.Println("^|       INSTITUT TEKNNOLOGI BOJONGSOANG       ^|")
	fmt.Println("^+---------------------------------------------^+")

	for opsi != "3" {
		fmt.Println("\n^+---------------------------------------------^+")
		fmt.Println("^|                HALAMAN UTAMA                ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 1. Daftar SMB IT Bojongosoang               ^|")
		fmt.Println("^| 2. Laporan Kelulusan                        ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 3. Keluar                                   ^|")
		fmt.Println("^+---------------------------------------------^+")

		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&opsi)

		for opsi != "1" && opsi != "2" && opsi != "3" {
			fmt.Print("Input tidak valid, Pilihan Anda : ")
			fmt.Scan(&opsi)
		}

		if opsi == "1" {
			daftar(P, &M, nP, &nM, &nMhs)
		} else if opsi == "2" {
			if nMhs == 0 {
				fmt.Println("Data tidak ditemukan. Silahkan lakukan pendaftaran pada menu Daftar")
			} else {
				kelulusan(&M, P, nM, nP)
			}
		}
	}
	fmt.Println("^+---------------------------------------------^+")
	fmt.Println("^|     ......Exiting from mahasiswa......     	^|")
	fmt.Println("^+---------------------------------------------^+")
}

func daftar(P arrJurusan, M *arrMhs, nP int, nM *int, nUser *int) {
	var opsi string

	for opsi != "3" {
		fmt.Println("\n^+---------------------------------------------^+")
		fmt.Println("^|            SELEKSI MAHASISWA BARU           ^|")
		fmt.Println("^|       INSTITUT TEKNNOLOGI BOJONGSOANG       ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^|               HALAMAN DAFTAR                ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 1. Input Data                               ^|")
		fmt.Println("^| 2. Edit Data                                ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 3. Keluar                                   ^|")
		fmt.Println("^+---------------------------------------------^+")

		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&opsi)

		for opsi != "1" && opsi != "2" && opsi != "3" {
			fmt.Print("Input tidak valid, Pilihan Anda : ")
			fmt.Scan(&opsi)
		}

		if opsi == "1" {
			if *nUser == 0 {
				viewProdi_mhs(P, nP)
				addMhs_byMhs(M, P, nM, nP, nUser)
				fmt.Println("^+---------------------------------------------^+")
				fmt.Println("Data telah berhasil diinputkan, silahkan lihat status kelulusan Anda.")
			} else {
				fmt.Println("Anda telah melakukan Pendaftaran.")
			}

		} else if opsi == "2" {
			if *nUser == 0 {
				fmt.Println("Data Tidak ditemukan. Silahkan Input Data terlebih dahulu")
			} else {
				editMahasiswa_byMhs(M, P, *nM, nP)
			}
		}
	}
}

func seqMhs(M arrMhs, nM int, x string) int {
	var found int = -1
	var i int

	for i <= nM && found == -1 {
		if x == M[i].nama {
			found = i
		}
		i++
	}
	return found
}

func addMhs_byMhs(M *arrMhs, P arrJurusan, nM *int, nP int, nUser *int) {
	var idx_prodi int

	*nM++

	fmt.Println("^+---------------------------------------------^+")
	fmt.Println("^|                INPUT DATA                   ^|")
	fmt.Println("^+---------------------------------------------^+")

	fmt.Print("Nama\t\t\t: ")
	fmt.Scan(&M[*nM].nama)

	fmt.Print("Nilai Matematika\t: ")
	fmt.Scan(&M[*nM].mtk)

	fmt.Print("Nilai B.Indonesia\t: ")
	fmt.Scan(&M[*nM].bindo)

	fmt.Print("Nilai B.Inggris\t\t: ")
	fmt.Scan(&M[*nM].bing)

	fmt.Print("Nilai Fisika\t\t: ")
	fmt.Scan(&M[*nM].fisika)

	fmt.Print("Pilih Jurusan\t\t: ")
	fmt.Scan(&M[*nM].prodi)

	idx_prodi = seqProdi(&P, &nP, M[*nM].prodi)
	for idx_prodi == -1 {
		fmt.Print("Jurusan invalid. Pilih jurusan: ")
		fmt.Scan(&M[*nM].prodi)

		idx_prodi = seqProdi(&P, &nP, M[*nM].prodi)
	}

	*nUser = *nM

}

func editMahasiswa_byMhs(M *arrMhs, P arrJurusan, nM int, nP int) {
	var opsi string

	for opsi != "5" {

		fmt.Println("\n^+---------------------------------------------^+")
		fmt.Println("^|           HALAMAN EDIT MAHASISWA            ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 1. Edit Nama                                ^|")
		fmt.Println("^| 2. Edit Nilai                               ^|")
		fmt.Println("^| 3. Edit Pilihan Jurusan                     ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 4. View Data                                ^|")
		fmt.Println("^| 5. Keluar                                   ^|")
		fmt.Println("^+---------------------------------------------^+")

		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&opsi)

		for opsi != "1" && opsi != "2" && opsi != "3" && opsi != "4" && opsi != "5" {
			fmt.Print("Input tidak valid, Pilihan Anda : ")
			fmt.Scan(&opsi)
		}

		if opsi == "1" {
			viewData_byMhs(*M, nM)
			editNama_byMhs(M, nM)
		} else if opsi == "2" {
			viewData_byMhs(*M, nM)
			editNilai_byMhs(M, nM)
		} else if opsi == "3" {
			viewData_byMhs(*M, nM)
			editPilihanJurusan_byMhs(M, P, nM, nP)
		} else if opsi == "4" {
			viewData_byMhs(*M, nM)
		}
	}
}

func viewData_byMhs(M arrMhs, nM int) {

	fmt.Println("\n^+---------------------------------------------^+")
	fmt.Println("^|             DATA CALON MAHASISWA            ^|")
	fmt.Println("^+---------------------------------------------^+")

	fmt.Println("Nama\t\t\t: ", M[nM].nama)
	fmt.Println("Nilai Matematika\t: ", M[nM].mtk)
	fmt.Println("Nilai B.Indonesia\t: ", M[nM].bindo)
	fmt.Println("Nilai B.Inggris\t\t: ", M[nM].bing)
	fmt.Println("Nilai Fisika\t\t: ", M[nM].fisika)
	fmt.Println("Jurusan\t\t\t: ", M[nM].prodi)
	fmt.Println("^+---------------------------------------------^+")

}

func dataNilai_byMhs(M arrMhs, nM int) {
	fmt.Println("\n^+---------------------------------------------^+")
	fmt.Printf("   Data Nilai %s   \n", M[nM].nama)
	fmt.Println("^+---------------------------------------------^+")

	fmt.Println("Nilai Matematika\t: ", M[nM].mtk)
	fmt.Println("Nilai B.Indonesia\t: ", M[nM].bindo)
	fmt.Println("Nilai B.Inggris\t\t: ", M[nM].bing)
	fmt.Println("Nilai Fisika\t\t: ", M[nM].fisika)
}

func editNama_byMhs(M *arrMhs, nM int) {
	fmt.Print("\nNama Baru\t\t: ")
	fmt.Scan(&M[nM].nama)
	fmt.Println("^+---------------------------------------------^+")
	fmt.Println("^| Nama mahasiswa berhasil diubah              ^|")
	fmt.Println("^+---------------------------------------------^+")
}

func editNilai_byMhs(M *arrMhs, nM int) {
	var opsi string

	for opsi != "5" {

		dataNilai_byMhs(*M, nM)
		fmt.Println("\n^+---------------------------------------------^+")
		fmt.Println("^|              HALAMAN EDIT NILAI             ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 1. Edit Nilai Matematika                    ^|")
		fmt.Println("^| 2. Edit Nilai B.Indonesia                   ^|")
		fmt.Println("^| 3. Edit Nilai B.Inggris                     ^|")
		fmt.Println("^| 4. Edit Nilai Fisika                        ^|")
		fmt.Println("^+---------------------------------------------^+")
		fmt.Println("^| 5. Keluar                                   ^|")
		fmt.Println("^+---------------------------------------------^+")

		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&opsi)

		for opsi != "1" && opsi != "2" && opsi != "3" && opsi != "4" && opsi != "5" {
			fmt.Print("Input tidak valid, Pilihan Anda : ")
			fmt.Scan(&opsi)
		}

		if opsi == "1" {
			fmt.Print("\nNilai Matematika baru\t: ")
			fmt.Scan(&M[nM].mtk)

			fmt.Println("^+---------------------------------------------^+")
			fmt.Println("^| Nilai Matematika berhasil diubah            ^|")
			fmt.Println("^+---------------------------------------------^+")
		} else if opsi == "2" {
			fmt.Print("\nNilai B.Indonesia baru\t: ")
			fmt.Scan(&M[nM].bindo)

			fmt.Println("^+---------------------------------------------^+")
			fmt.Println("^| Nilai B.Indonesia berhasil diubah           ^|")
			fmt.Println("^+---------------------------------------------^+")
		} else if opsi == "3" {
			fmt.Print("\nNilai B.Inggris baru\t: ")
			fmt.Scan(&M[nM].bing)

			fmt.Println("^+---------------------------------------------^+")
			fmt.Println("^| Nilai B.Inggris berhasil diubah             ^|")
			fmt.Println("^+---------------------------------------------^+")
		} else if opsi == "4" {
			fmt.Print("\nNilai Fisika baru\t: ")
			fmt.Scan(&M[nM].fisika)

			fmt.Println("^+---------------------------------------------^+")
			fmt.Println("^| Nilai Fisika berhasil diubah                ^|")
			fmt.Println("^+---------------------------------------------^+")
		}
	}
}

func editPilihanJurusan_byMhs(M *arrMhs, P arrJurusan, nM int, nP int) {
	var idx_prodi int

	viewProdi_mhs(P, nP)
	fmt.Print("Jurusan Baru\t: ")
	fmt.Scan(&M[nM].prodi)

	idx_prodi = seqProdi(&P, &nP, M[nM].prodi)
	for idx_prodi == -1 {
		fmt.Print("\nJurusan invalid.")
		fmt.Print("\nJurusan Baru\t: ")
		fmt.Scan(&M[nM].prodi)

		idx_prodi = seqProdi(&P, &nP, M[nM].prodi)
	}

	fmt.Println("^+---------------------------------------------^+")
	fmt.Println("^| Data jurusan berhasil diubah                ^|")
	fmt.Println("^+---------------------------------------------^+")
}

func kelulusan(M *arrMhs, P arrJurusan, nM int, nP int) {
	var idx_mhs int
	var idx_prodi int

	fmt.Println("\n^+---------------------------------------------^+")
	fmt.Println("^|      LAPORAN KELULUSAN CALON MAHASISWA      ^|")
	fmt.Println("^+---------------------------------------------^+")

	idx_mhs = seqMhs(*M, nM, M[nM].nama)
	idx_prodi = seqProdi(&P, &nP, M[nM].prodi)

	for i := 0; i <= nM; i++ {
		lulus(M, &P, nM, &nP)
	}

	if M[idx_mhs].status == true {
		fmt.Printf("\nSelamat saudara %s dinyatakan lulus di Prodi %s Institut Teknologi Bojongsoang\n", M[idx_mhs].nama, M[idx_mhs].prodi)

		P[idx_prodi].terisi++
	} else {
		fmt.Println("Mohon maaf anda tidak lulus. Tetap Semangat dan Jangan Putus Asa")
	}
}

func viewProdi_mhs(P arrJurusan, nP int) {
	var i int

	fmt.Println("\n^+---------------------------------------------^+")
	fmt.Println("^|        DAFTAR JURUSAN IT BOJONGSOANG        ^|")
	fmt.Println("^+---------------------------------------------^+")
	for i < nP {
		fmt.Println(P[i].namaProdi, P[i].akreditasi)
		i++
	}
	fmt.Println("")

}

func main() {
	menu()
}

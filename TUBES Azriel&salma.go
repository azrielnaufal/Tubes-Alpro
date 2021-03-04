package main

import "fmt"

type dataUser struct {
	nama       string
	no_telepon string
	umur       int
	profesi    string
}

const n int = 100

type Tuser struct {
	mUser [n]dataUser
	i     int
}

var user Tuser

///////////////////////////////////////////////////////////////////////////////////

type dataGroup struct {
	nama  string
	mUser [n]dataUser
	i     int
}

type Tgrup struct {
	mGrup [n]dataGroup
	y     int
}

var grup Tgrup

type Tpassword struct {
	email    [n]string
	password [n]string
	y        int
}

var pass Tpassword

var idx int
var nama, find string
var found bool

func tambah_user(userr *Tuser) {
	//MASUKKIN DATA USER BARU
	var eop, found bool
	var jwb string
	var nama, nomor, profesi string
	var umur, idx int

	eop = false
	for !(eop) {
		fmt.Print("\ntambah user : lanjut/tidak \n")
		fmt.Scan(&jwb)
		if jwb == "lanjut" {
			fmt.Print("masukkan nama :")
			fmt.Scan(&nama)
			fmt.Print("masukkan nomor telepon :")
			fmt.Scan(&nomor)
			fmt.Print("masukkan umur :")
			fmt.Scan(&umur)
			fmt.Print("masukkan profesi :")
			fmt.Scan(&profesi)
			cari_User1(*userr, nama, &idx, &found)
			if idx >= 0 {
				if nomor == userr.mUser[idx].no_telepon || (umur == userr.mUser[idx].umur && profesi == userr.mUser[idx].profesi) {
					fmt.Print("USER SUDAH ADA")
				} else {
					userr.mUser[userr.i].nama = nama
					userr.mUser[userr.i].no_telepon = nomor
					userr.mUser[userr.i].umur = umur
					userr.mUser[userr.i].profesi = profesi
					userr.i++
				}
			} else {
				userr.mUser[userr.i].nama = nama
				userr.mUser[userr.i].no_telepon = nomor
				userr.mUser[userr.i].umur = umur
				userr.mUser[userr.i].profesi = profesi
				userr.i++
			}
		} else if jwb == "tidak" {
			eop = true
		}
	}
	menu_utama()
}

func insert_buatGrup(grupp *Tgrup) {
	var eop bool
	var jwb string

	eop = false
	for !(eop) {
		fmt.Print("buat grup : lanjut/tidak \n")
		fmt.Scan(&jwb)
		if jwb == "lanjut" {
			fmt.Print("masukkan nama grup :")
			fmt.Scan(&grupp.mGrup[grupp.y].nama)
			grupp.y++
		} else if jwb == "tidak" {
			eop = true
		}
	}
	menu_utama()
}

func cari_User1(userr Tuser, nama string, index *int, found *bool) {
	var i int

	*index = -1
	*found = false
	i = 0
	for !(*found) && i <= userr.i {
		if userr.mUser[i].nama == nama {
			*found = true
			*index = i
		} else {
			i++
		}
	}
}

func cari_Grup(groupp Tgrup, cari string) bool {
	var found bool
	var i int

	found = false
	i = 0
	for !(found) && i <= groupp.y {
		if groupp.mGrup[i].nama == cari {
			found = true
		} else {
			i++
		}
	}
	return found
}

func viewUser(userr Tuser) {
	var i int

	i = 0
	for i < userr.i {
		fmt.Print("user ke-", i, ":\n")
		fmt.Print(userr.mUser[i].nama, "\n", userr.mUser[i].no_telepon, "\n", userr.mUser[i].umur, "\n", userr.mUser[i].profesi, "\n")
		fmt.Print("\n")
		i++
	}
	menu_utama()
}

func viewGrup(grupp Tgrup) {
	//////////////prosedur viewgrup untuk dipakai di menu utama
	var i int
	i = 0
	for i < grupp.y {
		fmt.Print("grup ke-", i, ":\n")
		fmt.Print(grupp.mGrup[i].nama, "\n")
		fmt.Print("\n")
		i++
	}
	menu_utama()
}

func viewGrup1(grupp Tgrup) {
	//////////////////prosedur viewgrup1 untuk dipakai di prsedur insert_inviteGrup
	var i int
	i = 0
	for i < grupp.y {
		fmt.Print("grup ke-", i, ":\n")
		fmt.Print(grupp.mGrup[i].nama, "\n")
		fmt.Print("\n")
		i++
	}
}

func view_User_dalamGrup(grupp Tgrup) {
	var x, y int

	x = 0
	for x < grupp.y {
		fmt.Print("grup ", grupp.mGrup[x].nama, " anggotanya :\n")
		y = 0
		for y < grupp.mGrup[x].i {
			if grupp.mGrup[x].mUser[y].nama != " " {
				fmt.Print(grupp.mGrup[x].mUser[y].nama, "\n")
			} else {
				fmt.Print("-\n")
			}
			y++
		}
		x++
	}
	menu_utama()
}

func insert_inviteGrup(Userr Tuser, Grupp *Tgrup) {
	var idxU, grp int
	var eof, found bool
	var jwb, nama string

	eof = false
	//Grupp.mGrup[grp].i = 0
	for !(eof) {
		fmt.Print("invite grup : lanjut/tidak \n")
		fmt.Scan(&jwb)
		if jwb == "lanjut" {
			fmt.Print("cari nama siapa :")
			fmt.Scan(&nama)
			cari_User1(Userr, nama, &idxU, &found)
			if found {
				fmt.Print(nama, " ada")
				fmt.Print("\nUndang ke mana: \n")
				viewGrup1(*Grupp)
				fmt.Print("Grup ke \n")
				fmt.Scan(&grp)
				Grupp.mGrup[grp].mUser[Grupp.mGrup[grp].i].nama = Userr.mUser[idxU].nama
				Grupp.mGrup[grp].mUser[Grupp.mGrup[grp].i].no_telepon = Userr.mUser[idxU].no_telepon
				Grupp.mGrup[grp].mUser[Grupp.mGrup[grp].i].umur = Userr.mUser[idxU].umur
				Grupp.mGrup[grp].mUser[Grupp.mGrup[grp].i].profesi = Userr.mUser[idxU].profesi
				//fmt.Print(Grupp.mGrup[grp].mUser[Grupp.mGrup[grp].i].nama)
				Grupp.mGrup[grp].i++
			} else {
				fmt.Print(nama, "tidak ada\n")
			}

		} else if jwb == "tidak" {
			eof = true
		}
	}
	menu_utama()
}

func sorting_umur(userr Tuser) {
	// memakai selection
	var pass, i, idx_max, idx_min int
	var temp, pilihan, temp1 int
	var Tnama, Tnama1 string

	fmt.Print("\n1. sorting umur descending(dari tertua)\n2. sorting umur ascending(dari termuda)\n")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		pass = 0
		fmt.Print("hasil sorting(descending umur) : \n")
		for pass <= userr.i-1 {
			idx_max = pass
			i = pass + 1
			for i <= userr.i {
				if userr.mUser[idx_max].umur <= userr.mUser[i].umur {
					idx_max = i
				}
				i++
			}
			temp = userr.mUser[idx_max].umur
			Tnama = userr.mUser[idx_max].nama
			userr.mUser[idx_max].umur = userr.mUser[pass].umur
			userr.mUser[idx_max].nama = userr.mUser[pass].nama
			userr.mUser[pass].umur = temp
			userr.mUser[pass].nama = Tnama
			//fmt.Print("hasil sorting :")
			fmt.Print(userr.mUser[pass].nama, " : ", userr.mUser[pass].umur, " tahun \n")
			pass++
		}
	} else if pilihan == 2 {
		pass = 0
		fmt.Print("hasil sorting(ascending umur) : \n")
		for pass <= userr.i-1 {
			idx_min = pass
			i = pass + 1
			for i <= userr.i {
				if userr.mUser[idx_min].umur > userr.mUser[i].umur {
					idx_min = i
				}
				i++
			}
			temp1 = userr.mUser[idx_min].umur
			Tnama1 = userr.mUser[idx_min].nama
			userr.mUser[idx_min].umur = userr.mUser[pass].umur
			userr.mUser[idx_min].nama = userr.mUser[pass].nama
			userr.mUser[pass].umur = temp1
			userr.mUser[pass].nama = Tnama1
			//fmt.Print("hasil sorting :")
			fmt.Print(userr.mUser[pass].nama, " : ", userr.mUser[pass].umur, " tahun \n")
			pass++
		}
		fmt.Print(userr.mUser[pass].nama, " : ", userr.mUser[pass].umur, " tahun \n")
	}
	menu_utama()
}

func sorting_namaUser(userr Tuser) {
	// memakai insertion  (dari a ke z)
	var pass, i, pilihan int
	var temp, temp1 string

	fmt.Print("1. Sorting nama user (ascending, dari a ke z) \n2. sorting nama user (descending dari z ke a) \n")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		pass = 0
		for pass <= userr.i-1 {
			i = pass + 1
			temp = userr.mUser[i].nama
			for i >= 1 && temp < userr.mUser[i-1].nama {
				userr.mUser[i].nama = userr.mUser[i-1].nama
				i = i - 1
			}
			userr.mUser[i].nama = temp
			pass++
		}
		fmt.Print("hasil sorting (dari a ke z) :\n")
		for y := 0; y <= userr.i; y++ {
			fmt.Print(userr.mUser[y].nama, "\n")
		}
	} else if pilihan == 2 {
		pass = 0
		for pass <= userr.i-1 {
			i = pass + 1
			temp1 = userr.mUser[i].nama
			for i >= 1 && temp1 > userr.mUser[i-1].nama {
				userr.mUser[i].nama = userr.mUser[i-1].nama
				i = i - 1
			}
			userr.mUser[i].nama = temp1
			pass++
		}
		fmt.Print("hasil sorting (dari z ke a) :\n")
		for y := 0; y <= userr.i; y++ {
			fmt.Print(userr.mUser[y].nama, "\n")
		}
	}
	menu_utama()
}

func edit(userr *Tuser) {
	var eof, found bool
	var idx int
	var jwb, nama string
	var jwb1 int

	eof = false
	for !(eof) {
		fmt.Print("lanjut/ tidak : \n")
		fmt.Scan(&jwb)
		if jwb == "lanjut" {
			fmt.Print("ingin edit data siapa : \n")
			fmt.Scan(&nama)
			cari_User1(*userr, nama, &idx, &found)
			if found {
				fmt.Print(nama, " ada")
				fmt.Print("\nedit: \n1.nama \n2.nomor telepon \n3.umur \n4.profesi \n")
				fmt.Scan(&jwb1)
				switch jwb1 {
				case 1:
					fmt.Print("ubah nama : \n")
					fmt.Scan(&userr.mUser[idx].nama)
				case 2:
					fmt.Print("ubah nomor telepon : \n")
					fmt.Scan(&userr.mUser[idx].no_telepon)
				case 3:
					fmt.Print("ubah umur : \n")
					fmt.Scan(&userr.mUser[idx].umur)
				case 4:
					fmt.Print("ubah profesi : \n")
					fmt.Scan(&userr.mUser[idx].profesi)
				}
			} else {
				fmt.Print(nama, "tidak ada\n")
			}
		} else if jwb == "tidak" {
			eof = true
		}
	}
	menu_utama()
}

func deleteArray(userr *Tuser) {
	var i, y int

	i = 0
	y = 0
	for i <= userr.i {
		if userr.mUser[i].umur != 0 {
			userr.mUser[y] = userr.mUser[i]
			y++
		}
		i++
	}
}

func delete(userr *Tuser) {
	var eof, found bool
	var idx int
	var jwb string

	eof = false
	for !(eof) {
		fmt.Print("lanjut/ tidak : \n")
		fmt.Scan(&jwb)
		if jwb == "lanjut" {
			fmt.Print("ingin delete data siapa : \n")
			fmt.Scan(&nama)
			cari_User1(*userr, nama, &idx, &found)
			if found {
				fmt.Print(nama, " sudah terhapus\n")
				userr.mUser[idx].nama = " "
				userr.mUser[idx].no_telepon = " "
				userr.mUser[idx].umur = 0
				userr.mUser[idx].profesi = " "
				deleteArray(&*userr)
			} else {
				fmt.Print(nama, "tidak ada\n")
			}

		} else if jwb == "tidak" {
			eof = true
		}
	}
	menu_utama()
}

func menu_insert() {
	var a int
	fmt.Print("\n===================menu insert========================\n")
	fmt.Print("1. invite user ke grup\n")
	fmt.Print("2. buat grup\n")
	fmt.Print("3. balik ke menu utama\n")
	fmt.Print("menu : ")
	fmt.Scan(&a)
	if a == 1 {
		insert_inviteGrup(user, &grup)
	} else if a == 2 {
		insert_buatGrup(&grup)
	} else if a == 3 {
		menu_utama()
	}
}

func menu_view() {
	var a int
	fmt.Print("\n=====================menu view==================================\n")
	fmt.Print("1. view user\n")
	fmt.Print("2. view grup\n")
	fmt.Print("3. view user dalam setiap grup\n")
	fmt.Print("4. sorting umur (menggunakan selection sort)\n")
	fmt.Print("5. sorting nama user (menggunakan insertion sort)\n")
	fmt.Print("6. edit\n")
	fmt.Print("7. delete\n")
	fmt.Print("8. cari user\n")
	fmt.Print("9. cari grup\n")
	fmt.Print("10. balik ke menu utama\n")
	fmt.Print("menu :")
	fmt.Scan(&a)
	switch a {
	case 1:
		viewUser(user)
	case 2:
		viewGrup(grup)
	case 3:
		view_User_dalamGrup(grup)
	case 4:
		sorting_umur(user)
	case 5:
		sorting_namaUser(user)
	case 6:
		edit(&user)
	case 7:
		delete(&user)
	case 8:
		fmt.Print("ingin cari siapa :")
		fmt.Scan(&nama)
		cari_User1(user, nama, &idx, &found)
		//fmt.Print("index = " , idx)
		if found {
			fmt.Print(nama, " ada")
		} else {
			fmt.Print("tidak ada")
		}
		menu_utama()
	case 9:
		fmt.Print("nama grup yang dicari :")
		fmt.Scan(&find)
		if cari_Grup(grup, find) {
			fmt.Print(find, " ada")
		} else {
			fmt.Print("tidak ada ")
		}
		menu_utama()
	case 10:
		menu_utama()
	}

}

func menu_utama() {
	var a int
	fmt.Print("\n========================================== welcome to main menu ==================================\n")
	fmt.Print("silahkan pilih: \n")
	fmt.Print("1. menu insert\n")
	fmt.Print("2. menu view\n")
	fmt.Print("3. menu tambah_user\n")
	fmt.Print("4. exit\n")
	fmt.Print("menu :")
	fmt.Scan(&a)
	if a == 1 {
		menu_insert()
	} else if a == 2 {
		menu_view()
	} else if a == 3 {
		tambah_user(&user)
	} else if a == 4 {
		menu_login()
	}
}

func menu_login() {
	var a int
	var email, pw string
	var hasil bool

	fmt.Print("\n=============================== welcome to login menu ==================================\n")
	fmt.Print("silahkan pilih\n")
	fmt.Print("1. login:\n")
	fmt.Print("2. daftar\n")
	fmt.Print("3. exit\n")
	fmt.Print("menu :\n")
	fmt.Scan(&a)
	if a == 1 {
		fmt.Print("masukkan email :\n")
		fmt.Scan(&email)
		fmt.Print("masukkan password: \n")
		fmt.Scan(&pw)
		cek_password(email, pw, &hasil, pass)
		if hasil {
			fmt.Print("email dan password and benar")
			menu_utama()
		} else {
			fmt.Print("email dan password salah, silahkan coba lagi\n")
			menu_login()
		}
	} else if a == 2 {
		daftar(&pass)
	} else if a == 3 {
		fmt.Print("aight mate.... thank u for using our apps, see you again\n")
	}
}

func cek_password(id string, pass string, eof *bool, Tpass Tpassword) {
	var i int

	*eof = false
	i = 0
	for !(*eof) && i <= Tpass.y {
		if Tpass.email[i] == id && Tpass.password[i] == pass {
			*eof = true
		} else {
			i++
		}
	}
}

func daftar(Tpass *Tpassword) {
	var eof bool
	var jwb string

	for !(eof) {
		fmt.Print("\ndaftar : lanjut/tidak\n")
		fmt.Scan(&jwb)
		if jwb == "lanjut" {
			fmt.Print("email :\n")
			fmt.Scan(&Tpass.email[Tpass.y])
			fmt.Print("password :\n")
			fmt.Scan(&Tpass.password[Tpass.y])
			Tpass.y++
		} else if jwb == "tidak" {
			eof = true
		}
	}
	menu_login()
}

func main() {

	/////////////////////////////////////////////////////////////////////////////
	user.mUser[0].nama = "azriel"
	user.mUser[0].no_telepon = "123"
	user.mUser[0].umur = 18
	user.mUser[0].profesi = "mahasiswa"
	//////////////////////////////////////////////////////////////////////////////////////////
	user.mUser[1].nama = "salma"
	user.mUser[1].no_telepon = "234"
	user.mUser[1].umur = 18
	user.mUser[1].profesi = "mahasiswi"
	/////////////////////////////////////////////////////////////////////////////
	user.mUser[2].nama = "khoiru"
	user.mUser[2].no_telepon = "345"
	user.mUser[2].umur = 20
	user.mUser[2].profesi = "mahasiswa"
	///////////////////////////////////////////////////////////////////////
	user.mUser[3].nama = "adhi"
	user.mUser[3].no_telepon = "456"
	user.mUser[3].umur = 19
	user.mUser[3].profesi = "mahasiswa"
	///////////////////////////////////////////////////////////////////////////////////////////////////////////
	user.mUser[4].nama = "haviza"
	user.mUser[4].no_telepon = "567"
	user.mUser[4].umur = 19
	user.mUser[4].profesi = "mahasiswi"
	///////////////////////////////////////////////////////////////////////////////////////////////////////
	user.mUser[5].nama = "rahmatia"
	user.mUser[5].no_telepon = "678"
	user.mUser[5].umur = 17
	user.mUser[5].profesi = "pelajar"
	//////////////////////////////////////////////////////////////////////////////////////////////////
	user.mUser[6].nama = "andra"
	user.mUser[6].no_telepon = "789"
	user.mUser[6].umur = 16
	user.mUser[6].profesi = "mahasiswa"
	///////////////////////////////////////////////////////////////////////
	user.mUser[7].nama = "aurora"
	user.mUser[7].no_telepon = "890"
	user.mUser[7].umur = 23
	user.mUser[7].profesi = "aktris"
	///////////////////////////////////////////////////////////////////////////////////////////////////////////
	user.mUser[8].nama = "kylie"
	user.mUser[8].no_telepon = "900"
	user.mUser[8].umur = 25
	user.mUser[8].profesi = "model"
	///////////////////////////////////////////////////////////////////////////////////////////////////////
	user.mUser[9].nama = "lebron"
	user.mUser[9].no_telepon = "678"
	user.mUser[9].umur = 37
	user.mUser[9].profesi = "atlit"
	//////////////////////////////////////////////////////////////////////////////////////////////////
	grup.mGrup[0].nama = "IF-43-01"
	grup.mGrup[1].nama = "IF-43-02"
	grup.mGrup[2].nama = "IF-43-03"
	grup.mGrup[3].nama = "IF-43-04"
	grup.mGrup[4].nama = "IF-43-05"
	grup.mGrup[5].nama = "IF-43-06"
	grup.mGrup[6].nama = "IF-43-07"
	grup.mGrup[7].nama = "IF-43-08"
	grup.mGrup[8].nama = "IF-43-09"
	grup.mGrup[9].nama = "IF-43-10"
	////////////////////////////////////////////////////////////////////////////////////////////////////////
	pass.email[0] = "azrielnaufal31"
	pass.password[0] = "12345"
	pass.email[1] = "salmasalsabila"
	pass.password[1] = "salma123"

	user.i = 10
	grup.y = 10
	pass.y = 2
	menu_login()
	//////////////////////////////////////////////////////////////////////////////////////////////////
	//tambah_user(&user)
	//insert_buatGrup(&grup)
	//cari_User(user)
	//cari_Grup(grup)
	//viewUser(user)
	//viewGrup(grup)
	//insert_inviteGrup(user,&grup)
	//view_User_dalamGrup(grup)
	//sorting_umur(user)
	//sorting_namaUser(user)
}

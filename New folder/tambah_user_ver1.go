package main	
import "fmt"

 type dataUser struct {
 	nama string
 	no_telepon string
 	umur int
 	profesi string
}
	const n int = 100

type Tuser struct{
	mUser [n]dataUser
	i int
}
	var user Tuser
///////////////////////////////////////////////////////////////////////////////////

type dataGroup struct {
	nama string
	mUser [n]dataUser
	i int
}

type Tgrup struct {
	mGrup [n]dataGroup
	y int
}
	var grup Tgrup
	var idx int


func tambah_user (userr *Tuser) {
//MASUKKIN DATA USER BARU
	var eop bool
	var jwb string

	eop = false 
	for !(eop) {
	fmt.Print("tambah user : lanjut/tidak \n")
	fmt.Scan(&jwb)
		if jwb == "lanjut"{
			fmt.Print("masukkan nama :")
			fmt.Scan(&userr.mUser[userr.i].nama)
			fmt.Print("masukkan nomor telepon :")
			fmt.Scan(&userr.mUser[userr.i].no_telepon)
			fmt.Print("masukkan umur :")
			fmt.Scan(&userr.mUser[userr.i].umur)
			fmt.Print("masukkan profesi :")
			fmt.Scan(&userr.mUser[userr.i].profesi)
			userr.i++
		} else if jwb == "tidak" {
			eop = true
		}
	}
	menu_utama()
}

func insert_buatGrup (grupp *Tgrup) {
	var eop bool
	var jwb string

	eop = false 
	for !(eop) {
	fmt.Print("buat grup : lanjut/tidak \n")
	fmt.Scan(&jwb)
		if jwb == "lanjut"{
			fmt.Print("masukkan nama grup :")
			fmt.Scan(&grupp.mGrup[grupp.y].nama)
			grupp.y++
		} else if jwb == "tidak" {
			eop = true
		}
}
	menu_utama()
}

func cari_User (userr Tuser , index *int) {
	var found bool
	var cari string
	var i int

	found = false
	i = 0
	fmt.Print("nama yang dicari :")
	fmt.Scan(&cari)
	for !(found) && i <= userr.i  {
		if userr.mUser[i].nama == cari {
			found = true
			*index = i
		} else {
			i++
		}
	}
	if found {
		fmt.Print(userr.mUser[*index].nama , " ada")
	} else {
		fmt.Print("tidak ada")
	}
	menu_utama()
}
func cari_User1 (userr Tuser , index *int) {
	var found bool
	var cari string
	var i int

	found = false
	i = 0
	fmt.Print("nama :")
	fmt.Scan(&cari)
	for !(found) && i <= userr.i  {
		if userr.mUser[i].nama == cari {
			found = true
			*index = i
		} else {
			i++
		}
	}
	if found {
		fmt.Print(userr.mUser[*index].nama , " ada")
	} else {
		fmt.Print("tidak ada")
	}
}

func cari_Grup (groupp Tgrup) {
	var found bool
	var cari string
	var idx, i int

	found = false
	i=0
	fmt.Print("grup yang dicari: ")
	fmt.Scan(&cari)
	for !(found) && i <= groupp.y{
		if groupp.mGrup[i].nama==cari {
			found=true
			idx=i
		} else{
			i++
		}
	}
	if found {
		fmt.Print(groupp.mGrup[idx].nama , " ada")
	} else {
		fmt.Print("tidak ada")
	}
	menu_utama()
}

func viewUser (userr Tuser) {
	var i int

	i = 0
	for  i < userr.i {
		fmt.Print("user ke-",i,":\n")
		fmt.Print(userr.mUser[i].nama,"\n",userr.mUser[i].no_telepon,"\n",userr.mUser[i].umur,"\n",userr.mUser[i].profesi,"\n")
		fmt.Print("\n")
	i++
	}
	menu_utama()
}

func viewGrup (grupp Tgrup) {
	var i int
	i = 0
	for  i < grupp.y {
		fmt.Print("grup ke-",i,":\n")
		fmt.Print(grupp.mGrup[i].nama,"\n")
		fmt.Print("\n")
	i++
	}
	menu_utama()
}

func viewGrup1 (grupp Tgrup) {
	var i int
	i = 0
	for  i < grupp.y {
		fmt.Print("grup ke-",i,":\n")
		fmt.Print(grupp.mGrup[i].nama,"\n")
		fmt.Print("\n")
	i++
	}
}

func view_User_dalamGrup (grupp Tgrup) {
		var x,y int

	x = 0
	for  x < grupp.y {
		fmt.Print("grup ",grupp.mGrup[x].nama," anggotanya :\n")
			y = 0
			for y < grupp.mGrup[x].i {
				if grupp.mGrup[x].mUser[y].nama != " " {
					fmt.Print(grupp.mGrup[x].mUser[y].nama,"\n")
				} else {
					fmt.Print("-\n")
				}
			y++
			}
	x++
	}
	menu_utama()
}

func insert_inviteGrup (Userr Tuser , Grupp *Tgrup) {	
	var idxU,grp int
	var eof bool
	var jwb string
	eof = false
	Grupp.mGrup[grp].i = 0
	for !(eof) {
	fmt.Print("invite grup : lanjut/tidak \n")
	fmt.Scan(&jwb)
		if jwb == "lanjut" {
			cari_User1(Userr,&idxU)
			fmt.Print("\nUndang ke mana: \n")
			viewGrup1(*Grupp)
			fmt.Print("Grup ke-")
			fmt.Scan(&grp)
			Grupp.mGrup[grp].mUser[Grupp.mGrup[grp].i].nama = Userr.mUser[idxU].nama
			Grupp.mGrup[grp].mUser[Grupp.mGrup[grp].i].no_telepon = Userr.mUser[idxU].no_telepon
			Grupp.mGrup[grp].mUser[Grupp.mGrup[grp].i].umur = Userr.mUser[idxU].umur
			Grupp.mGrup[grp].mUser[Grupp.mGrup[grp].i].profesi = Userr.mUser[idxU].profesi
			//fmt.Print(Grupp.mGrup[grp].mUser[Grupp.mGrup[grp].i].nama)
			Grupp.mGrup[grp].i++
		} else if jwb == "tidak" {
			eof = true
		}
	}
	menu_utama()
}

func sorting_umur (userr Tuser) {
// memakai selection dan descending (dari umur tertua ke termuda)
	var pass,i,idx_max int
	var temp int
	var Tnama string

	pass = 0
	fmt.Print("hasil sorting(descending umur) : \n")
	for pass <= userr.i - 1 {
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
	fmt.Print(userr.mUser[pass].nama ," : ", userr.mUser[pass].umur,"tahun \n")
	pass++
	}
	menu_utama()
}

func sorting_namaUser (userr Tuser) {
// memakai insertion dan ascending (dari a ke z)
	var pass,i int
	var temp string

	pass = 0
	fmt.Print("hasil sorting (dari a ke z) :\n")
	for pass <= userr.i - 1 {
		i = pass + 1
		temp = userr.mUser[i].nama
		for i > 1 && temp < userr.mUser[i-1].nama {
			userr.mUser[i].nama = userr.mUser[i-1].nama
			i = i - 1
		}
		userr.mUser[i].nama = temp
		pass++
	}
	for y:= 0 ; y <= userr.i ; y++ {
		fmt.Println(userr.mUser[y].nama)
	}
	menu_utama()
}


func edit (userr *Tuser) {
	var eof bool
	var idx int
	var jwb string
	var jwb1 int

	eof = false
	for !(eof) {
		fmt.Print("lanjut/ tidak : \n")
		fmt.Scan(&jwb)
		if jwb == "lanjut" {
			fmt.Print("ingin edit data siapa : \n")
			cari_User1(*userr,&idx)
			fmt.Print("\nedit: \n1.nama \n2.nomor telepon \n3.umur \n4.profesi \n")
			fmt.Scan(&jwb1)
			switch jwb1 {
			case 1 :
				fmt.Print("ubah nama : \n")
				fmt.Scan(&userr.mUser[idx].nama)
			case 2 :
				fmt.Print("ubah nomor telepon : \n")
				fmt.Scan(&userr.mUser[idx].no_telepon)
			case 3 :
				fmt.Print("ubah umur : \n")
				fmt.Scan(&userr.mUser[idx].umur)
			case 4 :
				fmt.Print("ubah profesi : \n")
				fmt.Scan(&userr.mUser[idx].profesi)
			}
		} else if jwb == "tidak" {
			eof = true
		}
	}
	menu_utama()
}

func delete (userr *Tuser) {
	var eof bool
	var idx int
	var jwb string
	var jwb1 int

	eof = false
	for !(eof) {
		fmt.Print("lanjut/ tidak : \n")
		fmt.Scan(&jwb)
		if jwb == "lanjut" {
			fmt.Print("ingin delete data siapa : \n")
			cari_User1(*userr,&idx)
			fmt.Print("\ndelete: \n1.nama \n2.nomor telepon \n3.umur \n4.profesi \n")
			fmt.Scan(&jwb1)
			switch jwb1 {
			case 1 :
				userr.mUser[idx].nama = " "
			case 2 :
				userr.mUser[idx].no_telepon = " "
			case 3 :
				userr.mUser[idx].umur = 0
			case 4 :
				userr.mUser[idx].profesi = " "
			}
		} else if jwb == "tidak" {
			eof = true
		}
	}
	menu_utama()
}

func menu_insert () {
	var a int
	fmt.Print("\n===================menu insert========================\n")
	fmt.Print("1. invite user ke grup\n")
	fmt.Print("2. buat grup\n")
	fmt.Print("3. balik ke menu utama\n")
	fmt.Print("menu : ")
	fmt.Scan(&a)
	if a == 1 {
		insert_inviteGrup(user,&grup)
	} else if a == 2 {
		insert_buatGrup(&grup)
	} else if a == 3 {
		menu_utama()
	}
}

func menu_view () {
	var a int
	fmt.Print("\n=====================menu view==================================\n")
	fmt.Print("1. view user\n")
	fmt.Print("2. view grup\n")
	fmt.Print("3. view user dalam setiap grup\n")
	fmt.Print("4. sorting umur user secara descending menggunakan selection sort\n")
	fmt.Print("5. sorting nama user secara ascending(a ke z) menggunakan insertion sort\n")
	fmt.Print("6. edit\n")
	fmt.Print("7. delete\n")
	fmt.Print("8. cari user\n")
	fmt.Print("9. cari grup\n")
	fmt.Print("10. balik ke menu utama\n")
	fmt.Print("menu :")
	fmt.Scan(&a)
	switch a {
	case 1 : 
		viewUser(user)
	case 2 :
		viewGrup(grup)
	case 3 :
		view_User_dalamGrup(grup)
	case 4 :
		sorting_umur(user)
	case 5 :
		sorting_namaUser(user)
	case 6 : 
		edit(&user)
	case 7 :
		delete(&user)
	case 8 :
		cari_User(user,&idx)
	case 9 :
		cari_Grup(grup)
	case 10 :
		menu_utama()
	} 

}

func menu_utama () {
	var a int
	fmt.Print("\n========================================== WELCOME ==================================\n")
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
	} else if  a == 3 {
		tambah_user(&user)
	} else if a == 4 {
		fmt.Print("aight mate.... see you again\n")
	}

}

func main () {

	menu_utama()
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
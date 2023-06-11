package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== To-Do List CLI ===")

	// Membaca input dari pengguna
	reader := bufio.NewReader(os.Stdin)

	// Buatlah slice untuk menyimpan tugas
	tasks := make([]string, 0)

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambahkan tugas")
		fmt.Println("2. Tampilkan daftar tugas")
		fmt.Println("3. Tandai tugas sebagai selesai")
		fmt.Println("4. Hapus tugas")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih menu: ")
		option, _ := reader.ReadString('\n')

		switch option {
		case "1\n":
			fmt.Print("Masukkan tugas: ")
			task, _ := reader.ReadString('\n')
			tasks = append(tasks, task)
			fmt.Println("Tugas berhasil ditambahkan.")

		case "2\n":
			if len(tasks) == 0 {
				fmt.Println("Tidak ada tugas dalam daftar.")
			} else {
				fmt.Println("Daftar Tugas:")
				for i, task := range tasks {
					fmt.Printf("%d. %s", i+1, task)
				}
			}

		case "3\n":
			if len(tasks) == 0 {
				fmt.Println("Tidak ada tugas dalam daftar.")
			} else {
				fmt.Print("Pilih nomor tugas yang telah selesai: ")
				var num int
				fmt.Scanln(&num)
				if num >= 1 && num <= len(tasks) {
					tasks = append(tasks[:num-1], tasks[num:]...)
					fmt.Println("Tugas berhasil ditandai sebagai selesai.")
				} else {
					fmt.Println("Nomor tugas tidak valid.")
				}
			}

		case "4\n":
			if len(tasks) == 0 {
				fmt.Println("Tidak ada tugas dalam daftar.")
			} else {
				fmt.Print("Pilih nomor tugas yang ingin dihapus: ")
				var num int
				fmt.Scanln(&num)
				if num >= 1 && num <= len(tasks) {
					tasks = append(tasks[:num-1], tasks[num:]...)
					fmt.Println("Tugas berhasil dihapus.")
				} else {
					fmt.Println("Nomor tugas tidak valid.")
				}
			}

		case "0\n":
			fmt.Println("Terima kasih telah menggunakan To-Do List CLI.")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

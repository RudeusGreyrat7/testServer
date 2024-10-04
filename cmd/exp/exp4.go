package main

import (
	"fmt"
	"os"

	"github.com/LENSLOCKED/models"
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Current working directory:", dir)

	g := models.GalleryService{
		ImagesDir: "/Users/matvejkuznecenkov/go/src/github.com/LENSLOCKED/images",
	}
	fmt.Println(g.Images(2))
}

// func main() {
// 	// Укажите директорию, где хранятся изображения
// 	g := models.GalleryService{
// 		ImagesDir: "/github.com/LENSLOCKED/images/gallery-2", // Замените на реальный путь
// 	}

// 	// Попробуйте вывести результат функции Images
// 	images, err := g.Images(2)
// 	fmt.Println("Images:", images)
// 	fmt.Println("Error:", err)
// }

package main

import "fmt"

type User struct {
	ID   int
	Name string
	Data []byte // имитация "тяжёлых" данных
}

// removeUser удаляет i-й элемент из слайса указателей на User
// и предотвращает утечку памяти.
func removeUser(users []*User, i int) []*User {
	if i < 0 || i >= len(users) {
		panic("index out of range")
	}

	// Сдвигаем хвост на одну позицию влево
	copy(users[i:], users[i+1:])

	// Обнуляем последний элемент, чтобы GC мог освободить память
	users[len(users)-1] = nil

	// Уменьшаем длину слайса
	return users[:len(users)-1]
}

func main() {
	users := []*User{
		{ID: 1, Name: "Alice", Data: make([]byte, 1024*1024)}, // 1 МБ
		{ID: 2, Name: "Bob", Data: make([]byte, 1024*1024)},
		{ID: 3, Name: "Charlie", Data: make([]byte, 1024*1024)},
	}

	fmt.Println("До удаления:", len(users), "пользователей")

	// Удаляем второго пользователя (индекс 1)
	users = removeUser(users, 1)

	fmt.Println("После удаления:", len(users), "пользователей")
	for _, u := range users {
		fmt.Printf("ID: %d, Name: %s\n", u.ID, u.Name)
	}
	// Теперь объект Bob больше не удерживается в памяти
}

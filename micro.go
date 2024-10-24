package main

import (
	"fmt"
	"github.com/R0n1ns/wireguard_go_ubuntu"
	"log"
	"os"
)

var filename = "data.json"
var wg_client = wireguard_go_ubuntu.WireGuardConfig{}

func init() {
	wg_client.LoadFromFile(filename)
}

// Функция для запроса целого числа от пользователя
func inputInt(prompt string) int {
	var value int
	fmt.Print(prompt)
	fmt.Scanln(&value)
	return value
}

// Функция для запроса строки от пользователя
func inputString(prompt string) string {
	var value string
	fmt.Print(prompt)
	fmt.Scanln(&value)
	return value
}

func handleStartWireguard(wg_client *wireguard_go_ubuntu.WireGuardConfig) {
	tp := inputInt("если хотите автоматическую настройку введите 0, ручной режим 1: ")
	if tp != 0 {
		wg_client.ListenPort = inputString("Введите порт, или 0 для порта 51820: ")
		if wg_client.ListenPort == "0" {
			wg_client.ListenPort = "51820"
		}
		wg_client.Endpoint = inputString("Введите ip сервера: ") + ":" + wg_client.ListenPort
		wg_client.InterName = inputString("Введите имя интерфейса, или 0 для eth0: ")
		if wg_client.InterName == "0" {
			wg_client.InterName = "eth0"
		}
		wg_client.BotToken = inputString("Введите токен бота: ")

		fmt.Println("Параметры успешно заданы")
		wg_client.CollectTraffic()
		wg_client.GenServerKeys()
		fmt.Printf("Созданный приватный ключ: %s\nСозданный публичный ключ: %s\n", wg_client.PrivateKey, wg_client.PublicKey)

		wg_client.GenerateWireGuardConfig()
		wg_client.WireguardStart()
		log.Printf("Соединение wireguard запущено")
	} else {
		wg_client.Autostart()
		wg_client.BotToken = inputString("Введите токен бота: ")
		fmt.Printf("Созданный приватный ключ: %s\nСозданный публичный ключ: %s\n", wg_client.PrivateKey, wg_client.PublicKey)
	}
}

func handleAddClient(wg_client *wireguard_go_ubuntu.WireGuardConfig, id *int) {
	client, _ := wg_client.AddWireguardClient(*id)
	fmt.Printf("Данные клиента:\nАдрес: %s\nПубличный ключ: %s\nПриватный ключ: %s\n",
		client.AddressClient, client.PublicClientKey, client.PrivateClientKey)

	tgID := inputInt("Введите тг id клиента или -1, если отправлять конфиг не нужно: ")
	if tgID != -1 {
		client.TgId = tgID
		wg_client.Clients[*id] = client
		wg_client.SendConfigToUserTg(*id)
		fmt.Println("Данные отправлены")
	}
	(*id)++
}

func handleClientAction(wg_client *wireguard_go_ubuntu.WireGuardConfig, action func(int), prompt string) {
	usID := inputInt(prompt)
	action(usID)
}

func printAllClients(wg_client *wireguard_go_ubuntu.WireGuardConfig) {
	clients := wg_client.AllClients()
	fmt.Println(clients)
}

func handleSendClientData(wg_client *wireguard_go_ubuntu.WireGuardConfig) {
	usID := inputInt("Введите id клиента или -1, если отправлять конфиг не нужно: ")
	if usID != -1 {
		wg_client.SendConfigToUserTg(usID)
		fmt.Println("Данные отправлены")
	}
}

func handleDropWireguard(wg_client *wireguard_go_ubuntu.WireGuardConfig) {
	wg_client.DropWireguard()
	fmt.Println("Конфигурации и скрипт удален")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			// Вывод ошибки
			fmt.Println("Произошла ошибка:", r)
			// Вызов функции снова
			wg_client.SaveToFile(filename)
			main()
		}
	}()
	id := 2
	firstMess := `Доступные команды:
1. Запуск wireguard
2. Добавление клиента
3. Остановка клиента
4. Активация клиента
5. Удаление клиента
6. Все клиенты
7. Отправить данные клиента
0. В меню
-1. Выйти
-2. Удалить и отключить wireguard, выйти`

	fmt.Println(firstMess)

	commands := map[int]func(){
		1: func() { handleStartWireguard(&wg_client) },
		2: func() { handleAddClient(&wg_client, &id) },
		3: func() {
			handleClientAction(&wg_client, wg_client.StopClient, "Введите id клиента для остановки: ")
		},
		4: func() {
			handleClientAction(&wg_client, wg_client.ActClient, "Введите id клиента для активации: ")
		},
		5: func() {
			handleClientAction(&wg_client, wg_client.DeleteClient, "Введите id клиента для удаления: ")
		},
		6:  func() { printAllClients(&wg_client) },
		7:  func() { handleSendClientData(&wg_client) },
		-2: func() { handleDropWireguard(&wg_client) },
		0:  func() { fmt.Println(firstMess) },
		-1: func() { fmt.Println("Скрипт остановлен"); os.Exit(0) },
	}

	for {
		comand := inputInt("Введите команду: ")
		if action, exists := commands[comand]; exists {
			action()
		} else {
			fmt.Println("Команда не найдена\n" + firstMess)
		}
	}
}

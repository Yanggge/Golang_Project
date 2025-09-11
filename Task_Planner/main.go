package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Zadacha struct {
	Theme   string
	Content string
	Time    string
	Status  bool
}

type Events struct {
	Input  string
	Errors string
	Time   string
}

func main() {
	fmt.Println("Планировщик задач")
	DoSlice := make([]Zadacha, 0, 20)
	EventSlice := make([]Events, 0, 50)
	for {
		fmt.Println("Выберите задачу:\n добавить задачу\n удалить задачу\n мои задачи\n изменить статус\n логи\n выход")
		fmt.Print("ввод команды:")
		scanner := bufio.NewScanner(os.Stdin)
		if ok := scanner.Scan(); !ok {
			Logevents(&EventSlice, "", "Ошибка ввода")
			fmt.Println("Ошибка ввода")
		}
		text := scanner.Text()

		Logevents(&EventSlice, text, "")
		if text == "выход" {
			fmt.Println("До встречи")
			break
		}

		switch text {
		case "добавить задачу":
			fmt.Print("Введите заголовок не больше 1 слова: ")
			scanner.Scan()
			theme := scanner.Text()
			fields := strings.Fields(theme)
			if len(fields) != 1 {
				Logevents(&EventSlice, theme, "Некоректный ввод!")
				fmt.Println("Некоректный ввод!")
				continue
			}

			fmt.Print("Введите содержание: ")
			scanner.Scan()
			content := scanner.Text()
			fields2 := strings.Fields(content)
			if len(fields2) < 1 {
				Logevents(&EventSlice, content, "Вы ничего не указали")
				fmt.Println("Вы ничего не указали")
				continue
			}

			status := false

			DoSlice = append(DoSlice, Zadacha{
				Theme:   theme,
				Content: content,
				Time:    time.Now().Format("2006-01-02 15:04:05"),
				Status:  status,
			})

			fmt.Println("========================")
			fmt.Println("Задача № ", len(DoSlice))
			fmt.Println("Заголовок: ", theme)
			fmt.Println("Содержание: ", content)
			fmt.Println("Время создания: ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Println("Статус: ", map[bool]string{true: "выполнена", false: "не выполнена"}[status])
			fmt.Println("========================")

		case "удалить задачу":
			fmt.Print("Удалить задачу № ")
			scanner.Scan()
			del, err := strconv.Atoi(scanner.Text())
			if err != nil || del > len(DoSlice) {
				Logevents(&EventSlice, strconv.Itoa(del), "Вы ввели несуществующую задачу")
				fmt.Println("Вы ввели несуществующую задачу")
			} else {
				DoSlice = append(DoSlice[:del-1], DoSlice[del:]...)
				fmt.Println("Вы успешно удалили задачу № ", del)
			}
		case "мои задачи":
			for i, _ := range DoSlice {
				fmt.Println("========================")
				fmt.Println("Задача № ", i+1)
				fmt.Println("Заголовок: ", DoSlice[i].Theme)
				fmt.Println("Содержание: ", DoSlice[i].Content)
				fmt.Println("Время создания: ", time.Now().Format("2006-01-02 15:04:05"))
				fmt.Println("Статус: ", map[bool]string{true: "выполнена", false: "не выполнена"}[DoSlice[i].Status])
				fmt.Println("========================")
			}
		case "изменить статус":
			StatusEchange(scanner, &DoSlice, &EventSlice)
		case "логи":
			for i, _ := range EventSlice {
				fmt.Println("#", i+1)
				fmt.Println("Input: ", EventSlice[i].Input)
				fmt.Println("Errors: ", EventSlice[i].Errors)
				fmt.Println("Time add: ", EventSlice[i].Time)
			}
		default:
			Logevents(&EventSlice, text, "Вы ввели неправильную команду")
			fmt.Println("Вы ввели неправильную команду")
		}
	}
}

// ////////////////////
func StatusEchange(scanner *bufio.Scanner, slice *[]Zadacha, EvSlice *[]Events) {
	fmt.Print("У какой задачи изменить статус? Задача № ")
	if ok := scanner.Scan(); !ok {
		Logevents(EvSlice, "", "Ошибка ввода")
		fmt.Println("Ошибка ввода")
	}
	num, err := strconv.Atoi(scanner.Text())
	if err != nil || num > len(*slice) {
		Logevents(EvSlice, strconv.Itoa(num), "Вы ввели несуществующую задачу")
		fmt.Println("Вы ввели несуществующую задачу")
		return
	}

	fmt.Print("Введите статус выполнена/не выполнена\nстатус: ")
	scanner.Scan()
	status := scanner.Text()
	if strings.TrimSpace(status) == "" {
		Logevents(EvSlice, status, "Вы ничего не указали!")
		fmt.Println("Вы ничего не указали!")
		return
	}

	switch status {
	case "выполнена":
		(*slice)[num-1].Status = true
		fmt.Println("Вы изменили статус у задачи № ", num)
		fmt.Println("Время изменения: ", time.Now().Format("2006-01-02 15:04:05"))
	case "не выполнена":
		(*slice)[num-1].Status = false
		fmt.Print("Вы изменили статус у задачи № ", num)
	default:
		Logevents(EvSlice, status, "Вы указали неверный статус!")
		fmt.Println("Вы указали неверный статус!")
	}
}

// ////////////////////
func Logevents(slice *[]Events, input string, err string) {
	*slice = append(*slice, Events{
		Input:  input,
		Errors: err,
		Time:   time.Now().Format("2006-01-02 15:04:05"),
	})
}

/*
### Задание на проект:
- Необходимо написать приложение, работающее с пользовательским вводом
- Приложение представляет из себя "Список дел", или же "ToDo list"
- Должна быть возможность:
    - Добавлять новые задачи:
        - У задачи должен быть заголовок из одного слова
        - У задачи должен быть основной текст задачи, который может состоять из любого количества слов
        - У задачи должно быть время её создания
        - Задача может быть помечена как НЕ выполненная, либо как выполненная
        - У задачи должно быть время её выполнения, если она была выполнена
        - Все добавленные задачи должны сохраняться в программе
    - Получать полный список всех добавленных задач
    - Отмечать (по заголовку задачи) каждую отдельную задачу как выполненную:
        - Необходимо запоминать какие задачи выполненные а какие нет
        - Необходимо запоминать время выполнения задачи
    - Удалять (по заголовку задачи) ранее добавленные задачи
- Каждое событие в программе должно запоминаться:
    - Событием считается пользовательский ввод
    - Текст пользовательского ввода необходимо сохранять в событии
    - Если пользовательский ввод закончился с ошибкой, то сохранять текст ошибки в описание события
    - Если ошибки не было, то в качестве описания события оставить пустую строку
    - Необходимо сохранять время создания события
- Необходимо поддержать возможность получать список всех произошедших за время работы программы событий
- Программа должна быть модулем, и быть зависима от какой-нибудь сторонней библиотеки (например github.com/k0kubun/pp для красивого вывода)
- Программа должна содержать в себе хотя бы один пакет, помимо main пакета
- Чем больше пройденных в первой части полного курса по Golang возможностей будет использовано при написании этого учебного проекта, тем лучше

---
- Список команд, которые должны быть доступны в приложении:
    - help — эта команда позволяет узнать доступные команды и их формат
    - add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов} — эта команда позволяет добавлять новые задачи в список задач
    - list — эта команда позволяет получить полный список всех задач
    - del {заголовок существующей задачи} — эта команда позволяет удалить задачу по её заголовку
    - done {заголовок существующей задачи} — эта команда позволяет отменить задачу как выполненную
    - events — эта команда позволяет получить список всех событий
    - exit — эта команда позволяет завершить выполнение программы
*/

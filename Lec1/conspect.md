# Лекция 1. Установка и запуск первой программы

## Шаг 1. Установка компилятора
* Переход по ссылке: https://golang.org/dl/
* Установили для своей ОС комплиятор GoLang
***Важно*** : на данном курсе желательно, чтобы у вас была версия комплиятора > 1.08

## Шаг 2. GoRoot GoPath
**Определение** : '''GOROOT''' - это файловый путь, указывающий расположение ***КОМПИЛЯТОРА*** Go.
**Определение** : '''GOPATH''' - это файловый путь, указывающий на расположение ***РАБОЧЕГО ОКРУЖЕНИЯ*** (Там где пишем код и мазюкаем проекты). По умолчанию, на курсе мы создали '''GOPATH''' по адресу '''C:\Users\ilyap\go'''

## Шаг 3. Инициализация рабочего окружения
Чтобы создать рабочее окружение нам надо в '''GOPATH''' определить 3 директории: 
* '''src''' - место, где лежат исходники проектов (скрипты.go)
* '''bin''' - место, где лежат скомпилированные бинарники, после выполнения компиляции проектов
* '''pkg''' - место, где живут сторонние пакеты для наших проектов

## Шаг 4. Первая программа
В '''GOPATH''' /src создадим файл '''main.go''' со следующей начинкой:
'''

package main

import "fmt"

func main() {
	fmt.Println("halo wourd!")
}

'''

## Шаг 5. Запуск и компиляция
Go - компилируемый язык.
Для того, чтобы скомпилировать исполняемый файл, можно выполнить команду
'''
go build <path/to/go/file.go>
'''
Данная команда создаёт исполняеый файл по месту её вызова. Это удобно, когда мы в пылу битвы и хотим посмотреть, скомпилируется ли оно вообще или в случе каких-либо тестов, это позволяет на месте всё проверять.

Другая команда, которая также позволяет создать исполняемый файл:
'''
go install <path/to/go/file.go>
'''

Данная команда создаёт исполняемый файл по пути '''GOPATH/bin""".

Третья команда, которая будет часто использоваться на курсе - '''go run'''
'''
go run <path/to/go/file.go>
'''

Данная команда делает следующие действия:
* Создаёт исполняемый файл в временном хранилище
* Запускает этот файл
* И устраняет его (зависит от ОС)

Для того, чтобы узнать, где это временное хранилище '''--work'''

## Шаг 6. Правильная структуризация рабочего окружения
Очень рекомендуется создать следующий путь '''GOPATH/src/github.com/<your_github_username>/<github_repo_name>'''


package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"wbtechl1tasks/taskutils"
)

/*
Будем использовать следующую схему остановки горутин
Для остановки отдельной горутины кажется
Интересным способ когда мы передаём сигнал в канал
Но в нашем случае всё не так-то просто, ведь нам
Нужно передать не 1 сигнал а n сигналов
То есть провести "бродкаст" скажем так

Источники на великом сайте
https://stackoverflow.com/questions/61007385/golang-pattern-to-kill-multiple-goroutines-at-once
Говорят использовать contrext+waitgroup в таком случае

в то же время Context - это структура которая применяется
Именно в построении различных API
Что подходит под специфику стажировки, так что а почему
бы не использовать её?

В целом, учитывая что текущая задачка
- это консольное приложение, было бы странно использовать
здесь context, это и при чтении документации
для пакета context звучит как что-то кривое
Плюс существует эта статья
https://dave.cheney.net/2017/08/20/context-isnt-for-cancellation
На которую ссылаются среди ответов
Но в то же время
"he std library almost exclusively uses Contexts for cancelation.
 Just looks at the use in the net package for example,
  along with os/exec, net/http"
И я проверил - os/exec не имеющий по-сути отношения
к каким-либо веб-апи использует context))))
Так что это ПЛОХО но поменяем потом
Свою реализацию для отмены писать не хочу, честно говоря
*/

// В Воркере будем хранить канал, группу ожидания и контекст
type Worker struct {
	c   chan string
	wg  *sync.WaitGroup
	ctx context.Context
}

// WaitGroup нельзя передавать по значению потому что в ней копи-лок
func NewWorker(c chan string, wg *sync.WaitGroup, ctx context.Context) Worker {
	w := new(Worker)
	w.c = c
	w.wg = wg
	w.ctx = ctx
	return *w
}

func (w *Worker) Run() {
	/*
		сразу обозначим что при выходе из функции нужно передать
		в WaitGroup что мы всё
	*/
	defer w.wg.Done()
	for {
		select {
		// если прилетает сигнал на вывод - печатаем его
		case output := <-w.c:
			fmt.Println(output)
			// если прилетает сигнал что контекст завершается - завершаемся
		case <-w.ctx.Done():
			fmt.Println("Done!")
			return
		default:
			// либо бездействуем
			continue
		}
	}
}

/* Функция для очистки ресурсов - вызываем функцию отмены контекста
Плюс дожидаемся пока горутины завершатся
*/
func waitToEnd(wg *sync.WaitGroup, cf context.CancelFunc) {
	cf()
	wg.Wait()
}

func main() {
	fmt.Print("Number of workers: ")
	var workers int
	fmt.Scan(&workers)

	wg := sync.WaitGroup{}

	/*Нам необходимо обеспечить ручное выключение
	Поэтому в отличие от следующей задачи выбираем WithCancel
	Тогда нам возвращается сам контекст + функция отмены */
	ctx, cancelFunction := context.WithCancel(context.Background())

	// теперь у нас есть все что нужно для того чтобы вызывать waitToEnd
	defer waitToEnd(&wg, cancelFunction)

	// 20 взято с потолка не судите строго
	// Создаём нужное количество воркеров, конкурентно запускаем их
	channelino := make(chan string, 20)
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		w := NewWorker(channelino, &wg, ctx)
		go w.Run()
	}

	// Привязываем обработку сигнала SIGINT (Ctrl+C передаёт именно его)
	// К каналу
	ossig := make(chan os.Signal, 1)
	signal.Notify(ossig, os.Interrupt)

	for {
		select {
		// Если прилетает Ctrl+C выходим
		case <-ossig:
			return
		default:
			// Иначе генерим строку и отправляем её воркерам
			taskutils.GenerateString(20, channelino)

		}
	}

}

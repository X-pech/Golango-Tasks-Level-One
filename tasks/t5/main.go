package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
	"wbtechl1tasks/taskutils"
)

/*
Задача по приходе схожа с t4, но тут нам нужно создавать контекст
Не с помощью WithCancel а с помощью WithTimeout
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

func waitToEnd(wg *sync.WaitGroup, cf context.CancelFunc) {
	cf()
	wg.Wait()
}

func main() {
	var timeout int64
	fmt.Print("Timeout (in seconds): ")
	fmt.Scan(&timeout)

	// Всё как в задаче 4, но тут передаём вторым аргументом количество секунд
	ctx, cf := context.WithTimeout(context.Background(), time.Duration(time.Second*time.Duration(timeout)))

	wg := sync.WaitGroup{}
	wg.Add(1)

	channelino := make(chan string, 1)
	w := NewWorker(channelino, &wg, ctx)
	defer waitToEnd(&wg, cf)

	// Привязываем обработку сигнала SIGINT (Ctrl+C передаёт именно его)
	// К каналу
	ossig := make(chan os.Signal, 1)
	signal.Notify(ossig, os.Interrupt)

	go w.Run()

	for {
		select {
		case <-ossig:
			return
		case <-ctx.Done():
			// Также приходится и в основном потоке ловить сигнал что контекст завершен
			// Потому что теперь таймаут управляет потоком, а не только основной поток
			return
		default:
			taskutils.GenerateString(20, channelino)

		}
	}
}

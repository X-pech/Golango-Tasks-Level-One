package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

type RoutineSignal int

const (
	HELLO       RoutineSignal = 0
	STOP_RETURN RoutineSignal = 1
	STOP_GOEXIT RoutineSignal = 2
)

func signalChannelStop(signals chan RoutineSignal) {
	for {
		select {
		case signal, status := <-signals:
			if !status {
				fmt.Println("Stop with channel closing")
				return
			}
			if signal == STOP_RETURN {
				fmt.Println("Stop with return")
				return
			}
			if signal == STOP_GOEXIT {
				fmt.Println("Stop with goexit")
				runtime.Goexit()
			}
			fmt.Println(signal)
		default:
			time.Sleep(250 * time.Millisecond)
			fmt.Println("I'm working")
		}
	}
}

func contextedGoroutine(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context said we're done")
			return
		default:
			fmt.Println("I'm working in context")
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func main() {
	/*
		Можно передавать специальный сигнал для остановки
		Саму остановку можно производить как с помощью return
		Так и с помощью Goexit
		Что особо не влияет на не-главные горутины
		(Останавливать с помощью Goexit main() не очень хорошо:
		>Calling Goexit from the main goroutine
		>terminates that goroutine without func main returning.
		>Since func main has not returned, the program continues
		>execution of other goroutines.
		>If all other goroutines exit, the program crashes.
	*/
	s := make(chan RoutineSignal)
	go signalChannelStop(s)
	time.Sleep(500 * time.Millisecond)
	s <- STOP_RETURN
	go signalChannelStop(s)
	time.Sleep(500 * time.Millisecond)
	s <- STOP_GOEXIT

	/*
		Также можно закрывать канал, тоже вполне себе сигнал
		Тогда даже не обязательно иметь канал-специально-для-сигналов
	*/
	go signalChannelStop(s)
	time.Sleep(500 * time.Millisecond)
	close(s)
	time.Sleep(500 * time.Millisecond)

	/*
		Мы можем завершать горутины с помощью контекстов!
		Они помогут в вещании того же "канала выключения"
		на множество горутин
		(ручками пришлось бы создавать множество каналов)
		WaitGroup добавлен чтобы дождаться конца их работы
	*/
	wg := sync.WaitGroup{}
	wg.Add(2)
	ctx, cf := context.WithCancel(context.Background())
	go contextedGoroutine(ctx, &wg)
	go contextedGoroutine(ctx, &wg)
	time.Sleep(500 * time.Millisecond)
	cf()
	wg.Wait()

	/*
		Ещё контексты можно создать с помощью таймаутов
		(относительное время)
		wg добавлен для того чтобы не писать новую функцию
	*/
	ctx, _ = context.WithTimeout(context.Background(), time.Second*1)
	wg.Add(2)
	go contextedGoroutine(ctx, &wg)
	go contextedGoroutine(ctx, &wg)
	time.Sleep(time.Millisecond * 1500)

	/*
		...И с помощью дедлайнов (абсолютное время)
	*/
	ctx, _ = context.WithDeadline(context.Background(), time.Now().Add(time.Second*1))
	wg.Add(2)
	go contextedGoroutine(ctx, &wg)
	go contextedGoroutine(ctx, &wg)
	time.Sleep(time.Millisecond * 1500)

	/*
		Кстати если main вернёт значение прежде окончания работы
		остальных горутин то они просто остановятся мгновенно
		Об это нужно помнить чтобы они не устроили нам
		инвалидацию чего-нибудь важного
		(Goexit() потому и не останавливает
		остальные горутины при вызове
		из main() что main() не сделает "return")
	*/
	s = make(chan RoutineSignal)
	go signalChannelStop(s)
	time.Sleep(1 * time.Second)
}

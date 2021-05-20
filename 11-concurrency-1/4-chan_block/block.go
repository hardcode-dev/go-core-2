package main

func main() {
	var ch = make(chan int)
	close(ch)
	// Попытка записи в закрытый канал вызывает панику.
	ch <- 10
	// Чтение из закрытого канала не блокирует и всегда возвращает
	// значение по умолчанию для типа данных канала.
	v := <-ch
	print(v)
	return
	var ch1 = make(chan int)
	// При достижении предела ёмкости канала
	// операция записи в него блокирует.
	ch1 <- 10
	n := <-ch1
	print(n)
	ch2 := make(chan string, 2)
	ch2 <- "S1"
	ch2 <- "S2"
	ch2 <- "S3"
	s := <-ch2
	print(s)
	var ch3 = make(chan bool)
	// Операция чтения из канала блокирует.
	b := <-ch3
	print(b)
	// Канал – это ресурс. Важно явно закрывать канал.
	close(ch1)
	close(ch2)
	close(ch3)
}

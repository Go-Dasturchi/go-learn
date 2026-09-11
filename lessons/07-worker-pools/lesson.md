# 07 — Worker Pools

## THEORY

Do'konda mijozlar soni yuzlab bo'lishi mumkin, lekin do'kon har bir mijoz uchun alohida kassir yollamaydi — o'rniga, **cheklangan sondagi** kassirlar bor, va mijozlar ular bo'shaganda birma-bir xizmat oladi. **Worker pool (ishchilar guruhi)** — dasturlashda aynan shu g'oya: "ishlar" ko'p bo'lishi mumkin, lekin ularni qayta ishlaydigan **goroutine'lar soni cheklangan**.

**Nega barcha ishlarga alohida-alohida goroutine ishga tushirmaslik kerak.** "Goroutines" darsida ko'rgan naqshda, har bir elementga bitta goroutine ishga tushirgan edik — bu kichik miqdordagi ishlar uchun yaxshi, lekin million ish bo'lsa-chi? Million goroutine bir vaqtda ishga tushishi (garchi goroutine'lar yengil bo'lsa ham) xotira va tizim resurslarini haddan tashqari band qilib yuborishi mumkin. Worker pool — ishlaydigan goroutine'lar sonini **nazorat ostida** ushlab turadi.

**Naqsh — ikkita kanal: `jobs` va `results`.**

```go
func worker(jobs <-chan int, results chan<- Natija, wg *sync.WaitGroup) {
	defer wg.Done()
	for son := range jobs {
		results <- Natija{Indeks: son, Qiymat: son * son}
	}
}
```

Bu yerda `jobs <-chan int` va `results chan<- Natija` — **yo'nalishli kanal turlari**: `<-chan int` — "bu funksiya bu kanaldan faqat **o'qiy oladi**", `chan<- Natija` — "bu funksiya bu kanalga faqat **yoza oladi**". Bu — Go kompilyatoriga funksiyaning kanal bilan qanday ishlashini aniq bildirib, tasodifiy xato (masalan, faqat o'qishi kerak bo'lgan kanalga yozib yuborish) qilishning oldini oladi.

**Butun naqsh — belgilangan sondagi worker ishga tushirish:**

```go
func workerHisoblash(sonlar []int, workerSoni int) []int {
	jobs := make(chan int, len(sonlar))
	results := make(chan Natija, len(sonlar))
	var wg sync.WaitGroup

	for i := 0; i < workerSoni; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	for _, son := range sonlar {
		jobs <- son
	}
	close(jobs) // endi yangi ish yo'q — worker'lar shuni ko'rib, tsiklni tugatadi

	wg.Wait()
	close(results)

	natija := make([]int, len(sonlar))
	for r := range results {
		natija[r.Indeks] = r.Qiymat
	}
	return natija
}
```

**Nima uchun `close(jobs)` muhim.** Har bir `worker` `for son := range jobs { ... }` orqali ishlaydi — bu tsikl `jobs` yopilib, bo'shagunicha to'xtamaydi ("Channels" darsida ko'rgan `range` xususiyatini eslang). `close(jobs)`ni chaqirmasangiz, worker'lar abadiy yangi ish kutib, dastur hech qachon tugamaydi.

**Nega natijalarni indeks orqali yig'amiz.** Worker'lar qaysi ishni qachon bajarishi **kafolatlanmagan** tartibda — shuning uchun "Goroutines" darsida ko'rgan naqshni takrorlaymiz: har bir natijaga uning asl "manzili"ni (bu holda, sonning o'zi indeks sifatida ishlatilmoqda — sodda misol uchun) biriktirib, yakunda to'g'ri tartibda joylashtiramiz.

## EXAMPLE

```go
package main

import (
	"fmt"
	"sync"
)

type Natija struct {
	Indeks int
	Qiymat int
}

func worker(jobs <-chan int, results chan<- Natija, wg *sync.WaitGroup) {
	defer wg.Done()
	for son := range jobs {
		results <- Natija{Indeks: son, Qiymat: son * son}
	}
}

func workerHisoblash(sonlar []int, workerSoni int) map[int]int {
	jobs := make(chan int, len(sonlar))
	results := make(chan Natija, len(sonlar))
	var wg sync.WaitGroup

	for i := 0; i < workerSoni; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	for _, son := range sonlar {
		jobs <- son
	}
	close(jobs)

	wg.Wait()
	close(results)

	natija := make(map[int]int)
	for r := range results {
		natija[r.Indeks] = r.Qiymat
	}
	return natija
}

func main() {
	fmt.Println(workerHisoblash([]int{1, 2, 3, 4, 5}, 2))
}
```

Natija (tartib har xil bo'lishi mumkin, chunki bu — map, lekin qiymatlar har doim to'g'ri):

```
map[1:1 2:4 3:9 4:16 5:25]
```

## TASK

`Natija` struct'i, `worker` funksiyasi (allaqachon yozilgan) va `workerHisoblash(sonlar []int, workerSoni int) map[int]int` funksiyasi berilgan. Funksiyani shunday to'ldiringki, u `workerSoni` ta worker goroutine ishga tushirib, `sonlar`dagi har bir sonning kvadratini hisoblasin, va natijani `map[son]kvadrat` ko'rinishida qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `jobs` va `results` kanallarini yarating, `workerSoni` marta `wg.Add(1)` va `go worker(jobs, results, &wg)` qiling, keyin `sonlar`ni `jobs`ga yuboring va `close(jobs)` qiling.
2. `wg.Wait()`, `close(results)`, so'ng `natija := make(map[int]int)` va `for r := range results { natija[r.Indeks] = r.Qiymat }`, oxirida `return natija`.

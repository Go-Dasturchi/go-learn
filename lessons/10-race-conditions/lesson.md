# 10 — Race Conditions

## THEORY

Ikki kishi bitta doskaga, bir vaqtning o'zida, muvofiqlashmasdan yozayotganini tasavvur qiling — natija nima bo'lishi oldindan aniq emas: bittasining yozuvi ikkinchisinikini "yopib" qo'yishi, yoki ikkalasi aralashib, chalkash natija chiqishi mumkin. **Musobaqa holati (race condition)** — dasturlashda aynan shu: ikki yoki undan ko'p goroutine **bir xil xotiraga**, hech qanday muvofiqlashtirishsiz, va kamida bittasi **yozish** bilan murojaat qilganda yuzaga keladi. Natija — **bashorat qilib bo'lmaydigan**, ba'zan har ishga tushirilganda boshqacha bo'ladi.

**Bu bo'lim davomida o'rgangan butun asboblar qutingiz — aynan shu muammoni hal qilish uchun edi:**

- **Channel** ("Channels", "Buffered Channels") — ma'lumotni **almashtirish** orqali, umumiy xotiraga bevosita murojaat qilmasdan muloqot qilish.
- **WaitGroup** ("WaitGroup") — bir nechta goroutine tugashini **kutish**.
- **Mutex** ("Mutex") — umumiy xotiraga murojaatni **navbat bilan**, bittadan ruxsat berish.
- **Atomic** ("sync.Once and Atomic") — oddiy sonlar ustida Mutex'siz, bo'linmas amallar.
- **Select** ("Select") va **Context** ("Context") — bir nechta manba/signalni boshqarish va bekor qilish.

Go'ning mashhur maqoli buni yaxshi jamlaydi: **"Do not communicate by sharing memory; instead, share memory by communicating"** ("Xotirani baham ko'rish orqali muloqot qilmang; aksincha, muloqot qilish orqali xotirani baham ko'ring") — imkon qadar channel orqali ma'lumot "uzatishga" harakat qiling, umumiy o'zgaruvchini har tomondan "ushlab turishdan" ko'ra.

**Race condition qanday ko'rinishda bo'ladi — misol:**

```go
// XAVFLI — himoyasiz umumiy o'zgaruvchi
son := 0
var wg sync.WaitGroup
for i := 0; i < 1000; i++ {
	wg.Add(1)
	go func() {
		defer wg.Done()
		son++ // bir nechta goroutine bir vaqtda o'qib-yozadi!
	}()
}
wg.Wait()
fmt.Println(son) // ko'pincha 1000dan kam chiqadi — natija barqaror emas
```

Va uni tuzatishning ikkita to'g'ri yo'li (allaqachon ko'rgan edingiz):

```go
// TO'G'RI — Mutex bilan
var mu sync.Mutex
son := 0
// ... mu.Lock(); son++; mu.Unlock() ...

// TO'G'RI — atomic bilan
var son atomic.Int64
// ... son.Add(1) ...
```

**Go'ning race detektori — `-race` bayrog'i.** Go standart kutubxonasi bilan birga **race detektori** ham keladi — bu, dasturni maxsus rejimda ishga tushirib, xavfli, himoyasiz umumiy xotiraga murojaatlarni **avtomatik aniqlab beradi**:

```bash
go test -race ./...
go run -race main.go
```

Bu — ishlab chiqarishda ishlatilmaydi (sekinroq ishlaydi), lekin **test va sozlash bosqichida** juda qimmatli vosita: u ko'zga tashlanmaydigan, faqat vaqti-vaqti bilan namoyon bo'ladigan xatolarni ishonchli tarzda topib beradi.

**Yakuniy mashq — hamma narsani birlashtiring.** Ushbu darsning mashqi — butun bo'lim davomida o'rgangan vositalardan (Goroutines, Channels, WaitGroup, Mutex yoki Atomic) **o'zingiz tanlagan** kombinatsiyasi bilan, sonlar ro'yxatini **parallel** ravishda, **xavfsiz** yig'indisini hisoblovchi funksiya yozish.

## EXAMPLE

```go
package main

import (
	"fmt"
	"sync"
)

// Xavfsiz — Mutex bilan himoyalangan parallel yig'indi
func parallelYigindi(sonlar []int, workerSoni int) int {
	var mu sync.Mutex
	var wg sync.WaitGroup
	natija := 0

	jobs := make(chan int, len(sonlar))
	for _, son := range sonlar {
		jobs <- son
	}
	close(jobs)

	for i := 0; i < workerSoni; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for son := range jobs {
				mu.Lock()
				natija += son
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	return natija
}

func main() {
	fmt.Println(parallelYigindi([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3))
}
```

Natija:

```
55
```

## TASK

`parallelYigindi(sonlar []int, workerSoni int) int` funksiyasi berilgan. Uni shunday to'ldiringki, u `sonlar`ning yig'indisini **`workerSoni` ta parallel goroutine** yordamida, **xavfsiz** (race condition'siz) hisoblab qaytarsin. Qanday vosita ishlatishni (Mutex, Atomic, yoki Channel orqali yig'ish) o'zingiz tanlashingiz mumkin — muhimi, natija har doim **to'g'ri va barqaror** bo'lishi.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Eng oddiy yo'l — barcha sonlarni `jobs` kanaliga joylab, `workerSoni` ta goroutine ishga tushiring, har biri `jobs`dan o'qib, umumiy `natija`ni **Mutex bilan himoyalab** oshirsin.
2. `wg.Wait()`dan keyin `natija`ni qaytaring — `WaitGroup` barcha worker'lar tugaganini kafolatlaydi, shundan keyingina yakuniy qiymatni o'qish xavfsiz.

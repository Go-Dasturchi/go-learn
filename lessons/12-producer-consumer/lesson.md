# 12 — Producer / Consumer

## THEORY

Pekarxonani tasavvur qiling: bir non yopuvchi (producer, ishlab chiqaruvchi) tinimsiz yangi non pishirib, javonga qo'yib boradi, sotuvchi (consumer, iste'molchi) esa javondan nonlarni olib, mijozlarga sotadi. Ular **turli tezlikda** ishlashi mumkin — bittasi tezroq, ikkinchisi sekinroq — lekin javon (kanal) ularni bog'lab, muvofiqlashtirib turadi. Bu — **producer/consumer (ishlab chiqaruvchi/iste'molchi)** naqshi.

**"Worker Pools" bilan farqi.** Worker pool'da ("Worker Pools" darsini eslang) bir nechta **bir xil** worker, umumiy ish navbatidan olib, **parallel ravishda bir xil turdagi ishni** bajarardi — maqsad, ishni tezroq tugatish uchun yukni taqsimlash edi. Producer/consumer esa boshqacha urg'u beradi: bu yerda muhimi — **ishlab chiqarish** va **iste'mol qilish** bosqichlarini **bir-biridan ajratish**, ular turli tezlikda, mustaqil ishlay olishi uchun. Producer'lar soni ham, consumer'lar soni ham — bittadan ko'p bo'lishi mumkin.

**Bir nechta producer, bitta consumer:**

```go
func producer(sonlar []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, son := range sonlar {
		ch <- son
	}
}

func ishlabChiqaruvchiIsteMolchi(sonlar1, sonlar2 []int) int {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(sonlar1, ch, &wg)
	go producer(sonlar2, ch, &wg)

	go func() {
		wg.Wait()
		close(ch) // ikkala producer ham tugagach, kanalni yopamiz
	}()

	natija := 0
	for son := range ch { // consumer — kanal yopilguncha o'qiydi
		natija += son
	}
	return natija
}
```

**Nega `close(ch)` alohida goroutine ichida chaqiriladi.** Bu — muhim, nozik joy: `consumer` (`for son := range ch`) kanal to'liq **yopilguncha** to'xtamaydi, shuning uchun agar `wg.Wait()`ni to'g'ridan-to'g'ri asosiy oqimda (consumer tsiklidan oldin) chaqirsangiz, dastur **abadiy muzlab qoladi** — `wg.Wait()` producer'lar tugashini kutadi, lekin producer'lar hali kanalga yuborishda davom etayotgan bo'lishi mumkin, va consumer ularni hali o'qib ulgurmagan. Yechim: `wg.Wait()` + `close(ch)`ni **alohida goroutine**ga chiqarib, asosiy oqim erkin ravishda `for range ch` bilan **darhol** o'qishni boshlaydi — ikkala producer ham tugagach, o'sha alohida goroutine kanalni yopadi, va `for range` o'zi to'xtaydi.

**Nega bu naqsh foydali.** Ishlab chiqarish va iste'mol qilishni ajratish, ikkala tomonni ham **mustaqil** rivojlantirish imkonini beradi — masalan, kelajakda yana bir producer qo'shish, consumer kodini o'zgartirmasdan mumkin bo'ladi. Bu naqsh veb-serverlarda so'rovlarni navbatga qo'yish, log yozish tizimlarida, va ma'lumotlarni oqim (stream) sifatida qayta ishlashda keng qo'llaniladi.

## EXAMPLE

```go
package main

import (
	"fmt"
	"sync"
)

func producer(sonlar []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, son := range sonlar {
		ch <- son
	}
}

func ishlabChiqaruvchiIsteMolchi(sonlar1, sonlar2 []int) int {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(sonlar1, ch, &wg)
	go producer(sonlar2, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	natija := 0
	for son := range ch {
		natija += son
	}
	return natija
}

func main() {
	fmt.Println(ishlabChiqaruvchiIsteMolchi([]int{1, 2, 3}, []int{4, 5, 6}))
}
```

Natija:

```
21
```

## TASK

`producer` funksiyasi berilgan (allaqachon yozilgan). `ishlabChiqaruvchiIsteMolchi(sonlar1, sonlar2 []int) int` funksiyasini shunday to'ldiringki, u ikkita `producer` goroutine (har bir slice uchun bittadan) ishga tushirsin, ular tugagach kanalni yopsin, va asosiy oqim (consumer) kanaldagi barcha sonlarning yig'indisini qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `ch := make(chan int)`, `wg.Add(2)`, `go producer(sonlar1, ch, &wg)`, `go producer(sonlar2, ch, &wg)`, keyin alohida `go func() { wg.Wait(); close(ch) }()`.
2. `natija := 0`, `for son := range ch { natija += son }`, oxirida `return natija`.

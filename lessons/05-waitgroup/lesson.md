# 05 — WaitGroup

## THEORY

O'qituvchi sinfga topshiriq berib, "hammangiz tugatgandan keyin daftaringizni yig'ib olaman" desa — u har bir o'quvchini alohida kuzatib o'tirmaydi, faqat "nechta o'quvchi hali tugatmadi" degan hisobni yuritadi, va bu hisob nolga tushganda daftarlarni yig'ishga o'tadi. `sync.WaitGroup` — Go'da aynan shu vazifani bajaradi: **bir nechta goroutine tugashini kutish**.

**Uchta asosiy method:**
- **`Add(n)`** — "yana `n` ta ish kutilyapti" deb hisobni oshiradi.
- **`Done()`** — "bitta ish tugadi" deb hisobni bittaga kamaytiradi (odatda `defer` bilan ishlatiladi, "Defer" darsini eslang — shunda goroutine qanday tugashidan qat'iy nazar, `Done()` albatta chaqiriladi).
- **`Wait()`** — hisob nolga tushguncha **kutadi** (bloklanadi).

```go
var wg sync.WaitGroup

for i := 0; i < 3; i++ {
	wg.Add(1) // "yana bitta ish bor"
	go func() {
		defer wg.Done() // ish tugagach, hisobni kamaytir
		fmt.Println("Ish bajarildi")
	}()
}

wg.Wait() // barcha 3 ta ish tugaguncha shu yerda kutadi
fmt.Println("Hammasi tugadi")
```

**Nega `WaitGroup` kerak — "Goroutines" darsidagi muammoni eslang.** O'sha darsda `main()` boshqa goroutine'larni **kutmasligini** ko'rgan edik. `WaitGroup` — aynan shu muammoni hal qiladi: u orqali dastur "hamma ish tugaguncha shu yerda tur" deb aniq ko'rsatma bera oladi.

**`WaitGroup` va natijalarni yig'ish — birga ishlatish.** `WaitGroup`ning o'zi faqat "tugadimi-yo'qmi"ni biladi, natijalarning **o'zini** tashimaydi — shuning uchun uni ko'pincha "Channels" darsida ko'rgan kanal bilan birga ishlatishadi: `WaitGroup` barcha "ishlab chiqaruvchi" goroutine'lar tugaganini bildirsa, shundan keyingina kanalni **xavfsiz yopish** mumkin bo'ladi:

```go
func jamiHisobla(sonlar []int) int {
	var wg sync.WaitGroup
	ch := make(chan int, len(sonlar))

	for _, son := range sonlar {
		wg.Add(1)
		go func(s int) {
			defer wg.Done()
			ch <- s * s
		}(son)
	}

	wg.Wait()  // barcha goroutine'lar ch ga yuborib bo'lguncha kutamiz
	close(ch)  // endi xavfsiz yopsa bo'ladi — hech kim yana yubormaydi

	natija := 0
	for v := range ch {
		natija += v
	}
	return natija
}
```

**Nega `wg.Wait()` `close(ch)`dan oldin turishi shart.** Agar kanalni goroutine'lar hali yuborib bo'lmasdan yopib qo'ysangiz, ular yopilgan kanalga yuborishga urinib, **panic** qilishadi ("Panic and Recover" darsini eslang — yopilgan kanalga yuborish jiddiy xato hisoblanadi). `wg.Wait()` — barcha yuborishlar tugaganiga **kafolat** beradi, shundan keyingina yopish xavfsiz bo'ladi.

## EXAMPLE

```go
package main

import (
	"fmt"
	"sync"
)

func jamiHisobla(sonlar []int) int {
	var wg sync.WaitGroup
	ch := make(chan int, len(sonlar))

	for _, son := range sonlar {
		wg.Add(1)
		go func(s int) {
			defer wg.Done()
			ch <- s * s
		}(son)
	}

	wg.Wait()
	close(ch)

	natija := 0
	for v := range ch {
		natija += v
	}
	return natija
}

func main() {
	fmt.Println(jamiHisobla([]int{1, 2, 3, 4}))
}
```

Natija:

```
30
```

## TASK

`jamiHisobla(sonlar []int) int` funksiyasi berilgan. Uni `sync.WaitGroup` va kanal yordamida shunday to'ldiringki, u har bir son uchun alohida goroutine ishga tushirib, uning **kvadratini** kanalga yuborsin; barcha goroutine'lar tugagach (WaitGroup orqali), kanalni yopib, barcha kvadratlarning **yig'indisini** qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `var wg sync.WaitGroup`, `ch := make(chan int, len(sonlar))` bilan boshlang. Har bir son uchun: `wg.Add(1)`, keyin `go func(s int) { defer wg.Done(); ch <- s * s }(son)`.
2. `wg.Wait()`, `close(ch)`, so'ng `natija := 0; for v := range ch { natija += v }`, oxirida `return natija`.

# 02 — Channels

## THEORY

"Goroutines" darsida channel'ni "quvur" sifatida tanishtirgan edik — endi uni chuqurroq ko'rib chiqamiz. Pochta qutisini tasavvur qiling: bittasi xat tashlaydi, boshqasi kelib oladi — ular bir-birlarini ko'rishlari shart emas, faqat qutining o'zi orqali "muloqot" qilishadi. **Channel** — Go'da goroutine'lar orasida shunday xavfsiz "muloqot" qilish vositasi.

**Yaratish va asosiy amallar:**

```go
ch := make(chan int) // int turidagi qiymatlar uchun kanal

ch <- 5     // kanalga yuborish (send)
qiymat := <-ch // kanaldan qabul qilish (receive)
```

**Muhim: oddiy (buferlanmagan) kanal — bloklovchi.** `make(chan int)` orqali yaratilgan kanal **buferlanmagan** (unbuffered): unga yuborish, boshqa tomonda **kimdir kutib turgan** bo'lishini talab qiladi — aks holda `ch <- 5` qatori shu yerda **to'xtab qoladi** (bloklanadi), toki kimdir `<-ch` orqali qabul qilguncha. Bu — jismoniy qo'lma-qo'l topshirishga o'xshaydi: ikkala tomon ham bir vaqtda "hozir" bo'lishi kerak.

**Kanalni yopish — `close()`.** Agar kanalga endi hech qachon yangi qiymat yuborilmasligini bildirmoqchi bo'lsangiz, uni yopish mumkin:

```go
close(ch)
```

Yopilgan kanaldan qabul qilish davom etadi (agar hali yuborilgan, lekin olinmagan qiymatlar bo'lsa, ular baribir olinadi), lekin kanal **butunlay bo'shab, yopilgan**dan keyin, undan qabul qilish darhol **standart (zero) qiymat** va **`false`** qaytaradi — bu "Maps" darsida ko'rgan comma-ok naqshiga o'xshaydi:

```go
qiymat, ok := <-ch
if !ok {
	fmt.Println("Kanal yopilgan va bo'sh")
}
```

**`range` bilan kanaldan o'qish.** Kanal yopilguncha undan barcha qiymatlarni birma-bir olish uchun, "For Loop" darsida ko'rgan `range`ni kanal ustida ham ishlatish mumkin — bu **avtomatik ravishda** kanal yopilganda to'xtaydi:

```go
for qiymat := range ch {
	fmt.Println(qiymat)
}
// kanal yopilib, bo'shagandan keyin tsikl o'zi to'xtaydi
```

**Naqsh: bitta "ishlab chiqaruvchi" goroutine + asosiy funksiyada yig'ish.** Bu — juda keng tarqalgan naqsh: bitta goroutine ma'lumotlarni kanalga "quyib" boradi va oxirida uni yopadi, asosiy kod esa `range` bilan hammasini yig'ib oladi:

```go
func yigindiChannel(sonlar []int) int {
	ch := make(chan int)
	go func() {
		for _, son := range sonlar {
			ch <- son
		}
		close(ch) // MUHIM: yuborish tugagach, kanalni yopish kerak
	}()

	natija := 0
	for son := range ch {
		natija += son
	}
	return natija
}
```

**Nega `close(ch)`ni unutmaslik kerak.** Agar kanal hech qachon yopilmasa, `for son := range ch` tsikli **abadiy kutib qoladi** — u yangi qiymat yoki "kanal yopildi" signalini kutadi, lekin hech biri kelmaydi. Bu — dasturni "abadiy muzlatib qo'yadigan" (deadlock) eng ko'p uchraydigan xatolardan biri, shuning uchun "ishlab chiqaruvchi" goroutine ishini tugatgach, kanalni yopishni unutmaslik juda muhim.

## EXAMPLE

```go
package main

import "fmt"

func yigindiChannel(sonlar []int) int {
	ch := make(chan int)
	go func() {
		for _, son := range sonlar {
			ch <- son
		}
		close(ch)
	}()

	natija := 0
	for son := range ch {
		natija += son
	}
	return natija
}

func main() {
	fmt.Println(yigindiChannel([]int{1, 2, 3, 4, 5}))
}
```

Natija:

```
15
```

## TASK

`yigindiChannel(sonlar []int) int` funksiyasi berilgan. Uni shunday to'ldiringki:

1. Alohida goroutine ishga tushirib, `sonlar` ichidagi har bir sonni kanalga yuborsin, so'ng kanalni yopsin.
2. Asosiy funksiya `range` yordamida kanaldagi barcha sonlarni yig'ib, ularning yig'indisini qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `ch := make(chan int)` bilan kanal yarating, keyin `go func() { for _, son := range sonlar { ch <- son }; close(ch) }()` bilan yuboruvchi goroutine'ni ishga tushiring.
2. `natija := 0`, so'ng `for son := range ch { natija += son }`, oxirida `return natija` — kanal yopilishi bilan tsikl o'zi to'xtaydi.

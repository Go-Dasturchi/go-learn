# 04 — Select

## THEORY

Bir vaqtning o'zida ikkita telefon liniyasini kutayotganingizni tasavvur qiling — qaysi biri **avval jiringlasa**, o'shani ko'tarasiz, ikkinchisini keyinroq. Go'da `select` operatori aynan shunday ishlaydi: u **bir nechta kanal amali**ni bir vaqtda "kutadi" va **qaysi biri birinchi tayyor bo'lsa**, o'shani bajaradi.

**Sintaksis — `switch`ga o'xshaydi, lekin kanallar bilan:**

```go
select {
case v := <-ch1:
	fmt.Println("ch1 dan keldi:", v)
case v := <-ch2:
	fmt.Println("ch2 dan keldi:", v)
}
```

`select` — har bir `case`dagi kanal amalini tekshiradi, va **qaysi biri tayyor bo'lsa** (ya'ni, o'sha kanaldan qiymat olish mumkin bo'lsa), o'sha `case`ni bajaradi. Agar hech biri hali tayyor bo'lmasa, `select` **kutadi** (bloklanadi), toki ulardan biri tayyor bo'lguncha.

**Ikkalasi ham bir vaqtda tayyor bo'lsa — tasodifiy tanlanadi.** Agar bir nechta `case` bir vaqtda tayyor bo'lib qolsa, Go ulardan birini **tasodifiy** tanlaydi — bu "adolatli" bo'lish uchun ataylab shunday qilingan (aks holda birinchi yozilgan `case` doim "ustunlik" qilib, boshqalari hech qachon tanlanmasligi mumkin edi).

**`default` — bloklanmasdan "urinib ko'rish".** Agar `select` ichida `default` bo'lsa, va hech qanday kanal hali tayyor bo'lmasa, `select` **kutmasdan**, darhol `default` blokini bajaradi:

```go
select {
case v := <-ch:
	fmt.Println("Qiymat keldi:", v)
default:
	fmt.Println("Hozircha hech narsa yo'q")
}
```

Bu naqsh — kanalni **bloklanmasdan tekshirib ko'rish** kerak bo'lganda foydali (masalan, "agar tayyor bo'lsa ol, bo'lmasa boshqa ish bilan shug'ullan").

**Nega `select` foydali.** U bir nechta manbadan (masalan, bir nechta ishchi goroutine'dan, yoki "Context" darsida ko'radigan bekor qilish signalidan) kelayotgan ma'lumotni **bitta joyda**, tartibli tarzda kutish imkonini beradi — har bir kanal uchun alohida `for` tsikli yozib, ularni qandaydir tarzda "aralashtirishga" hojat qoldirmaydi.

## EXAMPLE

```go
package main

import "fmt"

func birinchiTayyor(ch1, ch2 chan string) string {
	select {
	case v := <-ch1:
		return v
	case v := <-ch2:
		return v
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "salom"
	}()

	fmt.Println(birinchiTayyor(ch1, ch2))
}
```

Natija:

```
salom
```

## TASK

`birinchiTayyor(ch1, ch2 chan string) string` funksiyasi berilgan. Uni `select` yordamida shunday to'ldiringki, u `ch1` yoki `ch2`dan **qaysi biri birinchi bo'lib** qiymat yuborsa, o'sha qiymatni qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `select { case v := <-ch1: ... case v := <-ch2: ... }` — ikkala kanalni ham bitta `select` ichida kuting.
2. Har bir `case` ichida oddiygina `return v` yozing — qaysi kanal birinchi tayyor bo'lsa, o'sha `case` ishga tushadi.

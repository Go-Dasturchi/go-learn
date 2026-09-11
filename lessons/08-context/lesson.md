# 08 — Context

## THEORY

Fabrika boshlig'i ishchilarga "signal berilsa, darhol ishni to'xtatinglar" deb, hammaga bitta hushtak tarqatganini tasavvur qiling — boshliq hushtakni chalsa, hamma, qayerda bo'lishidan qat'iy nazar, buni eshitib ishni to'xtatadi. Go'da **`context.Context`** — aynan shu "hushtak"ning vazifasini bajaradi: u orqali bir nechta goroutine'ga bir vaqtda **"bekor qilindi" yoki "vaqt tugadi"** signalini yetkazish mumkin.

**Yaratish — `context.WithCancel`:**

```go
ctx, bekorQil := context.WithCancel(context.Background())
```

`context.Background()` — "bo'sh", boshlang'ich context (odatda dastur eng yuqori darajasida shu bilan boshlanadi). `WithCancel` esa undan ikkita narsa qaytaradi: yangi `ctx` (uni goroutine'larga uzatasiz) va `bekorQil` — bu funksiyani chaqirganingizda, `ctx` "bekor qilindi" holatiga o'tadi.

**Signalni "eshitish" — `ctx.Done()`.** `ctx.Done()` — bu maxsus kanal, va u `bekorQil()` chaqirilganda **yopiladi** ("Channels" darsida ko'rgan `close()` va yopilgan kanaldan o'qish naqshini eslang):

```go
select {
case <-ctx.Done():
	fmt.Println("Bekor qilindi!")
default:
	fmt.Println("Hali davom etyapti")
}
```

**Naqsh — uzun ishni bosqichma-bosqich bajarish, har safar bekor qilinganini tekshirib.** Agar biror ish bir nechta qadamdan iborat bo'lsa, har qadamdan oldin `ctx.Done()`ni tekshirib, agar signal kelgan bo'lsa, ishni **erta to'xtatish** mumkin:

```go
func ishlaGachaBekor(ctx context.Context, sonlar []int) []int {
	natija := []int{}
	for _, son := range sonlar {
		select {
		case <-ctx.Done():
			return natija // signal keldi — hozirgacha yig'ilganini qaytarib, to'xtaymiz
		default:
			natija = append(natija, son*son)
		}
	}
	return natija
}
```

**Nima uchun `select` + `default` kombinatsiyasi ishlatiladi.** "Select" darsida ko'rgan `default` bilan naqshni eslang: bu — `ctx.Done()`ni **bloklanmasdan** tekshirish usuli. Agar `default` bo'lmasa, `select` `ctx.Done()` "tayyor" bo'lguncha (ya'ni, bekor qilinguncha) shu yerda **kutib qolardi** — bizga esa "agar bekor qilingan bo'lsa to'xta, aks holda **darhol** ishni davom ettir" kerak.

**`context.WithTimeout` — vaqt bo'yicha avtomatik bekor qilish.** `WithCancel`dan tashqari, ma'lum vaqt o'tgach **avtomatik** bekor bo'ladigan context ham bor:

```go
ctx, bekorQil := context.WithTimeout(context.Background(), 5*time.Second)
defer bekorQil() // resurslarni tozalash uchun har doim chaqirish tavsiya etiladi
```

Bu, real vaqtga bog'liq bo'lgani uchun, avtomatik testlarda ishlatish qiyinroq (har safar aniq vaqt kutish kerak bo'lardi) — shuning uchun bu darsning mashqi `WithCancel`ga asoslangan, natijasi darhol va aniq bo'ladigan holatga qaratilgan.

## EXAMPLE

```go
package main

import (
	"context"
	"fmt"
)

func ishlaGachaBekor(ctx context.Context, sonlar []int) []int {
	natija := []int{}
	for _, son := range sonlar {
		select {
		case <-ctx.Done():
			return natija
		default:
			natija = append(natija, son*son)
		}
	}
	return natija
}

func main() {
	ctx := context.Background()
	fmt.Println(ishlaGachaBekor(ctx, []int{1, 2, 3}))

	bekorlangan, bekorQil := context.WithCancel(context.Background())
	bekorQil() // darhol bekor qilamiz
	fmt.Println(ishlaGachaBekor(bekorlangan, []int{1, 2, 3}))
}
```

Natija:

```
[1 4 9]
[]
```

## TASK

`ishlaGachaBekor(ctx context.Context, sonlar []int) []int` funksiyasi berilgan. Uni shunday to'ldiringki, u `sonlar` bo'ylab yurib, har bir sonning kvadratini natijaga qo'shsin — lekin har qadamdan oldin `ctx.Done()`ni (bloklanmasdan, `select`+`default` orqali) tekshirsin, va agar bekor qilingan bo'lsa, **hozirgacha yig'ilgan** natijani darhol qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir son uchun: `select { case <-ctx.Done(): return natija; default: natija = append(natija, son*son) }`.
2. `default` bo'lmasa, `select` bekor qilinishini kutib, bloklanib qolardi — `default` orqali "hozircha bekor qilinmagan bo'lsa, davom et" tekshiruvi bloklanmasdan amalga oshadi.

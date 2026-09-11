# 11 — Hash Table

## THEORY

Teatr garderobini tasavvur qiling: kiyimingizni topshirganda sizga raqamli chek berishadi, va garderobchi o'sha raqamga qarab, kiyimingiz turgan **aniq katakni** darhol topadi — butun garderobni "qidirib" o'tirmaydi. "Maps" darsida ko'rgan Go'ning o'rnatilgan `map`i aynan shunday, juda tez ishlaydi — va sabab shu: uning ichida **hash table (xesh jadvali)** degan tuzilma yotadi. Keling, uning ichida qanday ishlashini qo'lda qurib ko'rib tushunamiz.

**Asosiy g'oya — kalitni "raqamga" aylantirish.** Hash table kalitni (masalan, matnni) maxsus **hash funksiya** orqali songa aylantiradi, va shu son orqali qaysi "katakcha"da (bucket, chelak) saqlanishini aniqlaydi:

```go
func oddiyHash(kalit string) int {
	yigindi := 0
	for _, harf := range kalit {
		yigindi += int(harf) // har bir harfning kod qiymatini qo'shamiz
	}
	return yigindi % 10 // 10 ta chelakdan biriga tushiramiz
}
```

**To'qnashuv (collision) — ikkita kalit bir xil chelakka tushishi mumkin.** Turli kalitlar ba'zan bir xil hash qiymatini berishi mumkin (masalan, ikki xil so'zning harflari yig'indisi teng chiqishi mumkin). Buning yechimi — har bir chelakda **bir nechta** juftlikni saqlashga imkon berish (bunga **zanjirlash**, chaining, deyiladi):

```go
type Juftlik struct {
	Kalit  string
	Qiymat int
}

type HashJadval struct {
	chelaklar [10][]Juftlik // har bir chelak — Juftlik'lar slice'i
}

func (h *HashJadval) Qoy(kalit string, qiymat int) {
	indeks := oddiyHash(kalit)
	h.chelaklar[indeks] = append(h.chelaklar[indeks], Juftlik{Kalit: kalit, Qiymat: qiymat})
}
```

**Qidirish — to'g'ri chelakni topib, ichida qidirish.** Kalitni qidirish uchun, avval hash funksiya orqali **qaysi chelakda** ekanligini aniqlaymiz, keyin faqat **o'sha chelak ichida** (butun jadvalda emas!) kalitni qidiramiz:

```go
func (h *HashJadval) Ol(kalit string) (int, bool) {
	indeks := oddiyHash(kalit)
	for _, juftlik := range h.chelaklar[indeks] {
		if juftlik.Kalit == kalit {
			return juftlik.Qiymat, true
		}
	}
	return 0, false
}
```

**Nega bu tez.** Agar har bir chelakda kam sonli juftlik bo'lsa (yaxshi hash funksiya, ma'lumot yetarlicha "tarqalgan" bo'lsa), qidiruv deyarli **darhol** (`O(1)`ga yaqin) bo'ladi — chunki siz to'g'ridan-to'g'ri kerakli chelakka "sakraysiz", butun jadvalni aylanib chiqishga hojat yo'q. Aynan shu tamoyil, ancha murakkab va optimallashtirilgan ko'rinishda, Go'ning o'rnatilgan `map` turi ichida ishlaydi — shuning uchun amaliy kodda o'zingiz hash table qurishning hojati yo'q, oddiygina `map[K]V`dan foydalanaverasiz ("Maps" darsida ko'rganingizdek). Bu darsning maqsadi — o'sha "sehr" aslida qanday ishlashini tushunish.

## EXAMPLE

```go
package main

import "fmt"

type Juftlik struct {
	Kalit  string
	Qiymat int
}

type HashJadval struct {
	chelaklar [10][]Juftlik
}

func oddiyHash(kalit string) int {
	yigindi := 0
	for _, harf := range kalit {
		yigindi += int(harf)
	}
	return yigindi % 10
}

func (h *HashJadval) Qoy(kalit string, qiymat int) {
	indeks := oddiyHash(kalit)
	h.chelaklar[indeks] = append(h.chelaklar[indeks], Juftlik{Kalit: kalit, Qiymat: qiymat})
}

func main() {
	h := &HashJadval{}
	h.Qoy("olma", 100)
	h.Qoy("banan", 200)

	fmt.Println(h.Ol("olma"))
	fmt.Println(h.Ol("uzum"))
}
```

Natija:

```
100 true
0 false
```

## TASK

`HashJadval` struct'i, `oddiyHash` funksiyasi va `Qoy` methodi (hammasi allaqachon yozilgan) berilgan. `Ol(kalit string) (int, bool)` methodini shunday to'ldiringki, u berilgan `kalit`ga mos qiymatni qaytarsin (agar topilsa `qiymat, true`, topilmasa `0, false`).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Avval to'g'ri chelakni toping: `indeks := oddiyHash(kalit)`, keyin `for _, juftlik := range h.chelaklar[indeks] { ... }` bilan o'sha chelak ichida qidiring.
2. `if juftlik.Kalit == kalit { return juftlik.Qiymat, true }` tsikl ichida, tsikldan keyin `return 0, false`.

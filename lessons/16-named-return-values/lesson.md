# 16 — Named Return Values

## THEORY

Oldingi darsda funksiya bir nechta qiymat qaytarishini ko'rdik. Endi tasavvur qiling: bozor tarozisi og'irlik va narxni qaytarganda, ular allaqachon "og'irlik: ..." va "narx: ..." deb **yorliqlangan** bo'lsa, natijani o'qish yanada osonlashadi. Go'da ham qaytariladigan qiymatlarga oldindan **nom** berish mumkin — bular funksiya boshida avtomatik e'lon qilingan, standart (zero) qiymat bilan boshlanadigan oddiy o'zgaruvchidek ishlaydi:

```go
func bolishVaQoldiq(a, b int) (bolinma, qoldiq int) {
	bolinma = a / b
	qoldiq = a % b
	return
}
```

Diqqat qiling: `return` so'zidan keyin hech narsa yozilmagan — bu **"yalang'och" (naked) return** deyiladi. Go avtomatik ravishda `bolinma` va `qoldiq` nomli o'zgaruvchilarning **joriy qiymatlarini** qaytaradi — funksiya ichida qayerda bo'lishidan qat'iy nazar, chunki ular allaqachon funksiya boshida "e'lon qilingan" holatda.

Named return values ikki narsaga foydali:
1. Funksiya nima qaytarishini **hujjatlashtiradi** — `(bolinma, qoldiq int)` o'qigan odamga darhol tushunarli, oddiy `(int, int)`dan ko'ra (`(int, int)` — ikkalasi ham xuddi shu tur, qaysi biri nima ekanligi imzoning o'zidan ko'rinmaydi).
2. Uzun funksiyalarda, ayniqsa `defer` bilan birga ishlatilganda (keyingi darsda ko'ramiz), qulaylik yaratadi — bu haqda pastda batafsil.

**Named return va `defer` bog'liqligi — kelajakda foydali bo'ladigan bir "aha" fakt.** Named return qiymati funksiya ichida oddiy o'zgaruvchi bo'lgani uchun, uni funksiya haqiqatan tugashidan oldin ishga tushadigan `defer` (17-darsda ko'ramiz) ham **o'zgartirib qo'yishi** mumkin — bu odatda xatolikni ushlab, natijani "tuzatib qo'yish" kerak bo'lgan murakkabroq holatlarda ishlatiladi. Hozircha buni chuqur bilish shart emas, lekin bu — named return'lar nafaqat "chiroyliroq ko'rinish" uchun emas, balki haqiqiy imkoniyatlar ochib berishini bilib qo'yish uchun foydali fakt.

**Muhim ogohlantirish: uzun funksiyalarda "yalang'och" `return` o'qishni qiyinlashtirishi mumkin**, chunki qaysi qiymat qaytarilayotgani darhol ko'rinmaydi — o'quvchi funksiyaning butun tanasini o'qib chiqib, o'sha nomli o'zgaruvchi qayerda va nimaga tenglashtirilganini qidirishi kerak bo'ladi. Shuning uchun ko'p Go dasturchilari qisqa funksiyalardagina named return'dan foydalanadi, uzun funksiyalarda esa oddiy `return qiymat1, qiymat2` yozishni afzal ko'radi — bu ikkalasi ham to'g'ri, tanlov kodning uzunligi va aniqligiga bog'liq.

## EXAMPLE

```go
package main

import "fmt"

func kvadrat(son int) (natija int) {
	natija = son * son
	return
}

func minMax(royxat []int) (kichik, katta int) {
	kichik, katta = royxat[0], royxat[0]
	for _, son := range royxat {
		if son < kichik {
			kichik = son
		}
		if son > katta {
			katta = son
		}
	}
	return
}

func main() {
	fmt.Println(kvadrat(5))

	k, m := minMax([]int{4, 8, 1, 9, 3})
	fmt.Println(k, m)
}
```

Natija:

```
25
1 9
```

## TASK

`minMax(royxat []int) (kichik, katta int)` funksiyasi berilgan — qaytish qiymatlari allaqachon nomlangan (`kichik`, `katta`). Funksiya tanasini shunday to'ldiringki, ro'yxatdagi eng kichik va eng katta sonlarni hisoblab, **yalang'och `return`** orqali qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Boshlang'ich qiymat sifatida `kichik = royxat[0]` va `katta = royxat[0]` deb belgilang, keyin `for _, son := range royxat` bilan aylanib, kerak bo'lsa yangilang.
2. Oxirida faqat `return` yozing (hech narsa qo'shmasdan) — `kichik` va `katta` avtomatik qaytariladi.

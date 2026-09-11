# 04 — If / Else

## THEORY

Har kuni hayotda shartlarga qarab qaror qabul qilamiz: "agar yomg'ir yog'sa — soyabon olaman, aks holda — olmayman". Dasturlashda ham xuddi shu mantiq ishlatiladi, va u `if`/`else` deb ataladi.

**Asosiy ko'rinishi:**

```go
if age >= 18 {
	fmt.Println("Voyaga yetgan")
} else {
	fmt.Println("Voyaga yetmagan")
}
```

Bu yerda `age >= 18` — shart. U tekshiriladi, natija albatta `true` yoki `false` bo'ladi (Boolean darsida bu haqda batafsil gaplashamiz). Agar `true` bo'lsa — `if` blokidagi kod ishlaydi, agar `false` bo'lsa — `else` blokidagi kod ishlaydi.

Ikkita muhim sintaktik qoida:
- Shart atrofida qavs `()` **shart emas** (boshqa ko'plab tillardan farqli — u yerda `if (age >= 18)` deb yoziladi, Go'da esa shunchaki `if age >= 18`).
- Figurali qavslar `{}` esa **har doim majburiy**, hatto blok ichida bitta qator bo'lsa ham.

**Bir nechta shartni ketma-ket tekshirish — `else if`:**

```go
if ball >= 90 {
	fmt.Println("A'lo")
} else if ball >= 70 {
	fmt.Println("Yaxshi")
} else if ball >= 50 {
	fmt.Println("Qoniqarli")
} else {
	fmt.Println("Yomon")
}
```

Bu yerda shartlar **yuqoridan pastga** ketma-ket tekshiriladi, va birinchi `true` chiqqan shartning bloki ishlaydi, qolganlari tekshirilmaydi ham. Shuning uchun tartib muhim: agar `ball = 95` bo'lsa, birinchi shart (`ball >= 90`) darrov `true` bo'lgani uchun "A'lo" chiqadi, qolgan shartlar umuman tekshirilmaydi.

**`if` ichida `return` bilan erta chiqish.** Funksiya qiymat qaytarganda, ko'pincha `else` yozishning hojati yo'q — `if` bloki ichida `return` bo'lsa, funksiya o'sha yerdayoq tugaydi, va undan keyingi qator avtomatik ravishda "aks holda" ma'nosini beradi:

```go
func toifa(age int) string {
	if age >= 18 {
		return "Voyaga yetgan"
	}
	return "Voyaga yetmagan"
}
```

Bu ikkinchi `return` qatori faqat `age >= 18` **yolg'on** bo'lgandagina bajariladi — chunki agar rost bo'lsa, funksiya birinchi `return` bilan allaqachon tugab bo'lgan bo'ladi. Bu naqsh Go kodida juda ko'p uchraydi, chunki u ortiqcha `else` bloklarisiz kodni tekisroq va o'qishga qulayroq qiladi.

**Shart ichida qisqa boshlang'ich amal (initializer).** Go'da yana bir foydali imkoniyat bor: `if`dan oldin, xuddi shu qatorda, kichik bir amal bajarib, natijasini darhol shu shart doirasida ishlatish mumkin:

```go
if natija := hisobla(); natija > 100 {
	fmt.Println("Katta natija:", natija)
}
```

Bu yerda `natija` o'zgaruvchisi faqat shu `if` (va uning `else` qismi) doirasida yashaydi — undan keyin u "yo'qoladi" (bu haqda Scope darsida batafsil gaplashamiz). Bu naqsh, masalan, bir funksiyani chaqirib, uning natijasini darhol tekshirib ko'rishda juda qulay — vaqtinchalik o'zgaruvchini keraksiz joyda "sarg'aytirib" qo'ymaydi.

**Ichma-ich `if`lar.** `if` blokining ichiga yana `if` yozish mumkin — bu ikkita shartni ketma-ket, ichkariga qarab tekshirish kerak bo'lganda ishlatiladi:

```go
if yosh >= 18 {
	if haydovchilikGuvohi {
		fmt.Println("Haydashi mumkin")
	} else {
		fmt.Println("Guvohnoma yo'q")
	}
} else {
	fmt.Println("Yoshi yetmagan")
}
```

(Aslida bu misolni `&&` mantiqiy amali bilan bitta shartga ham birlashtirish mumkin — buni Boolean darsida ko'ramiz, lekin ba'zan ichma-ich `if` o'qishga aniqroq bo'ladi.)

## EXAMPLE

```go
package main

import "fmt"

func baho(ball int) string {
	if ball >= 90 {
		return "A'lo"
	} else if ball >= 70 {
		return "Yaxshi"
	}
	return "Qoniqarli"
}

func main() {
	fmt.Println(baho(95))
	fmt.Println(baho(75))
	fmt.Println(baho(60))
}
```

Natija:

```
A'lo
Yaxshi
Qoniqarli
```

## TASK

Quyida `toifa(age int) string` funksiyasi berilgan. Uni shunday to'ldiringki:

1. Agar `age` 18 yoki undan katta bo'lsa — `"Voyaga yetgan"` qaytarsin.
2. Aks holda — `"Voyaga yetmagan"` qaytarsin.

`main()` funksiyasini o'zgartirish shart emas — u sizning `toifa` funksiyangizni chaqirib, natijani ekranga chiqaradi.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `if age >= 18 { return "Voyaga yetgan" }` qatoridan boshlang.
2. `if` blokidan keyin alohida `else` yozish shart emas — `if` ichida `return` bo'lsa, funksiya davomida keyingi qatorni oddiy `return "Voyaga yetmagan"` deb yozsangiz ham yetarli.

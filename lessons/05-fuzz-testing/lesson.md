# 05 — Fuzz Testing

## THEORY

Yangi qulfni sinash uchun, siz o'zingiz o'ylab topgan bir nechta kalitni urinib ko'rasiz. Endi tasavvur qiling — maxsus mashina bor, u **minglab tasodifiy shakldagi** kalitlarni avtomatik yasab, hammasini urinib ko'radi, va agar birortasi kutilmagan tarzda qulfni ochib yuborsa (yoki uni buzib qo'ysa), sizga xabar beradi. **Fuzz testing** — dasturlashda aynan shu: siz o'zingiz **o'ylab topmagan**, g'alati yoki chekka (edge case) kirishlarni **avtomatik ravishda** generatsiya qilib, kodni sinash.

**Nega bu kerak — inson faqat o'zi o'ylagan holatlarni test qiladi.** Oldingi darslarda siz har doim **o'zingiz tanlagan** aniq holatlar bilan test yozdingiz (masalan, `5`, `0`, `-3`). Lekin ba'zi xatolar faqat **kutilmagan** kirishlarda (juda uzun matn, g'alati belgilar, bo'sh qiymatlar) namoyon bo'ladi — inson bularning **hammasini** o'ylab topa olmaydi. Fuzzing — aynan shu "ko'r nuqta"ni yopadi.

**Fuzz test qanday yoziladi.** Nomi `Fuzz` bilan boshlanadi, parametri `*testing.F`:

```go
func FuzzTeskariAylantir(f *testing.F) {
	f.Add("salom") // "urug' (seed)" — boshlang'ich, ma'lum misol
	f.Add("")

	f.Fuzz(func(t *testing.T, s string) {
		teskari := TeskariAylantir(s)
		yanaTeskari := TeskariAylantir(teskari)
		if yanaTeskari != s {
			t.Errorf("ikki marta teskari aylantirish asl holatga qaytmadi: %q", s)
		}
	})
}
```

`f.Add(...)` — Go'ga "shu misollardan boshla" deb aytadi (bular — **urug' (seed) qiymatlar**). `f.Fuzz(func(t *testing.T, s string) {...})` — ichkarida Go **avtomatik ravishda** minglab turli `s` qiymatlarini (urug'lardan "mutatsiya" qilib hosil qilingan) sinab ko'radi.

**Ishga tushirish — `go test -fuzz`.** Oddiy `go test` fuzz test'larni faqat **urug' qiymatlar** bilan (oddiy unit test kabi) ishga tushiradi — haqiqiy, uzoq davom etadigan tasodifiy fuzzing uchun maxsus bayroq kerak:

```bash
go test -fuzz=FuzzTeskariAylantir
```

**Klassik fuzzing misoli — UTF-8 va rune'lar.** "Strings" darsida qisqacha eslatilgan nozik joyni eslang: string'ning bayt uzunligi va undagi "belgilar" soni har doim bir xil emas (agar lotin-kirillik bo'lmagan belgilar bo'lsa). Agar `TeskariAylantir` funksiyasi string'ni **bayt-baytlab** (rune'larga e'tibor bermasdan) teskari aylantirsa, u lotin harflarida to'g'ri ishlaydi, lekin ko'p baytli belgilarda (masalan, kirill yoki emoji) **buziladi**. Aynan shunday xatolarni inson qo'lda yozgan bir nechta test holati bilan **ko'pincha topa olmaydi** — lekin fuzzing tasodifiy kirishlar orasidan buni tezda "ushlab" beradi.

## EXAMPLE

```go
package main

import (
	"fmt"
	"testing"
)

func TeskariAylantir(s string) string {
	runlar := []rune(s)
	uzunlik := len(runlar)
	natija := make([]rune, uzunlik)
	for i, r := range runlar {
		natija[uzunlik-1-i] = r
	}
	return string(natija)
}

func FuzzTeskariAylantir(f *testing.F) {
	f.Add("salom")
	f.Fuzz(func(t *testing.T, s string) {
		ikki := TeskariAylantir(TeskariAylantir(s))
		if ikki != s {
			t.Errorf("ikki marta teskari aylantirish %q ni qaytarmadi", s)
		}
	})
}

func main() {
	fmt.Println(TeskariAylantir("salom"))
}
```

Natija:

```
molas
```

## TASK

`TeskariAylantir(s string) string` funksiyasi berilgan. Uni shunday to'ldiringki, u `s`ni **rune'lar bo'yicha** (baytlar emas — ko'p baytli Unicode belgilar to'g'ri ishlashi uchun) teskari aylantirib qaytarsin (masalan, `TeskariAylantir("salom")` → `"molas"`).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `runlar := []rune(s)` — string'ni rune'lar slice'iga aylantiring (bu, har bir belgini to'g'ri "birlik" sifatida ajratadi, hatto ko'p baytli bo'lsa ham).
2. Yangi slice yarating va elementlarni teskari tartibda joylashtiring: `natija[uzunlik-1-i] = r`, so'ng `return string(natija)`.

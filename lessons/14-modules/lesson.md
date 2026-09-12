# 14 — Modules

## THEORY

"Packages" darsida ko'rgan paketlar — bu loyihaning **ichki** tashkiloti. Endi savol: agar loyihangiz **boshqa birovning** kodidan (masalan, internetdagi ochiq kutubxonadan) foydalanmoqchi bo'lsa-chi? Yoki, aksincha, sizning loyihangizni boshqalar ishlatmoqchi bo'lsa-chi? Buning uchun **module (modul)** tushunchasi bor — bu, butun loyihangizni **nomlangan, versiyalangan birlik** sifatida belgilaydi.

**`go.mod` fayli — modulning "pasporti".** Har bir Go moduli o'z ildiz papkasida `go.mod` nomli faylga ega:

```
module go-learn

go 1.27.0
```

- **`module go-learn`** — modulning nomi (yoki yo'li). Agar modul boshqalar tomonidan `import` qilinishi kerak bo'lsa, bu odatda haqiqiy manzil bo'ladi, masalan `github.com/foydalanuvchi/loyiha`.
- **`go 1.27.0`** — bu modul yozilgan/sinaldigan Go tilining minimal versiyasi.

**`go mod init` — yangi modul yaratish.** Yangi loyihani boshlaganda, terminalda:

```bash
go mod init loyiham
```

Bu, `go.mod` faylini avtomatik yaratadi, va shu paytdan boshlab loyihangiz **modul** hisoblanadi — uning ichidagi barcha paketlar bir-birini `loyiham/pastki-papka` kabi yo'llar orqali import qila oladi ("Packages" darsida ko'rgan naqsh).

**Tashqi kutubxona qo'shish.** Agar boshqa birovning kutubxonasidan foydalanmoqchi bo'lsangiz:

```bash
go get github.com/kimdir/kutubxona
```

Bu buyruq kerakli kodni yuklab oladi, va **`go.sum`** nomli ikkinchi faylni yaratadi (yoki yangilaydi) — bu fayl yuklab olingan kodning **aniq versiyasi va yaxlitligini** (hash orqali) qayd etadi, shunda loyihangiz boshqa kompyuterda ham **xuddi shu** kodni ishlatishini kafolatlaydi.

**Semantik versiyalash (semver).** Go kutubxonalari odatda `v1.2.3` kabi uch qismli raqam bilan versiyalanadi:

- **1-son (major)** — katta, orqaga mos kelmaydigan o'zgarishlar bo'lsa oshadi.
- **2-son (minor)** — yangi, lekin orqaga mos keladigan imkoniyat qo'shilsa oshadi.
- **3-son (patch)** — kichik tuzatishlar (bug fix) uchun oshadi.

Bu qoida — kutubxona yangilanganda, sizning kodingiz **kutilmaganda buzilib qolish** xavfini oldindan baholash imkonini beradi: agar faqat patch versiyasi o'zgargan bo'lsa, xavfsiz deb hisoblanadi; major versiya o'zgargan bo'lsa, ehtiyot bo'lish kerak.

**Bu kursning o'zi ham — bitta modul.** Loyihangizni ochib ko'rsangiz, uning ildizida ham aynan shunday `go.mod` fayli borligini ko'rasiz — bu kursning butun kodi (`internal/cli`, `internal/lesson` va h.k.) — bittagina, `go-learn` nomli modulning turli paketlari.

## EXAMPLE

```go
package main

import "fmt"

// versiyaTaqqosla ikkita "major.minor.patch" versiya satrini
// solishtiradi: v1 v2'dan katta bo'lsa true qaytaradi.
func versiyaTaqqosla(v1, v2 string) bool {
	return v1 > v2 // sodda, lug'aviy solishtirish (haqiqiy semver mantig'i murakkabroq)
}

func main() {
	fmt.Println(versiyaTaqqosla("2.0.0", "1.9.9"))
}
```

Natija:

```
true
```

## TASK

`versiyaEskimi(joriy, minimal string) bool` funksiyasi berilgan. Bu yerda `joriy` va `minimal` — `"1.20.0"` kabi versiya satrlari. Uni shunday to'ldiringki, u `joriy` versiya `minimal`dan **kichik** (ya'ni eskirgan) bo'lsa `true`, aks holda `false` qaytarsin. Solishtirish uchun oddiy **lug'aviy (string) solishtirish** yetarli deb hisoblang (`<` operatori).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `return joriy < minimal` — Go'da stringlarni `<` bilan to'g'ridan-to'g'ri solishtirish mumkin (alifbo/lug'aviy tartibda).
2. Funksiya tanasi shu bitta qatordan iborat bo'lishi mumkin.

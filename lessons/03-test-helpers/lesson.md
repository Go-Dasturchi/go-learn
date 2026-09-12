# 03 — Test Helpers

## THEORY

Kontsert oldidan sahna orqasidagi xodim har bir chiqishdan oldin mikrofonni sozlab, yorug'lik va pardani tekshirib qo'yadi — aktyorlarning o'zi buni har safar qaytadan qilishiga hojat yo'q. **Test helper (yordamchi funksiya)** — testlarda aynan shu vazifani bajaradi: bir nechta testda **takrorlanadigan** tekshiruv yoki tayyorgarlik kodini, alohida funksiyaga chiqarib qo'yish.

**Oddiy yordamchi funksiya:**

```go
func assertEqual(t *testing.T, got, kutilgan int) {
	if got != kutilgan {
		t.Errorf("Kutilgan %d, lekin %d keldi", kutilgan, got)
	}
}

func TestKvadrat(t *testing.T) {
	assertEqual(t, Kvadrat(5), 25)
	assertEqual(t, Kvadrat(0), 0)
	assertEqual(t, Kvadrat(-3), 9)
}
```

Bu yerda `assertEqual` — "tekshirish" mantig'ini bitta joyga jamlab, har bir testda `if ... { t.Errorf(...) }` deb qayta-qayta yozishning oldini oladi.

**`t.Helper()` — xato joyini to'g'ri ko'rsatish uchun muhim usul.** Agar `assertEqual` ichida xato yuz bersa, va `t.Helper()` chaqirilmagan bo'lsa, Go xato qatorini **`assertEqual` ichidan** ko'rsatadi — bu unchalik foydali emas, chunki chindan ham muammoli joy — uni **chaqirgan** test qatori. `t.Helper()` shu muammoni hal qiladi:

```go
func assertEqual(t *testing.T, got, kutilgan int) {
	t.Helper() // MUHIM: xato bo'lsa, chaqiruvchi qatorni ko'rsatadi
	if got != kutilgan {
		t.Errorf("Kutilgan %d, lekin %d keldi", kutilgan, got)
	}
}
```

`t.Helper()` — Go'ga "bu funksiya — yordamchi, xatolik chiqsa, mening emas, meni **chaqirgan** joyning qator raqamini ko'rsat" deb aytadi. Bu, katta test fayllarida, aynan qaysi test holatida muammo borligini tezda topishga yordam beradi.

**Nega bu foydali — "Higher-Order Functions" va "Closures" darslaridagi g'oyaning davomi.** Test helper'lar — bu, aslida, oddiy funksiyalar, va ular xuddi boshqa har qanday Go funksiyasi kabi — parametr qabul qiladi, natija qaytaradi (yoki hech narsa qaytarmaydi, `t.Errorf` orqali to'g'ridan-to'g'ri "shikoyat" qiladi). Ularni yozish — kodni DRY (Don't Repeat Yourself — o'zingizni takrorlamang) qilib, testlarni ham, oddiy dastur kodi kabi, toza va o'qish oson qilib saqlaydi.

## EXAMPLE

```go
package main

import (
	"fmt"
	"testing"
)

func Validatsiya(email string) bool {
	return len(email) > 3 && containsAt(email) && containsDot(email)
}

func containsAt(s string) bool {
	for _, r := range s {
		if r == '@' {
			return true
		}
	}
	return false
}

func containsDot(s string) bool {
	for _, r := range s {
		if r == '.' {
			return true
		}
	}
	return false
}

func assertValid(t *testing.T, email string, kutilgan bool) {
	t.Helper()
	got := Validatsiya(email)
	if got != kutilgan {
		t.Errorf("Validatsiya(%q) = %v, kutilgan %v", email, got, kutilgan)
	}
}

func TestValidatsiya(t *testing.T) {
	assertValid(t, "ali@mail.uz", true)
	assertValid(t, "notorgri", false)
}

func main() {
	fmt.Println(Validatsiya("ali@mail.uz"))
	fmt.Println(Validatsiya("notorgri"))
}
```

Natija:

```
true
false
```

## TASK

`Validatsiya(email string) bool` funksiyasi berilgan (yordamchi `containsAt` va `containsDot` funksiyalari allaqachon yozilgan). Uni shunday to'ldiringki, u `email` uzunligi 3dan katta **va** ichida `@` **va** `.` belgilari bo'lsa `true`, aks holda `false` qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Uchala shart ham `&&` bilan bog'lanadi: uzunlik, `@` mavjudligi, `.` mavjudligi.
2. `return len(email) > 3 && containsAt(email) && containsDot(email)`.

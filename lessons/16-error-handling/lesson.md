# 16 — Error Handling

## THEORY

"Errors" va "Custom Errors" darslarida xatolik yaratish va o'z turingizni belgilashni ko'rgan edingiz. Endi savol: agar funksiyalar zanjiri bo'ylab xatolik "yuqoriga" o'tkazilsa, va yo'l-yo'lakay unga **qo'shimcha ma'lumot** (masalan, "qaysi funksiyada, nima qilayotganda" yuz bergani) qo'shish kerak bo'lsa-chi? Bu — **error wrapping (xatolikni o'rash)** deb ataladi.

**Sentinel error — oldindan e'lon qilingan, "belgi" xatolik.** Ba'zan xatolikning **aynan qaysi turi** ekanligini keyinroq tekshirish kerak bo'ladi (masalan, "bu — 'topilmadi' xatosimi, yoki boshqasimi?"). Buning uchun paket darajasida, oldindan bitta **"sentinel" (nazoratchi)** xatolik e'lon qilinadi:

```go
var ErrTopilmadi = errors.New("topilmadi")

func qidir(royxat []string, maqsad string) (int, error) {
	for i, s := range royxat {
		if s == maqsad {
			return i, nil
		}
	}
	return -1, ErrTopilmadi
}
```

**Xatolikni "o'rash" — `fmt.Errorf` va `%w`.** Agar xatolikka qo'shimcha kontekst qo'shmoqchi bo'lsangiz, lekin asl xatolikning "izini" ham yo'qotmasangiz, `%w` (`%s`/`%v` emas) joy egallovchisi ishlatiladi:

```go
_, err := qidir(royxat, "Vali")
if err != nil {
	return fmt.Errorf("foydalanuvchini qidirishda xato: %w", err)
}
```

`%w` — bu xatolikni shunchaki matn qilib "yopishtirmaydi", balki asl xatolikni **ichida saqlab qoladi** ("o'raydi") — bu, keyinroq uni **qayta ochib ko'rish** imkonini beradi.

**`errors.Is` — o'ralgan xatolik ichida sentinel'ni topish.** Agar xatolik bir necha marta "o'ralgan" bo'lsa ham (funksiyalar zanjiri bo'ylab), `errors.Is` uning ichida **aynan qaysi sentinel xatolik** yashiringanini tekshira oladi:

```go
err := fmt.Errorf("tashqi qatlam: %w", fmt.Errorf("ichki qatlam: %w", ErrTopilmadi))

if errors.Is(err, ErrTopilmadi) {
	fmt.Println("Bu — 'topilmadi' xatosi edi, nechta qatlam bilan o'ralgan bo'lishidan qat'iy nazar")
}
```

Bu — oddiy `err == ErrTopilmadi` solishtirishdan ancha kuchliroq, chunki `==` faqat xatolik **hech qanday o'ralmagan, aynan o'zi** bo'lgandagina ishlaydi; `errors.Is` esa butun "o'rash zanjiri"ni ichkariga qarab tekshirib chiqadi.

**Umumiy qoida — Go'da xatolik "e'tiborsiz qoldirilmaydi".** "Errors" darsida ko'rgan `if err != nil` naqshini eslang — Go'da xatolikni **jim** o'tkazib yuborish (masalan, `_` bilan e'tiborsiz qoldirish) mumkin, lekin bu **yomon amaliyot** hisoblanadi. Yaxshi Go kodi har doim xatolikni **darhol tekshiradi** va unga mos munosabatda bo'ladi — qaytarish, log qilish, yoki (kamdan-kam) qayta urinish.

## EXAMPLE

```go
package main

import (
	"errors"
	"fmt"
)

var ErrTopilmadi = errors.New("topilmadi")

func qidir(royxat []string, maqsad string) (int, error) {
	for i, s := range royxat {
		if s == maqsad {
			return i, nil
		}
	}
	return -1, ErrTopilmadi
}

func foydalanuvchiniTop(ism string) error {
	royxat := []string{"Ali", "Vali"}
	_, err := qidir(royxat, ism)
	if err != nil {
		return fmt.Errorf("foydalanuvchini qidirishda xato: %w", err)
	}
	return nil
}

func main() {
	err := foydalanuvchiniTop("Guli")
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrTopilmadi))
}
```

Natija:

```
foydalanuvchini qidirishda xato: topilmadi
true
```

## TASK

`ErrTopilmadi` sentinel xatoligi va `qidir` funksiyasi berilgan. `foydalanuvchiniTop(ism string) error` funksiyasini shunday to'ldiringki, u `qidir`ni chaqirib, agar xatolik bo'lsa, uni `"foydalanuvchini qidirishda xato: %w"` formati bilan **o'rab** qaytarsin (bu yerda `%w` — `qidir`dan kelgan asl xatolik). Agar xatolik bo'lmasa, `nil` qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Avval `royxat := []string{"Ali", "Vali"}` yarating, keyin `_, err := qidir(royxat, ism)` — natijaning o'zi kerak emas, faqat xatolik.
2. `if err != nil { return fmt.Errorf("foydalanuvchini qidirishda xato: %w", err) }`, aks holda `return nil`.

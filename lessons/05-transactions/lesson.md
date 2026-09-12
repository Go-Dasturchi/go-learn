# 05 — Transactions

## THEORY

Bankomatdan pul o'tkazmasi qilayotganingizni tasavvur qiling: sizning hisobingizdan pul **ayiriladi**, qabul qiluvchining hisobiga **qo'shiladi**. Agar birinchi amal (ayirish) muvaffaqiyatli bo'lib, lekin ikkinchisi (qo'shish) xatolik tufayli bajarilmasa — pul "yo'qolib qoladi"! **Transaction (tranzaksiya)** — aynan shu muammoni hal qiladi: u bir nechta amalni **"hammasi yoki hech biri"** (all-or-nothing) qoidasi bilan bog'laydi.

**SQL'da tranzaksiya — `BEGIN`, `COMMIT`, `ROLLBACK`:**

```sql
BEGIN;

UPDATE hisoblar SET balans = balans - 100 WHERE id = 1;
UPDATE hisoblar SET balans = balans + 100 WHERE id = 2;

COMMIT; -- ikkala UPDATE ham "yakuniy" bo'ladi, birga
```

Agar ikkinchi `UPDATE` biror sababga ko'ra muvaffaqiyatsiz bo'lsa, `COMMIT` o'rniga `ROLLBACK` chaqiriladi — bu, **BEGIN'dan beri qilingan barcha o'zgarishlarni bekor qiladi**, xuddi hech narsa bo'lmagandek.

**Go'da tranzaksiya — `db.Begin()`:**

```go
tx, err := db.Begin()
if err != nil {
	return err
}

_, err = tx.Exec("UPDATE hisoblar SET balans = balans - 100 WHERE id = $1", 1)
if err != nil {
	tx.Rollback() // xatolik — hammasini bekor qilamiz
	return err
}

_, err = tx.Exec("UPDATE hisoblar SET balans = balans + 100 WHERE id = $1", 2)
if err != nil {
	tx.Rollback()
	return err
}

return tx.Commit() // ikkalasi ham muvaffaqiyatli — yakunlaymiz
```

`tx` — `db`ga o'xshaydi (`Exec`, `Query` methodlari bor), lekin uning barcha amallari **bitta tranzaksiya ichida** to'planadi, va faqat `Commit()` chaqirilgandagina "haqiqiy" bo'ladi.

**Nega bu muhim — ma'lumotlar izchilligi (consistency).** Tranzaksiyasiz, dastur qulab tushishi, tarmoq uzilishi, yoki oddiy dastur xatosi — ma'lumotlar bazasini **noto'g'ri, yarim bajarilgan** holatda qoldirishi mumkin (masalan, pul bir hisobdan yo'qoldi, lekin boshqa hisobga qo'shilmadi). Tranzaksiyalar bunday holatlarning **hech qachon** yuz bermasligini kafolatlaydi.

**Go'da bu naqshni, real bazasiz, xotirada sinash.** Xuddi "CRUD" darsida ko'rgan naqshga o'xshab, tranzaksiyaning **mantig'ini** (agar bir qadam muvaffaqiyatsiz bo'lsa, avvalgi qadamlarni "qaytarish") xotiradagi oddiy struct bilan mashq qilish mumkin:

```go
func pulOtkazish(hisoblar map[string]int, kimdan, kimga string, summa int) error {
	if hisoblar[kimdan] < summa {
		return errors.New("mablag' yetarli emas") // hech narsa o'zgarmaydi — "ROLLBACK"
	}
	hisoblar[kimdan] -= summa
	hisoblar[kimga] += summa
	return nil // ikkalasi ham bajarildi — "COMMIT"
}
```

Bu yerda "tranzaksiya" — shunchaki, **tekshiruvni amaldan oldin** qilish orqali amalga oshirilgan: agar mablag' yetarli bo'lmasa, hech qanday o'zgarish qilinmaydi (hisoblar butunlay o'zgarishsiz qoladi) — bu, natijada, "hammasi yoki hech biri" qoidasini ta'minlaydi.

## EXAMPLE

```go
package main

import (
	"errors"
	"fmt"
)

func pulOtkazish(hisoblar map[string]int, kimdan, kimga string, summa int) error {
	if hisoblar[kimdan] < summa {
		return errors.New("mablag' yetarli emas")
	}
	hisoblar[kimdan] -= summa
	hisoblar[kimga] += summa
	return nil
}

func main() {
	hisoblar := map[string]int{"Ali": 500, "Vali": 200}

	err := pulOtkazish(hisoblar, "Ali", "Vali", 300)
	fmt.Println(err, hisoblar)

	err = pulOtkazish(hisoblar, "Ali", "Vali", 1000)
	fmt.Println(err, hisoblar)
}
```

Natija:

```
<nil> map[Ali:200 Vali:500]
mablag' yetarli emas map[Ali:200 Vali:500]
```

## TASK

`pulOtkazish(hisoblar map[string]int, kimdan, kimga string, summa int) error` funksiyasi berilgan. Uni shunday to'ldiringki:

1. Agar `hisoblar[kimdan]` `summa`dan kichik bo'lsa, `errors.New("mablag' yetarli emas")` qaytarsin, **hech qanday o'zgartirish qilmasdan**.
2. Aks holda, `kimdan` hisobidan `summa`ni ayirib, `kimga` hisobiga qo'shib, `nil` qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `if hisoblar[kimdan] < summa { return errors.New("mablag' yetarli emas") }` — bu tekshiruv **amaldan oldin**, shuning uchun muvaffaqiyatsizlikda hech narsa o'zgarmaydi.
2. `hisoblar[kimdan] -= summa`, `hisoblar[kimga] += summa`, so'ng `return nil`.

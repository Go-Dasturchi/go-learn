# Defanging an IP Address

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **5 daqiqa**

## EXAMPLE

```
Input: address = "1.1.1.1"
Output: "1[.]1[.]1[.]1"
```

## TASK

Yaroqli (valid) IPv4 manzili `address` berilgan. Uning **"zararsizlantirilgan" (defanged)** versiyasini qaytaring, ya'ni har bir `"."` ni `"[.]"` bilan almashtiring.

Quyidagi funksiyani to'ldiring:

```go
func solve(address string) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu masalada tayyor satr almashtirish funksiyasidan foydalanish eng qulay yo'l.
2. strings paketida barcha '.' belgilarini '[.]' ga almashtiruvchi tayyor funksiyani qidiring va uni to'g'ridan-to'g'ri qo'llang.

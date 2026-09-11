# Longest valid parentheses

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: s = ")()())"
Output: 4
Tushuntirish: Eng uzun yaroqli qism-satr - "()()".
```

## TASK

Faqatgina `(` va `)` belgilaridan iborat `s` satri berilgan. Undagi to'g'ri (yaroqli) ochilib-yopilgan eng uzun qavslar ketma-ketligi uzunligini toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Qavslarning o'zini emas, balki ularning INDEKSLARINI saqlaydigan stek (stack) ishlatishni o'ylab ko'ring — stekning tepasi doim "hozirgacha yaroqli bo'lmagan oxirgi chegara" indeksini bildirsin.
2. Stekni -1 bilan boshlang (chegara sifatida). `(` uchraganda uning indeksini stekka qo'shing. `)` uchraganda stekdan bittani chiqaring: agar stek shundan keyin bo'sh qolsa, joriy indeksni yangi chegara sifatida stekka qo'shing; aks holda joriy indeks bilan stekning yangi tepasi orasidagi farq — shu nuqtada tugaydigan yaroqli qism-satr uzunligi, uni maksimal natija bilan solishtiring.

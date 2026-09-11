# Strong Password Checker

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **50 daqiqa**

## EXAMPLE

```
Input: password = "a"
Output: 5
Tushuntirish: 5 ta belgi qo'shish kerak: katta harf, raqam va yana 3 ta belgi.
```

## TASK

Kuchli parol uchun quyidagi shartlar bajarilishi lozim:

1. Uzunligi kamida 6 va ko'pi bilan 20 belgidan iborat bo'lsin.
2. Kamida bitta kichik harf, bitta katta harf va bitta raqam bo'lsin.
3. Qator-qator 3 yoki undan ortiq bir xil belgi bo'lmasin (masalan, `...aaa...` yoki `...BBB...`).

`password` satri berilgan. Uni kuchli parolga aylantirish uchun kerak bo'ladigan minimal amallar soni (qo'shish, o'chirish, almashtirish)ni toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(password string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

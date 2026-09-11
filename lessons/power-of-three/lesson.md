# Power of Three

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: n = 27
Output: true
Tushuntirish: 27 = 3^3
```

## TASK

Sizga butun son `n` berilgan. Agar u `3` ning darajasi bo'lsa `true` qaytaring, aks holda `false`.

Boshqacha aytganda, agar shunday bir butun `x` soni mavjud bo'lsaki, bunda $n == 3^x$ o'rinli bo'lsa, u holda `n` soni 3 ning darajasi hisoblanadi.

Sikllardan foydalanmasdan (rekursiya ham yo'q) buni hal qila olasizmi?

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

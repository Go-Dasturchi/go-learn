# Smallest Good Base

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **50 daqiqa**

## EXAMPLE

```
Input: n = "13"
Output: "3"
Tushuntirish: 3 asosida: 13 = 111 (3^2 + 3^1 + 3^0 = 9 + 3 + 1 = 13)
```

## TASK

Sizga satr ko'rinishida butun son `n` berilgan. `k` asosida `n` ni `111...1` (faqat 1 lardan tashkil topgan satr) sifatida ifodalash mumkin bo'lgan eng kichik `k` asosini satr sifatida qaytaring.

Boshqacha qilib aytganda, shunday `k` ni topingki, bunda:
$$n = k^0 + k^1 + k^2 + ... + k^m$$
tengligi $m > 1$ qiymat uchun o'rinli bo'lsin.

Quyidagi funksiyani to'ldiring:

```go
func solve(n string) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

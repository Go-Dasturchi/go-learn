# Count Primes

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: n = 10
Output: 4
Tushuntirish: 10 dan kichik tub sonlar: 2, 3, 5, 7 (jami 4 ta).
```

## TASK

Sizga butun son `n` berilgan. `n` dan qat'iy kichik (strictly less than `n`) bo'lgan tub sonlar (prime numbers) sonini qaytaring.

**Tub son (prime number)** — bu faqat 1 va o'ziga qoldiqsiz bo'linadigan 1 dan katta sondir.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

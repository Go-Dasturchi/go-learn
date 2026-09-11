# Super Ugly Number

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: n = 12, primes = [2,7,13,19]
Output: 32
Tushuntirish: Ro'yxat: [1,2,4,7,8,13,14,16,19,26,28,32,...]. 12-chi son 32.
```

## TASK

**O'ta xunuk son (Super Ugly Number)** deb tub ko'paytuvchilari berilgan `primes` to'plamidagi sonlardan iborat bo'lgan musbat butun songa aytiladi.

Sizga `n` va `primes` massivi berilgan. `n`-chi o'ta xunuk sonni toping va qaytaring.

Masalan, `primes = [2, 7, 13, 19]` bo'lsa, O'ta xunuk sonlar: `[1, 2, 4, 7, 8, 13, 14, 16, 19, 26, ...]`.

1 ning tub ko'paytuvchilari yo'q, u kichik holat sifatida kiradi.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int, primes []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

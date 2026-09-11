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

1. Har bir yangi "xunuk son" avvalgi xunuk sonlardan birini biror tub songa ko'paytirish orqali hosil bo'ladi — shuning uchun har bir tub son uchun alohida ko'rsatkich (pointer) saqlab boruvchi dinamik dasturlash (DP) haqida o'ylab ko'ring.
2. dp[0]=1 dan boshlang va har bir primes[j] uchun alohida ko'rsatkich (pointers[j]) yuriting; har qadamda barcha j lar bo'yicha `dp[pointers[j]] * primes[j]` qiymatlarining eng kichigini keyingi dp elementi sifatida oling, so'ng aynan shu eng kichik qiymatni bergan barcha j larning ko'rsatkichini bittaga oshiring (dublikatlarning oldini olish uchun).

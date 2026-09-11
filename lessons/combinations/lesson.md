# Combinations

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
```

## TASK

Sizga ikkita butun son `n` va `k` berilgan.

`[1, n]` oraliqdagi sonlardan tuzish mumkin bo'lgan barcha `k` o'lchamli kombinatsiyalarni qaytaring.

Javobni istalgan tartibda qaytarishingiz mumkin (ham kombinatsiyalar tartibi, ham kombinatsiya ichidagi sonlar tartibi muhim emas).

Quyidagi funksiyani to'ldiring:

```go
func solve(n int, k int) [][]int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

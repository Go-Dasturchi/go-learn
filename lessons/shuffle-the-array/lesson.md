# Shuffle the Array

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: nums = [2,5,1,3,4,7], n = 3
Output: [2,3,5,4,1,7]
Tushuntirish:
x1=2, x2=5, x3=1
y1=3, y2=4, y3=7
Natija: [2,3,5,4,1,7]
```

## TASK

Sizga `2n` uzunlikdagi `nums` massivi berilgan bo'lib, elementlari `[x1, x2, ..., xn, y1, y2, ..., yn]` ko'rinishida joylashgan.
Massivni aralashtiring (shuffle) va uning ko'rinishini quyidagiga keltiring: `[x1, y1, x2, y2, ..., xn, yn]`.
Hosil bo'lgan massivni qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int, n int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

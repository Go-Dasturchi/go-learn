# Search insert position

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: nums = [1,3,5,6], target = 5
Output: 2
```

## TASK

O'sish tartibida saralangan aniq butun sonlardan iborat `nums` massivi va `target` soni berilgan. Agar massivda `target` mavjud bo'lsa, uning indeksini qaytaring. Agar mavjud bo'lmasa, u tartibni saqlagan holda joylashtirilganda qanday indeksga ega bo'lishini qaytaring.

Siz `O(log n)` vaqt murakkabligida ishlaydigan algoritm yozishingiz kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int, target int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

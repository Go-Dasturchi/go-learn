# Count of Smaller Numbers After Self

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: nums = [5,2,6,1]
Output: [2,1,1,0]
Tushuntirish:
5 dan o'ng tomonda: 2 va 1 kichik (2 ta)
2 dan o'ng tomonda: 1 kichik (1 ta)
6 dan o'ng tomonda: 1 kichik (1 ta)
1 dan o'ng tomonda: hech narsa yo'q (0 ta)
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan. `counts` ro'yxatini qaytaring, bunda `counts[i]` — bu `nums[i]` dan o'ng tomonda turuvchi va `nums[i]` dan qat'iy kichik sonlar soni.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

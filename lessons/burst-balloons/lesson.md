# Burst Balloons

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: nums = [3,1,5,8]
Output: 167
Tushuntirish:
nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
```

## TASK

Sizga `n` ta sharning ustiga yozilgan raqamlar `nums` massivi berilgan. Siz barcha sharlarni portlatishingiz kerak.

Agar siz `i` indeksdagi sharni portlatsangiz, u holda `nums[i - 1] * nums[i] * nums[i + 1]` miqdorda tanga olasiz. Agarda `i - 1` yoki `i + 1` massiv chegarasidan tashqariga chiqib ketsa, ularning o'rniga `1` raqami bor deb hisoblashingiz mumkin.

Barcha sharlarni portlatib yig'ish mumkin bo'lgan maksimal tangalar (coins) miqdorini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

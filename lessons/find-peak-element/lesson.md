# Find Peak Element

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: nums = [1,2,3,1]
Output: 2
Tushuntirish: 3 cho'qqi element hisoblanadi va uning indeksi 2.
```

## TASK

Cho'qqi (peak) element deb o'zining o'ng va chap qo'shnilaridan qat'iy katta bo'lgan elementga aytiladi.

Sizga butun sonlardan iborat `nums` massivi berilgan. Uning ichidan ixtiyoriy bir cho'qqi elementning indeksini toping va qaytaring. Massivda bir nechta cho'qqilar bo'lishi mumkin, bunday holda ulardan ixtiyoriy birining indeksini qaytarish kifoya.

Massiv chegaralaridan tashqaridagi elementlar $-\infty$ ga teng deb faraz qilishingiz mumkin (ya'ni `nums[-1] = nums[n] = -\infty`).

Yechimingiz albatta `O(log n)` vaqt murakkabligida ishlashi kerak.

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

# Contains Duplicate II

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: nums = [1,2,3,1], k = 3
Output: true
Tushuntirish: nums[0] va nums[3] o'zaro teng va ularning indekslari orasidagi farq 3 ga teng.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi va bitta `k` butun soni berilgan.

Agar massivda shunday ikkita alohida `i` va `j` indekslar mavjud bo'lsaki, bunda `nums[i] == nums[j]` va ularning indekslari orasidagi masofa ko'pi bilan `k` ga teng bo'lsa (ya'ni `abs(i - j) <= k`), u holda `true` qaytaring. Aks holda `false` qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int, k int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

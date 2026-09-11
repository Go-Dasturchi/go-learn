# Intersection of Two Arrays

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
```

## TASK

Sizga ikkita butun sonlardan iborat `nums1` va `nums2` massivlari berilgan. Ularning kesishmasini (ikkalasida ham mavjud bo'lgan elementlarni) qaytaring. Natijadagi massivda har bir element faqat bir marta (takrorlanmasdan) ishtirok etishi kerak. Natijani istalgan tartibda qaytarishingiz mumkin.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums1 []int, nums2 []int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

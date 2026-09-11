# Intersection of Two Arrays II

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
```

## TASK

Sizga ikkita butun sonlardan iborat `nums1` va `nums2` massivlari berilgan. Ularning kesishmasini qaytaring. Natijadagi massivda har bir element ikkala massivda necha marta takrorlangan bo'lsa, xuddi shuncha marta ko'rinishi kerak (agar birinchi massivda 2 marta, ikkinchisida 3 marta qatnashgan bo'lsa, natijada 2 marta ishtirok etishi kerak). Natijani istalgan tartibda qaytarishingiz mumkin.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums1 []int, nums2 []int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikkinchi massivdagi har bir sonning necha marta uchraganini hisoblab, birinchi massivni aylanib shu hisobdan foydalanishni o'ylang.
2. nums2'dagi sonlarning uchrash sonini map'da saqlang; nums1'ni aylanib, agar son map'da mavjud va soni 0 dan katta bo'lsa natijaga qo'shib hisoblagichni kamaytiring.

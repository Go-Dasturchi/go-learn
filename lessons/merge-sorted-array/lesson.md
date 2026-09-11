# Merge Sorted Array

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Tushuntirish: [1,2,3] va [2,5,6] massivlari birlashtirildi.
```

## TASK

Sizga ikkita kamaymaydigan (non-decreasing) tartibda joylashgan `nums1` va `nums2` butun sonlar massivlari, hamda mos ravishda ulardagi elementlar sonini bildiruvchi `m` va `n` butun sonlari berilgan.

`nums1` va `nums2` ni yagona tartiblangan massiv sifatida birlashtiring.

Yangi tartiblangan massiv funksiya tomonidan qaytarilmasdan, uning o'rniga `nums1` massivi ichida saqlanishi kerak (in-place). Bunga moslashish uchun `nums1` ning uzunligi `m + n` ga teng qilingan. Dastlabki `m` ta element saqlanishi kerak bo'lgan elementlar bo'lsa, oxirgi `n` ta element `0` qilib belgilangan va e'tiborga olinmasligi lozim.

Izoh: Bizning test tizimimizda in-place o'zgartirish o'rniga, siz birlashtirilgan massivni (`[Int]`) qaytaradigan funksiya yozishingiz talab etiladi.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums1 []int, m int, nums2 []int, n int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

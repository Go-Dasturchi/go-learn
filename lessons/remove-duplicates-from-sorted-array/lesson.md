# Remove Duplicates from Sorted Array

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: nums = [1,1,2]
Output: 2
Tushuntirish: Funksiyangiz k = 2 ni qaytarishi kerak, nums dagi dastlabki ikkita element 1 va 2 bo'ladi.
```

## TASK

Sizga kamaymaydigan (non-decreasing) tartibda joylashgan `nums` butun sonlar massivi berilgan. Har bir element faqat bir marta paydo bo'lishi uchun takrorlanuvchi elementlarni **joyida (in-place)** olib tashlang. Elementlarning nisbiy tartibi bir xil bo'lishi kerak. So'ngra `nums` dagi yagona elementlar sonini qaytaring.

Agar yagona elementlar soni `k` bo'lsa, sizning yechimingiz quyidagilarni bajarishi kerak:

- `nums` massivini shunday o'zgartiringki, uning dastlabki `k` ta elementi yagona elementlarni saqlasin. Dastlabki `k` ta elementdan keyin nima turganining ahamiyati yo'q.
- `k` ni qaytaring.

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

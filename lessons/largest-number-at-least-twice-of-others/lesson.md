# Larger Number At Least Twice of Others

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: nums = [3,6,1,0]
Output: 1
Tushuntirish: Eng katta element 6 (indeks 1). 6 >= 2*3, 6 >= 2*1, 6 >= 2*0. Shart bajarilgan.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan. Agar massivdagi eng katta element boshqa barcha elementlardan kamida ikki marta katta bo'lsa, uning indeksini qaytaring. Aks holda `-1` qaytaring.

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

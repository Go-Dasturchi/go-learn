# Minimum Operations to Reduce X to Zero

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: nums = [1,1,4,2,3], x = 5
Output: 2
Tushuntirish: Oxiridan ikki element olamiz: 3 va 2. 5 - 3 - 2 = 0. 2 operatsiya.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi va butun son `x` berilgan.

Bir operatsiyada siz massivning eng chapidagi yoki eng o'ngidagi elementni olib, `x` dan ayirasiz. Maqsad: `x` ni aynan `0` ga teng qilish.

Agar imkon bo'lsa, bu ishni qilish uchun zarur bo'lgan **minimal** operatsiyalar sonini qaytaring, aks holda `-1` qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int, x int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

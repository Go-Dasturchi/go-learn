# Two sum

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: nums = [2,7,11,15], target = 9
Output: [0, 1]
Tushuntirish: nums[0] + nums[1] == 9 bo'lgani uchun [0, 1] qaytariladi.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi va bitta `target` butun soni berilgan. Massiv ichidan yig'indisi `target` ga teng bo'ladigan ikkita sonni toping va ularning indekslarini qaytaring.

Har bir test uchun faqat bitta aniq yechim mavjud deb hisoblashingiz mumkin va bitta elementdan ikki marta foydalana olmaysiz. Javobni istalgan tartibda qaytarishingiz mumkin.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int, target int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

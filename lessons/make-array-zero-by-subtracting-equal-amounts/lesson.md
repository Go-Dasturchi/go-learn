# Make Array Zero by Subtracting Equal Amounts

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: nums = [1,5,0,3,5]
Output: 3
Tushuntirish:
1-qadam: noldan katta eng kichik son 1. Ayiramiz: [0,4,0,2,4].
2-qadam: noldan katta eng kichik son 2. Ayiramiz: [0,2,0,0,2].
3-qadam: noldan katta eng kichik son 2. Ayiramiz: [0,0,0,0,0].
Jami 3 operatsiya.
```

## TASK

Sizga manfiy bo'lmagan butun sonlardan iborat `nums` massivi berilgan.

Bir operatsiyada siz massivdagi noldan katta bo'lgan **eng kichik** elementni tanlaysiz va uni massivdagi barcha noldan katta bo'lgan elementlardan ayirib chiqasiz.

Massivning barcha elementlari `0` ga aylanguncha necha marta operatsiya qilish kerakligini toping.

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

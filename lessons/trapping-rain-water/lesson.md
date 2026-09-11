# Trapping rain water

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
```

## TASK

`n` ta manfiy bo'lmagan butun sonlar (har biri kengligi 1 bo'lgan ustun balandligini bildiradi) `height` massivi orqali berilgan. Yomg'ir yoqqanidan so'ng ushbu hududda qancha miqdorda suv to'planib qolishini hisoblang.

Quyidagi funksiyani to'ldiring:

```go
func solve(height []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

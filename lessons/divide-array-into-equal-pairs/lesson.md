# Divide Array Into Equal Pairs

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: nums = [3,2,3,2,2,2]
Output: true
Tushuntirish:
6 ta element bor (3 juftlik kerak).
Biz ularni shunday guruhlaymiz: (2, 2), (2, 2), (3, 3).
Barcha juftliklar hosil qilindi, shuning uchun true.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan bo'lib, uning uzunligi juft sondir (ya'ni `2 * n`).

Massivni har birida o'zaro bir xil ikkita element bo'lgan `n` ta juftlikka ajratish mumkin bo'lsa `true` qaytaring, aks holda `false` qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

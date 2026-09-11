# Largest Rectangle in Histogram

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: heights = [2,1,5,6,2,3]
Output: 10
Tushuntirish: Eng katta to'rtburchak 5 va 6 balandlikdagi ustunlar orqali hosil qilinadi, uning yuzasi 5 * 2 = 10.
```

## TASK

Sizga gistogrammaning har bir ustuni balandligini ifodalovchi, butun sonlardan iborat `heights` massivi berilgan. Har bir ustunning kengligi 1 ga teng. Ushbu gistogramma ichiga chizilishi mumkin bo'lgan eng katta to'rtburchakning yuzini toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(heights []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

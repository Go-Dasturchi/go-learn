# Two Sum II - Input Array Is Sorted

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Tushuntirish: 2 + 7 = 9. Ularning indekslari 1 va 2 (1-indexed).
```

## TASK

Sizga o'sish tartibida joylashgan (tartiblangan) butun sonlardan iborat `numbers` massivi berilgan. Ushbu massivda yig'indisi `target` ga teng bo'lgan ikkita sonni toping va ularning **1 dan boshlanadigan** (1-indexed) indekslarini qaytaring — ya'ni `[index1, index2]` ko'rinishida, bunda `index1 < index2`.

Siz massiv xotirasidan tashqarida `O(1)` qo'shimcha xotiradan foydalaningsingiz kerak (ya'ni `O(n)` xotira imkoni yo'q).

Quyidagi funksiyani to'ldiring:

```go
func solve(numbers []int, target int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

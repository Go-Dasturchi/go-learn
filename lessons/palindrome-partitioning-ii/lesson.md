# Palindrome Partitioning II

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: s = "aab"
Output: 1
Tushuntirish: "aab" satrini ["aa", "b"] qilib bo'lish uchun faqat 1 marta kesish kifoya. Barcha qismlar palindrom.
```

## TASK

Sizga `s` satri berilgan. `s` satrini shunday qismlarga bo'lingki (partition), bunda har bir qism satr palindrom bo'lsin.

Bunday bo'lishlarni amalga oshirish uchun kerak bo'ladigan **eng kam (minimum)** kesishlar sonini toping va qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

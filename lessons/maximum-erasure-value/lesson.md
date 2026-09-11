# Maximum Erasure Value

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: nums = [4,2,4,5,6]
Output: 17
Tushuntirish: Eng yuqori ball [2,4,5,6] yig'indisiga teng: 2+4+5+6=17.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan. Siz massivdan takrorlanmaydigan elementlar bilan iborat **uzluksiz** qism-massivni o'chirish orqali ball yig'aysiz. Ball sifatida o'chirilgan elementlar **yig'indisi** hisoblanadi.

Bir ishda yig'ish mumkin bo'lgan maksimal ballni toping.

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

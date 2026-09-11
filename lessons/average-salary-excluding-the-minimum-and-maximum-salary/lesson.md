# Average Salary Excluding the Minimum and Maximum Salary

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: salary = [4000,3000,1000,2000]
Output: 2500.00000
Tushuntirish: Minimal 1000, maksimal 4000. Qoldi: 3000, 2000. O'rtacha: 2500.
```

## TASK

Sizga xodimlarning oylik maoshlaridan iborat `salary` massivi berilgan. Unda kamida ikkita element bor.

Eng kichik va eng katta maoshlarni olib tashlagandan so'ng qolgan maoshlarning o'rtacha qiymatini (average) qaytaring. `10^-5` aniqlikdagi javob to'g'ri hisoblanadi.

Quyidagi funksiyani to'ldiring:

```go
func solve(salary []int) float64 {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Massivni ikki marta aylanish shart emas — bitta o'tishda kerakli barcha ma'lumotni (min, max, yig'indi) yig'ib olish mumkin.
2. Bitta pastadan sonlarni aylanib eng kichik va eng katta qiymatlarni hamda umumiy yig'indini toping, so'ng yig'indidan min va maxni ayirib qolgan elementlar soniga (len-2) bo'ling.

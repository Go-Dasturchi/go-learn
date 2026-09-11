# Fizz Buzz

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: n = 15
Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
```

## TASK

Sizga `n` butun soni berilgan. Siz 1 dan `n` gacha bo'lgan sonlar asosida quyidagi qoidalar bilan satrlar massivini (string array) qaytarishingiz kerak:

1. Agar son ham 3 ga, ham 5 ga qoldiqsiz bo'linsa `"FizzBuzz"` ni qo'shing.
2. Agar son faqat 3 ga qoldiqsiz bo'linsa `"Fizz"` ni qo'shing.
3. Agar son faqat 5 ga qoldiqsiz bo'linsa `"Buzz"` ni qo'shing.
4. Aks holda sonning o'zini satr ko'rinishida qo'shing.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) []string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

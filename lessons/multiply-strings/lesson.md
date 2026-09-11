# Multiply Strings

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **30 daqiqa**

## EXAMPLE

```
Input: num1 = "2", num2 = "3"
Output: "6"
```

## TASK

Sizga ikkita manfiy bo'lmagan butun sonlarni ifodalovchi `num1` va `num2` satrlari berilgan. Ularning ko'paytmasini ham satr ko'rinishida qaytaring.

Eslatma: Tayyor BigInteger kutubxonalaridan yoki to'g'ridan-to'g'ri butun songa aylantiruvchi tayyor funksiyalardan foydalanmang. Ikkala son ham juda katta bo'lishi va standart butun son toifalariga sig'masligi mumkin.

Quyidagi funksiyani to'ldiring:

```go
func solve(num1 string, num2 string) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Maktabda qo'lda ko'paytirish usulini eslang: har bir raqamni boshqasining har bir raqamiga alohida ko'paytirib, natijalarni mos pozitsiyalarga qo'shib chiqish mumkin — bu uchun natijani xonalar (digits) massivi sifatida saqlang.
2. Uzunligi m+n bo'lgan natija massivi tuzing; num1[i] va num2[j] ko'paytmasi natija massivining `i+j` va `i+j+1` pozitsiyalariga ta'sir qiladi — har bir ko'paytmani mos joyga qo'shib, ortiqcha (carry) qismini chapdagi xonaga o'tkazing; oxirida boshidagi ortiqcha nollarni olib tashlab, raqamlarni satrga aylantiring.

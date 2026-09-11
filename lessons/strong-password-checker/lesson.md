# Strong Password Checker

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **50 daqiqa**

## EXAMPLE

```
Input: password = "a"
Output: 5
Tushuntirish: 5 ta belgi qo'shish kerak: katta harf, raqam va yana 3 ta belgi.
```

## TASK

Kuchli parol uchun quyidagi shartlar bajarilishi lozim:

1. Uzunligi kamida 6 va ko'pi bilan 20 belgidan iborat bo'lsin.
2. Kamida bitta kichik harf, bitta katta harf va bitta raqam bo'lsin.
3. Qator-qator 3 yoki undan ortiq bir xil belgi bo'lmasin (masalan, `...aaa...` yoki `...BBB...`).

`password` satri berilgan. Uni kuchli parolga aylantirish uchun kerak bo'ladigan minimal amallar soni (qo'shish, o'chirish, almashtirish)ni toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(password string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Uchta shartni (uzunlik, belgi turlari, ketma-ket takrorlanish) mustaqil emas, balki uzunlikka qarab uchta ALOHIDA holatga (n<6, 6<=n<=20, n>20) bo'lib yeching — chunki uzunlik juda qisqa yoki juda uzun bo'lganda o'chirish/qo'shish amallari takroriy belgilar muammosini "bepul" hal qilib yuborishi mumkin.
2. Avval 3 va undan uzun ketma-ket bir xil belgili guruhlarni toping va har biri uchun length/3 ta almashtirish kerakligini hisoblang (bu `replaces`). n<6 bo'lsa javob = max(missingTypes, 6-n). 6<=n<=20 bo'lsa javob = max(missingTypes, replaces). n>20 bo'lsa avval n-20 ta belgini o'chirishingiz kerak — bu o'chirishlarni birinchi navbatda uzunligi 3k ga teng bo'lgan guruhlardan (har birida 1 tadan o'chirish 1 ta almashtirishni yo'q qiladi), keyin 3k+1 guruhlardan (2 tadan o'chirish kerak), so'ng qolganini 3 tadan o'chirib `replaces` ni kamaytirib boring, va yakuniy javob = deletes + max(missingTypes, replaces).

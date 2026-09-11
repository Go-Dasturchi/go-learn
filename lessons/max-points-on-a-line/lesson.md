# Max Points on a Line

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: points = [[1,1],[2,2],[3,3]]
Output: 3
Tushuntirish: Barcha 3 ta nuqta y = x to'g'ri chizig'i ustida yotadi.
```

## TASK

Sizga 2D tekisligidagi bir nechta nuqtalar berilgan `points`, bunda har bir `points[i] = [xi, yi]` ni bildiradi.

Bitta to'g'ri chiziq ustida yotuvchi eng ko'p nuqtalar sonini toping va qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(points [][]int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir nuqtani navbat bilan "bazaviy nuqta" sifatida olib, undan boshqa har bir nuqtagacha bo'lgan yo'nalishni (og'ish/slope) guruhlash orqali yeching — bir xil yo'nalishga ega nuqtalar bitta to'g'ri chiziqda yotadi.
2. Suzuvchi son xatoliklaridan qochish uchun og'ishni kasr sifatida emas, balki (dx, dy) ni ularning eng katta umumiy bo'luvchisiga (GCD) bo'lib, normallashtirilgan butun sonlar juftligi sifatida xesh-kalit qilib saqlang (ishorani ham izchil qilib belgilang). Har bir bazaviy nuqta uchun shu kalitlar bo'yicha hisoblagichni yuritib, eng ko'p uchragan kalitning soniga bazaviy nuqtaning o'zini va bir xil koordinatali (duplicate) nuqtalarni qo'shib, natijani yangilang.

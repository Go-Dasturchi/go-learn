# Trapping rain water

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
```

## TASK

`n` ta manfiy bo'lmagan butun sonlar (har biri kengligi 1 bo'lgan ustun balandligini bildiradi) `height` massivi orqali berilgan. Yomg'ir yoqqanidan so'ng ushbu hududda qancha miqdorda suv to'planib qolishini hisoblang.

Quyidagi funksiyani to'ldiring:

```go
func solve(height []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir ustun ustidagi suv miqdori "chapdagi eng baland ustun" va "o'ngdagi eng baland ustun"ning kichigi bilan cheklanadi — buni oldindan hisoblab qo'yish o'rniga ikkita ko'rsatkich (chapdan va o'ngdan) bilan bitta o'tishda hisoblash mumkin.
2. `left` va `right` ko'rsatkichlarini massiv chetlaridan boshlang, `leftMax` va `rightMax` ni kuzatib boring. Har qadamda height[left] va height[right] ni solishtirib, qaysi tomon kichik bo'lsa o'sha tomonni ishlating: agar height[left] < height[right] bo'lsa, height[left] ni leftMax bilan solishtirib (agar undan katta bo'lsa leftMax ni yangilang, aks holda leftMax-height[left] miqdorda suv qo'shing) left++ qiling; aks holda xuddi shunday right tomon uchun bajaring. Bu ishlaydi, chunki kichikroq tomonning suv sathi albatta ikkala tomon maksimumining kichigi bilan belgilanadi.

# 3Sum

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rtacha** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1, -1, 2], [-1, 0, 1]]
```

## TASK

Butun sonlardan iborat `nums` massivi berilgan. Uning ichidan `nums[i] + nums[j] + nums[k] == 0` va `i != j`, `i != k`, `j != k` shartlarni qanoatlantiruvchi barcha har xil uchliklarni toping. Javobda dublikat uchliklar bo'lmasligi kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) [][]int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.

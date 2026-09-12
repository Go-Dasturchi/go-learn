# 06 — Docker

## THEORY

"Bu mening kompyuterimda ishlaydi" — dasturchilar orasida mashhur muammo: sizning kompyuteringizda Go versiyasi, kutubxonalar, operatsion tizim sozlamalari boshqacha bo'lishi mumkin, va dastur boshqa serverga ko'chirilganda **ishlamay qolishi** mumkin.

**Yechim — Docker.** Docker, dasturingizni, unga kerakli **hamma narsa bilan birga** (operatsion tizim qismlari, kutubxonalar, sozlamalar) — **konteyner** deb ataladigan yopiq muhitga joylaydi. Konteyner, qayerda ishga tushirilishidan qat'i nazar (sizning noutbukingiz, hamkasbingiz kompyuteri, yoki production server) — **bir xil** ishlaydi.

### 1-qadam: Docker'ni o'rnatish

**macOS:**

```bash
brew install --cask docker
```

O'rnatilgach, Docker Desktop ilovasini oching (u fonda ishlab turishi kerak).

**Ubuntu/Debian:**

```bash
curl -fsSL https://get.docker.com | sh
sudo usermod -aG docker $USER
```

(Oxirgi buyruqdan keyin, terminalni qayta oching yoki tizimga qayta kiring.)

O'rnatilganini tekshirish:

```bash
docker --version
docker run hello-world
```

Agar "Hello from Docker!" degan xabar chiqsa — hammasi tayyor.

### 2-qadam: Go dasturi uchun Dockerfile yozish

Loyihangiz papkasida `Dockerfile` nomli fayl yarating:

```dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o server .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
```

Bu — **multi-stage build** (ko'p bosqichli qurish): birinchi bosqich (`builder`) — Go kodini kompilyatsiya qiladi (bu bosqich katta, chunki Go kompilyatori va kutubxonalar bor). Ikkinchi bosqich — faqat **tayyor binary faylni** olib, juda kichik `alpine` image ichiga joylaydi. Natijada, yakuniy image — bir necha megabayt, Go kompilyatorining o'zi (yuzlab megabayt) unga **kirmaydi**.

### 3-qadam: Image qurish va konteynerni ishga tushirish

```bash
docker build -t mening-dasturim .
docker run -p 8080:8080 mening-dasturim
```

- `docker build -t mening-dasturim .` — joriy papkadagi `Dockerfile` asosida, `mening-dasturim` nomli **image** (qolip) quradi.
- `docker run -p 8080:8080 mening-dasturim` — shu image'dan **konteyner** ishga tushiradi. `-p 8080:8080` — kompyuteringizdagi 8080-portni, konteyner ichidagi 8080-portga **bog'laydi** (host:container).

Foydali qo'shimcha buyruqlar:

```bash
docker ps                 # ishlab turgan konteynerlar ro'yxati
docker logs <konteyner_id> # konteyner logini ko'rish
docker stop <konteyner_id> # konteynerni to'xtatish
```

**Nega port bog'lash muhim.** Konteyner — o'zining **izolyatsiyalangan tarmog'ida** ishlaydi; `-p host:container` bo'lmasa, tashqaridan (sizning brauzeringizdan) konteyner ichidagi serverga **umuman ulanib bo'lmaydi**.

**Go'da bu g'oyani sinash — `docker run` buyrug'ini qurish.** Real Docker ishga tushirilmasa ham, buyruq **qanday tuzilishini** Go orqali mashq qilish mumkin:

```go
func dockerBuyrugiQur(image string, hostPort, konteynerPort int) string {
	return fmt.Sprintf("docker run -p %d:%d %s", hostPort, konteynerPort, image)
}
```

## EXAMPLE

```go
package main

import "fmt"

func dockerBuyrugiQur(image string, hostPort, konteynerPort int) string {
	return fmt.Sprintf("docker run -p %d:%d %s", hostPort, konteynerPort, image)
}

func main() {
	fmt.Println(dockerBuyrugiQur("mening-dasturim", 8080, 8080))
	fmt.Println(dockerBuyrugiQur("nginx", 3000, 80))
}
```

Natija:

```
docker run -p 8080:8080 mening-dasturim
docker run -p 3000:80 nginx
```

## TASK

`dockerBuyrugiQur(image string, hostPort, konteynerPort int) string` funksiyasini to'ldiring — u, `"docker run -p <hostPort>:<konteynerPort> <image>"` ko'rinishidagi qatorni qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

**Endi, terminalingizni oching va Docker'ni haqiqatan sinab ko'ring** (bu qism checker tomonidan tekshirilmaydi, lekin bilim uchun juda muhim):

```bash
docker --version
docker run hello-world
```

Agar Docker o'rnatilgan bo'lsa, oddiy Go dasturi uchun yuqoridagi `Dockerfile`ni yaratib, `docker build` va `docker run` buyruqlarini bajarib ko'ring.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `fmt.Sprintf("docker run -p %d:%d %s", hostPort, konteynerPort, image)`.
2. Argumentlar tartibiga e'tibor bering: avval `hostPort`, keyin `konteynerPort`, oxirida `image`.

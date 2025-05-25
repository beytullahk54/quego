# 🕒 Delayed HTTP Job Queue (Go)

**Delayed HTTP Job Queue**, belirlediğiniz bir adrese, belirlediğiniz zamanda HTTP isteği gönderen basit ve genişletilebilir bir zamanlı kuyruk sistemidir.

## 🚀 Ne İşe Yarar?

Bu servis sayesinde:

- Belirli bir **URL'ye**
- Belirli bir **zaman sonra**
- Belirli bir **HTTP metodu**, **header** ve **body** ile

otomatik olarak HTTP isteği gönderebilirsiniz.

## 🔧 Özellikler

- 🕰️ Gecikmeli HTTP istek planlama (`POST /jobs`)
- 🔐 API Token doğrulama
- ✅ Job durumlarını takip edebilme (`pending`, `done`, `failed`, `cancelled`)
- ♻️ Basit retry (yeniden deneme) desteği
- 🗃️ SQLite veya PostgreSQL desteği (esnek yapı)

## 📝 Yapılacaklar

- [ ] 🕰️ İstenilen tam tarih ve saat getirilmesi
- [ ] 📡 Postman örnek dosyası paylaşılması
- [ ] ✅ Validation sisteminin geliştirilmesi
- [ ] 🎨 Basit bir UI hazırlanması
- [ ] 🔑 API token oluşturulması ve bu token'a göre ait olunan işlerin listelenmesinin sağlanması

## 📦 Kurulum

```bash
git clone https://github.com/kullanici/delayed-http-job-queue.git
cd delayed-http-job-queue
go run cmd/server/main.go
```

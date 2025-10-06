# 🔔 Sentiric Notification Service - Mantık ve Akış Mimarisi

**Stratejik Rol:** SMS ve E-posta gibi tek yönlü bildirim kanallarını platforma soyutlayan merkezi bir geçit. Bu, Agent'ların doğrudan harici API'ler yerine tek bir RPC'ye güvenmesini sağlar.

---

## 1. Temel Akış: Bildirim Gönderme (SendSMS/SendEmail)

Notification Service, gelen isteği işler, yapılandırılmış adaptörünü (Twilio, SendGrid vb.) seçer ve mesajı gönderir.

```mermaid
graph TD
    A[Agent Service / Task Service] -- gRPC: SendSMS --> B(Notification Service)
    
    Note over B: 1. Adaptör Seçimi (SMS Adapter)
    B --> C{Twilio Adaptörü};
    C -- HTTP API Çağrısı --> Twilio[Harici Twilio API];
    Twilio -- Response --> C;
    
    C --> B;
    B -- Response --> A;
```

## 2. Adaptör Mimarisi

Notification Service, SmsAdapter ve EmailAdapter gibi yapılandırma değişkenlerine göre uygun harici sağlayıcıyı seçer.

* SMS: Twilio, Vonage
* E-posta: SendGrid, Mailgun

Bu mimari, entegrasyon mantığını servisin içine hapseder ve Agent'ın basit kalmasını sağlar.
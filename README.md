# Contact Form & Email API - Development Plan

## 🎯 Project Overview

A **production-ready microservice** that handles contact form submissions with email notifications, spam protection, and analytics. This API can be integrated into any website (your personal site, client projects, landing pages).

### **Key Value Proposition**

- **Reusable across all your projects** (personal website + client sites)
- **Revenue opportunity** (offer as a service to clients)
- **No frontend coupling** (works with any framework)
- **Enterprise-grade features** (rate limiting, spam protection, analytics)

---

## 🏗️ API Endpoints Design

### **Core Endpoints**

```
POST   /api/v1/contact              # Submit contact form
POST   /api/v1/contact/validate     # Validate form without sending
GET    /api/v1/health               # Health check
GET    /api/v1/stats                # Basic analytics (admin only)
```

### **Request/Response Structure**

```json
// POST /api/v1/contact
{
  "name": "John Doe",
  "email": "john@example.com",
  "subject": "Project Inquiry",
  "message": "I'd like to discuss a project...",
  "website_id": "personal_site", // To track which site sent it
  "honeypot": "" // Anti-spam field
}

// Response
{
  "success": true,
  "message": "Message sent successfully",
  "id": "msg_123456789"
}
```

---

## 📋 Development Phases

### **Phase 1: Core Functionality (Week 1)**

**Goal**: Basic working API that accepts forms and sends emails

#### **Day 1-2: Project Setup**

- [ ] Initialize Go module
- [ ] Set up project structure
- [ ] Configure environment variables
- [ ] Create basic HTTP server with Gorilla Mux
- [ ] Add health check endpoint

#### **Day 3-4: Contact Form Handler**

- [ ] Create contact form models/structs
- [ ] Implement JSON request parsing
- [ ] Add basic input validation
- [ ] Create contact form handler

#### **Day 5-7: Email Integration**

- [ ] Integrate with email service (SMTP or service like Resend/SendGrid)
- [ ] Create email templates
- [ ] Test email sending functionality
- [ ] Add error handling and logging

**Deliverable**: Working API that accepts contact forms and sends emails

---

### **Phase 2: Security & Validation (Week 2)**

**Goal**: Production-ready security features

#### **Security Features**

- [ ] **Rate Limiting**: Prevent spam (max requests per IP/hour)
- [ ] **Input Sanitization**: Clean HTML/JS from inputs
- [ ] **Honeypot Fields**: Catch basic bots
- [ ] **CORS Configuration**: Control which domains can use your API
- [ ] **Request Size Limits**: Prevent large payload attacks

#### **Advanced Validation**

- [ ] **Email Validation**: Verify email format and DNS
- [ ] **Profanity Filter**: Block inappropriate content
- [ ] **Required Fields**: Ensure all needed data is present
- [ ] **Field Length Limits**: Prevent extremely long inputs

**Deliverable**: Secure API ready for production use

---

### **Phase 3: Analytics & Admin Features (Week 3)**

**Goal**: Business intelligence and management features

#### **Analytics Features**

- [ ] **Submission Tracking**: Count messages per day/week/month
- [ ] **Website Analytics**: Track which sites generate most leads
- [ ] **Response Metrics**: Email delivery success rates
- [ ] **Spam Detection Stats**: Track blocked attempts

#### **Admin Endpoints**

- [ ] **Stats Dashboard**: GET /api/v1/stats (protected)
- [ ] **Message History**: List recent submissions
- [ ] **Configuration**: Update settings without redeployment

**Deliverable**: Full-featured service with business intelligence

---

### **Phase 4: Advanced Features (Week 4)**

**Goal**: Enterprise-grade features for client projects

#### **Multi-tenant Support**

- [ ] **Website Configuration**: Different settings per client site
- [ ] **Custom Email Templates**: Branded emails per client
- [ ] **Webhook Support**: Notify external systems of new submissions
- [ ] **API Keys**: Secure access per client

#### **Integration Features**

- [ ] **Slack Notifications**: Send alerts to team channels
- [ ] **Database Storage**: Optional message persistence
- [ ] **Auto-responses**: Confirmation emails to form submitters
- [ ] **File Attachments**: Support for uploaded files

**Deliverable**: Enterprise-ready service for client projects

---

## 🛠️ Technical Implementation Plan

### **Project Structure**

```
contact-form-api/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── internal/
│   ├── handlers/
│   │   ├── contact.go              # Contact form handlers
│   │   ├── health.go               # Health check
│   │   └── stats.go                # Analytics endpoints
│   ├── models/
│   │   ├── contact.go              # Contact form models
│   │   └── response.go             # API response models
│   ├── services/
│   │   ├── email.go                # Email sending service
│   │   ├── validation.go           # Input validation
│   │   └── analytics.go            # Usage tracking
│   ├── middleware/
│   │   ├── ratelimit.go            # Rate limiting
│   │   ├── cors.go                 # CORS handling
│   │   └── logging.go              # Request logging
│   └── config/
│       └── config.go               # Configuration management
├── templates/
│   ├── email_notification.html     # Email template for you
│   └── email_confirmation.html     # Auto-response template
├── .env.example
├── docker-compose.yml              # For local development
├── Dockerfile                      # For deployment
├── go.mod
└── README.md
```

### **Key Dependencies**

```go
// go.mod
require (
    github.com/gorilla/mux v1.8.0           // HTTP routing
    github.com/joho/godotenv v1.5.1         // Environment variables
    github.com/go-playground/validator/v10   // Input validation
    golang.org/x/time v0.3.0                // Rate limiting
    github.com/resend/resend-go v2.0.0       // Email service (or SMTP)
    github.com/microcosm-cc/bluemonday      // HTML sanitization
)
```

---

## 🚀 Deployment Strategy

### **Development Environment**

- [ ] Local development with hot reload
- [ ] Environment variable management
- [ ] Local email testing (MailHog or similar)

### **Production Deployment Options**

1. **Railway/Render**: Simple deployment for MVP
2. **DigitalOcean App Platform**: Managed with custom domains
3. **Docker + VPS**: Full control for advanced features
4. **Vercel Edge Functions**: Integrate directly with Next.js

### **Domain & SSL**

- [ ] Custom subdomain: `api.yourdomain.com`
- [ ] SSL certificate (Let's Encrypt or provider SSL)
- [ ] CDN setup for global performance

---

## 🧪 Testing Strategy

### **Testing Phases**

1. **Unit Tests**: Individual function testing
2. **Integration Tests**: Email sending, validation flows
3. **Load Testing**: Rate limiting, concurrent requests
4. **Security Testing**: SQL injection, XSS attempts, spam testing

### **Test Cases**

- [ ] Valid form submission → email received
- [ ] Invalid email format → proper error response
- [ ] Rate limiting → blocked after threshold
- [ ] Honeypot triggered → silent rejection
- [ ] Large payload → request blocked
- [ ] CORS validation → only allowed domains work

---

## 💰 Monetization Opportunities

### **Service Tiers**

- **Free Tier**: 100 submissions/month (for your personal use)
- **Basic**: $10/month - 1,000 submissions, basic analytics
- **Pro**: $25/month - 10,000 submissions, custom templates, webhooks
- **Enterprise**: Custom pricing - unlimited, white-label, SLA

### **Client Integration**

- Easy integration code snippets
- WordPress plugin potential
- Shopify app opportunity
- Documentation site with examples

---

## 🎯 Success Metrics

### **Technical Metrics**

- [ ] **99.9% uptime** (use monitoring service)
- [ ] **< 200ms response time** for form submissions
- [ ] **99%+ email delivery rate**
- [ ] **0 security incidents**

### **Business Metrics**

- [ ] **Personal website**: Contact form working
- [ ] **Client adoption**: 3+ client sites using the API
- [ ] **Revenue generation**: $100+/month from service
- [ ] **Portfolio piece**: Case study for your software house

---

## 🚦 Getting Started - Next Steps

### **This Week**

1. **Day 1**: Set up Go project structure and basic HTTP server
2. **Day 2**: Create contact form models and basic handler
3. **Day 3**: Integrate email sending (start with SMTP)
4. **Day 4**: Add input validation and error handling
5. **Day 5**: Test with your personal website
6. **Weekend**: Deploy to staging environment

### **Ready to Start?**

1. Create new directory: `mkdir contact-form-api && cd contact-form-api`
2. Initialize Go module: `go mod init contact-form-api`
3. Let's build the foundation together!

---

_This API will become a cornerstone service for your software house - powering contact forms across all your projects while teaching you production Go development patterns._

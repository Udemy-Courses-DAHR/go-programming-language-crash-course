Perfect! Let's build a realistic Go application from requirement to delivery, showing how these principles create maintainable, testable code.

## 📋 Business Requirement: **User Notification System**

**User Story:** 
> As a product manager, I want to notify users about important events (welcome, password reset, order confirmation) via multiple channels (email, SMS, Slack) so that users receive timely updates through their preferred medium.

**Definition of Done:**
- [ ] Support email, SMS, and Slack notifications
- [ ] Unit tests with 90%+ coverage
- [ ] Easy to add new notification channels
- [ ] Configurable per user which channels they prefer
- [ ] Proper error handling and logging

---

## 🏗️ Step 1: Define Abstractions (DIP)

First, we define the stable abstractions - this is our **contract** that won't change frequently.

```go
// domain/notifications.go - Core business abstractions
package domain

// Notification represents a message to be sent
type Notification struct {
    UserID    string
    Title     string
    Message   string
    Type      NotificationType
}

type NotificationType string

const (
    WelcomeNotification       NotificationType = "welcome"
    PasswordResetNotification NotificationType = "password_reset"
    OrderConfirmation         NotificationType = "order_confirmation"
)

// Notifier is the core abstraction - business logic depends on this
type Notifier interface {
    Send(notification Notification) error
    CanSend(notificationType NotificationType) bool
    GetName() string
}

// UserPreferenceRepository abstraction
type UserPreferenceRepository interface {
    GetUserChannels(userID string) ([]string, error)
}
```

---

## 🔧 Step 2: Implement Concrete Details

Now we implement the low-level details that can change frequently.

```go
// infrastructure/email_notifier.go
package infrastructure

import (
    "fmt"
    "your-app/domain"
)

type EmailNotifier struct {
    SMTPHost string
    SMTPPort int
}

func NewEmailNotifier(host string, port int) *EmailNotifier {
    return &EmailNotifier{
        SMTPHost: host,
        SMTPPort: port,
    }
}

func (e *EmailNotifier) Send(notification domain.Notification) error {
    // Real email sending logic would go here
    fmt.Printf("Sending email to user %s: %s - %s\n", 
        notification.UserID, notification.Title, notification.Message)
    return nil
}

func (e *EmailNotifier) CanSend(notificationType domain.NotificationType) bool {
    // Email can send all notification types
    return true
}

func (e *EmailNotifier) GetName() string {
    return "email"
}

// infrastructure/sms_notifier.go
type SMSNotifier struct {
    APIKey string
}

func NewSMSNotifier(apiKey string) *SMSNotifier {
    return &SMSNotifier{APIKey: apiKey}
}

func (s *SMSNotifier) Send(notification domain.Notification) error {
    fmt.Printf("Sending SMS to user %s: %s\n", 
        notification.UserID, notification.Message)
    return nil
}

func (s *SMSNotifier) CanSend(notificationType domain.NotificationType) bool {
    // SMS only for urgent notifications
    return notificationType == domain.PasswordResetNotification
}

func (s *SMSNotifier) GetName() string {
    return "sms"
}

// infrastructure/slack_notifier.go
type SlackNotifier struct {
    WebhookURL string
}

func NewSlackNotifier(webhookURL string) *SlackNotifier {
    return &SlackNotifier{WebhookURL: webhookURL}
}

func (sl *SlackNotifier) Send(notification domain.Notification) error {
    fmt.Printf("Sending Slack message to user %s: %s\n", 
        notification.UserID, notification.Message)
    return nil
}

func (sl *SlackNotifier) CanSend(notificationType domain.NotificationType) bool {
    // Slack for non-urgent, informative notifications
    return notificationType != domain.PasswordResetNotification
}

func (sl *SlackNotifier) GetName() string {
    return "slack"
}
```

---

## 💼 Step 3: Business Logic (High-Level Policy)

This is our core business logic that depends on abstractions, not concrete implementations.

```go
// application/notification_service.go - Business logic
package application

import (
    "fmt"
    "log"
    "your-app/domain"
)

type NotificationService struct {
    notifiers  []domain.Notifier
    userPrefRepo domain.UserPreferenceRepository
}

// Dependency Injection through constructor
func NewNotificationService(
    notifiers []domain.Notifier, 
    userPrefRepo domain.UserPreferenceRepository,
) *NotificationService {
    return &NotificationService{
        notifiers:    notifiers,
        userPrefRepo: userPrefRepo,
    }
}

// Core business logic - doesn't know about SMTP, SMS APIs, etc.
func (ns *NotificationService) SendNotification(notification domain.Notification) error {
    // Get user's preferred channels
    preferredChannels, err := ns.userPrefRepo.GetUserChannels(notification.UserID)
    if err != nil {
        return fmt.Errorf("failed to get user preferences: %w", err)
    }

    // Determine which notifiers to use based on user preferences and capability
    var errors []error
    for _, notifier := range ns.notifiers {
        if ns.shouldUseNotifier(notifier, preferredChannels, notification.Type) {
            if err := notifier.Send(notification); err != nil {
                log.Printf("Failed to send via %s: %v", notifier.GetName(), err)
                errors = append(errors, err)
            } else {
                log.Printf("Successfully sent notification via %s", notifier.GetName())
            }
        }
    }

    if len(errors) > 0 {
        return fmt.Errorf("some notifications failed: %v", errors)
    }
    return nil
}

func (ns *NotificationService) shouldUseNotifier(
    notifier domain.Notifier, 
    preferredChannels []string, 
    notificationType domain.NotificationType,
) bool {
    // Check if user prefers this channel
    userPrefers := false
    for _, channel := range preferredChannels {
        if channel == notifier.GetName() {
            userPrefers = true
            break
        }
    }

    // Check if notifier can send this type of notification
    canSend := notifier.CanSend(notificationType)

    return userPrefers && canSend
}
```

---

## 🗄️ Step 4: Infrastructure Implementations

```go
// infrastructure/user_repository.go
package infrastructure

import "your-app/domain"

type UserPreferenceRepositoryImpl struct {
    // This could be a database connection, but we'll use in-memory for example
    userPreferences map[string][]string // userID -> channels
}

func NewUserPreferenceRepository() *UserPreferenceRepositoryImpl {
    return &UserPreferenceRepositoryImpl{
        userPreferences: map[string][]string{
            "user123": {"email", "slack"},    // User prefers email and slack
            "user456": {"sms", "email"},      // User prefers SMS and email
        },
    }
}

func (repo *UserPreferenceRepositoryImpl) GetUserChannels(userID string) ([]string, error) {
    channels, exists := repo.userPreferences[userID]
    if !exists {
        // Default channels if user not found
        return []string{"email"}, nil
    }
    return channels, nil
}
```

---

## 🧪 Step 5: Comprehensive Testing

This is where our design pays off - easy testing!

```go
// application/notification_service_test.go
package application

import (
    "testing"
    "your-app/domain"
)

// MockNotifier for testing
type MockNotifier struct {
    Name            string
    SendCalled      bool
    LastNotification domain.Notification
    ShouldError     bool
    CanSendResponse bool
}

func (m *MockNotifier) Send(notification domain.Notification) error {
    m.SendCalled = true
    m.LastNotification = notification
    if m.ShouldError {
        return fmt.Errorf("mock error")
    }
    return nil
}

func (m *MockNotifier) CanSend(notificationType domain.NotificationType) bool {
    return m.CanSendResponse
}

func (m *MockNotifier) GetName() string {
    return m.Name
}

// MockUserRepository for testing
type MockUserRepository struct {
    UserChannels map[string][]string
    ShouldError  bool
}

func (m *MockUserRepository) GetUserChannels(userID string) ([]string, error) {
    if m.ShouldError {
        return nil, fmt.Errorf("repository error")
    }
    return m.UserChannels[userID], nil
}

func TestNotificationService_SendNotification(t *testing.T) {
    // Setup
    emailNotifier := &MockNotifier{Name: "email", CanSendResponse: true}
    smsNotifier := &MockNotifier{Name: "sms", CanSendResponse: true}
    
    userRepo := &MockUserRepository{
        UserChannels: map[string][]string{
            "user123": {"email", "sms"},
        },
    }

    service := NewNotificationService(
        []domain.Notifier{emailNotifier, smsNotifier},
        userRepo,
    )

    notification := domain.Notification{
        UserID:  "user123",
        Title:   "Welcome",
        Message: "Welcome to our platform!",
        Type:    domain.WelcomeNotification,
    }

    // Execute
    err := service.SendNotification(notification)

    // Verify
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    if !emailNotifier.SendCalled {
        t.Error("Expected email notifier to be called")
    }

    if !smsNotifier.SendCalled {
        t.Error("Expected SMS notifier to be called")
    }

    if emailNotifier.LastNotification.UserID != "user123" {
        t.Error("Email notifier received wrong notification")
    }
}

func TestNotificationService_UserPreferencesError(t *testing.T) {
    userRepo := &MockUserRepository{ShouldError: true}
    service := NewNotificationService([]domain.Notifier{}, userRepo)

    err := service.SendNotification(domain.Notification{UserID: "user123"})

    if err == nil {
        t.Error("Expected error when user repository fails")
    }
}
```

---

## 🚀 Step 6: Application Setup (Dependency Injection)

```go
// cmd/api/main.go - Dependency Injection and Composition Root
package main

import (
    "log"
    "your-app/application"
    "your-app/domain"
    "your-app/infrastructure"
)

func main() {
    // Initialize all concrete implementations (details)
    emailNotifier := infrastructure.NewEmailNotifier("smtp.example.com", 587)
    smsNotifier := infrastructure.NewSMSNotifier("abc123-api-key")
    slackNotifier := infrastructure.NewSlackNotifier("https://hooks.slack.com/...")
    
    userPrefRepo := infrastructure.NewUserPreferenceRepository()

    // Dependency Injection - wiring everything together
    notifiers := []domain.Notifier{emailNotifier, smsNotifier, slackNotifier}
    notificationService := application.NewNotificationService(notifiers, userPrefRepo)

    // Use the service (this could be in HTTP handlers, CLI commands, etc.)
    welcomeNotification := domain.Notification{
        UserID:  "user123",
        Title:   "Welcome to Our App",
        Message: "Thank you for signing up!",
        Type:    domain.WelcomeNotification,
    }

    if err := notificationService.SendNotification(welcomeNotification); err != nil {
        log.Printf("Failed to send notification: %v", err)
    } else {
        log.Println("Notification sent successfully")
    }
}
```

---

## 🔄 Step 7: Adding a New Channel (Requirement Change)

Now the product manager says: *"We need to add WhatsApp notifications!"*

**With our architecture, this is trivial:**

```go
// infrastructure/whatsapp_notifier.go
package infrastructure

import "your-app/domain"

type WhatsAppNotifier struct {
    PhoneNumber string
    APIKey      string
}

func NewWhatsAppNotifier(phoneNumber, apiKey string) *WhatsAppNotifier {
    return &WhatsAppNotifier{
        PhoneNumber: phoneNumber,
        APIKey:      apiKey,
    }
}

func (w *WhatsAppNotifier) Send(notification domain.Notification) error {
    fmt.Printf("Sending WhatsApp to user %s: %s\n", 
        notification.UserID, notification.Message)
    return nil
}

func (w *WhatsAppNotifier) CanSend(notificationType domain.NotificationType) bool {
    return true // WhatsApp can send all types
}

func (w *WhatsAppNotifier) GetName() string {
    return "whatsapp"
}
```

**Only change needed in main.go:**
```go
func main() {
    // ... existing code ...
    whatsappNotifier := infrastructure.NewWhatsAppNotifier("+123456789", "whatsapp-key")
    
    // Just add to the existing list - no other changes needed!
    notifiers := []domain.Notifier{emailNotifier, smsNotifier, slackNotifier, whatsappNotifier}
    
    // ... rest unchanged ...
}
```

---

## ✅ Definition of Done - CHECKLIST

- [x] **Support multiple channels**: Email, SMS, Slack (and easily added WhatsApp)
- [x] **Unit tests**: Comprehensive testing with mocks
- [x] **Easy to add new channels**: Just implement `Notifier` interface
- [x] **Configurable per user**: User preference system implemented
- [x] **Error handling**: Proper error handling and logging
- [x] **Low coupling**: Business logic doesn't depend on concrete implementations
- [x] **High cohesion**: Each component has a single responsibility

## 🎯 Benefits Demonstrated

1. **Testable**: Mock dependencies easily
2. **Maintainable**: Changes to one channel don't affect others
3. **Extensible**: New channels added without modifying existing code
4. **Reliable**: Proper error handling and logging
5. **Business-focused**: Core logic expresses business rules clearly

This is exactly how professional Go applications are structured to meet real business requirements while maintaining code quality!
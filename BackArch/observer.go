package main

import "fmt"

// Observer interface defines the contract for all observers
type Observer interface {
	Update(news string)
	GetName() string
}

// Subject interface defines the contract for subjects
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyAll(news string)
}

// NewsSubscriber implements Observer interface
type NewsSubscriber struct {
	name string
}

func (ns *NewsSubscriber) Update(news string) {
	fmt.Printf("📰 %s received news: %s\n", ns.name, news)
}

func (ns *NewsSubscriber) GetName() string {
	return ns.name
}

// EmailSubscriber implements Observer interface
type EmailSubscriber struct {
	name  string
	email string
}

func (es *EmailSubscriber) Update(news string) {
	fmt.Printf("📧 %s (%s) received news via email: %s\n", es.name, es.email, news)
}

func (es *EmailSubscriber) GetName() string {
	return es.name
}

// MobileSubscriber implements Observer interface
type MobileSubscriber struct {
	name        string
	phoneNumber string
}

func (ms *MobileSubscriber) Update(news string) {
	fmt.Printf("📱 %s (%s) received news via mobile: %s\n", ms.name, ms.phoneNumber, news)
}

func (ms *MobileSubscriber) GetName() string {
	return ms.name
}

// NewsAgency implements Subject interface
type NewsAgency struct {
	observers []Observer
	news      string
}

func NewNewsAgency() *NewsAgency {
	return &NewsAgency{
		observers: make([]Observer, 0),
	}
}

func (na *NewsAgency) Register(observer Observer) {
	na.observers = append(na.observers, observer)
	fmt.Printf("✅ %s subscribed to news agency\n", observer.GetName())
}

func (na *NewsAgency) Unregister(observer Observer) {
	for i, obs := range na.observers {
		if obs.GetName() == observer.GetName() {
			na.observers = append(na.observers[:i], na.observers[i+1:]...)
			fmt.Printf("❌ %s unsubscribed from news agency\n", observer.GetName())
			return
		}
	}
}

func (na *NewsAgency) NotifyAll(news string) {
	fmt.Printf("📢 Broadcasting news to %d subscribers\n", len(na.observers))
	for _, observer := range na.observers {
		observer.Update(news)
	}
}

func (na *NewsAgency) SetNews(news string) {
	fmt.Printf("🔥 BREAKING NEWS: %s\n", news)
	na.news = news
	na.NotifyAll(news)
}

func (na *NewsAgency) GetNews() string {
	return na.news
}

// StockMarket implements Subject interface for stock updates
type StockMarket struct {
	observers []Observer
	price     float64
	symbol    string
}

func NewStockMarket(symbol string) *StockMarket {
	return &StockMarket{
		observers: make([]Observer, 0),
		symbol:    symbol,
	}
}

func (sm *StockMarket) Register(observer Observer) {
	sm.observers = append(sm.observers, observer)
	fmt.Printf("✅ %s subscribed to %s stock updates\n", observer.GetName(), sm.symbol)
}

func (sm *StockMarket) Unregister(observer Observer) {
	for i, obs := range sm.observers {
		if obs.GetName() == observer.GetName() {
			sm.observers = append(sm.observers[:i], sm.observers[i+1:]...)
			fmt.Printf("❌ %s unsubscribed from %s stock updates\n", observer.GetName(), sm.symbol)
			return
		}
	}
}

func (sm *StockMarket) NotifyAll(update string) {
	fmt.Printf("📊 Broadcasting stock update to %d subscribers\n", len(sm.observers))
	for _, observer := range sm.observers {
		observer.Update(update)
	}
}

func (sm *StockMarket) SetPrice(price float64) {
	sm.price = price
	update := fmt.Sprintf("%s stock price updated to $%.2f", sm.symbol, price)
	sm.NotifyAll(update)
}

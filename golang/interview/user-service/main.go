package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// domain

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// -------------------------------------------------------------------------

type IIdGenerator interface {
	Next() int
}

// простой, линейный генератор
// "поток-безопасный"
type IdGenerator struct {
	mu        sync.Mutex
	currentId int // можно ли "хитрить" с atomic?
}

func NewIdGenerator(startValue int) *IdGenerator {
	res := &IdGenerator{
		mu:        sync.Mutex{},
		currentId: startValue,
	}

	return res
}

func (self *IdGenerator) Next() int {
	self.mu.Lock()
	defer self.mu.Unlock()

	copy := self.currentId
	self.currentId += 1

	return copy
}

// -------------------------------------------------------------------------

// название лучше дать БОЛЕЕ конкретное, но название сервисе не будем менять
// поэтому текущий набор под интерфейсом назовем так...
type IServiceMetrics interface {
	IncCreate()
	IncCreateNotificationsSent()

	WriteAsJson(w io.Writer) error
}

type ServiceMetrics struct {
	mu      sync.RWMutex
	metrics map[string]int
}

func NewServiceMetrics() *ServiceMetrics {
	return &ServiceMetrics{
		mu:      sync.RWMutex{},
		metrics: make(map[string]int),
	}
}

func (self *ServiceMetrics) IncCreate() {
	self.mu.Lock()
	defer self.mu.Unlock()

	self.metrics["creates"]++
}

func (self *ServiceMetrics) IncCreateNotificationsSent() {
	self.mu.Lock()
	defer self.mu.Unlock()

	self.metrics["notifications_sent"]++
}

func (self *ServiceMetrics) WriteAsJson(w io.Writer) error {
	self.mu.RLock()
	defer self.mu.RUnlock()

	// тут надо вспомнить, а портит ли он w?
	jsonEncoder := json.NewEncoder(w)
	return jsonEncoder.Encode(self.metrics) // тут можно свой уровень ошибок
}

// -------------------------------------------------------------------------

// + тут смешали, наверно можно Service отдельно, Http "ручки" отдельно
type Service struct {
	userIdGenerator IIdGenerator
	metrics         IServiceMetrics

	// вообще можно заменить на sync.Map
	// ИЛИ, что более правильно, сделать отдельный класс IUserStorage +-.
	usersMu sync.RWMutex
	users   map[int]User
}

func NewService(userIdGenerator IIdGenerator, metrics IServiceMetrics) (*Service, error) {
	if userIdGenerator == nil {
		return nil, fmt.Errorf("userIdGenerator is nil")
	}

	return &Service{
		userIdGenerator: userIdGenerator,
		metrics:         metrics,

		users: make(map[int]User),
	}, nil
}

func (s *Service) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// ***

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	io.Copy(io.Discard, r.Body) // вычитать остальное, полезно или мусор?
	defer r.Body.Close()

	// ***

	user.ID = s.userIdGenerator.Next()

	s.usersMu.Lock()
	s.users[user.ID] = user
	s.usersMu.Unlock()

	s.metrics.IncCreate()

	// ***

	// асинхронно отправляем уведомление
	go func() {
		// имитация сетевого вызова
		time.Sleep(50 * time.Millisecond)

		// обновляем метрику
		s.metrics.IncCreateNotificationsSent()
	}()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (s *Service) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	io.Copy(io.Discard, r.Body)
	r.Body.Close()

	// ***

	// ВОТ это можно вынести в отдельный слой.
	// БУДЕТ как HTTPHandler ---> Service ---> UserStorage.
	// ПОКА пусть так...

	var id string = r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id required", http.StatusBadRequest)
		return
	}

	var userID int
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "id is not number", http.StatusBadRequest)
		return
	}

	// ***

	s.usersMu.RLock()
	user, ok := s.users[userID]
	s.usersMu.RUnlock()

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user) // игнор ошибки, по сути невозможно...
}

// -------------------------------------------------------------------------

func (s *Service) Metrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	io.Copy(io.Discard, r.Body)
	r.Body.Close() // всегда ли надо? или смысла нет?

	if err := s.metrics.WriteAsJson(w); err != nil {
		http.Error(w, "write metrics as json err: "+err.Error(),
			http.StatusInternalServerError)

		return
	}
}

// TODO:
/*
Управление горутинами — нет WaitGroup и context для корректного завершения фоновых уведомлений.
Graceful shutdown — нет корректной остановки HTTP-сервера и ожидания текущих операций.
Валидация входных данных — не проверяется Name на пустоту/длину.
Ограничение размера тела — нет http.MaxBytesReader.
Игнорирование ошибок записи ответа — json.NewEncoder(w).Encode в CreateUser и GetUser без проверки.
Разделение слоёв — HTTP-обработчики смешаны с бизнес-логикой, нет отдельного Handler.
Отсутствие интерфейса для хранилища — users map напрямую в сервисе, лучше вынести в UserStorage.
Генератор ID не возвращает ошибку — сигнатура Next() без error.
Проверка metrics == nil — в конструкторе не подставляются дефолтные метрики.
Логирование ошибок — нет логирования при сбоях в ответах или фоновых задачах.
*/

// NOTE:
/*
	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	type Service struct {
		mu      sync.Mutex
		users   map[int]User
		nextID  int
		metrics map[string]int
	}

	func NewService() *Service {
		return &Service{
			users:   make(map[int]User),
			metrics: make(map[string]int),
		}
	}

	func (s *Service) CreateUser(w http.ResponseWriter, r *http.Request) {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		s.mu.Lock()
		user.ID = s.nextID
		s.nextID++
		s.users[user.ID] = user
		s.metrics["creates"]++
		s.mu.Unlock()

		// асинхронно отправляем уведомление
		go func() {
			// имитация сетевого вызова
			time.Sleep(50 * time.Millisecond)
			// обновляем метрику
			s.mu.Lock()
			s.metrics["notifications_sent"]++
			s.mu.Unlock()
		}()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}

	func (s *Service) GetUser(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "id required", http.StatusBadRequest)
			return
		}

		// конвертируем без проверки ошибок
		var userID int
		fmt.Sscanf(id, "%d", &userID)

		s.mu.Lock()
		user, ok := s.users[userID]
		s.mu.Unlock()

		if !ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)
	}

	func (s *Service) Metrics(w http.ResponseWriter, r *http.Request) {
		s.mu.Lock()
		defer s.mu.Unlock()

		json.NewEncoder(w).Encode(s.metrics)
	}
*/

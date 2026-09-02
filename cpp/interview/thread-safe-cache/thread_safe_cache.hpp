#pragma once

#include <mutex>
#include <shared_mutex>
#include <stdexcept>
#include <string>
#include <unordered_map>
#include <functional>

// NOTE: оригинал
/*
template <typename K, typename V>
class ThreadSafeCache {
public:
    // Вставка или обновление значения по ключу
    void put(const K& key, const V& value) {
        std::lock_guard<std::mutex> lock(mutex_);
        data_[key] = value;
    }

    // Получение значения по ключу
    V& get(const K& key) {
        std::lock_guard<std::mutex> lock(mutex_);
        auto it = data_.find(key);
        if (it == data_.end()) {
            throw std::runtime_error("Key not found");
        }
        return it->second;
    }

    // Проверка наличия ключа
    bool contains(const K& key) {
        std::lock_guard<std::mutex> lock(mutex_);
        return data_.count(key) > 0;
    }

private:
    std::unordered_map<K, V> data_;
    mutable std::mutex mutex_;
};
*/

template <typename K, typename V>
class ThreadSafeCache {
public:
    // Вставка или обновление значения по ключу
    void put(const K& key, const V& value) {
        std::lock_guard<std::shared_mutex> lock(mutex_); // уникальная блокировка

        data_[key] = value;
    }

    // Получить копию значения по ключу
    V get_copy(const K& key) const {
        std::shared_lock<std::shared_mutex> lock(mutex_);

        auto it = data_.find(key);
        if (it == data_.end()) {
            throw std::runtime_error("Key not found");
        }

        return it->second;
    }

    void process_value_under_shared(
        const K& key,
        std::function<void(std::optional<V*>)> process_value
    ) {
        // проблема, отдавать ССЫЛКУ на изменение, опасно!

        std::shared_lock<std::shared_mutex> lock(mutex_);


    }

    // Проверка наличия ключа
    bool contains(const K& key) const {
        std::shared_lock<std::shared_mutex> lock(mutex_);

        return data_.count(key) > 0;
    }

private:
    std::unordered_map<K, V> data_;
    mutable std::shared_mutex mutex_;
};

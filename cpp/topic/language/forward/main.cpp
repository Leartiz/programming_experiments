#include <chrono>
#include <iostream>
#include <utility>
#include <vector>

// --- 1) корректность: & vs && ---

struct F {
    void operator()(int x) & { std::cout << "lvalue F, x=" << x << '\n'; }
    void operator()(int x) && { std::cout << "rvalue F, x=" << x << '\n'; }
};

template <class Fn>
void call_forward(Fn&& fn) {
    std::forward<Fn>(fn)(1);
}

template <class Fn>
void call_plain(Fn&& fn) {
    fn(1);
}

// --- 2) скорость: move vs copy ---

struct Heavy {
    std::vector<int> data;

    static int copies;
    static int moves;

    Heavy() : data(1'000'000) {}

    Heavy(const Heavy& other) : data(other.data) { ++copies; }
    Heavy(Heavy&& other) noexcept : data(std::move(other.data)) { ++moves; }
};

int Heavy::copies = 0;
int Heavy::moves = 0;

template <class T>
void take_forward(T&& x) {
    Heavy local = std::forward<T>(x);  // rvalue -> move
    (void)local;
}

template <class T>
void take_plain(T&& x) {
    Heavy local = x;  // имя x - lvalue -> всегда copy
    (void)local;
}

template <class Fn>
long long bench_ms(Fn fn, int n) {
    using clock = std::chrono::steady_clock;
    const auto t0 = clock::now();
    for (int i = 0; i < n; ++i) {
        fn();
    }
    const auto t1 = clock::now();
    return std::chrono::duration_cast<std::chrono::milliseconds>(t1 - t0).count();
}

int main() {
    std::cout << "=== & / && ===\n";
    F f;
    call_forward(f);
    call_forward(F{});
    call_plain(f);
    call_plain(F{});

    std::cout << "\n=== move vs copy (Heavy, 1e6 ints) ===\n";
    constexpr int N = 50;

    Heavy::copies = Heavy::moves = 0;
    const auto ms_fwd = bench_ms([] { take_forward(Heavy{}); }, N);
    std::cout << "forward: " << ms_fwd << " ms, copies=" << Heavy::copies
              << " moves=" << Heavy::moves << '\n';

    Heavy::copies = Heavy::moves = 0;
    const auto ms_plain = bench_ms([] { take_plain(Heavy{}); }, N);
    std::cout << "plain:   " << ms_plain << " ms, copies=" << Heavy::copies
              << " moves=" << Heavy::moves << '\n';
}

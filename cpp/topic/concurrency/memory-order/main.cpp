#include <atomic>
#include <iostream>
#include <thread>
#include <vector>

// 1) relaxed: atomicity only (counter).
// 2) release/acquire: publish non-atomic payload via a flag.
// 3) seq_cst: single global order of seq_cst ops (store buffering).

static void demo_relaxed_counter()
{
    constexpr int kThreads = 8;
    constexpr int kPerThread = 100'000;

    std::atomic<int> counter{0};

    std::vector<std::thread> threads;
    threads.reserve(kThreads);

    for (int t = 0; t < kThreads; ++t) {
        threads.emplace_back([&] {
            for (int i = 0; i < kPerThread; ++i) {
                counter.fetch_add(1, std::memory_order_relaxed);
            }
        });
    }

    for (auto& th : threads) {
        th.join();
    }

    std::cout << "[relaxed] counter = " << counter.load(std::memory_order_relaxed)
              << " (expected " << (kThreads * kPerThread) << ")\n";
}

static void demo_release_acquire_publish()
{
    int data = 0;
    std::atomic<bool> ready{false};

    std::thread writer([&] {
        data = 42;
        // release: everything written before this store is visible
        // to a thread that does acquire-load on the same atomic.
        ready.store(true, std::memory_order_release);
    });

    std::thread reader([&] {
        while (!ready.load(std::memory_order_acquire)) {
            // spin until published
        }
        // acquire saw release-store => data == 42 is guaranteed
        std::cout << "[release/acquire] data = " << data << " (expected 42)\n";
    });

    writer.join();
    reader.join();
}

// Store buffering (SB):
//   A: x=true;  r1=y
//   B: y=true;  r2=x
// Can both see false?  seq_cst: no.  release+acquire: formally yes.
template <std::memory_order StoreOrder, std::memory_order LoadOrder>
static int count_both_false(int iterations)
{
    int both_false = 0;

    for (int i = 0; i < iterations; ++i) {
        std::atomic<bool> x{false};
        std::atomic<bool> y{false};
        bool r1 = true;
        bool r2 = true;

        std::thread a([&] {
            x.store(true, StoreOrder);
            r1 = y.load(LoadOrder);
        });
        std::thread b([&] {
            y.store(true, StoreOrder);
            r2 = x.load(LoadOrder);
        });

        a.join();
        b.join();

        if (!r1 && !r2) {
            ++both_false;
        }
    }

    return both_false;
}

static void demo_seq_cst_vs_acq_rel()
{
    constexpr int kIters = 50'000;

    const int seq =
        count_both_false<std::memory_order_seq_cst, std::memory_order_seq_cst>(kIters);
    const int acq_rel =
        count_both_false<std::memory_order_release, std::memory_order_acquire>(kIters);

    std::cout << "[seq_cst] both loads false: " << seq
              << " / " << kIters << " (must be 0)\n";
    std::cout << "[release/acquire] both loads false: " << acq_rel
              << " / " << kIters
              << " (allowed; often 0 on x86, may be >0 on ARM)\n";
}

int main()
{
    demo_relaxed_counter();
    demo_release_acquire_publish();
    demo_seq_cst_vs_acq_rel();
}

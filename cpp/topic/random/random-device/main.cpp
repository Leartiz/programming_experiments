#include <iostream>
#include <random>

int main() {
    std::random_device rd;
    for (int i = 0; i < 3; ++i) {
        std::cout << "random_device: " << rd() << '\n';
    }

    std::mt19937 gen(rd());
    std::uniform_int_distribution<int> dist(1, 6);

    std::cout << "random_device seed sample: " << rd() << '\n';
    std::cout << "dice: " << dist(gen) << '\n';
}

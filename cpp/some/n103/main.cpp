#include <iostream>
#include <iterator>
#include <functional>
#include <algorithm>
#include <random>

namespace utils
{

namespace compare
{

template<typename T>
using compare_t = std::function<bool(T, T)>;

template<typename T>
bool gt(T lhs, T rhs) {
    return lhs > rhs;
}

template<typename T>
bool lt(T lhs, T rhs) {
    return lhs < rhs;
}

// ***

template<typename T>
struct Gt {
    bool operator()(T lhs, T rhs) {
        return lhs > rhs;
    }
};

template<typename T>
struct Lt {
    bool operator()(T lhs, T rhs) {
        return lhs < rhs;
    }
};

} // compare

namespace find
{

template<typename T>
T min(T val) {
    std::cout << "T min(T val)" << std::endl;

    return val;
}

template<typename T, typename ...Args>
T min(T lhs, Args... restArgs) {
    std::cout << "T min(T lhs, Args... restArgs)" << std::endl;

    const T rhs = min(restArgs...); // ?
    return lhs > rhs ? rhs : lhs;
}

// ***

template<typename T, typename Compare = compare::Lt<T>>
T num(T val) {
    return val;
}

template<typename T, typename Compare = compare::Lt<T>, typename ...Args>
T num(T lhs, Args... restArgs) {
    const T rhs = num<T, Compare>(restArgs...); // ?
    const T result = Compare()(lhs, rhs) ? lhs : rhs;
    return result;
}

} // find

namespace ds
{

struct LinkedList {
    struct Node {
        Node* next;
        std::string val;

        explicit Node(std::string val) : next{ nullptr },
            val{ std::move(val) } {}
    };

    LinkedList() : beg{ nullptr }, end{ nullptr } {}

    ~LinkedList() {
        clear();
    }

    void clear() {
        while (beg != nullptr) {
            auto temp = beg;
            beg = beg->next;
            delete temp;
        }

        beg = nullptr;
        end = nullptr;
    }

    void insert(std::string val) {
        Node* newNode = new Node(std::move(val));
        if (beg == nullptr) {
            beg = newNode;
            end = newNode;
        } else {
            end->next = newNode;
            end = newNode;
        }
    }

    void println() const {
        auto temp = beg;
        while (temp != nullptr) {
            std::cout << temp->val << "\n";
            temp = temp->next;
        }

        std::cout << std::endl;
    }

    int size() const {
        auto temp = beg;
        int sz = 0;
        while (temp != nullptr) {
            temp = temp->next;
            ++sz;
        }
        return sz;
    }

    Node *beg, *end;
};

} // ds

namespace algo
{

std::vector<int> gen_vec(size_t n, int min = 0, int max = 1000)
{
    std::default_random_engine e;
    static std::uniform_int_distribution<> dist(min, max);
    std::vector<int> vec; vec.reserve(n);
    for (decltype(n) i = 0; i < n; ++i)
        vec.push_back(dist(e));
    return vec;
}

template<typename T>
void println_vec(std::vector<T>& v, const std::string& sep = " ") {
    std::copy(v.begin(), v.end(),
              std::ostream_iterator<T>(std::cout, sep.c_str()));
    std::cout << std::endl;
}

template<typename T, typename Compare = compare::Lt<T>>
void bubble_sort(std::vector<T>& vec) {

    bool swapped = true;
    while (swapped) {
        swapped = false;

        for (size_t i = 0; i < vec.size() - 1; ++i) {

            if (Compare()(vec[i], vec[i+1])) {
                std::swap(vec[i], vec[i+1]);
                swapped = true;
            }
        }
    }
}

template<typename T, typename Compare = compare::Lt<T>>
void quick_sort(std::vector<T>& vec, size_t beg = 0, size_t end = -1) {

}

} // algo

} // utils

// -----------------------------------------------------------------------

int main()
{
    {
        std::cout << "less: "    << (std::less<int>()(1, 2) ? 1 : 2) << "\n";
        std::cout << "greater: " << (std::greater<int>()(1, 2) ? 1 : 2) << "\n";

        // ***

        std::cout << "num greater: " <<
            utils::find::num<int, std::greater<int>>(1, 2, 3, 4) << std::endl;

        std::cout << "num less: " <<
            utils::find::num<int, std::less<int>>(1, 2, 3, 4) << std::endl;

        // ***

        std::cout << utils::find::min(1, 2, 3, 4) << std::endl;
    }
    std::cout << std::endl;
    {
        utils::ds::LinkedList ll;
        std::cout << "ll.size: " << ll.size() << std::endl;
        ll.println();

        ll.insert("123");
        ll.insert("hii");
        ll.insert("c++");
        std::cout << "ll.size: " << ll.size() << std::endl;
        ll.println();

        ll.clear();
        std::cout << "ll.size: " << ll.size() << std::endl;
        ll.println();
    }
    std::cout << std::endl;
    {
        std::vector<int> v{ 1, 3, 2, 4, 1, -100 };
        utils::algo::bubble_sort<int>(v);

        std::copy(v.begin(), v.end(),
                  std::ostream_iterator<int>(std::cout, " "));
        std::cout << std::endl;

        utils::algo::bubble_sort<int, utils::compare::Gt<int>>(v);
        utils::algo::println_vec(v);
    }
    return 0;
}

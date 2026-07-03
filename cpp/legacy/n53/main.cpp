#include <iostream>
#include <vector>
#include <map>
#include <exception>
#include <functional>
#include <chrono>
#include <numeric>
#include <cassert>
#include <cmath>

bool isEqual(double a, double b, double eps = 0.00001) {
    return abs(a-b) < eps;
}

double sum_vec(std::vector<double>& vec) {
    return std::accumulate(vec.begin(), vec.end(), 0);
}

enum class FunctionType {
    NoException, Normal, FullException
};

//! решение уравнения при a, b, c одновременно не равных нулю
std::vector<double> solve_correct_equation(double a, double b, double c) noexcept {
    assert(!(isEqual(a, 0) && isEqual(b, 0) && isEqual(c, 0)));

    if (isEqual(a, 0) && isEqual(b, 0)) {
        return std::vector<double>(0);
    }

    if (isEqual(a,0)) {
        return std::vector<double> {0, -c/b};
    }

    double discriminant = (b*b) - (4*a*c);

    if (isEqual(discriminant, 0)) {
        return std::vector<double> {0, -b/(2*a)};
    }
    if (discriminant < 0) {
        return std::vector<double>(0);
    }

    return std::vector<double> {
        (-b + sqrt(discriminant))/(2*a),
        (-b - sqrt(discriminant))/(2*a)
    };
}

std::vector<double> solve_no_exception(double a, double b, double c, bool& ok) noexcept {
    ok = true;

    if (isEqual(a, 0) && isEqual(b, 0) && isEqual(c, 0)) {
        ok = false;
        return std::vector<double>(0);
    }

    return solve_correct_equation(a, b, c);
}

std::vector<double> solve(double a, double b, double c) {
    if (isEqual(a, 0) && isEqual(b, 0) && isEqual(c, 0)) {
        throw std::runtime_error("root is any value");
    }

    return solve_correct_equation(a, b, c);
}

void solve_full_exception(double a, double b, double c) {
    if (isEqual(a, 0) && isEqual(b, 0) && isEqual(c, 0)) {
        throw std::runtime_error("root is any value");
    }

    throw solve_correct_equation(a, b, c);
}

double roots_sum_no_exception(double a, double b, double c) noexcept {
    bool ok;
    auto roots = solve_no_exception(a, b, c, ok);
    if (true == ok) {
        return sum_vec(roots);
    }
    return 0;
}

double roots_sum(double a, double b, double c) noexcept {
    try {
        auto roots = solve(a, b, c);
        return sum_vec(roots);
    }
    catch(std::runtime_error&) {
        return 0;
    }
}

double roots_sum_full_exception(double a, double b, double c) {
    try {
        solve_full_exception(a, b, c);
    }
    catch(std::runtime_error&) {
        return 0;
    }
    catch(std::vector<double>& roots) {
        return sum_vec(roots);
    }
    throw std::runtime_error("wrong exception type");
}

//! вызывает функцию в соответствии с типом
double call_solver(FunctionType type, double a, double b, double c) noexcept {
    static const std::map<
        FunctionType,
        std::function<double(double, double, double)>
        > type2function = {
            {FunctionType::NoException, roots_sum_no_exception},
            {FunctionType::Normal, roots_sum},
            {FunctionType::FullException, roots_sum_full_exception}
        };

    return type2function.at(type)(a, b, c);
}

void run(unsigned long n, FunctionType type) noexcept {
    auto begin = std::chrono::steady_clock::now();

    double sum = 0;
    //#pragma omp parallel for reduction (+: sum)
    for (unsigned long i = 0; i < n; i++) {
        double a = ((i%2000)-1000)/33.0;
        double b = ((i%200)-100)/22.0;
        double c = ((i%20)-10)/11.0;
        sum += call_solver(type, a, b, c);
    }

    auto end = std::chrono::steady_clock::now();

    auto elapsed_ms = std::chrono::duration_cast<std::chrono::milliseconds>(end - begin);

    std::cout << n << "\t" << elapsed_ms.count() << std::endl;
}

int main() {
    const int equation_count = 4000000;

    std::cout << "n\ttime (ms)" << std::endl;
    std::cout << "\t\t\t--------No exception-----------" << std::endl;
    run(equation_count, FunctionType::NoException);

    std::cout << "\t\t\t--------Normal-----------" << std::endl;
    run(equation_count, FunctionType::Normal);

    std::cout << "\t\t\t--------Full exception-----------" << std::endl;
    run(equation_count, FunctionType::FullException);
}

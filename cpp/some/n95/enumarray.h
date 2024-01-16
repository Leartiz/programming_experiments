#ifndef ENUMARRAY_H
#define ENUMARRAY_H

#include <utility>
#include <iostream>

#include <initializer_list>
#include <array>

template<typename Enum, typename T>
class EnumArray
{
public:
    EnumArray(std::initializer_list<std::pair<Enum, T>> values)
    {
        std::cout << values.size() << "\n";
        for (const auto& val : values) {
            data[static_cast<std::size_t>(val.first)] =
                val.second;
        }
    }

    T& operator[](Enum key)
    {
        return data[static_cast<std::size_t>(key)];
    }
    const T& operator[](Enum key) const
    {
        return data[static_cast<std::size_t>(key)];
    }

private:
    // ---> use std::to_underlying
    static constexpr std::size_t N =
        static_cast<std::size_t>(Enum::Count);
    std::array<T, N> data;
};

#endif // ENUMARRAY_H

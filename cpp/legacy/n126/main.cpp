#include <cstddef>
#include <ostream>
#include <sstream>
#include <fstream>
#include <string_view>
#include <variant>

#include <unistd.h>
#include <limits.h>

#include <iostream>
#include <functional>
#include <memory>
#include <string>
#include <thread>
#include <unordered_map>

#include <ppconsul/coordinate.h>

#include <boost/uuid/random_generator.hpp>
#include <boost/uuid/uuid.hpp>
#include <boost/uuid/uuid_io.hpp>

#include <boost/asio/dispatch.hpp>
#include <boost/asio/strand.hpp>
#include <boost/filesystem.hpp>
#include <boost/asio.hpp>
#include <utility>
#include <vector>

using NodeRtt = double;
using NodeName = std::string;

using NodeRttStorage = std::unordered_map<NodeName, NodeRtt>;

struct ICoordinates {};

// -----------------------------------------------------------------------

struct Coordinate
{
    double adjustment;
    double error;
    double height;
    std::vector<double> vec;
};

struct Node
{
    std::string node;
    std::string segment;
    Coordinate coord;
};

struct ConsulCoordinates : ICoordinates {
    std::vector<Node> internal;
};

using NodeRttQuery = std::function<std::unique_ptr<ICoordinates>()>;
using NodeRttCalculator = std::function<NodeRttStorage(const std::string & self_node_name,
                                                       const ICoordinates& coordinates)>;

// -----------------------------------------------------------------------

struct NodeRttGetterImpl
{
    NodeRttQuery coordinate_query;
    NodeRttCalculator rtt_calc;
};

// -----------------------------------------------------------------------

std::string get_hostname()
{
    std::string hostname(HOST_NAME_MAX + 1, 0);
    if (gethostname(hostname.data(), hostname.size()) == 0) {
        return hostname;
    }

    return "localhost";
}

std::string gen_id(const std::string& prefix = "prefix") {
    return prefix + boost::uuids::to_string(
        boost::uuids::random_generator()()
    );
}

// -----------------------------------------------------------------------

class Base {
public:
    void foo() {
        std::cout << "foo" << std::endl;
    }
};

class Derived : protected Base {
public:
    operator Base&() {
        return static_cast<Base&>(*this);
    }
};

class External {
public:
    External(Base& b) {
        b.foo();
    }
};

int get_int(std::shared_ptr<std::pair<int, int>> p)
{
    std::cout << "use count: " << p.use_count() << std::endl;
    return p->first + p->second;
}

inline constexpr auto message_flags_offset = 48;
inline constexpr auto message_id_mask = 0x0000FFFFFFFFFFFF;

using MessageId = uint64_t;
using MessageFlags = uint8_t;
using ModuleId = uint16_t;

#pragma pack(push, 1)
struct Header
{
    uint8_t command      = 0;
    uint8_t dynamic_data_len = 0;
    uint16_t src_sme     = 0;
    uint16_t dst_sme     = 0;
    uint16_t ttl         = 0;
    uint64_t ref1        = 0;
    uint32_t ref2        = 0;
    uint32_t param       = 0;
    uint32_t len         = 0;
};
#pragma pack(pop)

inline void set_message_meta(Header & header,
    MessageId msg_id, MessageFlags flags) noexcept
{
    header.ref1 = (MessageId)flags << message_flags_offset
                  | (msg_id & message_id_mask);
}

using MyCollection = std::variant<
    std::string,
    std::vector<int>
>;

// -----------------------------------------------------------------------

struct Storage {
    virtual int make_number() const = 0;
};

class SecretStorage : public Storage {
private:
    int make_number() const override {
        return 111;
    }
public:
    int get_number() const {
        return make_number();
    }
};

int main(int /*argc*/, char** /*argv*/)
{
    {
        SecretStorage* ss{ new SecretStorage  };
        std::cout << "make number: " << ss->make_number() << std::endl;

        Storage* s{ ss };
        std::cout << "make number: " << s->make_number() << std::endl;
    }
    {
        MyCollection mc = std::vector<int>{ 1, 2, 3, 5, 100, 999 };
        std::visit([](const auto& sv) -> void {

            using T = std::decay_t<decltype(sv)>;

            if constexpr (std::is_same_v<T, std::string>) {
                std::cout << "Str: " << std::string{ sv } << std::endl;
            }
            else {
                const auto& vec = std::vector<int>{ sv };
                for (size_t i = 0; i < vec.size(); ++i) {
                    if (i == vec.size() - 1) {
                        std::cout << vec[i];
                    }
                    else {
                        std::cout << vec[i] << ", ";
                    }
                }
                std::cout << std::endl;
            }
        },
        mc);

        try
        {
            std::cout << std::get<std::string>(mc) << std::endl;

            std::string& ref_str = std::get<std::string>(mc);
            ref_str = "555";

            std::cout << std::get<std::string>(mc) << std::endl;

            auto& v = std::get<0>(mc);
            v.resize(100);
        }
        catch (const std::bad_variant_access& ex)
        {
            std::cout << ex.what() << '\n';
        }
    }
    {
        Header header;
        header.ref1 >> message_flags_offset;
    }
    {
        std::exception_ptr eptr;
        if (eptr == nullptr) {
            std::cout << "eptr is nullptr" << std::endl;
        }
    }

    {
        auto p = std::make_shared<std::pair<int, int>>(2, 2);
        std::cout << "use count: " << p.use_count() << std::endl;
        get_int(std::move(p));
        std::cout << "use count: " << p.use_count() << std::endl;
    }
    {
        {
            std::ifstream fin("1.txt", std::ifstream::ios_base::app);
            if (!fin.is_open()) {
                std::string error_message = "Error opening file: " + std::string(strerror(errno));
                std::cout << error_message << std::endl;
            }
        }
        {
            std::ifstream fin("2.txt");
            if (!fin.is_open()) {
                std::string error_message = "Error opening file: " + std::string(strerror(errno));
                std::cout << error_message << std::endl;
            }
        }
        //...
    }
    /*
    {
        std::unordered_map<std::string, std::string> m;
        {
            auto [it, ignored] = m.try_emplace("key", "value1");
            std::cout << it->first << ": " << it->second << std::endl;
            std::cout << "ignored: " << ignored << std::endl;
        }
        {
            auto [it, ignored] = m.try_emplace("key", "value2");
            std::cout << it->first << ": " << it->second << std::endl;
            std::cout << "ignored: " << ignored << std::endl;
        }
        {
            std::string str{ "123" };
            std::istringstream ist{ std::move(str) };
            std::cout << "str: " << str << std::endl;
        }
    }
    {
        boost::system::error_code ec;
        const auto t1 = boost::filesystem::last_write_time("123.txt", ec);
        std::cout << ec << std::endl;
        std::cout << t1 << std::endl;

        const auto t2 = boost::filesystem::last_write_time("123.txt");
        std::cout << t2 << std::endl;
    }
    {
        const auto ep = boost::asio::ip::tcp::endpoint();
        std::cout << ep << std::endl;
    }
    {
        std::cout << get_hostname() << std::endl;

        std::cout << gen_id() << std::endl;
        std::cout << gen_id() << std::endl;
        std::cout << gen_id() << std::endl;
    }
    {
        boost::asio::io_context ioc;
        boost::asio::io_context::strand strand{ ioc };

        for (int i = 0; i < 10; ++i) {
            boost::asio::post(strand, [i](){
                std::cout << i << std::endl;
            });
        }

        const auto job_count = 4;
        std::vector<std::thread> jobs;
        jobs.reserve(job_count);
        for (int i = 0; i < job_count; ++i) {
            jobs.emplace_back([&ioc]{
                ioc.run();
            });
        }

        for (int i = 0; i < job_count; ++i) {
            jobs[i].join();
        }
    }
    */
	return 0;
}
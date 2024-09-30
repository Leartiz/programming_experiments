#include <string>
#include <vector>
#include <bitset>
#include <iostream>
#include <winsock2.h>
#include <ws2tcpip.h>

class Solution {
public:
    vector<string> readBinaryWatch(int num) {
        vector<string> rs;
        for (int h = 0; h < 12; h++)
            /*
                0 = 0000
                1 = 0001
                2 = 0010
                3 = 0011
                4 = 0100
                _
                5 = 0101
                6 = 0110
                7 = 0111
                8 = 1000
                _
                9  = 1001
                10 = 1010
                11 = 1011
                12 = 1100
            */
            for (int m = 0; m < 60; m++)
                if (bitset<10>(m << 4 | h).count() == size_t(num))
                    rs.emplace_back(to_string(h) + (m < 10 ? ":0" : ":") + to_string(m));
        return rs;
    }
};

void printVec(const std::vector<string>& vec) {
    for (size_t i = 0; i < vec.size(); ++i)
        std::cout << vec[i] << " ";
    std::cout << std::endl;
}

int main()
{
<<<<<<< HEAD
    Solution s;
    printVec(s.readBinaryWatch(1));
    // printVec(s.readBinaryWatch(2));
    // printVec(s.readBinaryWatch(3));
    // printVec(s.readBinaryWatch(4));
=======
    WSAData wsaData;
    if (WSAStartup(MAKEWORD(1, 1), &wsaData) != 0) {
        std::cerr << "WSAStartup failed." << std::endl;
        return 1;
    }

    // ***

    std::cout << "iMaxSockets: " << wsaData.iMaxSockets << std::endl;
    std::cout << "iMaxUdpDg: " << wsaData.iMaxUdpDg << std::endl;

    // With error!
    //std::cout << "lpVendorInfo: " << wsaData.lpVendorInfo << std::endl;

    std::cout << "szDescription: " << wsaData.szDescription << std::endl;
    std::cout << "szSystemStatus: " << wsaData.szSystemStatus << std::endl;
    std::cout << "wHighVersion: " << wsaData.wHighVersion << std::endl;
    std::cout << "wVersion: " << wsaData.wVersion << std::endl;

    // ***

    {
        std::cout << "sizeof(u_long): " << sizeof(u_long) << std::endl;
        u_long a = 1;

        std::cout << "a: " << a << std::endl;
        a = htonl(a);
        std::cout << "a: " << a << std::endl;
        a = ntohl(a);
        std::cout << "a: " << a << std::endl;
    }

    // ***

    {
        addrinfo *res;

        /*
            Функция getaddrinfo в языке программирования C
            используется для получения информации
            об адресе сетевого узла, соответствующего
            заданному имени хоста и номеру порта.
        */
        int error = getaddrinfo("www.example.com", NULL, NULL, &res); // ???
        if (error != 0) {
            perror("get addr info failed!"); // !
            return EXIT_FAILURE;
        }

        auto step = res;
        while (step != nullptr) {
            step = step->ai_next;

            std::cout << "res->ai_flags: "     << res->ai_flags << std::endl;
            std::cout << "res->ai_family: "    << res->ai_family << std::endl;
            std::cout << "res->ai_socktype: "  << res->ai_socktype << std::endl;
            std::cout << "res->ai_protocol: "  << res->ai_protocol << std::endl;

            // With error!
            //std::cout << "res->ai_canonname: " << res->ai_canonname << std::endl;

            sockaddr *sa = res->ai_addr;
            if (sa != nullptr) {
                std::cout << "res->ai_addrlen: "         << res->ai_addrlen << std::endl;
                std::cout << "res->ai_addr->sa_data: "   << sa->sa_data << std::endl;
                std::cout << "res->ai_addr->sa_family: " << sa->sa_family << std::endl;
            } else {
                std::cout << "sa == <null>" << std::endl;
            }
        }
    }

    // ***

    {
        in_addr ia;
        ia.S_un.S_addr = inet_addr("127.0.0.1");

        std::cout << "ia.S_un.S_un_b.s_b1: " << int(ia.S_un.S_un_b.s_b1) << std::endl;
        std::cout << "ia.S_un.S_un_b.s_b2: " << int(ia.S_un.S_un_b.s_b2) << std::endl;
        std::cout << "ia.S_un.S_un_b.s_b3: " << int(ia.S_un.S_un_b.s_b3) << std::endl;
        std::cout << "ia.S_un.S_un_b.s_b4: " << int(ia.S_un.S_un_b.s_b4) << std::endl;
        //...
    }

    // ***

    {
        struct sockaddr_in sa;
        struct sockaddr_in6 sa6;

        // Новый способ!

        if (inet_pton(AF_INET, "192.0.2.1", &(sa.sin_addr)) <= 0)
            std::cerr << "inet printable to network v4 failed!" << std::endl;
        if (inet_pton(AF_INET6, "2001:db8:63b3:1::3490", &(sa6.sin6_addr)) <= 0)
            std::cerr << "inet printable to network v6 failed!" << std::endl;
    }

    // ***

    {
        // IPv4
        {
            std::cout << "INET_ADDRSTRLEN: " << INET_ADDRSTRLEN << std::endl;

            in_addr ia;
            inet_pton(AF_INET, "192.0.2.2", &(ia));

            // equal!
            std::cout << "&ia: " << &ia << std::endl;
            std::cout << "&(ia.S_un.S_addr): " << &(ia.S_un.S_addr) << std::endl;

            char ip4[INET_ADDRSTRLEN];
            inet_ntop(AF_INET, &(ia), ip4, INET_ADDRSTRLEN);
            std::cout << "IPv4 addr: " << ip4 << std::endl;
        }

        // IPv6
        {
            std::cout << "INET6_ADDRSTRLEN: " << INET6_ADDRSTRLEN << std::endl;

            in6_addr ia6;
            inet_pton(AF_INET6, "2001:db8:63b3:1::3490", &(ia6));

            std::cout << "&ia6: " << &ia6 << std::endl;

            char ip6[INET6_ADDRSTRLEN];
            inet_ntop(AF_INET6, &(ia6), ip6, INET6_ADDRSTRLEN);
            std::cout << "IPv6 addr: " << ip6 << std::endl;
        }
    }

    // ***

    {
        struct sockaddr_in sa;
        struct sockaddr_in6 sa6;

        std::cout << "INADDR_ANY: " << INADDR_ANY << std::endl;

        // ?
        //std::cout << "in6addr_any: " << in6addr_any << std::endl;

        sa.sin_addr.S_un.S_addr = INADDR_ANY;
        sa6.sin6_addr = in6addr_any;

        std::cout << "sa.sin_addr.S_un.S_addr: " << sa.sin_addr.S_un.S_addr << std::endl;

        for (int i = 0; i < 8; ++i) {
            std::cout << "sa6.sin6_addr.u.Word[" << i << "]: "
                      << sa6.sin6_addr.u.Word[i] << std::endl;
        }
    }

    // ***

    {
        auto he = gethostbyname("www.ya.ru");
        if (!he) {
            std::cout << "get host by name failed" << std::endl;
        } else {
            std::cout << "he->h_addrtype: " << he->h_addrtype << std::endl;
            std::cout << "he->h_length: " << he->h_length << std::endl;
            std::cout << "he->h_name: " << he->h_name << std::endl;

            for (short i = 0; i < he->h_length; ++i) {

                const auto bufSize = 255*sizeof(char); // !
                char* buf = static_cast<char*>(malloc(bufSize));
                memset(buf, 0, bufSize);

                sprintf(buf, "he->h_addr_list[%d]: ", i);
                std::cout << buf << he->h_addr_list[i] << std::endl;

                memset(buf, 0, bufSize);
                sprintf(buf, "he->h_aliases[%d]: ", i);
                std::cout << buf << he->h_aliases[i] << std::endl;

                free(buf); // !
            }
        }
    }

    // ***

    {
        struct addrinfo *res;
        /*
            - поиск имён DNS;
            - заполняет нужные структуры.
        */
        getaddrinfo("www.ya.ru", "http", nullptr, &res);
        std::cout << "res->ai_addrlen: " << res->ai_addrlen << std::endl;
    }

    // ***

    std::cout << "[OK]" << std::endl;

    const auto cleanupCode = WSACleanup();
    std::cout << "cleanup code: "
              << cleanupCode << std::endl;

>>>>>>> d26a70116fb70142de5c81e6998d458a135dc7b7
    return 0;
}

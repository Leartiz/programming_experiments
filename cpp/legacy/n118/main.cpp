#include <cstdio>
#include <iostream>

#include <winsock2.h>
#include <ws2tcpip.h>

int main() {
    WSADATA wsaData;
    if (WSAStartup(MAKEWORD(2, 2), &wsaData) != 0) {
        std::cerr << "WSAStartup failed: " << WSAGetLastError() << '\n';
        return 1;
    }

    const char* ipStr = "192.168.1.1";
    struct sockaddr_in addr;

    if (inet_pton(AF_INET, ipStr, &addr.sin_addr) <= 0) {
        std::cerr << "inet_pton error: " << WSAGetLastError() << '\n';
        WSACleanup();
        return 1;
    }

    std::cout << "Binary IP address: " << addr.sin_addr.S_un.S_addr << '\n';

    char ipBuffer[INET_ADDRSTRLEN];
    if (inet_ntop(AF_INET, &addr.sin_addr, ipBuffer, sizeof(ipBuffer)) == nullptr) {
        std::cerr << "inet_ntop error: " << WSAGetLastError() << '\n';
        WSACleanup();
        return 1;
    }

    std::cout << "String IP address: " << ipBuffer << '\n';

    WSACleanup();
    return 0;
}

#include "usrsctp.h"
#include <unistd.h> // для close() на POSIX

#include <iostream>
#include <cstring>
#include <cstdlib>

#define SERVER_PORT 2905
#define BUFFER_SIZE 1024

// Функция обратного вызова для обработки событий SCTP
int receive_callback(struct socket *sock, union sctp_sockstore addr, void *data, size_t datalen, struct sctp_rcvinfo rcv, int flags, void *ulp_info) {
    if (data) {
        // Выводим данные, переданные клиентом
        std::cout << "Получено сообщение (" << datalen << " байт): " << std::string(static_cast<char *>(data), datalen) << std::endl;

        // Отправляем ответ клиенту
        std::string reply = "Message received successfully!";
        usrsctp_sendv(sock, reply.c_str(), reply.length(), nullptr, 0, nullptr, 0, SCTP_SENDV_NOINFO, 0);

        // Освобождаем ресурс
        free(data);
    }
    return 1;
}

int main() {
    // Инициализация usrsctp
    usrsctp_init(0, nullptr, nullptr);

    // Создание серверного сокета
    struct socket *server_socket = usrsctp_socket(AF_INET6, SOCK_SEQPACKET, IPPROTO_SCTP, receive_callback, nullptr, 0, nullptr);
    if (!server_socket) {
        std::cerr << "Не удалось создать SCTP сокет!" << std::endl;
        return -1;
    }

    // Установка параметров сокета
    struct linger linger_opt = {1, 0};
    usrsctp_setsockopt(server_socket, SOL_SOCKET, SO_LINGER, &linger_opt, sizeof(linger_opt));

    // Создаем адрес для сокета
    struct sockaddr_in6 server_addr{};
    memset(&server_addr, 0, sizeof(server_addr));
    server_addr.sin6_family = AF_INET6;
    server_addr.sin6_port = htons(SERVER_PORT);
    server_addr.sin6_addr = in6addr_any;

    // Привязываем сокет
    if (usrsctp_bind(server_socket, reinterpret_cast<struct sockaddr *>(&server_addr), sizeof(server_addr)) < 0) {
        std::cerr << "Ошибка привязки сокета." << std::endl;
        return -1;
    }

    // Слушаем входящие соединения
    if (usrsctp_listen(server_socket, 5) < 0) {
        std::cerr << "Ошибка при попытке прослушивания." << std::endl;
        return -1;
    }

    std::cout << "RUN sctp server: " << SERVER_PORT << std::endl;

    // Останавливаем обработку с задержкой (например, для тестов - 60 секунд)
    sleep(60);

    // Завершаем работу
    usrsctp_close(server_socket);
    usrsctp_finish();

    return 0;
}

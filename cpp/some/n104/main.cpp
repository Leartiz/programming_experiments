#include <stdio.h>
#include <winsock2.h>

#define PORT 18080

int main() {

    WSADATA wsa;
    SOCKET server_socket, client_socket;

    struct sockaddr_in server, client;
    int addrlen = sizeof(server);
    char hello[] = "hi man!"; // Отправляем один байт на клиента

    // Initialize Winsock
    if (WSAStartup(MAKEWORD(2, 2), &wsa) != 0) {
        printf("WSAStartup failed\n");
        return 1;
    }

    // Create a socket
    if ((server_socket = socket(AF_INET, SOCK_STREAM, 0)) == INVALID_SOCKET) {
        printf("Socket creation failed\n");
        return 1;
    }

    server.sin_family      = AF_INET;
    server.sin_addr.s_addr = INADDR_ANY;
    server.sin_port        = htons(PORT);

    // Bind the socket
    if (bind(server_socket, (struct sockaddr*)&server, sizeof(server)) == SOCKET_ERROR) {
        printf("Bind failed\n");
        return 1;
    }

    // Listen for incoming connections
    if (listen(server_socket, 3) == SOCKET_ERROR) {
        printf("Listen failed\n");
        return 1;
    }

    printf("Server started, waiting for incoming connections...\n");

    while(true) {
        // Accept incoming connection
        if ((client_socket = accept(server_socket, (struct sockaddr*)&client, &addrlen)) == INVALID_SOCKET) {
            printf("Accept failed\n");
            return 1;
        }

        // Send data to client
        send(client_socket, hello, sizeof(hello), 0);
        printf("Byte successfully sent to client\n");

        closesocket(client_socket);
    }

    closesocket(server_socket);
    WSACleanup();

    return 0;
}

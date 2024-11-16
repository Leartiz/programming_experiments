#include <boost/asio.hpp>

#include <iostream>
#include <sstream>

#include <functional>
#include <memory>


using boost::asio::ip::tcp;

class Session : public std::enable_shared_from_this<Session> {
public:
    using toggle_handler = std::function<void(const std::string&)>;

public:
    Session(tcp::socket socket,
                     const toggle_handler& on_about_start,
                     const toggle_handler& on_about_stop)

        : socket_(std::move(socket))
        , on_about_start_{ on_about_start }
        , on_about_stop_{ on_about_stop }
    {

        const auto ep = socket_.remote_endpoint();

        std::ostringstream oss;
        oss << ep.address().to_string() << ":" << ep.port();
        addr_ = oss.str();

        // ***

        on_about_start_(addr_);
    }

    void start() { do_read(); }
    const std::string& get_addr() const { return addr_; }

private:
    void do_read() {
        auto self(shared_from_this());
        socket_.async_read_some(
            boost::asio::buffer(data_),
            [this, self](boost::system::error_code ec, std::size_t length) {
                if (!ec) {
                    std::cout << "Received: " << std::string(data_, length) << std::endl;
                    do_write(length);
                } else if (ec == boost::asio::error::eof){

                    on_about_stop_(addr_);
                }
            });
    }

    void do_write(std::size_t length) {
        auto self(shared_from_this());
        boost::asio::async_write(
            socket_, boost::asio::buffer(data_, length),
            [this, self](boost::system::error_code ec, std::size_t) {
                if (!ec) {
                    do_read();
                }
            });
    }

private:
    std::string addr_;
    tcp::socket socket_;
    char data_[1024] { 0 };

private:
    toggle_handler on_about_start_;
    toggle_handler on_about_stop_;
};

class Server {
public:
    Server(boost::asio::io_context& io_context, short port)
        : acceptor_(io_context, tcp::endpoint(tcp::v4(), port)) {
        do_accept();
    }

private:
    void do_accept() {
        acceptor_.async_accept(
            [this](boost::system::error_code ec, tcp::socket socket) {
                if (!ec) {
                    auto on_about_start =
                        [](const std::string& addr) -> void {
                        std::cout << "Client connected: " << addr << std::endl;
                    };
                    auto on_about_stop =
                        [](const std::string& addr) -> void {
                        std::cout << "Client disconnected: " << addr << std::endl;
                    };

                    const auto session = std::make_shared<Session>(
                        std::move(socket), on_about_start, on_about_stop);

                    // ***

                    session->start();

                } else {
                    std::cerr << ec.message() << std::endl;
                }

                do_accept(); // registering a new task!
            });
    }

    tcp::acceptor acceptor_;
};

int main() {
    const int16_t port{ 12345 };
    {
        const auto ep = tcp::endpoint(tcp::v4(), port);

        std::cout << "endpoint address: " << ep.address() << std::endl;
        std::cout << "endpoint port: " << ep.port() << std::endl;

        std::cout << "endpoint size: " << ep.size() << std::endl;
        std::cout << "endpoint capacity: " << ep.capacity() << std::endl;

        std::cout << "endpoint data: " << ep.data() << std::endl;

        // ***

        const auto ep_protocol = ep.protocol();
        std::cout << "ep_protocol type: " << ep_protocol.type() << std::endl;
        std::cout << "ep_protocol family: " << ep_protocol.family() << std::endl;
        std::cout << "ep_protocol protocol: " << ep_protocol.protocol() << std::endl;

        std::cout << "ep_protocol v4 protocol: " << ep_protocol.v4().protocol() << std::endl;
        std::cout << "ep_protocol v6 protocol: " << ep_protocol.v6().protocol() << std::endl;
    }

    // ***

    try {
        boost::asio::io_context io_context;
        Server server(io_context, port);  // Порт 12345
        io_context.run();

    } catch (std::exception& e) {
        std::cerr << "Exception: " << e.what() << "\n";
    }

    return 0;
}

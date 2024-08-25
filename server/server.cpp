#include "server.hpp"
#include <iostream>
#include <thread>
#include <chrono>
#include <atomic>
#include <csignal>

std::atomic<bool> keep_running(true);

void signalHandler(int signum) {
    keep_running = false;
}

int main() {
    SomeIPServer server;

    std::cout << "Server running" << std::endl;

    std::signal(SIGINT, signalHandler);  // Handle Ctrl+C for graceful shutdown

    while (keep_running) {
        std::this_thread::sleep_for(std::chrono::seconds(1));
    }

    std::cout << "\nServer shutting down..." << std::endl;

    return 0;
}

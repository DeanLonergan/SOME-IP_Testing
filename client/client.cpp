#include "client.hpp"
#include <iostream>

SomeIPClient::SomeIPClient(SomeIPServer& server) : server_(server) {
    server_.subscribe([this](const std::string& name, const std::string& value) {
        std::cout << "Signal received: " << name << " = " << value << std::endl;
    });
}

void SomeIPClient::requestBooleanSignal() {
    std::string signal = server_.getBooleanSignal();
    std::cout << "Boolean signal: " << signal << std::endl;
}

void SomeIPClient::requestVariableSignal() {
    int signal = server_.getVariableSignal();
    std::cout << "Variable signal: " << signal << std::endl;
}

void SomeIPClient::changeBooleanSignal(const std::string& signal) {
    server_.setBooleanSignal(signal);
}

void SomeIPClient::changeVariableSignal(int signal) {
    server_.setVariableSignal(signal);
}

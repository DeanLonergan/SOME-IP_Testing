#ifndef CLIENT_HPP
#define CLIENT_HPP

#include "server.hpp"

class SomeIPClient {
public:
    SomeIPClient(SomeIPServer& server);

    void requestBooleanSignal();
    void requestVariableSignal();

    void changeBooleanSignal(const std::string& signal);
    void changeVariableSignal(int signal);

private:
    SomeIPServer& server_;
};

#endif // CLIENT_HPP

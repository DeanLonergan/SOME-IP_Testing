#ifndef SERVER_HPP
#define SERVER_HPP

#include <string>
#include <functional>

class SomeIPServer {
public:
    using SignalCallback = std::function<void(const std::string&, const std::string&)>;

    void subscribe(SignalCallback callback);
    void setBooleanSignal(const std::string& signal);   // "on" or "off"
    void setVariableSignal(int signal);                 // 0, 1, 2, or 3

    std::string getBooleanSignal() const;
    int getVariableSignal() const;

private:
    SignalCallback callback_;

    std::string boolean_signal_ = "off";  // Default to "off"
    int variable_signal_ = 0;             // Default to 0

    void notify(const std::string& signal_name, const std::string& signal_value);
};

#endif // SERVER_HPP

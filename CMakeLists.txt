cmake_minimum_required(VERSION 3.10)
project(SOME_IP_Testing)

set(CMAKE_CXX_STANDARD 17)

include_directories(${CMAKE_SOURCE_DIR}/server ${CMAKE_SOURCE_DIR}/client)

# Build the server executable
add_executable(server server/server.cpp)

# Build the client executable
add_executable(client client/client.cpp)

# Link the client directly with the server's object file
target_link_libraries(client PRIVATE server/server.cpp)

add_executable(client ${ROOT_PATH}/src/client/client.cpp)
target_link_libraries(client PUBLIC
                vsomeip3 
                ${Boost_LIBRARIES})
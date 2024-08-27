add_executable(client client.cpp)
target_link_libraries(client PUBLIC
                vsomeip3 
                ${Boost_LIBRARIES})

#${ROOT_PATH}/src/client/
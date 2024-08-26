

add_executable(service ${ROOT_PATH}/src/service/service.cpp)
target_link_libraries(service PUBLIC 
                vsomeip3 
                ${Boost_LIBRARIES})
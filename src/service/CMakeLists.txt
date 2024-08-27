add_executable(service service.cpp)
target_link_libraries(service PUBLIC 
                vsomeip3 
                ${Boost_LIBRARIES})
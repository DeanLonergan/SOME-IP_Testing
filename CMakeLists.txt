cmake_minimum_required (VERSION 3.13)

set (CMAKE_CXX_FLAGS "-g -std=c++0x")

find_package (vsomeip3 3.4.10 REQUIRED)
find_package( Boost 1.55 COMPONENTS system thread log REQUIRED )
set(ROOT_PATH ${CMAKE_CURRENT_SOURCE_DIR})
add_subdirectory(src/service)
add_subdirectory(src/routingmanagerd)
cmake_minimum_required(VERSION 3.29)
project(algorithms C)

set(CMAKE_C_STANDARD 11)

include_directories(skiena-4)

add_executable(algorithms
  skiena-4/unit_test.c
)

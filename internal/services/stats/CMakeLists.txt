# Specifies the required CMake version.
cmake_minimum_required(VERSION 3.22 FATAL_ERROR)

# Defines the project name.
project(stats
        VERSION 1.0.0
        DESCRIPTION "Shortlink service for work with stats"
        HOMEPAGE_URL "https://github.com/shortlink-org/shortlink"
)

set(CMAKE_CXX_STANDARD 14)

include(GenerateExportHeader)
include(GNUInstallDirs)

# Includes the contents of the conanbuildinfo.cmake file.
include(${CMAKE_BINARY_DIR}/conanbuildinfo.cmake)
# Prepares the CMakeList.txt for Conan.
conan_basic_setup()

# add all files of source folder
file(GLOB SOURCES
    src/*.h
    src/*.cpp
)

# Specifies the executable to build.
add_executable(${PROJECT_NAME} ${SOURCES})
# Specifies what libraries to link, using Conan.
target_link_libraries(${PROJECT_NAME} ${CONAN_LIBS})


def load_conan_dependencies():
    
    native.new_local_repository(
        name="cli11",
        path="/Users/user/.conan/data/cli11/1.9.1/_/_/package/5ab84d6acfe1f23c4fae0ab88f26e3a396351ac9",
        build_file="/Users/user/myprojects/shortlink/internal/services/stats/deps/cli11/BUILD",
    )
    
    
    native.new_local_repository(
        name="fmt",
        path="/Users/user/.conan/data/fmt/7.1.3/_/_/package/80138d4a58def120da0b8c9199f2b7a4e464a85b",
        build_file="/Users/user/myprojects/shortlink/internal/services/stats/deps/fmt/BUILD",
    )
    
    
    native.new_local_repository(
        name="prometheus-cpp",
        path="/Users/user/.conan/data/prometheus-cpp/1.0.0/_/_/package/b39e1754fc610f750a6d595455854696692ec5bc",
        build_file="/Users/user/myprojects/shortlink/internal/services/stats/deps/prometheus-cpp/BUILD",
    )
    
    
    native.new_local_repository(
        name="civetweb",
        path="/Users/user/.conan/data/civetweb/1.15/_/_/package/77e8df9f2be98ef80d2a9f31ea49eb14597b20b0",
        build_file="/Users/user/myprojects/shortlink/internal/services/stats/deps/civetweb/BUILD",
    )
    
    
    native.new_local_repository(
        name="libcurl",
        path="/Users/user/.conan/data/libcurl/7.86.0/_/_/package/a097455223234e250d01a2687cf7c15446fbd5d5",
        build_file="/Users/user/myprojects/shortlink/internal/services/stats/deps/libcurl/BUILD",
    )
    
    
    native.new_local_repository(
        name="zlib",
        path="/Users/user/.conan/data/zlib/1.2.13/_/_/package/6841fe8f0f22f6fa260da36a43a94ab525c7ed8d",
        build_file="/Users/user/myprojects/shortlink/internal/services/stats/deps/zlib/BUILD",
    )
    
    
    native.new_local_repository(
        name="openssl",
        path="/Users/user/.conan/data/openssl/1.1.1s/_/_/package/6841fe8f0f22f6fa260da36a43a94ab525c7ed8d",
        build_file="/Users/user/myprojects/shortlink/internal/services/stats/deps/openssl/BUILD",
    )


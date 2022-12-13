from conans import ConanFile


class MyConanFile(ConanFile):
    requires = "cli11/1.9.1", "fmt/7.1.3", "prometheus-cpp/1.0.0"
    generators = "cmake", "BazelDeps", "BazelToolchain"


load("@rules_cc//cc:defs.bzl", "cc_import", "cc_library")


cc_import(
    name = "prometheus-cpp-pull_precompiled",
    static_library = "lib/libprometheus-cpp-pull.a",
)

cc_import(
    name = "prometheus-cpp-push_precompiled",
    static_library = "lib/libprometheus-cpp-push.a",
)

cc_import(
    name = "prometheus-cpp-core_precompiled",
    static_library = "lib/libprometheus-cpp-core.a",
)




cc_library(
    name = "prometheus-cpp",
    
    hdrs = glob(["include/**"]),
    
    
    includes = ["include"],
    
    
    
    visibility = ["//visibility:public"],
    
    deps = [
        # do not sort
    
    ":prometheus-cpp-pull_precompiled",
    
    ":prometheus-cpp-push_precompiled",
    
    ":prometheus-cpp-core_precompiled",
    
    
    
    "@civetweb",
    
    "@libcurl",
    
    "@zlib",
    
    ],
    
)

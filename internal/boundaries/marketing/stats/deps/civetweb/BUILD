
load("@rules_cc//cc:defs.bzl", "cc_import", "cc_library")


cc_import(
    name = "civetweb-cpp_precompiled",
    static_library = "lib/libcivetweb-cpp.a",
)

cc_import(
    name = "civetweb_precompiled",
    static_library = "lib/libcivetweb.a",
)




cc_library(
    name = "civetweb",
    
    hdrs = glob(["include/**"]),
    
    
    includes = ["include"],
    
    
    
    visibility = ["//visibility:public"],
    
    deps = [
        # do not sort
    
    ":civetweb-cpp_precompiled",
    
    ":civetweb_precompiled",
    
    
    
    "@openssl",
    
    ],
    
)

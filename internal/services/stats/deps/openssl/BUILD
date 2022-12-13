
load("@rules_cc//cc:defs.bzl", "cc_import", "cc_library")


cc_import(
    name = "ssl_precompiled",
    static_library = "lib/libssl.a",
)

cc_import(
    name = "crypto_precompiled",
    static_library = "lib/libcrypto.a",
)




cc_library(
    name = "openssl",
    
    hdrs = glob(["include/**"]),
    
    
    includes = ["include"],
    
    
    
    visibility = ["//visibility:public"],
    
    deps = [
        # do not sort
    
    ":ssl_precompiled",
    
    ":crypto_precompiled",
    
    
    
    ],
    
)

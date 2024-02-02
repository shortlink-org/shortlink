
load("@rules_cc//cc:defs.bzl", "cc_import", "cc_library")


cc_import(
    name = "z_precompiled",
    static_library = "lib/libz.a",
)




cc_library(
    name = "zlib",
    
    hdrs = glob(["include/**"]),
    
    
    includes = ["include"],
    
    
    
    visibility = ["//visibility:public"],
    
    deps = [
        # do not sort
    
    ":z_precompiled",
    
    
    
    ],
    
)

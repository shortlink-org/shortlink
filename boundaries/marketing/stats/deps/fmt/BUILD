
load("@rules_cc//cc:defs.bzl", "cc_import", "cc_library")


cc_import(
    name = "fmt_precompiled",
    static_library = "lib/libfmt.a",
)




cc_library(
    name = "fmt",
    
    hdrs = glob(["include/**"]),
    
    
    includes = ["include"],
    
    
    
    visibility = ["//visibility:public"],
    
    deps = [
        # do not sort
    
    ":fmt_precompiled",
    
    
    
    ],
    
)

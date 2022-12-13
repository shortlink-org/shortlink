
load("@rules_cc//cc:defs.bzl", "cc_import", "cc_library")


cc_import(
    name = "curl_precompiled",
    static_library = "lib/libcurl.a",
)




cc_library(
    name = "libcurl",
    
    hdrs = glob(["include/**"]),
    
    
    includes = ["include"],
    
    
    defines = ["CURL_STATICLIB=1"],
    
    
    visibility = ["//visibility:public"],
    
    deps = [
        # do not sort
    
    ":curl_precompiled",
    
    
    
    "@zlib",
    
    ],
    
)

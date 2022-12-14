module.exports = {
    rules: {
        "check-author": {
            create: function (context) {
                return {
                    Program: function (node) {
                        const comments = context.getAllComments();
                        for (const comment of comments) {
                            if (comment.type === "Block" && comment.value.includes("@author")) {
                                return;
                            }
                        }

                        context.report(node, "Missing @author tag in file comments.");
                    }
                };
            }
        },
        "check-empty-file": {
            create: function (context) {
                return {
                    Program: function (node) {
                        const sourceCode = context.getSourceCode();
                        const text = sourceCode.getText();

                        if (text.trim() === "") {
                            context.report(node, "File is empty.");
                        }
                    }
                };
            }
        }
    }
};
import { Rule } from "eslint";

const rules: Record<string, Rule.RuleModule> = {
  "check-author": {
    meta: {
      type: "problem",
      docs: {
        description: "Check if a file has an author tag",
        category: "Possible Errors",
        recommended: true
      },
      schema: []
    },
    create: (context: Rule.RuleContext) => {
      return {
        Program: (node: any) => {
          // @ts-ignore
          const comments = context.getAllComments();
          for (const comment of comments) {
            if (comment.type === "Block" && comment.value.includes("@author")) {
              return;
            }
          }

          context.report({ node, message: "Missing @author tag in file comments." });
        }
      };
    }
  },
  "check-empty-file": {
    meta: {
      type: 'problem',
      docs: {
        description: 'Check if a file is empty',
        category: 'Possible Errors',
        recommended: false,
      },
      schema: [],
    },
    create: (context: Rule.RuleContext) => {
      return {
        Program: (node: any) => {
          const sourceCode = context.getSourceCode();
          const text = sourceCode.getText();

          if (text.trim() === '') {
            context.report({ node, message: 'File is empty.' });
          }
        },
      };
    },
  }
};

export default rules;

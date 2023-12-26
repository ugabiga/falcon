import type {CodegenConfig} from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: "../api/graph/*.graphql",
  documents: "src/graph/query/*.graphql",
  generates: {
    "src/graph/generated/generated.ts": {
      plugins: [
        'typescript',
        'typescript-operations',
        'typed-document-node',
      ],
    }
  }
};

export default config;